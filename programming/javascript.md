# JavaScript

## Lexical scope closures and variable lookup

- JS pipeline = tokenization (stateless) | lexing (stateful) => parsing (AST +
  per-scope hoisting of `var`aible and `function` declarations) => optimization
  => code generation (JIT) => execution (variable assignment, function call)
- Compiler (code generator) = variable creation in the appropriate scope
- Engine (orchestrator) = variable lookup for variable / parameter assignment
  (LHS container) and variable / parameter referencing (RHS value)
- **Scope** (variable storage tree) = variable storage and retrieval + shadowing
- **Lexical scope** (closures) is defined statically at write-time (the scope
  chain is based on the source code)
- **Closure** = a returned function can access its lexical scope even when the
  function is executing outside its lexical scope
- **Dynamic scope** (`this`) defined at execution-time and depends on the
  execution path (the scope chain is based on the call stack)
- **Block scope** (`const`, `let`) = declare variables as close as possible to
  where they are used
    - `var` => function scope + hoisting (of variable and function declaration)
    - Function declarations are hoisted before variable declarations
    - `const`, `let` => block scope at any `{ ... }` even explicitly defined
    - `try/catch(e)` => block scope
- **Module pattern** = function creates a new nested function scope not
  accessible from the outside

    ```js
    function module(a) {
      let b = a // private state
      function f() { return ++b } // closure over the private state
      return { f } // public interface
    }
    const m = module(10) // module instance
    console.log(m.f()) // 11
    ```

## `this` dynamic binding rules

- `this` is dynamically defined for every function at runtime (late binding, not
  write-time lexical scope); the value of `this` depends on the location of a
  function call (not the location of function declaration) and how a funciton is
  called; `this` implicitly passes the execution context object (like dynamic
  scope) to the function
- Binding rules for `this` (from the highest to the lowest precedence)
    - `new` **binding** = construction call of a funciton with the `new`
      operator `new f()`. `this` points to a brand new object, which is
      automatically returned from the function (unless the function returns
      another object). The `new` operator ignores `this` hard binding with
      `bind()`
    - **Explicit binding** = function invocation through `f.call(this, args,
      ...)` or `f.apply(this, [args])` including the **hard binding** `const ff
      = f.bind(this, args, ...)` (partial application + currying). `this` points
      to the first argument
    - **Implicit binding** = function invocation `o.f()` through a containing
      context object `const o = { f }`. `this` points to the containing context
      object
    - **Default binding** = standalone function invocation `f()` including
      callback invocation. `this` == `undefined` as the global object is not
      eligible for the default binding in the `strict mode`
- **Lexical `this`** (`bind` alternative) = **arrow function** `(...) => { ... }`
  discards all the traditional rules for `this` binding and instead uses the
  lexical `this` from the immediate lexical enclosing scope. Arrow function is a
  syntactic replacement for `self = this` closures. Lexical `this` binding
  of an arrow function cannot be overrided even with the `new` operator

## `object` property descriptor and accessor descriptor

- **Type** vs `object`
    ```js
    const s = "a" // string type, immutable value, automatic coercion to object
    console.log(typeof s, s instanceof String) // string, false
    const s2 = new String("a") // String object, allows operations
    console.log(typeof s2, s2 instanceof String) // object, true
    ```
- `object` = container for named references to properties (values and functions),
  functions never belong to objects. Syntactic property access `o.p` vs
  programmatic key access `o["p"]`
    ```js
    const o = { a: 1 }
    console.log("a" in o) // true
    for (const p in o) { console.log(p, o[p]) } // a 1
    delete o.a
    console.log(o) // { }
    ```
- **Property descriptor** vs **accessor descriptor**
    ```js
    const o = { }
    Object.defineProperty( // property descriptor
      o, "a", { value: 1, writable: true, enumerable: true, configurable: true }
    )
    o.a = 2
    console.log(o.a) // 2

    const o = { }
    Object.defineProperty(o, "a", { // accessor descriptor
      set: function(v) { this._a = v },
      get: function() { return this._a * 2 }
    })
    o.a = 1
    console.log(o.a) // 2

    const o = { // object literal setter and getter
      set a(v) { this._a = v },
      get a() { return this._a * 2 }
    }
    o.a = 1
    console.log(o.a) // 2
    ```
- `[[Get]]` own property lookup => prototype chain lookup => return `undefined`
- `[[Put]]` accessor descriptor (`set`) => property descriptor (`writable`) =>
  prototype chain lookup => assign value directly to the object
- Object immutability `Object.preventExtensions()`, `Object.seal()`,
  `Object.freeze()`

## Iteration (`[Symbol.iterator]`)

- **Custom iterator** = iterates over arrays (indexing) and objects
  (properties). Ordered, sequential, pull-based consumption of data
  `iterator = { next() => { value, done } }` closure over iterator state +
  `for/of`
    ```js
    const a = [1, 2]
    for (let i = 0; i < a.length; ++i) { console.log(a[i]) } // 1, 2
    for (const e of a) { console.log(e) } // 1, 2
    const o = { a: 1, b: 2 }
    for (const p in o) { console.log(p, o[p]) } // a, 1, b, 2
    Object.defineProperty(o, Symbol.iterator, {
      writable: false, enumerable: false, configurable: true,
      value: function() {
        const o = this
        const keys = Object.keys(o)
        let i = 0 // iterator state
        function next() {
          return { value: o[keys[i++]], done: (i > keys.length) }
        }
        return { next }
      }
    })
    for (const e of o) { console.log(e) } // 1, 2
    ```
- **Iterable interface**
  `iteratble = { [Symbol.iterator]() { return { next() } } }` returns an
  iterator
    ```js
    function iterator(n) { // iterator configuration
      let i = 0; // iterator state
      const next = () => ({ value: i < n ? i : undefined, done: ++i > n })
      // iterable object + iterator function
      return { [Symbol.iterator]() { return this }, next }
    }
    for (const i of iterator(3)) { console.log(i) } // 0, 1, 2
    ```
- Array destructuring `[a, b] = it` and the spread operator `f(...it)` can
  consume an iterator

## Object `prototype` and property lookup

- **Prototype chain** = every object has an `o.[[Prototype]]` link to another
  object ending at `Object.prototype` (kind of global scope for variables)
    ```js
    const o = { a: 1 }
    // new object o2.[[Prototype]] = o (prototype chain)
    const o2 = Object.create(o)
    console.log("a" in o2) // true
    console.log(o2.a) // 1
    for (const p in o2) { console.log(p, o2[p]) } // a, 1
    ```
- **Prototypal inheritance** = all functions get by default a public,
  non-enumerable property `F.prototype` pointing to an object; each object
  created via `new F()` operator is linked to the `F.prototype` effectively
  delegating access to `F.prototype` properties
    ```js
    function F() { this.a = 1 } // constructor
    F.prototype.b = function() { return 2 } // method
    const o = new F()
    console.log(o.a, o.b()) // 1, 2
    function G() { F.call(this); this.c = 3 } // call parent constructor
    // Prototypal inheritance Option 1. Overwrite G.prototype
    G.prototype = Object.create(F.prototype)
    // Prototypal inheritance Option 2. Update G.prototype
    Object.setPrototypeOf(G.prototype, F.prototype)
    G.prototype.d = function() {
      return F.prototype.b.call(this) + 2 // call parent method
    }
    const o2 = new G()
    console.log(o2.a, o2.b(), o2.c, o2.d()) // 1, 2, 3, 4
    ```
- Purely flat data storage without `prototype` delegation
  `o = Object.create(null)`
- **Prototypal behavior delegation** = objects are linked to other objects
  forming a network of peers, not a vertical hierarchy as with classes
- Mutual delegation of two objects to each other forming a cycle is disallowed
- **ES6 class** = syntactic sugar on top of prototypal inheritance and
  prototypal behavior delegation
    ```js
    class F {
      constructor() { this.a = 1 } // constructor + property
      b() { return 2 } // method
    }
    const o = new F()
    console.log(o.a, o.b()) // 1, 2
    class G extends F { // prototypal inheritance
      constructor() { super(); this.c = 3 } // call parent constructor
      d() { return super.b() + 2 } // call parent method
    }
    const o2 = new G()
    console.log(o2.a, o2.b(), o2.c, o2.d()) // 1, 2, 3, 4
    ```
- Function chaining via `return this`
    ```js
    function N(x) { this.a = x }
    N.prototype.add = function add(x) { this.a += x; return this }
    console.log(new N(1).add(2).add(3).a) // 6
    ```

## Types

- Types are related to values, not to variables, which may store values of
  different types over time
- `function` is a `[[Call]]`able `object`
- The type of value determines whether the value will be **assigned by copy**
  (primitives `boolean`, `number`, `string`, `symbol`) or **assigned by
  reference** (`object`, `array`, `function`, automatically boxed values)
- `symbol` special unique primitive type used for **collision-free internal
  properties** on objects
    ```js
    const sym = Symbol("a")
    const o = { [sym]: 1 }
    console.log(o[sym]) // 1, collision-free property
    ```

## Coercion

- Coercion always results in one of the scalar primitive types
- Both `==` (implicit coercion) and `===` (no coercion) compare two `object`s by
  reference (not by value) `{ a: 1 } ==(=) { a: 1 } // false`
- Use `===` (no coercion) instead of `==` with `true`, `false`, `0`, `""`, `[]`

## JS grammar

- **Assignment expression** returns the assigned value `a = 1 // 1`
- `continue <label>` continues a labeled outer loop `label: while/for(...)`
- `break <label>` breaks out of an inner loop or a labeled block `label: { ... }`
- `let a = b || <default value>` vs `a && a.b()` guarded operation + short
  circuiting
- Right-associative: `=`, `?:`

## Callbacks

- **Single-threaded event loop** (sequential execution on every tick)
    ```js
    let events = [] // queue, FIFO
    while(true) {
      if (events.length) { // tick
        let event = events.shift()
        try { event() } // atomic unit of work run to completion
        catch (e) { console.log(e) }
      }
    }
    ```
- **Concurrency** = split 2 or more tasks into atomic steps, schedule steps from
  all tasks to the event loop (interleave steps from different tasks), execute
  steps in the event loop in order to progress simultaneously on all tasks
- **Callbacks** = strict separation between now (current code) and later
  (callback). Non-linear definition of a sequential control flow and error
  handling, trust issues due to control delegation (inversion of control,
  continuation)
    ```js
    function timeoutify(fun, timeout) {
      let id = setTimeout(() => {
        id = null
        fun(new Error("timeout"))
      }, timeout)
      return (...a) => {
        if (id) {
          clearTimeout(id)
          fun(null, ...a)
        }
      }
    }
    function f(e, d) {
      if (e) { console.error(e) } else { console.log(d) }
    }
    const tf = timeoutify(f, 500)
    setTimeout(() => tf(1), 400) // 1
    ```

## Promises

- **Promises** = placeholder/proxy for a future eventual value (trustable,
  composable, time consistent) that is guaranteed to be async (both now and
  later always are async). `resolve` and `reject` callabcks are guaranteed to
  be invoked async at most once and exclusively (even if a Promise is resolved
  sync with a value, even if `then()` is called on already settled Promise)
    ```js
    Promise.resolve(1).then(console.log) // next tick
    console.log(2) // 2, 1
    ```
- **Promises** = async flow control (multiple consumers subscribe to a
  completion event of a producer) that separates consumers from a producer
- Thrown error in either `resolve` or `reject` callbacks is automatically
  propagated through the chain of Promises as a rejection (`throw` is usable
  with Promises)
- Once a pending promise is settled, the `resolve()`d value or `reject()`ed
  reason becomes immutable. Repeated calls to `resolve()` and `reject()` are
  ignored. Promise must be `return`ed to form a valid promise chain
    ```js
    function timeoutPromise(timeout) {
      return new Promise((_, reject) =>
        setTimeout(() => reject("timeout"), timeout)
      )
    }
    function f(x, timeout) {
      return new Promise((resolve) =>
        setTimeout(() => resolve(x), timeout)
      )
    }
    Promise.race([f(1, 400), timeoutPromise(500)])
      .then(console.log).catch(console.error) // 1
    ```
- Promises solve the trust issues of callbacks by inverting the callback control
  delegation. Promises don't get rid of callbacks, they just let the caller
  control callbacks locally via `p.then(cb)` instead of passing callabcks to a
  third party code as in case of callbacks only approach
- `Promise.resolve(x)` normilizes values and misbehaving thenables to trustable
  and compliant Promises
- `p.then()` automatically and synchronously creates a new Promise in a chain
  either resolved with a value or rejected with an error/reason of the
  unwrapped Promise
    ```js
    Promise.resolve(1)
      .then(x => x + 1)
      .then(x => new Promise(resolve => setTimeout(_ => resolve(x * 2), 100)))
      .then(console.log) // 4
    ```
- `p.catch()`ed rejection restores the Promise chain back to normal
    ```js
    Promise.resolve(1)
      // default rejection handler: e => { throw e } for incoming errors
      .then(_ => { throw new Error("oh") })
      // default resolution handler: x => { return x } for incoming values
      .catch(e => { console.error(e.message); return 2 }) // for outgoing errors
      .then(console.log) // oh, 2 (back to normal)
    ```
- `Promise.all([])` a gate that resolves with the array of results of all
  concurrent, unordered resolved Promises or rejects with the first rejected
  Promise
- `Promise.race([])` a latch that either resolves or rejects with the first
  settled Promise (the other Promises cannot be canceled due to immutability,
  hense are settled and just ignored)
- Promises API
    - `new Promise((resolve, reject) => {...})`
    - `Promise.resolve(x)`, `Promise.reject(x)`
    - `p.then(success, [failure])`, `p.catch(failure)`, `p.finally(always)`
    - `Promise.all([]) => [all success] | first failure`
    - `Promise.allSettled([]) => [all success | failure]`
    - `Promise.race([]) => first success | first failure`
    - `Promise.any([]) => first success | all failures`
- Callback => Promise
    ```js
    function f(x, cb) {
      setTimeout(_ => { if (x >= 0) { cb(null, "ok") } else { cb("oh") } }, 100)
    }
    f(1, console.log) // null, ok
    f(-1, console.error) // oh
    function promisify(f) {
      return function(...args) {
        return new Promise((resolve, reject) => {
          f.apply(null, args.concat(
            function(e, x) { if (e) { reject(e) } else { resolve(x) } }))
        })
      }
    }
    const ff = promisify(f)
    ff(1).then(console.log) // ok
    ff(-1).catch(console.error) // oh
    ```
- Sequential composition of promises
    ```js
    const f = x => new Promise((resolve) => setTimeout(_ => resolve(x + 1), 100))
    const p = [f, f].reduce((p, f) => p.then(f), Promise.resolve(1))
    p.then(console.log) // 3
    let r = 1
    for (const ff of [f, f]) { r = await ff(r) }
    console.log(r) // 3
    ```

## Generators

- **Generators** = a new type of function that does not run to completion (as
  regular functions do), but creates an iterator that controls execution of the
  generator, suspends maintaining the iternal state at every `yield` and resumes
  on each iteration call `it.next()`. `yield` is right-associative like `=`
  assignment
- Generator use cases
    - On demand production of series of vaules through iteration
    - Async flow control through two-way message passing
- **Two-way message passing**
    - Generator `const y = yield x` yields `x` to the caller before suspending
      and receives `y` from the caller after resuming
    - Caller `const { value: x } = it.next(y)` receives `x` from a suspended
      generator, resumes the generator and passes `y` into the generator
- Generators implement **cooperative multitasking** by `yield`ing control,
  (not preemptive multitasking by extractal context switch). Gnerator suspends
  itself via `yield`, iterator call `it.next()` resumes the generator
    ```js
    function* g(x) {
      console.log(x++)
      // yield waits for a value passed by it.next(v)
      const y = yield "a" // yield requires 2 iterations: start + resume
      console.log(y, x) // implicitly returns undefined
    }
    const it = g(1) // creates a generator + receives an interator
    // starts the generator (must always be empty)
    const { value } = it.next() // 1
    console.log(value) // a
    // resumes the generator + passes message to the generator
    const r = it.next("b") // b, 2
    console.log(r) // { value: undefined, done: true }
    ```
- Initial `it.next()` + `yield` + message `it.next(v)`
    ```js
    function* g() {
      const a = yield "a"
      const b = yield "b"
      console.log(a, b)
    }
    const it = g() // creates the controlling iterator
    const { value: a } = it.next() // starts the generator
    const { value: b } = it.next(1)
    it.next(2) // 1, 2 (finishes the generator)
    console.log(a, b) // a, b
    ```
- **Early termination** via `break`, `return`, `throw` from the `for/of` loop
  automatically terminates generator's iterator (or manually via `it.return()`)
    ```js
    const iterator = function(n) {
      let v = 0 // state
      return {
        [Symbol.iterator]() { return this },
        next() {
          return v < n ? { value: ++v, done: false } :
          { value: undefined, done: true } }
      }
    }
    for (const i of iterator(3)) { console.log(i) } // 1, 2, 3

    const generator = function*(n) {
      let v = 0
      while (v < n) { yield ++v }
    }
    for (const i of generator(3)) { console.log(i) } // 1, 2, 3

    const generator = function*() {
      let v = 0
      try {
        while (true) { yield ++v }
      } finally { console.log("finally") }
    }
    const gen = generator()
    for (const i of gen) {
      if (i > 2) {
        const { value } = gen.return("return")
        console.log(value)
      }
      console.log(i) // 1, 2, finally, return, 3
    }
    ```
- Generators express **async flow control** in sequential, sync-like form
  through async iteration (`it.next()`) of a generator
    ```js
    function f(x, cb) {
      setTimeout(_ => x === "oh" ? cb(new Error(x)) : cb(null, x), 100)
    }
    function cb(err, data) { if (err) { it.throw(err) } else { it.next(data) } }
    function* g() {
      try {
        const a = yield f(1, cb)
        console.log(a)
        const b = yield f("oh", cb)
      } catch (e) { console.error(e.message) }
    }
    const it = g()
    it.next() // 1, oh
    ```
- **Promise-yielding generators** are basis for `async/await`
    ```js
    function f(x) {
      return new Promise((resolve, reject) =>
        setTimeout(_ => x === "oh" ? reject(new Error(x)) : resolve(x), 100)
      )
    }
    function* g() {
      try {
        const a = yield f(1)
        console.log(a)
        const b = yield f("oh")
        // const b = yield f(2)
        // console.log(b)
      } catch(e) { console.error(e.message) }
    }
    const it = g()
    it.next().value
      .then(a => it.next(a).value.then(b => it.next(b)))
      .catch(e => console.log(e.message)) // 1, oh
    ```
- `yield *` delegation for **composition of generators**. `yield *` requires an
  iterable, it then invokes that iterable's iterator and delegates generator's
  control to that iterator until it is exhausted
    ```js
    function* a() { yield 1; yield* b(); yield 4 }
    function* b() { yield 2; yield 3 }
    for (const i of a()) { console.log(i) } // 1, 2, 3, 4
    ```
- Error handling `try/catch` inside and outside of generators
    ```js
    function* g() {
      try {
        yield 1
      } catch (e) { console.error(e.message) } // uh
      throw new Error("oh")
    }
    const it = g()
    try {
      const { value: a } = it.next()
      console.log(a) // 1
      it.throw(new Error("uh"))
    } catch (e) { console.error(e.message) } // oh
    ```

# ES6+

- `let` **block scoped** variable (vs `var` function scoped + hoisting)
- `const` **block scoped** variable that must be initialized and cannot be
  reassigned (constant reference), while the content of reference types can
  still be modified
- **Spread arguments** `f(...[1, 2, 3])` => `f.apply(null, [1, 2, 3])`
- **Gather parameters** `function f(...args) {...}` => `[args]`
- **Object/array destructuring/transformation**
    ```js
    const o = { a: 1, b: 2, c: 3 }
    const a = [10, 20, 30]
    let o2 = { }
    let a2 = [];
    ({ a: o2.A, b: o2.B, c: o2.C } = o)  // object => object
    console.log(o2); // { A: 1, B: 2, C: 3 }
    [a2[2], a2[1], a2[0]] = a  // array => array
    console.log(a2); // [ 30, 20, 10 ]
    ({ a: a2[0], b: a2[1], c: a2[2] } = o) // object => array
    console.log(a2); // [ 1, 2, 3 ]
    [o2.A, o2.B, o2.C] = a // array => object
    console.log(o2) // { A: 10, B: 20, C: 30 }
    ```
- **Spread/gather destructuring**
    ```js
    const [x, ...y] = a
    console.log(x, y, [x, ...y]) // 10, [ 20, 30 ], [ 10, 20, 30 ]
    const { a, ...x } = o
    console.log(a, x, { a, ...x }) // 1 { b: 2, c: 3 } { a: 1, b: 2, c: 3 }
    ```
- **Default values destructuring** vs default parameters
    ```js
    const [p, q, r, s = 0] = a
    console.log(p, q, r, s) // 10, 20, 30, 0
    const { a: p, d: s = 0 } = o
    console.log(p, s) // 1, 0
    f({ x = 10 } = { }, { y } = { y: 10 }) { ... }
    ```
- **Concise methods** `{ f() { ... } }` imply anonymous function expression
  `{ f: function() { ... } }`
- **Getter/setter**
    ```js
    const o = {
      _a: 1,
      get a() { return this._a },
      set a(v) { this._a = v }
    }
    o.a++
    console.log(o.a) // 2
    ```
- **Computed property name**
    ```js
    const p = "a"
    const o = { [p]: 1 }
    console.log(o.a, o[p]) // 1, 1
    ```
- **Tagged template literal**
    ```js
    function tag(strings, ...values) {
      return `${strings[1].trim()} ${values[0] + 1} ${strings[0]}`
    }
    const a = 1
    console.log(tag`A ${a + 1} B`) // B 3 A
    ```
- **Arrow functions** are always anonymous (no named reference for recursion or
  event binding/unbinding) function expressions (there is no arrow function
  declaration) + parameters destructuring, default values, and spread/gather
  operator
- Inside an arrow function `this` is lexical (not dynamic). Arrow function is a
  nicer alternative to `const self = this` or `f.bind(this)`
- Array indexing `for (index; condition; increment)`
- Oject properties `for property in object`
- Iterator `for element of iterator`
- `RegExp` sticky `y` flag restricts the pattern to match just at the position
  of the `lastIndex` which is set to the next character beyond the end of the
  previous match (`y` flag implies a virtual anchor at the beginning of the
  pattern) vs non-sticky patterns are free to move ahead in their matching
- `Symbol("desc")` = primitive type with immutable, unique, hidden value used
  for collision-free object properties (e. g. `Symbol.iterator` => `for/of`)
    ```js
    function Singleton() {
      // const instance = Symbol("instance")
      const instance = Symbol.for("singleton.instance") // Symbol registry
      if (!Singleton[instance]) {
        this.a = 1
        Singleton[instance] = this
      }
      return Singleton[instance]
    }
    const s1 = new Singleton()
    const s2 = new Singleton()
    console.log(s1, s2, s1 === s2, new Number(1) === new Number(1))
    // Singleton { a: 1 }, Singleton { a: 1 }, true, false
    ```
- **Promise-yielding generator** `*function() { yield Promise } => iterator` is
  the basis for an **async function** `async function() { await Promise } =>
  Promise`

# Modules

- **Modules** = static, resolved at compile time with read-only, one-way live
  bindings (not copies) to exported values. One module per file, module is a
  cached singleton, there is no global scope inside a module (`this` is
  `undefined`), circular imports are correctly handled regardless of import
  order
- Module identifier (constant string) = relative path `../module.js`, absolute
  path `file:///module.js`, core modules or `node_modules` `module` or
  `core/module`, module URL `https://module.js`
- Export (not `export`ed object are private to the module)
    - **Named exports** `export var | const | let | function | class | { a, b as
      B }` of named object defined in the module
    - **Default export** `export default { a, b }` or `export { a as default }`
      not mutually exclusive with named exports `import def, { named } from
      "module"` that rewards with a simpler `import def` syntax. Default export
      is considered unnamed (internally `default` name is used) and can be
      imported under any name
    - **Re-export** from another module `export * | { a, b as B } from "module"`
- Import (all imported bindings are immutable and hoisted)
    - **Named import** `import { a, b as B } from "module"` binds to top-level
      identifiers in the current scope
    - **Default import** `import m | { default as m } from "module"`
    - **Wildcard import** to a single namespace `import * as ns from "module"`
- **Dynamic async import** = `import(module) => Promise` at runtime

# Classes

- **Class** = a macro that populates a `constructor` function, a `prototype`
  with methods and defines a `prototype` chain through `extends`
    ```js
    class A {
      constructor(a) { this._a = a }
      // property getter and setter
      get a() { return this._a }
      set a(v) { this._a = v }
    }
    class B extends A { // prototype delegation
      constructor(a, b) {
        super(a) // parent constructor
        this.b = b
      }
      static c = 10 // statics are on the constructor function, not the prototype
      sum() { return super.a + this.b } // parent object
    }
    const b = new B(1, 2)
    b.a += 3
    console.log(b.a, b.sum(), B.c) // 4, 6, 10
    ```

# Metaprogramming

- `Proxy` + `Reflect` intercepts at the proxy, extends in the proxy and forwards
  to the target object `get`, `set`, `delete`, `apply`, `construct` operations
  among others
- **Proxy first** design pattern
    ```js
     const o = { a: 1 }
     const handlers = {
       get(target, key, context) {
         if (Reflect.has(target, key)) {
           console.log("get key", key)
           // forward operation from context (proxy) to target (object)
           return Reflect.get(target, key, context)
         } else {
           throw new Error(`${key} does not exist`)
         }
       }
     }
     const p = new Proxy(o, handlers)
     console.log(p.a) // get key a, 1
    ```
- **Proxy last** design pattern
    ```js
    const o = { a: 1 }
    const handlers = {
      get(target, key, context) { throw new Error(`${key} does not exits`) }
    }
    const p = new Proxy(o, handlers)
    Object.setPrototypeOf(o, p)
    console.log(o.a, o.b) // 1, Error
    ```
- **Tail-call optimization** (TCO)
    ```js
    function rmap(a, f = e => e, r = []) {
      if (a.length > 1) {
        const [h, ...t] = a
        return rmap(t, f, r.concat(f(h)))
      } else {
        return r.concat(f(a[0]))
      }
    }
    const a = new Array(9999)
    console.log(rmap(a.fill(0), e => e + 1)) // Maximum call stack size exceeded
    ```
- **Trampoline** converts recursion => loop
    ```js
    function trampoline(f) { // factors out recursion into loop
      // stack depth remains constant (stack frames are reused)
      while (typeof f === "function") { f = f() }
      return f
    }
    function tmap(a, f = e => e, r = []) {
      if (a.length > 1) {
        // no recursive call to tmap(), just return the partial() function
        return function partial() { // executed by trampoline
          const [h, ...t] = a
          return tmap(t, f, r.concat(f(h)))
        }
      } else {
        return r.concat(f(a[0]))
      }
    }
    const a = new Array(9999)
    console.log(trampoline(tmap(a.fill(0), e => e + 1))) // no RangeError
    ```

# Node.js

## Reactor pattern

- Modern OSes provide async, non-blocking IO syscalls through the **Sync Event
  Demultiplexer** (SED) e. g. `epoll` on Linux. SED watches (sync) for a set of
  async operations to complete and allows multiple async operations to be
  processed in a single thread (events demultiplexing)
- **Reactor pattern** = executes (async) a handler (app callback) for each async
  operation (non-blocking IO)
    - App requests async operations [resource (file), opeartion (read), handler
      (callbback)] at SED (non-blocking)
    - SED watches requested async operations for completion (blocking)
    - SED enqueues [event (operation), handler (callback)] to the event queue
      (EQ)
    - Single-threaded event loop (EL) reads the EQ and executes handlers (app
      callbacks) to completion (no race conditions). App callbacks request more
      async operations at SED (when a task requests a new async operations it
      gives control back to the event loop; invoking an async operation always
      unwinds the stack back to the event loop, leaving it free to handle other
      requests, IO-bound)
    - EL blocks again (next EL cycle) at SED for new async operations to
      complete
- **libuv** = SED + reactor (event loop + event queue) = cross-platform
  low-level IO engine of Node.js
    - High resolution clock and timers
    - Async FS, TCP/UDP, DNS
    - Async thread pool and synchronization
    - Async child processes and signals handling
    - Async IPC via shared UNIX domain sockets
- **Node.js** = libuv (SED + reactor) + V8 (JavaScript runtime) + modules
- In Node.js only async, non-blocking IO operations are **executed in parallel**
  by the libuv internal IO threads. Sync code inside callbacks (until next async
  operation) is **executed concurrently** to completion without race conditions
  by the single-threaded event loop

## Callback pattern

- **Callback** (Continuation-Passing Stype, CPS = control is passed explicitly
  in the form foa continuation/callback, `return`/`throw` => `callback(error,
  resutl)`) = first-class function (to pass a callback) + closure (to retain the
  caller contenxt) that is passed to an async function (which returns
  immediately) and is executed on completion of the requested async operation
- Callback communicates **async result once** and has trust issues (tight
  coupling = callback is passed to the async function for execution). Callback
  is expected to be invoked exactly once either with result or error
- Async result/error is propagated through a chain of nested `callbacks(error,
  data | cb)` that accept `error | null` first, then the `result` or the next
  `callback` that comes always last. Never `return` a result or `throw` an error
  from a callback
- Sync operation (CPU-bound, discouraged) blocks the event loop and puts on hold
  all concurrent processing blocking the whole application
    - Interleave each step of a CPU-bound sync operation with `setImmediate()`
      to let pending IO tasks to be processed by the event loop (not efficient
      due to context switches, interleaved algorithm)
    - Pool of reusable external processes `child_process.fork()` with a
      communication channel `process/worker.send().on()` leaves the event loop
      unblocked (efficient: reusable processes run to completion, unmodified
      algorithm)
    - `new Worker()` threads with communication chanels
      `parentPort/worker.postMessage().on()`, per-thread own event loop and V8
      instance (small memory footprint, fast startup time, safe: no
      syncronization, no resource sharing)
- Avoid mixing sync/async calback behavior under the same inferface. To convert
  sync call to async use
    - Mictrotask `process.nextTick()` is executed just after the current
      operaiton, before other pending IO tasks
    - Immediate task `setImmediate()` is executed after all other pending IO
      tasks in the same cycle of the event loop
    - Regular task `setTimeout()` is executed on the next cycle of the event
      loop
- Uncaught error thrown from an async callback propagates to the stack of the
  event loop (not to the next callback, not to the stack of the caller that
  triggered the async operation), `process.on(unchaughtException, err)` is
  emitted and process exits with a non-zero exit code
- **Callback hell** = deeply nested code as a result of in-place nested
  anonymous callbacks that unnecessary consume memory because of closures
    - Do not abuse in-place nested anonymous callbacks
    - Early return principle = favor `return/continue/break` over nested
      `if/else`
    - Create named callbacks with clearly defined interface (no unnecessary
      closures)
- **Sequential iteration** (recursion) = applies an async operation to each
  element of an array one element at a time
    ```js
    function cbTask(x, cb) {
      setTimeout(() => {
        if (x === -1) { return cb("oh") }
        console.log(x); cb(null)
      }, 500)
    }
    function cbIterate(task, arr, cb) {
      let index = 0
      function iterate() {
        if (index === arr.length) {
          return process.nextTick(() => cb(null)) // always async
        }
        // Recurse for the next interation after completion of the previous one
        task(arr[index++], iterate)
      }
      iterate()
    }
    asyncIterate(cbTask, [], console.log) // null
    asyncIterate(cbTask, [1, 2, 3], console.log) // 1, 2, 3, null
    ```
- **Parallel execution** (loop for all tasks) = executes tasks in parallel
  (unlimited) until all complete or first error
    ```js
    function cbParallel(tasks, cb) {
      let completed = 0
      let failed = false
      function done(error) {
        if (error) { failed = true; return cb(error) }
        if (++completed === tasks.length && !failed) { return cb(null) }
      }
      for (const task of tasks) { task(done) }
    }
    // 1, 3, 2, null
    cbParallel([1, 2, 3].map(i => (done) => cbTask(i, done)), console.log)
    // 1, oh, 3
    cbParallel([1, -1, 3].map(i => (done) => cbTask(i, done)), console.log)
    ```
- **Limited parallel execution** (loop until limit) = executes at most N tasks
  in parallel until all complete or first error
    ```js
    function cbParallelLimit(tasks, limit, cb) {
      let completed = 0
      let failed = false
      let index = 0
      let running = 0 // Queue can be used instead
      function done (error) {
        --running
        if (error) { failed = true; return cb(error) }
        if (++completed === tasks.length && !failed) { return cb(null) }
        if (running < limit && !failed) { parallel() }
      }
      function parallel() {
        while (index < tasks.length && running < limit) {
          const task = tasks[index++]
          task(done)
          ++running
        }
      }
      parallel()
    }
    cbParallelLimit(
      [1, 2, 3, 4, 5].map(i => (done) => cbTask(i, done)), 2, console.log
    ) // 1, 2 | 3, 4 | 5, null
    cbParallelLimit(
      [1, 2, 3, -1, 5].map(i => (done) => cbTask(i, done)), 2, console.log
    ) // 1, 2 | 3, oh | 5, null
    ```

## Observer pattern

- **Observer** = object/subject notifies its observers on its state changes
- `EventEmitter` = registers multiple observing `listeners(args, ...)` for
  specific event types `ee.on|once(event, listener)`, `ee.emit(event, args,
  ...)`, `ee.removeListener(event, listener)` unsubscribe listeners when they
  are no longer needed to avoid memory leaks due to captured context in listener
  closures
- `EventEmitter` continuously notifies **multiple observers on different types
  of recurrent events** and does not have trust issues (loose coupling =
  callback is controlled by the caller). En event can be fired multiple times or
  not fired at all. Combining `EventEmitter` with a callback interface is an
  elegant and flexible solution
- Async result/error is proparaged through `emit` events. Never `return` a
  result or `throw` an error from the `EventEmitter`. Always register a listener
  for the `error` event
- Never mix sync and async events in the same `EventEmitter`. Sync event
  listeners must be registered before the status change. Async events allow
  registering listeners even after the status change but within the current
  cycle of the event loop (because events are guaranteed to fire in the next
  cycle of the event loop)

## Promise pattern

- `promisify` = converts a callback-based function into a Promise-returning
  function
    ```js
    function promisify(f) {
      return (...args) => {
        return new Promise((resolve, reject) => {
          const argsCb = [...args, (error, result) => {
            if (error) { return reject(error) }
            resolve(result)
          }]
          f(...argsCb)
        })
      }
    }
    const taskP = promisify(ctTask)
    taskP(1).then(console.log) // 1, undefined
    taskP(-1).catch(console.error) // oh
    ```
- **Sequential iteration** (dynamic promise chaining in a loop)
    ```js
    function promiseIterate(task, arr) {
      let p = Promise.resolve()
      for (const e of arr) { p = p.then(() => task(e)) }
      return p
    }
    promiseIterate(taskP, []).then(console.log) // undefined
    promiseIterate(taskP, [1, 2, 3]).then(console.log) // 1, 2, 3, undefined
    ```
- **Parallel execution** (loop for all tasks)
    ```js
    function promiseParallel(tasks) {
      let completed = 0
      return new Promise((resolve, reject) => {
        function done() {
          if (++completed === tasks.length) { resolve() }
        }
        for (const task of tasks) { task().then(done, reject) }
      })
    }
    promiseParallel([1, 2, 3].map(i => () => taskP(i)))
      .then(console.log) // 1, 2, 3, undefined
    promiseParallel([1, -1, 3].map(i => () => taskP(i)))
      .then(console.log, console.error) // 1, oh, 3
    ```
- **Limited parallel execution** (loop until limit)
    ```js
    function promiseParallelLimit(tasks, limit) {
      let completed = 0
      let index = 0
      let running = 0
      return new Promise((resolve, reject) => {
        function done() {
          --running
          if (++completed === tasks.length) { resolve() }
          if (running < limit) { parallel() }
        }
        function parallel() {
          while (index < tasks.length && running < limit) {
            const task = tasks[index++]
            task().then(done, reject)
            ++running
          }
        }
        parallel()
      })
    }
    promiseParallelLimit(
      [1, 2, 3, 4, 5].map(i => () => taskP(i)), 2
    ).then(console.log) // 1, 2 | 3, 4 | undefined
    promiseParallelLimit(
      [1, 2, 3, -1, 5].map(i => () => taskP(i)), 2
    ).then(console.log, console.error) // 1, 2 | 3, oh | 5
    ```

## Async/await pattern

- `async` function always returns a Promise immediately and synchronously.
  `try/catch/throw` inside an `async` function works for both sync and async
  code. Use `return await` to prevent error on the caller side and `catch`
  errors locally
    ```js
    async function localError() {
      try { return await taskP(-1) }
      catch (error) { console.error(`Local: ${error}`) }
    }
    localError().catch(error => console.error(`Caller: ${error}`)) // Local: oh
    ```
- On each `await` expression that always returns a Promise, an `async` function
  is put on hold, its state is saved and the control is retuned to the event
  loop (generator). When the Promise that has been `await`ed settles, the
  control is given back to the `async` function
- **Sequential iteration**
    ```js
    async function asyncIterate(task, arr) {
      for (const e of arr) { await task(e) }
    }
    await asyncIterate(taskP, [1, 2, 3]) // 1, 2, 3
    ```
- **Parallel execution** (await loop)
    ```js
    async function asyncParallel(tasks) {
      const promises = tasks.map(task => task())
      // Problem: unnecesary wait for all promosises in the array
      // preceding the rejected promise. Solution: use Promise.all()
      for (const promise of promises) { await promise }
    }
    await asyncParallel([1, 2, 3].map(i => () => taskP(i))) // 1, 2, 3
    try {
      await asyncParallel([1, -1, 3].map(i => () => taskP(i))) // 1, oh, 3
    } catch (error) { console.error(error) }
    ```
- **Limited parallel execution**
    ```js
    async function asyncParallelLimit(tasks, limit) {
      let completed = 0
      let index = 0
      let running = 0
      let promises = []
      function parallel() {
        while (index < tasks.length && running < limit) {
          const task = tasks[index++]
          promises.push(task())
          ++running
        }
      }
      parallel()
      while (completed !== tasks.length) {
        if (promises.length !== 0) {
          await promises.shift()
          ++completed, --running
          parallel()
        }
      }
    }
    asyncParallelLimit(
      [1, 2, 3, 4, 5].map(i => () => taskP(i)), 2
    ) // 1, 2 | 3, 4 | 5
    try {
      await asyncParallelLimit(
        [1, 2, 3, -1, 5].map(i => () => taskP(i)), 2
      ) // 1, 2 | 3, oh | 5
    } catch (error) { console.error(error) }
    ```
- **Infinite recursive promise chain** = creates memory leaks
    ```js
    // Problem: recursive, memory leak vs lost rejection
    async function leakingRecursion(i = 0) {
      await taskP(i)
      // Memory leak: hain of dependent Promises that never resolve
      return leakingRecursion(i + 1)
      // No memory leak (GC collected Promises), but lost rejection
      leakingRecursion(i + 1)
    }
    leakingRecursion() // 0, 1, 2, ...
    // Solution: loop with await
    async function nonLeakingLoop() {
      let i = 0
      while (true) {
        await taskP(i++) // No memory leak + correct error handling
      }
    }
    nonLeakingLoop() // 0, 1, 2, ...
    ```

## Stream pattern

- **Streaming** (parallel pipeline vs sequential buffering) = staged parallel
  processing of data in chunks as soon as it arrives with no delays due to
  buffering (reactive, modular, composable, constant memory)
- `Stream` extends `EventEmitter`
    - Binary mode (IO processing)
    - Object mode (function composition)
- `Readable` (source of data) = async iterator (`for await`)
    - Non-flowing mode (default) = `on(readable) + .read()`/`.pause()` pulls
      data in a controlled way
        ```js
        function nonFlowingReadable() {
          process.stdin
            .on("readable", () => { // new data is available
              let chunk // read() pulls data in a loop from an internal buffer
              // Flexible control over data consumption
              while ((chunk = process.stdin.read()) !== null) { // read() is sync
                  console.log(chunk.toString())
              }
            })
            .on("end", () => { console.log("done") }) // end of stream
        }
        ```
    - Flowing mode = `on(data)`/`.resume()` pushes data as soon as it arrives
        ```js
        function flowingReadable() {
          process.stdin
            // on(data) or resume() switches to the flowing mode that pushes data
            // .pause() switches back to the non-flowing mode (default)
            .on("data", chunk => console.log(chunk.toString()))
            .on("end", () => { console.log("done") })
        }
        ```
    - Implementing Readable
        ```js
        class RandomReadable extends Readable {
          #reads = 0
          async _read(size) {
            const chunk = await randomBytes(10)
            // push() => false for backpressure when the internal buffer is full
            // on(drain) resume pushing
            this.push(chunk) // push data to the internal buffer
            if (++this.#reads === 3) { this.push(null) } // end of stream
          }
        }
        async function randomReadable() {
            const rr = new RandomReadable()
            for await (const chunk of rr) {
              console.log(chunk.toString("hex"))
            }
        }
        ```
- `Writable` (sink of data)
    ```js
    function serverWritable() {
      const server = createServer(async (req, res) => {
        res.writeHead(200, { "Content-Type": "text/plain" })
        const body = await randomBytes(10)
        // write() => false for backpressure when the internal buffer is full
        // on(drain) resume writing
        res.write(body)
        res.end("\n\n")
        res.on("finish", () => console.log("done"))
      })
      const port = 9876
      server.listen(port, () => console.log("Listening on port ${port}"))
    }
    ```
- `Duplex` = `Readable` (source of data) + `Writable` (sink of data) where
  written and read data is not related
- `Transform` (data transformation) = `Duplex` where written and read data is
  related through a transformation (`map`, `filter`, `reduce`)
- `PassThrough` = `Transform` without transformation for observability,
  instrumentation, late piping, `Readable` to `Writable` change of interface,
  and lazy streams to delay through a proxy expensive resource initialization
  until actual stream consumption
- `readable.pipe(writableDest) => writableDest` switches `Readable` into the
  flowing mode, controls backpressure automatically, returns `Writable`
  destination for chaining (it must be `Duplex`, `Transform`, or `PassThrough`).
  Errors are not propagated automatically through a `pipe()`. `on(error)`
  handlers must be attached for every step
- `pipeline(straeam, ..., cb)` automatically attaches `on(error)` and
  `on(close)` handlares for each stream to correctly destroys streams on
  pipeline success or failure
- Both `pipe()` and `pipeline()` return only the last stream (not the combined
  stream)
- **Sequential iteration** = `Stream` always processes async operations in
  sequence one at a time e. g. `on(data)`
    ```js
    // Sequential iteration
    function streamIterate(task, arr) {
      const taskTr = new Transform({
        objectMode: true,
        async transform(chunk, encoding, cb) {
          try { await task(chunk); cb() }
          catch (error) { cb(error) }
        }
      })
      return new Promise((resolve, reject) => {
        // Sequential interation
        Readable.from(arr).pipe(taskTr)
          .on("error", reject)
          .on("finish", resolve)
      })
    }
    try {
      await streamIterate(taskP, [1, 2, 3]) // 1, 2, 3
      await streamIterate(taskP, [1, -1, 3]) // 1, oh
    } catch (error) { console.error(error) }
    ```
- **Parallel execution**
    ```js
    function streamParallel(tasks) {
      let completed = 0
      let done = null
      let fail = null
      const taskTr = new Transform({
        objectMode: true,
        transform(task, encoding, cb) {
          // Start all tasks in parallel
          task().catch(fail) // Global reject on first failure
            .finally(() => {
              // flush() stram only when all tasks are actually done
              if (++completed === tasks.length) { done() }
            })
          cb() // all tasks are done immediately
        },
        flush(cb) { done = cb }
      })
      return new Promise((resolve, reject) => {
        fail = reject
        Readable.from(tasks).pipe(taskTr)
          .on("error", reject)
          .on("finish", resolve)
      })
    }
    try {
      await streamParallel([1, 2, 3].map(i => () => taskP(i))) // 1, 2, 3
      await streamParallel([1, -1, 3].map(i => () => taskP(i))) // 1, oh, 3
    } catch(error) { console.error(error) }
    ```
- **Limited parallel execution**
    ```js
    function streamParallelLimit(tasks, limit) {
      let completed = 0
      let running = 0
      let done = null
      let fail = null
      let resume = null
      const taskTr = new Transform({
        objectMode: true,
        transform(task, encoding, cb) {
          task().catch(fail)
            .finally(() => {
              if (++completed === tasks.length) { return done() }
              --running
              // Resume the stream when a task is completed
              if (resume) { const r = resume; resume = null; r() }
            })
          if (++running < limit) { cb() }
          else { resume = cb } // Suspend the stream until a task is completed
        },
        flush(cb) { done = cb }
      })
      return new Promise((resolve, reject) => {
        fail = reject
        Readable.from(tasks).pipe(taskTr)
          .on("error", reject)
          .on("finish", resolve)
      })
    }
    try {
      // 1, 2 | 3, 4 | 5
      await streamParallelLimit([1, 2, 3, 4, 5].map(i => () => taskP(i)), 2)
      // 1, 2 | 3, oh | 5
      await streamParallelLimit([1, 2, 3, -1, 5].map(i => () => taskP(i)), 2)
    } catch(error) { console.error(error) }
    ```
- **Composition of streams** = creates a new combined stream (first `Writable`
  and last `Readable`)
- **Fork of a stream** = pipes a single `Readable` stream into multiple
  `Writable` streams with automatic backpressure of the slowest branch of the
  fork
- **Merge of streams** = pipes multiple `Readable` into a single `Writable`
