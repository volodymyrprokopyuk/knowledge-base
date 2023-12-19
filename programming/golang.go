/* Primitives */
// bool, int = int64, uint = uint64, float64
// byte = uint8, rune = int32, string
// type alias = attach different methods to the same underlying type
type Year int // nominal typing, block-level scope
var year = Year(2023)

/* Constants, enumerations */
const i int = 1 // compile-time typed constant only for primitive types
// no immutable definitions at runtime => use call-by-value
// for mutable definitions at runtime => use pointers
const i = 1 // compile-time untyped constant = named literal
type BayanModel int // enumeration = type + typed const list + iota
const (Nextra BayanModel = iota; Omnia; Selecta; Prime; Spectrum)

/* Variables */
var i, j int = 1, 2
var i, f = 1, 1.2 // default type of literals int, float64
var b, s, r = false, "ok", 'a' // type inference
// default initialization to zero
var (b bool; i int; f float64; r rune; s string)
// in-function inferred initialization/redefinition
b, i, f, r, s := true, 1, 1.2, 'a', `ok`

/* Pointers */
// call-by-value preserves immutability of original data
// use of pointers implies mutable data
// pointer initialization & address and * dereferencing
var i int = 1; var p *int = &i; i++; *p++ // i == 3; *p == 3
var p *int = new(int) // p* == 0

/* Strings, runes */
var bytes = "€12.34" // string as bytes + ASCII
fmt.Println(bytes[0], bytes[1:]) // 226 12.34
var runes = []rune("€12.34") // Unicode code points
fmt.Println(string(runes[0]), string(runes[1:])) // € 12.34
for _, r := range runes { fmt.Println(string(r)) } // Unicode code points

/* Arrays */
// allocated on stack as array size is known at compile-time
var a [3]int // [0 0 0] automatic zero initialization
// explicit initialization
b, c := [3]int{}, [...]int{1, 2, 3} // [0 0 0], [1 2 3]
d := [...]int{1, 2: 2, 4: 3} // [1 0 2 0 3] sparse array

/* Slices */
// slice = multiple views onto an array, change, but not append
// slice = single view per array, change and append
var a []int // nil slice = dynamic array
a = append(a, 1, 2) // [1 2] dynamic reallocation
b := []int{1, 2, 3} // [1 2 3] slice initialization
c := []int{1, 2: 2, 4: 3} // [1 0 2 0 3] sparse slice
d := make([]int, 3) // [0 0 0] zero initialized slice
e := make([]int, 0, 3) // [] empty slice with capacity
a := [...]int{1, 2, 3}
var b []int = a[:] // array => slice, shared memory
var c []int = make([]int, len(a))
copy(c, a[:]) // array => slice, separate copies
for i, v := range a { fmt.Println(i, v) }

/* Maps */
var a = map[string]int{} // empty map
var b = map[string]int{"a": 1, "b": 2} // initialized map
var c = make(map[string]int, 3) // map with capacity
c["a"] = 1
if v, ok := c["a"]; ok { fmt.Println(v) }
for k, v := range b { fmt.Println(k, v) }
delete(c, "a")

/* Structs, types */
// nominal typing, block-level scope
  var a, b Bayan = Bayan{model: "Nextra", year: 2022}, Bayan{"Omnia", 2023}
  var c, d *Bayan = new(Bayan), &Bayan{"Selecta", 2024}
  fmt.Println(d.model) // (*d).model => d.model
  e := []Bayan{{model: "Nextra", year: 2022}, {"Omnia", 2023}}
// constructor function = correct and uniform initialization of structs
// constructor function is preferred over a struct literal
// local value is allocated on heap when its pointer is returned from a function
func NewBayan(model string, year int) *Bayan { return &Bayan{model, year} }
// structural typing, anonymous struct
var a Bayan = struct { model string; year int }{"Nextra", 2022}
// struct embedding = composition of concrete types
// multiple types can be embedded into a new type
type Product struct { Bayan; price float64 }
p := Product{Bayan{"Nextra", 2022}, 2e4}
// Bayan fields are either directly available on Product through promotion
// or indirectly accessible through a Product type
// fields of nested embedded types are promoted to a top-level containing type
fmt.Println(p.model, p.price, p.Bayan.year)

/* Blocks */
// package block = definitions outside a function
// file block = imported definitions
// function block = top-level function definitions and parameters
// syntax block = in-function {...} and control structures
// shadowing = inner block same-name definition, unaccessible outer definition

/* if/else */
// allows for arbitrary conditions in each clause
// each clause has its own scope
if i := rand.Intn(10); i < 3 { fmt.Println(i, "low")
} else if i < 8 { fmt.Println(i, "mid")
} else { fmt.Println(i, "high") }

/* for + break/continue [label] */
// best for controlled iteration defining start, end, and step
for i := 0; i < 3; i++ { fmt.Println(i) } // 0 1 2
i := 0 // best for dynamic exit condition
for i < 3 { fmt.Println(i); i++ } // 0 1 2
i = 0 // unconditional first iteration
for { fmt.Println(i); i++; if i > 2 { break } } // 0 1 2
// best for arrays, slices, strings, maps
for _, v := range []int{1, 2, 3} { fmt.Println(v) } // 1 2 3

/* switch + break [label] */
a := []string{"one", "eleven", "thousand"}
outer: for _, v := range a {
  switch l := len(v); l { // equality check == in each clause
  case 1, 2, 3: fmt.Println(v, "small")
  case 4, 5, 6: fmt.Println(v, "medium"); break outer
  default: fmt.Println(v, "large")
  }
}
for _, v := range a {
  switch l := len(v); { // arbitrary conditions in clause clause
  case l < 4: fmt.Println(v, "small")
  case l < 7: fmt.Println(v, "medium")
  default: fmt.Println(v, "large")
  }
}

/* goto label */
// prefer goto over flow control flags or code duplication
for _, v := range []int{1, 2, 3} {
  if v == 2 { goto print }
  v *= 10
  print: fmt.Println(v) // 10, 2, 30
}

/* functions */
// call-by-value = all function arguments are copies, no need for immutability
// primitive and composite types are value types, slices and maps are pointers
// named return values
func quoteRem(a, b int) (quote, rem int) {
  quote, rem = a / b, a % b; return
}
// variadic parameters
func sum(args ...int) (sum int) {
  for _, arg := range args { sum += arg }; return
}
fmt.Println(sum([]int{1, 2, 3}...))
// function type, functions are values, block-level scope
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
// defer closures are evaluated after return in the reverse order
defer f.Close()
defer func() {
  if err == nil { err = tx.Commit() } else { tx.Rollback() }
}() // defer must end with ()

/* methods */
// method = a function that operates on a type pointer or a type value
// method modifies a receiver => must use a pointer receiver (p *T)
// method does not modify a receiver => may use a value receiver (v T)
type Bayan struct { model string; year int }
func (b *Bayan) show() { fmt.Printf("Bayan{%s %d}\n", b.model, b.year) }
var b = Bayan{"Nextra", 2022}; var p = &b
b.show(); p.show() // b.show() => (&b).show()
bShow := b.show; bShow() // method value = closure over its instance
BShow := (*Bayan).show; BShow(&b) // method expression => function(receiver)
// struct embedding = composition of concrete types
// fields and methods of an embedded type are promoted to a containing type
type Product struct { Bayan; price float64 }
p := Product{Bayan{"Nextra", 2023}, 1e3}
fmt.Println(p.model); p.show()

/* interfaces */
// interface (the only abstract type) = accept interfaces, return structs
// static type (interface type, abstract type) => interface{} +
// dynamic type (value type, concrete implementation) => nil
// interface = implicit type-safe structural typing when a method set of a
// concrete type including promoted methods from embedded types contains a
// method set of an interface

// empty interface accepts a value of any type
var a interface{}
a = Int(1); fmt.Println(a)
a = Flo(1.2); fmt.Println(a)

// assignment to an interface variable
var i = Int(1)
var v1, v2 View = i, &i // copy, pointer
i = 2; fmt.Println(i); v1.show(); v2.show() // 2 1 2

// process incompatible types through a uniform interface
type View interface { show() }
type Int int
// a type is decoupled from the implicit interface
func (i Int) show() { fmt.Println(i) }
type Flo float64
func (f Flo) show() { fmt.Println(f) }
// only a client specifies the required interface
vs := []View{Int(1), Flo(1.2)}
for _, v := range vs { v.show() } // 1, 1.2

// type assertion applied to an interface at runtime v.(Type)
// vs type conversions applied to concrete types at compile-time Type(v)
// type assertion to access a dynamic type of an interface
var v View = Int(1)
if i, ok := v.(Int); ok { i.show() }

// type switch to access a dynamic type of an interface
vs := []View{Int(1), Flo(1.2)}
for _, v := range vs {
  switch v.(type) {
  case Int: fmt.Print("Int "); v.show()
  case Flo: fmt.Print("Flo "); v.show()
  default: fmt.Println("unknown type")
  }
}

// a function can implement a one-method interface
type Logger interface { log(msg string) } // one-method interface
type LogFunc func(msg string) // function type
// function type implements a Logger interface
func (lf LogFunc) log(msg string) { lf(msg) }
func log(msg string) { fmt.Println(msg) } // log function
// log function => function type => one-method interface
var logger Logger = LogFunc(log)
logger.log("ok") // ok

// interface embedding = composition of abstract types
type Negate interface { View; neg() }
func (i *Int) neg() { *i = -*i }
func (f *Flo) neg() { *f = -*f }
i, f := Int(1), Flo(1.2)
var in, fn Negate = &i, &f
// embedded View.show() is directly accessible through the Negate interface
in.neg(); in.show(); fn.neg(); fn.show()

/* errors */
// sentinel errors = signal that processing cannot continue
// sentinel errors are checked using == equality comparison
type error interface { Error() string } // built-in error interface
// always return error interface => return different error types from a function
func quoteRem(a, b int) (quote, rem int, err error) {
  if b == 0 { err = errors.New("divide by zero"); return }
  quote, rem = a / b, a % b; return
}
if quote, rem, err := quoteRem(5, 3); err != nil { fmt.Println(err)
} else { fmt.Println(quote, rem) }

// custom error
type Status int
const (BadRequest Status = iota + 1; NotFound)
type CustomErr struct { status Status; err error }
func (ce CustomErr) Error() string {
  switch ce.status {
  case BadRequest: return fmt.Sprintf("400 Bad Request: %s", ce.err)
  case NotFound: return fmt.Sprintf("404 Not Found: %s", ce.err)
  default: return fmt.Sprintf("000 Unknown Error: %s", ce.err)
  }
}
func (ce CustomErr) Unwrap() error { return ce.err }

// wrap/unwrap error = build an error chain
func wrapError() error {
  _, err := os.Open("")
  // create new error with a message from a previous error
  return fmt.Errorf("New error: %v", err)
  return fmt.Errorf("Wrap error: %w", err) // wrap a standard error
  return CustomErr{NotFound, err} // wrap a custom error
}
// use defer to wrap errors at multiple returns
func deferWrapError() (err error) {
  defer func() {
    if err != nil { err = fmt.Errorf("Wrap error: %w", err) }
  }()
  _, err = os.Open(""); return
}
err := wrapError(); fmt.Println(err)
if werr := errors.Unwrap(err); werr != nil { fmt.Println(werr) }
// check for a specific error value using == in an error chain
if errors.Is(err, os.ErrNotExist) { fmt.Println("Not Exist")}
// check for a specific error type using reflection in an error chain
var ce CustomErr
if errors.As(err, &ce) { fmt.Println("Custom Error")}

