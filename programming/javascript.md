# JavaScript

## Lexical scope and closures (inside-out variable lookup and resolution)

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

## This (dynamic binding rules)

- `this` is dynamically defined at runtime (late binding) for every function,
  depends on the location where a function is called (call-site), not where a
  function is declared (lexical scope) and how a funciton is called, implicitly
  passes an execution context object (like dynamic scope) to a function and is
  not related to lexical scope
- Binding rules for `this` (from the highest to the lowest precedence)
    - `new` binding = construction call of regular funciton with the `new`
      operator `new f()`, `this` points to a brand new obect, which is
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

## Object (property and accessor descriptors)

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
  (`writable`) => assign value
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
## Object prototype (behavior delation via linked objects)
