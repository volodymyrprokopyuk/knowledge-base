/* Primitives */
// byte = uint8, int = int64, uint = uint64
// bool, float64, rune = int32, string

/* Constants */
const i int = 1 // typed constant
// literal values are untyped constants
const i = 1 // untyped constant is an alias to a literal
const (a = iota; b; c) // 0, 1, 2

/* Variables */
var i, j int = 1, 2
var b, s, c = false, "ok", 'a' // type inference
// default type of literals int, float64
var i, f = 1, 1.2
// default initialization to zero
var (b bool; i int; f float64; c rune; s string)
// in-function inferred initialization/redefinition
b, i, f, c, s := true, 1, 1.2, 'a', `ok`

/* Pointers */
// pointer initialization, dereferencing
i := 1; p := &i; i++; *p++ // i == 3; *p == 3

/* Arrays */
var a [3]int // zero initialized array, copy by value
var b = [3]int{1, 2, 3} // initialized array
var c = [...]int{1, 5: 2, 3, 9: 4} // sparse array
for i, v := range b { fmt.Println(i, v) }

/* Slices */
// slice = multiple views onto array, change, but not append
// slice = single view per array, change and append
var a []int // nil slice = dynamic array
a = append(a, 1, 2) // dynamic reallocation
var b = []int{1, 2, 3} // initialized slice
var c = []int{1, 5: 2, 3, 9: 4} // sparse slice
d := make([]int, 5) // zero initialized slice [0:4]
e := make([]int, 0, 5) // empty slice with capacity
var a = [...]int{1, 2, 3, 4, 5}
b := a[:] // array => slice (shared memory), [i:j)
c := make([]int, 5) // copy destination must be initialized
copy(c, a[:]) // array => slice (two copies), [:] == [0:len(a)]

/* Strings and runes */
var s = []rune("€12.34") // Unicode code points, default bytes
fmt.Println(string(s[0]), string(s[1:]))
for i, r := range s { fmt.Println(i, string(r)) }

/* Maps */
var a = map[string]int{} // empty map
var b = map[string]int{"a": 1, "b": 2} // initialized map
var c = make(map[string]int, 5) // map with capacity
c["a"] = 1
if v, ok := c["a"]; ok { fmt.Println(v) }
for k, v := range b { fmt.Println(k, v) }
delete(c, "a")

/* Structs */
// nominal typing, block-level scope, copy by value
type Bayan struct{model string; year int}
var a = Bayan{model: "Selecta", year: 2023}
var b = Bayan{"Selecta", 2023} // all fields in order
var c *Bayan = new(Bayan) // equivalent to &Bayan{}
fmt.Println(c.model) // (*c).model => c.model
var d = []Bayan{{model: "Selecta", year: 2023}}
func NewBayan(model string, year int) *Bayan {
	return &Bayan{model, year} // constructor function
}
// structural typing, anonymous struct
var e = struct{model string; year int}{"Nextra", 2023}
type Product struct {Bayan; price float64} // embedded field
var p = Product{Bayan: Bayan{"Selecta", 2023}, price: 2e4}
