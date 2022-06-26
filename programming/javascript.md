# JavaScript

## Lexical scope and closures for variable lookup

- Tokenization (stateless) | lexing (stateful) => parsing (AST + per-scope
  hoisting of `var`aible and `function` declarations) => optimization => code
  generation (JIT) => execution (variable assignment, function call)
- Engine (orchestrator) = variable lookup for variable / parameter assignment
  (LHS container) and variable / parameter referencing (RHS value)
- Compiler (code generator) = variable creation in the appropriate scope
- Scope (nested storage tree) = variable storage and retrieval + shadowing
- Function creates a new nested function scope not accessible from the outside
  (module pattern)

    ```js
    function module(a) {
      let b = a // private state
      function f() { return ++b } // closure over state
      return { f } // public interface
    }
    const m = module(10) // module instance
    console.log(m.f()) // 11
    ```
- Lexical scope (closures) is defined statically at write-time (the scope chain
  is based on the source code)
    - Dynamic scope (`this`-like) defined at execution-time and depends on the
      execution path (the scope chain is based on call stack)
- Block scope (`const`, `let`) = declare variables as close as possible to where
  they are used
    - `var` => function scope + hoisting (of variable and function declarations)
    - `const`, `let` => block scope at any `{ ... }` even explicitly defined
    - `try/catch(e)` => block scope
- Function declarations are hoisted before variable declarations
- Closure = a returned function can access its lexical scope even when the
  function is executing outside its lexical scope

## `this` dynamic binding rules

- `this` is dynamically defined at runtime (late binding) for every function,
  depends on the location where a function is called (call-site), not where a
  function is declared (lexical scope) and how a funciton is called, implicitly
  passes an execution context object (like dynamic scope) to a function and is
  not related to lexical scope
- Binding rules for `this` (from the highest to the lowest precedence)
    - `new` binding = construction call of regular funciton with the `new`
      operator `new f()`, `this` points to a brand new object, which is
      automatically returned from the function (unless the function returns its
      own alternate object). The `new` operator ignores `this` hard binding with
      `bind()`
    - Explicit binding = function invocation through `f.call(this, args, ...)`
      or `f.apply(this, [args])` including the hard binding `ff = f.bind(this,
      args, ...)` (partial application + currying), `this` points to the first
      argument
    - Implicit binding = function invocation `o.f()` through a containing
      context object `{ f: f }`, `this` points to the containing context object
    - Default binding = standalone function invocation `f()` including callback
      invocation, `this` = `undefined` as the global object is not eligible for
      the default binding in the `strict mode`
- Lexical `this` (`bind` alternative) = arrow functions `(...) => { ... }`
  discard all the traditional rules for `this` binding and instead use the
  lexical `this` from the immediate lexical enclosing scope. Arrow function is a
  syntactic replacement for `self = this` closures. The lexical `this` binding
  of an arrow function cannot be overrided even with `new`

## `object` property and accessor descriptors

- Type vs object
    ```js
    const s = "a" // string type, immutable value, automatic coercion to object
    console.log(typeof s, s instanceof String) // string, false
    const s2 = new String("a") // String object, allows operations
    console.log(typeof s2, s2 instanceof String) // object, true
    ```
- Object = container of named references to properties (values and functions),
  however functions never belong to objects. Syntactic property access `o.p` vs
  programmatic key access `o["p"]`
    ```js
    const o = { a: 1 }
    console.log("a" in o) // true
    for (const p in o) { console.log(p, o[p]) } // a 1
    delete o.a
    console.log(o) // { }
    ```
- Property descriptor vs accessor descriptor
    ```js
    const o = { }
    Object.defineProperty( // property descriptor
      o, "a", { value: 1, writable: true, enumerable: true, configurable: true }
    )
    o.a = 2
    console.log(o) // 2

    const o = { }
    Object.defineProperty(o, "a", { // accessor descriptor
      set: function(val) { this._a = val },
      get: function() { return this._a * 2 }
    })
    o.a = 1
    console.log(o.a) // 2

    const o = { // object literal setter and getter
      set a(val) { this._a = val },
      get a() { return this._a * 2 }
    }
    o.a = 1
    console.log(o.a) // 2
    ```
- `[[Get]]` own property lookup => prototype chain lookup => return `undefined`
- `[[Put]]` accessor descriptor (`set`, `get`) => property descriptor
  (`writable`) => prototype chain lookup => assign value directly to the object
- Object immutability `Object.preventExtensions()`, `Object.seal()`,
  `Object.freeze()`

## Iteration (`[Symbol.iterator]`)

- Iteration over arrays (indexing) and objects (properties). Custom iterator
  `{ [Symbol.iterator]: function () { return { next } } }` + `for/of`
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
## Object `prototype` for property lookup

- Prototype chain = every object has an `o.prototype` link to another object
  ending at `Object.prototype` (kind of global scope for variables)
    ```js
    const o = { a: 1 }
    // new object o2.[[Prototype]] = o (prototype chain)
    const o2 = Object.create(o)
    console.log("a" in o2) // true
    console.log(o2.a) // 1
    for (const p in o2) { console.log(p, o2[p]) } // a, 1
    ```
- All `function`s get by default a public, non-enumerable property `prototype`
  pointing to an object => each object created via `new F()` operator is linked
  to the `F.prototype` effectively delegating access to `F.prototype`'s
  properties (prototypal inheritance = objects are linked to other objects)
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
    G.prototype.d =
      function() { return F.prototype.b.call(this) + 2 } // call parent method
    const o2 = new G()
    console.log(o2.a, o2.b(), o2.c, o2.d()) // 1, 2, 3, 4
    ```
- Purely flat data storage `o = Object.create(null)` without `prototype`
  delegation
- Behavior delegation via `prototype` links / chain (objects are linked to other
  objects forming a network of peers, not a vertical hierarchy as with classes)
- Mutual delegation of two objects to each other forming cycle is disallowed
- ES6 class = syntax sugar built on top of prototypal inheritance and behavior
  delegation
    ```js
    class F {
      constructor() { this.a = 1 } // constructor + property
      b() { return 2 } // method
    }
    const o = new F()
    console.log(o.a, o.b()) // 1, 2
    class G extends F { // prototypal ihheritance
      constructor() { super(); this.c = 3 }// call parent constructor
      d() { return super.b() + 2 } // call parent method
    }
    const o2 = new G()
    console.log(o2.a, o2.b(), o2.c, o2.d()) // 1, 2, 3, 4
    ```
- Function chaining
    ```js
    function N(x) { this.a = x }
    N.prototype.add = function add(x) { this.a += x; return this }
    console.log(new N(1).add(2).add(3).a) // 6
    ```

## Types

- Types are related to values, not variables (which may store any value)
- `let x; typeof x === "undefined"` vs `typeof <undeclared> === "undefined"`
- `function` is a `[[Call]]`able `object`
- The type of value determines whether the value will be assigned by copy
  (primitives `boolean`, `number`, `string`, `symbol`) or by reference
  (`object`s, `array`, `function`, automatically boxed values)
- `symbol` special unique primitive type used for collision-free internal
  properies on objects
    ```js
    const sym = Symbol("a")
    const o = { [sym]: 1 }
    console.log(o[sym]) // 1, collision-free property
    ```

## Coercion

- Coercion always results in one of the scalar primitive types
- `null == undefined // true`
- Both `==` (implicit coercion) and `===` (no coercion) compare two `object`s by
  reference (not by value) `{ a: 1 } ==(=) { a: 1 } // false`
- Use `===` (to avoid coercion) instead of `==` with `true`, `false`, `0`, `""`,
  `[]`

## JS grammar

- Assignment expression returns assigned value `a = 1 // 1`
- `continue <label>` continues an outer loop `label: for(...)`
- `break <label>` breaks out of an inner loop or a block `label: { ... }`
- `let a = b || <default value>` vs `a && a.b()` guarded operation + short
  circuiting
- `=`, `?:` right-associative

## Async

- Single-threaded event loop (sequential execution on every tick)
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
- Concurrency = split 2 or more tasks into atomic steps, schedule steps from all
  tasks to the event loop (interleave steps from different tasks), execute steps
  in the event loop in order to progress simultaneously on all tasks
- Callbacks = strict separation between now (current code) and later (callback).
  Non-linear definition of a sequential control flow and error handling, trust
  issues due to control delegation (inversion of control, continuation)
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
- Trustable promises = composable, time consistent, future, eventual value
  placeholder
  (proxy) that behaves the same across now and later (by making both of them
  always async later). Async flow control completion event (promise object) to
  subscribe to (separation of consumers from producer) possibly multiple
  consumers. Solves the trust issues of callbacks by inverting callback control
  delegation. Once a promise is settled, the resolved value or rejected reason
  becomes immutable
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
      .then(console.log).catch(console.error)
    ```
- Promises are guaranteed to be async. Promises don't get rid of callbacks, they
  just let the caller control callbacks locally via `Promise.then(cb)` instead
  of passing callabcks to a third party code as in case of callbacks only
  approach. Repeated calls to `resolve` and `reject` are ignored
- `Promise.resolve(x)` normilizes values and misbehaving thenable to trustable
  and compliant Promises
- `p.then()` automatically creates a new Promise in a chain resolved with the
  `return`ed value or the unwrapped `return`ed Promise or rejected with the
  `throw`n error
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
  concurrent, unordered Promises or rejects with the first rejected Promise
- `Promise.race([])` a latch that resolves or rejects with the first settled
  Promise (the other Promises cannot be canceled due to immutability, hense are
  settled and just ignored)
- Promises API
    - `Promise.resolve(x)`, `Promise.reject(x)`
    - `p.then(success, [failure])`, `p.catch(failure)`, `p.finally(always)`
    - `Promise.race([]) => first success / first failure`
    - `Promise.any([]) => first success / all failure`
    - `Promise.allSettled([]) => [all either success or failure]`
    - `Promise.all([]) => [all success] / first failure`
- Callback => promise
    ```js
    function f(x, cb) {
      setTimeout(_ => { if (x >= 0) { cb(null, "ok") } else { cb("oh") } }, 100)
    }
    f(1, console.log)
    f(-1, console.error)
    function promisify(f) {
      return function(...args) {
        return new Promise((resolve, reject) => {
          f.apply(null, args.concat(
            function(e, x) { if (e) { reject(e) } else { resolve(x) } }))
        })
      }
    }
    const ff = promisify(f)
    ff(1).then(console.log)
    ff(-1).catch(console.error)
    ```
