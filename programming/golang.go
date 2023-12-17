/* Primitives */
// byte = uint8, int = int64, uint = uint64
// bool, float64, rune = int32, string
type Year int // attach different sets of methods to the same underlying data

/* Constants */
// no immutable definitions at run time => use copy, otherwise use pointer
// compile-time typed constant for primitive types, not structs
const i int = 1
// literal values are untyped constants
const i = 1 // untyped constant is an alias to a literal
type Model int // enumeration = const list + iota
const (Nextra Model = iota; Omnia; Selecta; Prime; Spectrum)

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
// call-by-value ensures immutability of original data
// pointers imply mutable data
// local data is stored on heap when its pointer is returned from a function
// limit use of pointers to store more data on stack
// pointer initialization and dereferencing
var i int = 1; var p *int = &i
i++; *p++ // i == 3; *p == 3
var p *int = new(int) // p* == 0

/* Arrays */
// allocated on stack as array size is know at compile-time
var a [3]int // zero initialized array, copy-by-value
var b = [3]int{1, 2, 3} // initialized array
var c = [...]int{1, 5: 2, 3, 9: 4} // sparse array

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
for i, v := range c { fmt.Println(i, v) }

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
// nominal typing, block-level scope, copy-by-value
type Bayan struct{model string; year int}
var a = Bayan{model: "Selecta", year: 2023}
var b = Bayan{"Selecta", 2023} // all fields in order
var c *Bayan = new(Bayan)
var d *Bayan = &Bayan{}
fmt.Println(c.model) // (*c).model => c.model
var e = []Bayan{{model: "Selecta", year: 2023}}
// constructor function
func NewBayan(model string, year int) *Bayan {
	return &Bayan{model, year}
}
// structural typing, anonymous struct
var e = struct{model string; year int}{"Nextra", 2023}
type Product struct {Bayan; price float64} // embedded field
var p = Product{Bayan: Bayan{"Selecta", 2023}, price: 2e4}

/* Blocks */
// package block = definitions outside a function
// file block = imported definitions
// function block = top-level function definitions and parameters
// syntax block = in-function {...} and control structures
// shadowing = inner block same-name definition, unaccessible outer definition

/* if/else */
// each clause has its own scope and allows for arbitrary conditions
if i := rand.Intn(10); i < 3 { // statement-wide definitions
	fmt.Println(i, "low")
} else if i < 8 {
	fmt.Println(i, "mid")
} else {
	fmt.Println(i, "high")
}

/* for + break/continue [label] */
// best for controlled iteration start, end, and step
for i := 0; i < 5; i++ { fmt.Println(i) }
i := 0 // dynamic exit conditions
for i < 5 { fmt.Println(i); i++ }
i = 0 // unconditional first iteration
for { fmt.Println(i); i++; if i > 4 { break } }
// best for arrays, slices, strings, maps
for _, v := range []int{1, 2, 3, 4, 5} { fmt.Println(v) }

/* switch + break [label] */
outer: for _, v := range a {
	switch l := len(v); l { // equality check on a value
	case 1, 2, 3: fmt.Println(v, "small")
	case 4, 5, 6: fmt.Println(v, "medium"); break outer
	default: fmt.Println(v, "large")
	}
}
for _, v := range a {
	switch l := len(v); { // arbitrary conditions in clauses
	case l < 4: fmt.Println(v, "small")
	case l < 7: fmt.Println(v, "medium")
	default: fmt.Println(v, "large")
	}
}

/* goto label */
// prefer goto over flow control flags or code duplication
for _, v := range []int{1, 2, 3, 4, 5} {
	if v == 3 { goto print }
	v *= 10
	print: fmt.Println(v)
}

/* functions */
// function = logic that depends only on input parameters
// call-by-value = all function arguments are copies (no need for immutability)
// all primitive and composite types are value types
// slice and map values are pointers
// named return values
func quotRem(a, b int) (quot, rem int) {
	quot, rem = a / b, a % b
	return
}
// variadic parameters
func sum(args ...int) (sum int) {
	for arg := range args { sum += arg }
	return
}
fmt.Println(sum([]int{1, 2, 3}...))
// function type, functions are values
type Op func(int, int) int
ops := map[string]Op{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
}
// anonymous functions (function literals) are closures
// closed on variables are evaluated every time when a closure is invoked
sort.Slice(bayans, func(i, j int) bool {
	return bayans[i].model < bayans[j].model
})
// defer closures are evaluated after return in reverse order
defer f.Close()
defer func() {
	if err == nil { err = tx.Commit() } else { tx.Rollback() }
}() // defer must end with ()

/* methods */
// method = logic that depends on a value of a type
// method modifies a receiver => must use a pointer receiver (p *T)
// method does not modify a receiver => can use a value receiver (v T)
// value receiver panics on nil
func (b *Bayan) String() string {
	return fmt.Sprintf("Bayan(%s %d)", b.model, b.year)
}
b := Bayan{"Selecta", 2023}
p := &b
fmt.Println(b.String(), p.String()) // b.String() => (&b).String()
bstr := b.String // method value = closure over its instance
fmt.Println(bstr())
bstr := (*Bayan).String // method expression => function(receiver)
fmt.Println(bstr(&b))
// composition = filed embedding
// fields or methods of an embedded type are promoted to a containing type
// methods of an embedded type are in a method set of a containing type
type Product struct {Bayan; price float64}
p := Product{Bayan{"Nextra", 2023}, 1e3}
fmt.Println(p.model, p.String())

/* interfaces */
// interface = type-safe structural typing when a method set of a concrete type
// contains a method set of an interface
// an interface can be embedded into another interface
// accept interfaces, return structs, not interfaces
type Presenter interface {show()}
// a type is decoupled from the implicit Presenter interface
func (b Bayan) show() { fmt.Println(b) }
// Only a client specifies the required Presenter interface
var p Presenter = Bayan{"Selecta", 2023}
p.show()
var a interface{} // empty interface can store a value of any type
a = Bayan{"Nextra", 2023}
// nominal type assertion applied to an interface at runtime v.(Type)
// vs type conversions on concrete types at compile-time Type(v)
if b, ok := a.(Bayan); ok { fmt.Println(b) } // type assertion
switch b := a.(type) { // type switch
case Bayan: fmt.Println(b)
default: fmt.Println("unknown type")
}
