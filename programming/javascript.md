# JavaScript language

## Types, coercion, and operators

- Types are related to values, not to variables, which may store values of
  different types over time
- The type of value determines whether the value will be **assigned by copy**
  (primitives `boolean`, `number`, `string`, `symbol`) or **assigned by
  reference** (`object`, `array`, `function`, automatically boxed values)
- Coercion always results in one of the scalar primitive types
- Both `==` (**implicit coercion**) and `===` (no coercion) compare two
  `object`s **by reference** (not by value) `{ a: 1 } ==(=) { a: 1 } // false`
- Use `===` (no coercion) instead of `==` with `true`, `false`, `0`, `""`, `[]`
- **Assignment expression** returns the assigned value `a = 1 // 1`
- `continue <label>` continues a labeled outer loop `label: while/for(...)`
- `break <label>` breaks out of an inner loop or a labeled block `label: { ... }`
- Right-associative: `=`, `?:`
- Array indexing `for (let index; condition; increment)`
- Oject properties `for (const property in object)`
- Iterator `for (const element of iterator)`
- `function` is a `[[Call]]`able `object`
- `let a = b || <default>` vs `a && a.b()` **guarded operation** + short
  circuiting
- `symbol` special unique primitive type used for **collision-free properties**
  on objects
    ```js
    const sym = Symbol("a")
    const o = { [sym]: 1 }
    console.log(o[sym]) // 1, collision-free property
    ```
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
- **Tagged template literal**
    ```js
    function tag(strings, ...values) {
      return `${strings[1].trim()} ${values[0] + 1} ${strings[0]}`
    }
    const a = 1
    console.log(tag`A ${a + 1} B`) // B 3 A
    ```
- `RegExp` sticky `y` flag restricts the pattern to match just at the position
  of the `lastIndex` which is set to the next character beyond the end of the
  previous match (`y` flag implies a virtual anchor at the beginning of the
  pattern) vs non-sticky patterns are free to move ahead in their matching

## `object` property descriptor and accessor descriptor

- **Type** vs `object`
    ```js
    const s = "a" // string type, immutable value, automatic coercion to object
    console.log(typeof s, s instanceof String) // string, false
    const s2 = new String("a") // String object, allows operations
    console.log(typeof s2, s2 instanceof String) // object, true
    ```
- `object` = container for named references to properties (values, objects, and
  functions). Functions never belong to objects. Syntactic property access `o.p`
  vs programmatic property access `o["p"]`
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

    class C { // class setter and getter
      set a(v) { this._a = v }
      get a() { return this._a * 2 }
    }
    const c = new C()
    c.a = 1
    console.log(c.a) // 2
    ```
- `[[Get]]` read lookup: own property lookup => prototype chain lookup => return
  `undefined`
- `[[Put]]` write lookup: accessor descriptor (`set`) => property descriptor
  (`writable`) => prototype chain lookup => assign value directly to the object
- Object immutability `Object.preventExtensions()`, `Object.seal()`,
  `Object.freeze()`
- **Object concise method** `{ f() { ... } }` => `{ f: function() { ... } }`
  implies anonymous function expression

## Lexical scope, closures, and variable lookup

- JS pipeline = tokenization (stateless) | lexing (stateful) => parsing (AST +
  per-scope hoisting of `var`aible and `function` declarations) => optimization
  => code generation (JIT) => execution (variable assignment, function call)
- Compiler (code generator) = variable creation in an appropriate scope
- Engine (orchestrator) = variable lookup for variable/parameter assignment
  (LHS container) and variable/parameter referencing (RHS value)
- **Scope** (storage tree for variables) = storage and referencing of variable,
  shadowing
- **Lexical scope** (closures) is defined statically at write-time (a scope
  chain is based on a source code)
- **Closure** = a returned function can access its lexical scope even when the
  function is executing outside its lexical scope
- **Dynamic scope** (`this`) defined at execution-time and depends on the
  execution path (a scope chain is based on a call stack)
- **Block scope** (`const`, `let`) = declare variables as close as possible to
  where they are used
    - `var` => function scope + hoisting (of variable and function declarations)
    - Function declarations are hoisted before variable declarations
    - `const/let` => block scope at any `{ ... }` even explicitly defined
    - `let` block scoped variable (vs `var` function scoped + hoisting)
    - `const` block scoped variable that must be initialized and cannot be
      reassigned (constant reference), while the content of reference types can
      still be modified
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

## `this` late binding, lexical `this`, and arrow functions

- `this` is dynamically defined for every function call at runtime (late
  binding, not write-time lexical scope). The value of `this` depends on the
  location of a function call (not the location of a function declaration) and
  how a funciton is called. `this` implicitly passes the execution context
  object (like dynamic scope) to a function
- Late binding rules for `this` (from the highest to the lowest precedence)
    - `new` **binding** = construction call of a funciton with the `new`
      operator `new f()`. `this` points to a brand new object, which is
      automatically returned from the function (unless the function returns
      another object). The `new` operator ignores `this` hard binding with
      `bind()`
    - **Explicit binding** = function invocation through `f.call(this, args,
      ...)` or `f.apply(this, [args])` including the **hard binding** `const ff
      = f.bind(this, args, ...)` for partial application, currying. `this`
      points to the first argument
    - **Implicit binding** = function invocation `o.f()` through a containing
      context object `const o = { f }`. `this` points to the containing context
      object
    - **Default binding** = standalone function invocation `f()` including
      callback invocation. `this` == `undefined` as the global object is not
      eligible for the default binding in the `strict mode`
- **Lexical `this`** (`bind` alternative) = **arrow function** `(...) => { ...
  }` discards all the traditional rules for `this` binding and instead uses the
  lexical `this` from the **immediate lexical enclosing scope**. Arrow function
  is a syntactic replacement for `self = this` closures. Lexical `this` binding
  of an arrow function cannot be overrided even with the `new` operator
- **Arrow function** = anonymous (no named reference for recursion or event
  bind/unbind) **function expressions** (there is no arrow function declaration)
  that support parameters destructuring, default values, and spread/gather
  operators. Inside an arrow function `this` is lexical (not dynamic). An arrow
  function is a nicer alternative to `const self = this` or `f.bind(this)`

## Prototype chain, prototypal ihneritance, and `class`

- **Prototype chain** = every object has an `o.[[Prototype]]` link to another
  object ending at `Object.prototype` (kind of global scope for variables)
    ```js
    const o = { a: 1 }
    // new object o2.[[Prototype]] = o (prototype chain)
    const o2 = Object.create(o)
    console.log("a" in o2, o2.a) // true, 1
    for (const p in o2) { console.log(p, o2[p]) } // a, 1
    ```
- **Prototypal inheritance** = all functions get by default a public,
  non-enumerable property `F.prototype` pointing to an object. Each object
  created via `new F()` operator is linked to the `F.prototype` effectively
  delegating access to `F.prototype` properties
    ```js
    function F() { this.a = 1 } // constructor, property
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
- `class` = syntactic sugar on top of prototypal inheritance and prototypal
  behavior delegation
    ```js
    class F {
      constructor() { this.a = 1 } // constructor, property
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
    ```js
    class A {
      constructor(a) { this._a = a } // property
      // property setter and getter
      set a(v) { this._a = v }
      get a() { return this._a }
    }
    class B extends A { // prototypal inheritance
      constructor(a, b) {
        super(a) // parent constructor
        this.b = b
      }
      // statics are on the constructor function, not the prototype
      static c = 10
      sum() { return super.a + this.b } // parent object
    }
    const b = new B(1, 2)
    b.a += 3
    console.log(b.a, b.sum(), B.c) // 4, 6, 10
    ```
- **Method chaining** via `return this`
    ```js
    function N(x) { this.a = x }
    N.prototype.add = function add(x) { this.a += x; return this }
    console.log(new N(1).add(2).add(3).a) // 6
    ```

## Spread/gather, object/array destructuring/transformation

- **Spread arguments** `f(...[1, 2, 3])` => `f.apply(null, [1, 2, 3])`
- **Gather parameters** `function f(...args) { ... }` => `[args]`
- **Object/array destructuring/transformation**
    ```js
    const o = { a: 1, b: 2, c: 3 }, a = [10, 20, 30], o2 = { }, a2 = [];
    ({ a: o2.A, b: o2.B, c: o2.C } = o)  // object => object
    console.log(o2); // { A: 1, B: 2, C: 3 }
    [a2[2], a2[1], a2[0]] = a  // array => array
    console.log(a2); // [30, 20, 10]
    ({ a: a2[0], b: a2[1], c: a2[2] } = o) // object => array
    console.log(a2); // [1, 2, 3]
    [o2.A, o2.B, o2.C] = a // array => object
    console.log(o2) // { A: 10, B: 20, C: 30 }
    ```
- **Spread/gather object/array destructuring**
    ```js
    const { a, ...x } = o
    console.log(a, x, { a, ...x }) // 1 { b: 2, c: 3 } { a: 1, b: 2, c: 3 }
    const [x, ...y] = a
    console.log(x, y, [x, ...y]) // 10, [ 20, 30 ], [ 10, 20, 30 ]
    ```
- **Default values destructuring** vs default parameters
    ```js
    const { a: p, d: s = 0 } = o
    console.log(p, s) // 1, 0
    const [p, q, r, s = 0] = a
    console.log(p, q, r, s) // 10, 20, 30, 0
    f({ x = 10 } = { }, { y } = { y: 10 }) { ... }
    ```

## Modules, `export`, `import`

- **Modules** = static, resolved at compile=time, read-only, **one-way live
  bindings to exported values** (not copies). One module per file, module is a
  **cached singleton**, there is no global scope inside a module (`this` is
  `undefined`), circular imports are correctly handled regardless of the
  `import` order
- **Module identifier** (a constant string) = relative path `"../module.js"`,
  absolute path `"/module.js"`, core modules (`node_modules`) `"core/module"`,
  module URL `"https://module.js"`
- Export (not `export`ed objects are private to a module)
    - **Named exports**
      `export var | const | let | function | class | { a, b as B }` of (re)named
      objects defined in a module
    - **Default export** `export default { a, b }` or `export { a as default }`
      not mutually exclusive with the named exports `import dflt, { named } from
      "module"` that rewards with a simpler `import dflt` syntax. Default export
      is unnamed and can be `import`ed under any name
    - **Re-export** from another module `export * | { a, b as B } from "module"`
- Import (all imported bindings are immutable and hoisted)
    - **Named import** `import { a, b as B } from "module"` binds to top-level
      identifiers in the current scope
    - **Default import** `import dflt | { default as dflt } from "module"`
    - **Wildcard import** to a single namespace `import * as ns from "module"`
- **Dynamic async import** = `import("module") => Promise` at runtime

## Iterator closure and iterable interface `[Symbol.iterator]`

- **Iterator closure** = iterates over arrays (indexing) and objects
  (properties). Ordered, sequential, pull-based consumption of data
  `iterator = { next() => { value, done } }` closure over iterator state through
  interface of `for/of`
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
- **Iterable interface** =
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

# Async JavaScript

## Callback

- **Single-threaded event loop** = sequential execution **run to completion**
  on every tick
    ```js
    const events = [] // queue (FIFO)
    while(true) {
      if (events.length > 0) { // tick
        const event = events.shift()
        try { event() } // atomic unit of work run to completion
        catch (e) { console.error(e) }
      }
    }
    ```
- **Concurrency** = split two or more compound tasks into atomic steps, schedule
  steps from all tasks to the event loop (interleave steps from different
  tasks), execute steps in the event loop in order to progress simultaneously on
  all tasks
- **Callback** = strict separation between now (the current code) and later
  (callback, control delegation). Non-linear definition of a sequential control
  flow and error handling, **trust issues due to control delegation** and to
  **inversion of control** (continuations). A thrown error is not automatically
  propagated through a chain of callbacks (`throw` is not usable with callbacks)
    ```js
    function timeoutify(f, timeout) {
      let id = setTimeout(() => {
        id = null; f(new Error("timeout"))
      }, timeout)
      return (...args) => {
        if (id) { clearTimeout(id); f(null, ...args) }
      }
    }
    function f(e, v) { e ? console.error(e) : console.log(v) }
    const tf = timeoutify(f, 200)
    setTimeout(() => tf(1), 100) // 1
    ```
- Callback testing
    ```js
    export function f(v, done) {
      setTimeout(() => v ? done(null, v) : done(new Error("oh")), 200)
    }
    f(true, console.log) // null, true, control delegation = trust issue
    f(false, console.error) // oh
    describe("f", () => {
      test("success", () => new Promise(done =>
        f(true, (e, v) => { expect(v).toBe(true); done() })
      ))
      test("failure", () => new Promise(done =>
        f(false, (e, v) => { expect(e.message).toBe("oh"); done() })
      ))
    })
    ```

## Promise

- **Promise** = a placeholder/proxy for a **future eventual value** (trustable,
  composable, time consistent, always async) that is **guaranteed to be async**
  (both now and later are always async). `resolve` and `reject` callabcks are
  guaranteed to be invoked **async at most once and exclusively** (even if a
  Promise is resolved sync with a value, even if `then()` is called on an
  already settled Promise)
    ```js
    Promise.resolve(1).then(console.log) // next tick
    console.log(2) // 2, 1
    ```
- **Promise** = **async composable flow control** (multiple consumers subscribe
  to a completion event of a producer) that separates consumers from a producer
- **Thrown error** in either `resolve` or `reject` callback is **automatically
  propagated** through a chain of Promises as a rejection (`throw` is usable
  with Promises)
- Once a pending promise is settled, a `resolve()`d value or a `reject()`ed
  error becomes immutable. Repeated calls to `resolve()` and `reject()` are
  ignored. A promise must be `return`ed to form a valid promise chain
    ```js
    function timeout(timeout) {
      return new Promise((_, reject) =>
        setTimeout(() => reject("timeout"), timeout)
      )
    }
    function f(x, timeout) {
      return new Promise(resolve =>
        setTimeout(() => resolve(x), timeout)
      )
    }
    Promise.race([f(1, 400), timeout(500)])
      .then(console.log).catch(console.error) // 1
    ```
- **Promises solve the trust issues** of callbacks by **inverting the callback
  control delegation**. Promises don't get rid of callbacks, but they **let the
  caller to control callbacks locally** via `p.then(cb)` instead of passing
  callabcks to a third party code as in case of callbacks only approach
- `Promise.resolve(x)` normilizes values and misbehaving thenables to trustable
  and compliant Promises
- `p.then()` automatically and synchronously **creates a new Promise in a
  chain** either resolved with a value or rejected with an error
    ```js
    Promise.resolve(1)
      .then(x => x + 1)
      .then(x => new Promise(resolve => setTimeout(() => resolve(x * 2), 100)))
      .then(console.log) // 4
    ```
- `p.catch()`ed rejection restores the Promise chain back to normal
    ```js
    Promise.resolve(1)
      // default rejection handler: e => { throw e } for incoming errors
      .then(() => { throw new Error("oh") })
      // default resolution handler: x => { return x } for incoming values
      .catch(e => { console.error(e.message); return 2 }) // for outgoing errors
      .then(console.log) // oh, 2 (back to normal)
    ```
- Promise testing
    ```js
    export function f(v) {
      return new Promise((resolve, reject) =>
        setTimeout(() => v ? resolve(v) : reject(new Error("oh")), 200)
      )
    }
    f(true).then(console.log) // true, caller controls, no trust issues
    f(false).catch(console.error) // oh
    describe("f", () => {
      test("success", () =>
        f(true).then(v => expect(v).toBe(true))
      )
      test("failure", () =>
        f(false).catch(e => expect(e.message).toBe("oh"))
      )
    })
    ```
- `Promise.all([Promise])` a gate that resolves with an array of all
  concurrently resolved Promises or rejects with the first rejected Promise
- `Promise.race([Promise])` a latch that either resolves or rejects with the
  first settled Promise (the other Promises cannot be canceled due to
  immutability, hense are settled and just ignored)
- Promises API
    - `new Promise((resolve, reject) => {...})`
    - `Promise.resolve(x)`, `Promise.reject(x)`
    - `p.then(success, [failure])`, `p.catch(failure)`, `p.finally(always)`
    - `Promise.all([]) => [all success] | first failure`
    - `Promise.allSettled([]) => [all success | failure]`
    - `Promise.race([]) => first success | first failure`
    - `Promise.any([]) => first success | all failures`
- Callback => Promise = converts a callback-based function into a
  `Promise`-returning function
    ```js
    function f(v, done) {
      setTimeout(() => v ? done(null, v) : done(new Error("oh")), 200)
    }
    function promisify(f) {
      return function(...args) {
        return new Promise((resolve, reject) => {
          args = [...args, (e, v) => e ? reject(e) : resolve(v)]
          f(...args)
        })
      }
    }
    const ff = promisify(f)
    ff(true).then(console.log) // true
    ff(false).catch(console.error) // oh
    ```
- Sequential composition of promises
    ```js
    function f(v) { return new Promise(
      resolve => setTimeout(() => resolve(v + 1), 100)
    )}
    const p = [f, f].reduce((p, f) => p.then(f), Promise.resolve(1))
    p.then(console.log) // 3
    let r = 1
    for (const ff of [f, f]) { r = await ff(r) }
    console.log(r) // 3
    ```

## Generators

- **Generators** = a new type of function that **does not run to completion**
  (as a regular function does), but **creates an iterator that controls
  execution of a generator**, **suspends maintaining an iternal state** at
  every `yield` and resumes on each iteration call to `it.next()`. `yield` is
  right-associative like `=` and `?:`
- Generator use cases
    - On demand production of series of vaules through iteration maintaining an
      internal state
    - Async flow control through a two-way message passing
- **Two-way message passing**
    - Generator `const y = yield x` yields `x` to the caller before suspending
      and receives `y` from the caller after resuming
    - Caller `const { value: x } = it.next(y)` receives `x` from a suspended
      generator, resumes the generator and passes `y` into the generator
- Generator implements the **cooperative multitasking** by `yield`ing control,
  (not preemptive multitasking by external context switch). Gnerator suspends
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
    const it = g() // creates a controlling iterator
    const { value: a } = it.next() // starts the generator, must always be empty
    const { value: b } = it.next(1)
    it.next(2) // 1, 2 (finishes the generator)
    console.log(a, b) // a, b
    ```
- **Early termination** via `break`, `return`, `throw` from the `for/of` loop
  automatically terminates generator's iterator (or manually via `it.return()`
  or `it.throw()`)
    ```js
    function* g(n) {
      let i = 0
      while (i < n) { yield i++ }
    }
    for (const el of g(3)) { console.log(el) } // 0, 1, 2

    const infinite = function*() {
      let v = 0
      try { while (true) { yield v++ } }
      finally { console.log("finally") }
    }
    const inf = infinite()
    for (const i of inf) {
      if (i > 1) {
        const { value } = inf.return("return")
        console.log(value)
      }
      console.log(i) // 0, 1, finally, return, 2
    }
    ```
- Generator expresses **async flow control** in sequential, sync-like form
  through async iteration (`it.next()`) of a generator
    ```js
    function f(v, done) {
      setTimeout(() => v ? done(null, v) : done(new Error("oh")), 200)
    }
    function done(e, v) { if (e) { gen.throw(e) } else { gen.next(v) } }
    function* g() {
      try {
        const a = yield f(true, done)
        console.log(a)
        const _ = yield f(false, done)
      } catch (e) { console.error(e.message) }
    }
    const gen = g()
    gen.next() // true, oh
    ```
- **Promise-yielding generator** is a basis for `async/await`
    ```js
    function f(v) {
      return new Promise((resolve, reject) =>
        setTimeout(() => v ? resolve(v) : reject(new Error("oh")), 200)
      )
    }
    function* g() {
      try {
        const a = yield f(true)
        console.log(a)
        const _ = yield f(false)
      } catch (e) { console.error(e.message) }
    }
    const gen = g()
    gen.next().value.then(x => gen.next(x).value.then(y => gen.next(y)))
      .catch(e => console.log(e.message)) // true, oh
    ```
- `yield*` delegation for **composition of generators**. `yield*` requires an
  iterable `[Symbol.iterator]`, it then invokes that iterable's iterator
  `it.next()` and delegates generator's control to that iterator until it is
  exhausted
    ```js
    function* inner() { yield 2; yield 3 }
    function* outer() { yield 1; yield* inner(); yield 4 }
    for (const el of outer()) { console.log(el) } // 1, 2, 3, 4
    ```
- **Error handling** `try/catch` inside and outside of generators
    ```js
    function* g() {
      try { yield 1 }
      catch (e) { console.error("inside", e.message) } // inside uh
      throw new Error("oh")
    }
    const gen = g()
    try {
      const { value } = gen.next()
      console.log(value) // 1, inside uh, outside oh
      gen.throw(new Error("uh"))
    } catch (e) { console.error("outside", e.message) } // outside oh
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
    - Async thread pool and synchronization
    - Async child processes and signal handling
    - Async FS, TCP/UDP, DNS
    - Async IPC via shared UNIX domain sockets
- **Node.js** (async IO operations in parallel by libuv threads, sync callback
  code to completion in the event loop) = libuv (SED + reactor) + V8 (JavaScript
  runtime) + modules
- In Node.js only async, non-blocking IO operations are **executed in parallel**
  by the **libuv internal IO threads**. Sync code inside callbacks (until next
  async operation) is **executed concurrently** to completion without race
  conditions by the single-threaded event loop

## Callback pattern

- **Callback** (Continuation-Passing Stype, CPS = control is passed explicitly
  in the form fo a continuation/callback, `return`/`throw` => `callback(error,
  result)`) = **first-class function** (to pass a callback) + **closure** (to
  retain the caller contenxt) that is passed to an async function (which
  **returns immediately**) and **is executed on completion** of the requested
  async operation
- Callback communicates an async result once and has **trust issues** (tight
  coupling = callback is passed to an async function for execution). Callback
  **is expected to be invoked exactly** once either with result or error
- Async result/error is propagated through a chain of nested `callbacks(error |
  null, data | cb)` that accept an `error | null` **first**, then a `result |
  callback` that comes always **last**. Never `return` a result or `throw` an
  error from a callback (`return` and `throw` are unusable from a callback)
- Sync operation (CPU-bound, discouraged) blocks the event loop and puts on hold
  all concurrent processing blocking the whole application
    - Interleave each step of a CPU-bound sync operation with `setImmediate()`
      to let pending IO tasks to be processed by the event loop (not efficient
      due to context switching, more complex interleaving algorithm)
    - Pool of reusable external processes `child_process.fork()` with a
      communication channel `process/worker.send().on()` leaves the event loop
      unblocked (efficient: reusable processes run to completion, unmodified
      algorithm)
    - `new Worker()` threads with communication chanels
      `parentPort/worker.postMessage().on()`, per-thread own event loop and V8
      instance (small memory footprint, fast startup time, safe: no
      syncronization, no resource sharing)
- Avoid mixing sync/async callback behavior under the same inferface. To convert
  sync call to async use
    - Mictrotask `process.nextTick()` is executed just after the current
      operaiton, before other pending IO tasks
    - Immediate task `setImmediate()` is executed after all other pending IO
      tasks in the same cycle of the event loop
    - Regular task `setTimeout()` is executed in the next cycle of the event
      loop
- Uncaught error thrown from an async callback propagates to the stack of the
  event loop (not to the next callback, not to the stack of the caller that
  triggered the async operation), `process.on(unchaughtException, err)` is
  emitted and process exits with a non-zero exit code
- **Callback hell** = deeply nested code as a result of **in-place nested
  anonymous callbacks** that unnecessary consume memory because of closures
    - Do not abuse in-place nested anonymous callbacks
    - Early return principle = favor `return/continue/break` over nested
      `if/else`
    - **Create named callbacks** with clearly defined interface (and without
      unnecessary closures)
- **Sequential iteration** (recursion) = executes async tasks in sequence one
  task at a time until all complete or the first error
    ```js
    function task(v, done) {
      setTimeout(
        () => v ? (console.log(v), done(null, v)) : done(new Error("oh")), 500
      )
    }
    function sequence(tasks, done) {
      let i = 0
      function next(err) {
        if (err) { return done(err.message) }
        i < tasks.length ? tasks[i++](next) : process.nextTick(() => done(null))
      }
      next()
    }
    const tasks = [1, 2, 3].map(el => done => task(el, done)) // 1, 2, 3, null
    const tasks = [1, 2, 0, 3].map(el => done => task(el, done)) // 1, 2, oh
    sequence(tasks, console.log)
    sequence([], console.log) // null
    ```
- **Parallel execution** (sync loop through all tasks) = executes tasks in
  parallel (unlimited) until all complete or the first error
    ```js
    function parallel(tasks, done) {
      let completed = 0
      function check(err) {
        if (err) { return done(err.message) }
        if (++completed === tasks.length) { return done(null) }
      }
      if (tasks.length === 0) { return process.nextTick(() => done(null)) }
      for (const task of tasks) { task(check) }
    }
    const tasks = [1, 2, 3].map(el => done => task(el, done)) // 1, 2, 3, null
    const tasks = [1, 2, 0, 3].map(el => done => task(el, done)) // 1, 2, oh, 3
    parallel(tasks, console.log)
    parallel([], console.log) // null
    ```
- **Limited parallel execution** (sync loop until a limit) = executes at most N
  tasks in parallel until all complete or the first error
    ```js
    function parallelLimit(tasks, limit, done) {
      let i = 0, running = 0, completed = 0
      function check(err) {
        if (err) { return done(err.message) }
        if (++completed === tasks.length) { return done(null) }
        if (--running < limit) { parallel() }
      }
      function parallel() {
        while (i < tasks.length && running < limit) {
          tasks[i++](check); ++running
        }
      }
      if (tasks.length === 0) { return process.nextTick(() => done(null)) }
      parallel()
    }
    // 1, 2 | 3, 4 | 5, null
    const tasks = [1, 2, 3, 4, 5].map(el => done => task(el, done))
    // 1, 2 | 3, oh | 4
    const tasks = [1, 2, 3, 0, 4].map(el => done => task(el, done))
    parallelLimit(tasks, 2, console.log)
    parallelLimit([], 2, console.log) // null
    ```

## Observer pattern

- **Observer** = an object/subject notifies its observers on its state changes
- `EventEmitter` = registers multiple observing `listeners(args, ...)` for
  specific event types `ee.on|once(event, listener)`, `ee.emit(event, args,
  ...)`, `ee.removeListener(event, listener)` unsubscribe listeners when they
  are no longer needed to avoid memory leaks due to captured context in listener
  closures
- `EventEmitter` continuously notifies **multiple observers on different types
  of recurrent events** and **does not have trust issues** (loose coupling =
  callback is controlled by the caller). En event can be fired multiple times or
  not fired at all. Events are **guaranteed to fire async** in the next cycle of
  the event loop. Combining `EventEmitter` with the **callback** interface is an
  elegant and flexible solution
- Async result/error is proparaged through `emit` events. Never `return` a
  result or `throw` an error from an `EventEmitter` (`return` and `throw` are
  unusaable with `EventEmitter`). Always register a listener for the `error`
  event

## Promise pattern

- **Sequential iteration** (dynamic promise chain in a sync loop)
    ```js
    function task(v) {
      return new Promise((resolve, reject) =>
        setTimeout(
          () => v ? (console.log(v), resolve()) : reject(new Error("oh")), 500
        )
      )
    }
    function sequence(tasks) {
      let p = Promise.resolve()
      for (const task of tasks) { p = p.then(task) }
      return p
    }
    // 1, 2, 3, undefined
    const tasks = [1, 2, 3].map(el => done => task(el, done))
    // 1, 2, oh
    const tasks = [1, 2, 0, 3].map(el => done => task(el, done))
    sequence(tasks).then(console.log, console.error)
    sequence([]).then(console.log, console.error) // undefined
    ```
- **Parallel execution** (sync loop for all tasks)
    ```js
    function parallel(tasks) {
      let completed = 0
      return new Promise((resolve, reject) => {
        function check() {
          if (++completed === tasks.length) { return resolve() }
        }
        if (tasks.length === 0) { return resolve() }
        for (const task of tasks) { task().then(check, reject) }
      })
    }
    // 1, 2, 3, undefined
    const tasks = [1, 2, 3].map(el => done => task(el, done))
    // 1, 2, oh, 3
    const tasks = [1, 2, 0, 3].map(el => done => task(el, done))
    parallel(tasks).then(console.log, console.error)
    parallel([]).then(console.log, console.error) // undefined
    ```
- **Limited parallel execution** (sync loop until limit)
    ```js
    function parallelLimit(tasks, limit) {
      return new Promise((resolve, reject) => {
        let i = 0, running = 0, completed = 0
        function check() {
          if (++completed === tasks.length) { return resolve() }
          if (--running < limit) { parallel() }
        }
        function parallel() {
          while (i < tasks.length && running < limit) {
            tasks[i++]().then(check, reject); ++running
          }
        }
        if (tasks.length === 0) { return resolve() }
        parallel()
      })
    }
    // 1, 2 | 3, 4 | 5, undefined
    const tasks = [1, 2, 3, 4, 5].map(el => done => task(el, done))
    // 1, 2, | 3, oh | 4
    const tasks = [1, 2, 3, 0, 4].map(el => done => task(el, done))
    parallelLimit(tasks, 2).then(console.log, console.error)
    parallelLimit([]).then(console.log, console.error) // undefined
    ```

## Async/await pattern

- `async` function **always returns a Promise** immediately and synchronously.
  `try/catch/throw` inside an `async` function works for both sync and async
  code. Use `return await` to prevent errors on the caller side and `catch`
  errors locally
    ```js
    async function localError() {
      try { return task(false) } // Caller oh
      try { return await task(false) } // Local oh
      catch (e) { console.error("Local", e.message) }
    }
    localError().catch(e => console.error("Caller", e.message))
    ```
- `async` function is a `Promise`-`yield`ing generator that on each `await`
  expression `yield`s a Promise and suspends the generator, its internal state
  is saved and the control is retuned to the event loop. When the Promise that
  has been `await`ed settles, the control is given back to the `async` function
  and the generator is resumed
- `Promise` abstraction and `async/await` syntax is used to manage async
  operations in a sync-like manner. However, if async operations are unrelated,
  `await` introduces unnecessary blocking. Do not use `await` inside a loop, use
  `Promise.all()` instead
    ```js
    console.log(await task(1), await task(2)) // sequence, slow
    console.log(await Promise.all([task(1), task(2)])) // parallel, fast
    ```
- **Sequential iteration** (sync loop with await)
    ```js
    async function sequence(tasks) {
      for (const task of tasks) { await task() }
    }
    // 1, 2, 3, undefined
    const tasks = [1, 2, 3].map(el => done => task(el, done))
    // 1, 2, oh
    const tasks = [1, 2, 0, 3].map(el => done => task(el, done))
    try {
      console.log(await sequence(tasks))
      console.log(await sequence([])) // undefined
    } catch (e) { console.error(e.message) }
    ```
- **Parallel execution** (await loop)
    ```js
    // Use the prallel() from the Promise pattern above
    // 1, 2, 3, undefined
    const tasks = [1, 2, 3].map(el => done => task(el, done))
    // 1, 2, oh, 3
    const tasks = [1, 2, 0, 3].map(el => done => task(el, done))
    try {
      console.log(await parallel(tasks))
      console.log(await parallel([])) // undefined
    } catch (e) { console.error(e.message) }
    ```
- **Limited parallel execution**
    ```js
    // Use the parallelLimit() from the Promise pattern above
    // 1, 2 | 3, 4 | 5, undefined
    const tasks = [1, 2, 3, 4, 5].map(el => done => task(el, done))
    // 1, 2, | 3, oh | 4
    const tasks = [1, 2, 3, 0, 4].map(el => done => task(el, done))
    try {
      console.log(await parallelLimit(tasks, 2))
      console.log(await parallelLimit([], 2)) // undefined
    } catch (e) { console.error(e.message) }
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
