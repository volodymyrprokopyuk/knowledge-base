/* Primitives */
// bool, int = int64, uint = uint64, float64
// byte = uint8, rune = int32, string
// type alias = attach different methods to the same underlying type
type Year int // nominal typing, block-level scope
type Year2 = Year // type alias for backward compatibility
var a Year = Year2(2023)
// comparable types (==) bool, int/float64, string, pointer, array,
// struct of comparable types, chan, interface (identical dynamic values)
// slices and maps are not comparable => use reflect.DeepEqual or a custom equal
// integer overflows and underflows are silent

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
// default initialization to zero for primitives
var (b bool; i int; f float64; r rune; s string)
// default initialization to nil for slice/map/chan, pointer/interface/func
// short variable declaration = in-function inferred initialization/redefinition
b, i, f, r, s := true, 1, 1.2, 'a', `ok`

/* Pointers */
// call-by-value preserves immutability of original data
// use of pointers implies mutable data
// pointer initialization & address and * dereferencing
var i int = 1; var p *int = &i; i++; *p++ // i == 3; *p == 3
var p *int = new(int) // p* == 0

/* Strings, runes */
// string = immutable sequence of arbitrary bytes; pointer to a byte array + len
// charset Unicode code points => variable length encoding UTF-8 (up to 4 bytes)
s := "Добро"
// byte index 0 2 4 6 8
for i, r := range s { fmt.Printf("%d: %c, ", i, r) }
s[:2] // first two bytes
// rune index 0 1 2 3 4
for i, r := range []rune(s) { fmt.Printf("%d: %c, ", i, r) }
string([]rune(s)[:2]) // first two runes

/* Arrays */
// allocated on stack as array size is known at compile-time
var a [3]int // [0 0 0] automatic zero initialization
// explicit initialization
b, c := [3]int{}, [...]int{1, 2, 3} // [0 0 0], [1 2 3]
d := [...]int{1, 2: 2, 4: 3} // [1 0 2 0 3] sparse array

/* Slices */
// slice = pointer to an array + len + cap; dynamic array
// slice = multiple views onto an array; update, but not append
// slice = single view per array; update and append
// slices are not comparable
var (a []int, b []int(nil)) // nil slice without allocation => nil
a = append(a, 1, 2) // [1 2] increments length, dynamic reallocation
b := []int{1, 2, 3} // [1 2 3] slice initialization
c := []int{1, 2: 2, 4: 3} // [1 0 2 0 3] sparse slice
d := make([]int, 3) // [0 0 0] zero initialized slice with capacity 3
e, f := []int{}, make([]int, 0, 3) // allocated empty slice len = 0, cap 3
a := [...]int{1, 2, 3}
var b []int = a[:] // array => slice, shared memory
var c []int = make([]int, len(a))
copy(c, a[:]) // array => slice, separate copies; destination must be init
// [i:j:cap] = [i, j), cap limits side effects by reallocating when len == cap
// make a slice copy instead of slicing big arrays to dispose unused memory
for i, v := range a { fmt.Println(i, v) }

/* Maps */
// maps are not comparable
var a = map[string]int{} // empty map
var b = map[string]int{"a": 1, "b": 2} // initialized map
var c = make(map[string]int, 3) // map with capacity
c["a"] = 1
if v, ok := c["a"]; ok { fmt.Println(v) }
for k, v := range b { fmt.Println(k, v) }
delete(c, "a") // maps never shrink, even after deletion buckets are not removed

/* Structs, types */
// nominal typing, block-level scope
  var a, b Bayan = Bayan{model: "Nextra", year: 2022}, Bayan{"Omnia", 2023}
  var c, d *Bayan = new(Bayan), &Bayan{"Selecta", 2024}
  fmt.Println(d.model) // (*d).model => d.model
  e := []Bayan{{model: "Nextra", year: 2022}, {"Omnia", 2023}}
// constructor function = correct and uniform initialization of structs
// constructor function is preferred over a struct literal
// local value is allocated on heap when its pointer is returned from a function
// constructor function
func NewBayan(model string, year int) *Bayan { return &Bayan{model, year} }
// accessor functions = exported methods for unexported struct fields for
// data validation, computed values, access serialization
func (b *Bayan) Model() { return p.model }
func (b *Bayan) SetModel(model string) { p.model = model }
// structural typing, anonymous struct
var a Bayan = struct { model string; year int }{"Nextra", 2022}
// struct embedding = composition of concrete types
// embedding makes exported identifiers of an embedded type public
// multiple types can be embedded into a new type
// composition = an embedded type is a receiver, a new type is not an embedded type
// inheritance = a subclass is a receiver and can subsitute a superclass
type Product struct { Bayan; price float64 } // embedded field without a name
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
// shadowing = re-definition of an identifier in an inner block

/* if/else */
// allows for arbitrary conditions in each clause
// each clause has its own scope
if i := rand.Intn(10); i < 3 { fmt.Println(i, "low")
} else if i < 8 { fmt.Println(i, "mid")
} else { fmt.Println(i, "high") }
// align the happy path to the left; favor early return
// the happy path goes down, while edge cases are handled on the right
// prefer vertical if error { return } over nested if/else

/* for + break/continue [label] */
// best for controlled iteration defining start, end, and step
for i := 0; i < 3; i++ { fmt.Println(i) } // 0 1 2
for i, l := 0, len(a); i < l; i++ { fmt.Println(a[i]) }
i := 0 // best for dynamic exit condition
for i < 3 { fmt.Println(i); i++ } // 0 1 2
i = 0 // unconditional first iteration
for { fmt.Println(i); i++; if i > 2 { break } } // 0 1 2
// best for arrays, slices, strings, maps (single container only)
// for/range returns a copy of values
for _, v := range []int{1, 2, 3} { fmt.Println(v) } // 1 2 3
// range expression is a copy evaluated once before a loop
a := []int{1, 2, 3}
for range a { a = append(a, 10) } // [1 2 3 10 10 10]
// all values are assigned to a single variable
type Acc struct { bal float64 }
accs := []Acc{{1}, {2}, {3}}
m := make(map[int]*Acc, len(accs))
for i, acc := range accs { m[i] = &acc } // 3 3 3
for i, acc := range accs { a := acc; m[i] = &a } // 1 2 3
for i := range accs { m[i] = &accs[i] } // 1 2 3

/* switch + break [label] */
a := []string{"one", "eleven", "thousand"}
outer: for _, v := range a {
  switch l := len(v); l { // equality check == in each clause
  case 1, 2, 3: fmt.Println(v, "small")
  // use break label in a switch within a for loop
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
// named return values = for documentation (interface), initialization, defer
// avoid naked returns for clarity
func quoteRem(a, b int) (quote, rem int) {
  quote, rem = a / b, a % b; return quote, rem
}
// variadic parameters
func sum(args ...int) (sum int) {
  for _, arg := range args { sum += arg }; return sum
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

/* defer */
// defer closures are evaluated after function return in the reverse order
// receivers are evaluated when declared, not when executed
defer f.Close()
defer func() {
  if err == nil { err = tx.Commit() } else { tx.Rollback() }
}() // defer must end with ()
// arguments to defer closures are evaluated when declared, not when executed
var i, j int
defer func(i int) { fmt.Println(i, j) }(i) // 0 1
i++; j++
// capture and return an error from defer; must use a named return + closure
func f() (err error) {
  err = fmt.Errorf("function oh")
  defer func() {
    // wrap a function error in a defer error
    if err != nil { err = fmt.Errorf("defer oh: %w", err)
    } else { err = fmt.Errorf("defer oh") }
  }()
  return err
}

/* methods */
// method = a function that operates on a type pointer or a type value
// function(receiver, args...) => receiver.method(args...)
// method modifies a receiver => must use a pointer receiver (p *T), large objects
// method does not modify a receiver => may use a value receiver (v T), primitives
// value receiver = a copy of a receiver is passed to a method
// method can be invoked through a nil pointer/receiver = valid receiver
// mixing receiver types should be avoided
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
// interface (the only abstract type)
// static type (interface type, abstract type) => interface{} +
// dynamic type (value type, concrete implementation) => nil
// nil interface == nil; nil receiver converted to interface != nil
// when returning an interface, return nil directly, not a nil pointer
type Err struct { }
func (e Err) Error() string { return "oh" }
func f(a bool) error {
  var err *Err
  if a { err = &Err{} }
  return err // nil pointer converted to the error interface != nil
}
func g(a bool) error {
  if a { return Err{} }
  return nil // nil converted to the error interface == nil
}
if err := f(false); err != nil { fmt.Println(err) } // <nil>
if err := f(true); err != nil { fmt.Println(err) } // oh
if err := g(false); err != nil { fmt.Println(err) } // no error
if err := g(true); err != nil { fmt.Println(err) } // oh
// interface = implicit type-safe structural typing when a method set of a
// concrete type including promoted methods from embedded types contains a
// method set of an interface
// common behavior = across distinct types e. g. sort.Interface
type Interface interface { len() int; less(i, j int) bool; swap(i, j int) }
// decoupling = rely on an abstraction, not an implementation (Dependency
// Inversion Principle)
// restriction = an interface restricts available operations e. g. read-only
// the bigger the interface, the weaker the abstraction
// abstractions should be discovered, not created: struct => interface
// do not force an interface on a producer side, let a consumer discover the
// right abstraction (Interface Segregation Principle)
// accept interfaces (flexible input), return structs (compliant output)
// (Robustness Principle)
// do not return an interface defined on a consumer side (unnecessary circular
// dependency)

// empty interface and any type accepts a value of any type
type any = interface{}
var a any // should be minimized
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

// type switch to access a dynamic type of an interface at runtime
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

/* generics */
// type parameeters cannot be used with method arguments, only with
// function arguments or method receivers
type Node[T any] struct { Val T; next *Node[T] } // container type
func (n *Node[T]) Add(node *Node[T]) { n.next = node }
// type constraint is an interface = set of methods or concrete types
type intStr interface { ~int | ~string } // ~int types derived from int
func keys[K intStr, V any](m map[K]V) (ks []K) { // common algorithm
  for k := range m { ks = append(ks, k) }; return
}
m := map[string]int{"a": 1, "b": 2}
ks := keys[string, int](m)
ks := keys(m) // inferred type parameters

/* errors */
// error = error type that signals an unexpected yet recoverable situation
// sentinel error = signals an expected, recoverable error (empty dataset, EOF)
// sentinel error = a value assigned to a global variable and compared with ==
// error type = unexpected error; switch err.(type); errors.As(err, &AnError{})
// sintinel error value = expected error; err == ErrA; errors.Is(err, ErrA)
// handle an error only once, provide additional context using wrapping, produce
// single error log entry (either log or return an error, but not both)
var ErrSentinel = errors.New("Sentinel Error")
type error interface { Error() string } // built-in error interface
// always return error interface => return different error types from a function
func quoteRem(a, b int) (quote, rem int, err error) {
  if b == 0 { err = errors.New("divide by zero"); return quote, rem }
  quote, rem = a / b, a % b; return quote, rem
}
if quote, rem, err := quoteRem(5, 3); err != nil { fmt.Println(err)
} else { fmt.Println(quote, rem) }

// custom error
type Status int
const (BadRequest Status = iota + 1; NotFound)
// specific error type
type CustomError struct { status Status; err error }
func (ce CustomError) Error() string {
  switch ce.status {
  case BadRequest: return fmt.Sprintf("400 Bad Request: %s", ce.err)
  case NotFound: return fmt.Sprintf("404 Not Found: %s", ce.err)
  default: return fmt.Sprintf("000 Unknown Error: %s", ce.err)
  }
}
func (ce CustomError) Unwrap() error { return ce.err }

// wrap/unwrap errors = build an error chain with additional context wrapped in
// a specific error type
func wrapError() error {
  _, err := os.Open("")
  // include an error message into a new error; a source error is  not available
  return fmt.Errorf("New error context: %v", err)
  // wrap a standard error; a source error is available (coupling of a client)
  return fmt.Errorf("Wrap error context: %w", err)
  return CustomError{NotFound, err} // wrap a custom error
}
// use defer to wrap errors at multiple returns
func deferWrapError() (err error) {
  defer func() {
    if err != nil { err = fmt.Errorf("Wrap error: %w", err) }
  }()
  _, err = os.Open(""); return
}
// explicit check for errors from a limited scope
err := wrapError(); fmt.Println(err)
if werr := errors.Unwrap(err); werr != nil { fmt.Println(werr) }
// check for a specific sentinel error value using == in an error chain
if errors.Is(err, os.ErrNotExist) { fmt.Println("Not Exist") }
// check for a specific error type using reflection in an error chain
if errors.As(err, &CustomErr{}) { fmt.Println("Custom Error") }

// panic = signals a termination of a program due to an unrecoverable situation
// panic = programming error, unavailable dependency
// panic = unwinds a stack only to the top of a current goroutine
// recover must be within the code of a goroutine
// recover must be called from defer as only defer functions are executed on panic
// app: use recover to gracefully handle shutdown (log panic message)
// lib: use recover to convert a panic to an error at a public API boundary
func panicRecover() {
  defer func() {
    if msg := recover(); msg != nil { fmt.Println(msg) }
  }()
  panic("oh")
}
panicRecover(); fmt.Println("continue") // oh continue

/* modules, packages */
// Go programs build from source code into a self-contained executable
// module = commands/packages root that consists of packages in a repository
// commands => go install ..., packages => go get -u ..., import "..."
// module = unit of versioning identified by a repository path (module ID)
// go mod init <scm/user/mod> => go.mod
// package name should match package directory
// name a package after what it provides, not what it contains
// package = noun, exports = nouns/verbs
// every source file in a directory must have the same package name
// package level Capitalized identifiers are exported (public API)
// unexported identifiers are accessible from different files of a package
// top-level identifiers in all package files must be unique
// merge packages or create a new package to resolve circular dependencies
package pkgname
import "scm/user/mod/pkgdir" // pkgname ~= pkgdir, absolute import always
pkgname.Identifier
package main; func main() { }
// import alias
import (crand "crypto/rand"; "math/rand")
// shares identifiers between parent and sibling packages without exporting them
package internal // special package recognized by go tooling
// automatic singleton initialization of package state through global variables
// prefer encapsulated variables + errors over global variables + panic
// init function runs after initialization of package variables
// inside an init function only panic is available to signal an error
func init() { }
// go get scm/user/mod@version # upgrades a module to a specific version
// go get -u scm/user/mod # upgrades a module to the most recent version
// go mod tidy # removes unused versions from go.mod
// git tag v1.2.3 # for backward compatible versions
// mkdir v2; git branch v2 # for backward incompatible versions
import "scm/user/mod/v2/pkgdir" // new module import path
