# Gauche Scheme

- Small core + powerful tools for language extension
- Nested block (`let`) structure + shared namespace for variables and procedures
- Lexcal scope + closures
- Tail-call optimizaiton + recursion (`let name`)
- First-class continuations (`call/cc`) for (non-local exit, exceptions, generators,
  coroutines, backtracking, actors)
- Delayed evaluation (`delay`, `force`, `lazy`, `eager`) + promises (thunk + memoization)
- Hygienic macros (`define-syntax`, `syntax-rules`) + recursive macro expansion
- Fundamental and derived forms
    - `lambda` + `:optional`, `:key`, `:rest` (procedure (primitieve, closure), binding
      block, recursion)
        - `case-lambda` procedure with variable number of arguments
        - `let` (block), `let*` (nested), `letrec` (set!), `let name` (recursion)
        - `and-let*` sequential `*` guarded `and` binded `let` expressions
        - `values` construct values, `receive` access values
            - `define-values`, `set!-values`, `let-values`
    - `define` (variable definition + initialization), `define-constant`
    - `set!` (variable update + assignment), generalized `set!`, `~` universal accessor,
      `push!`, `pop!`, `inc!`, `dec!`, `update!`
    - `if` (conditionals)
        - `and`, `or`, `cond`, `case`, `when`, `unless`
    - `quote`=`'` (do not evaluate)
        - `quasiquote`=`` ` ``, `unquote`=`,`, `unquote-splicing`=`,@`
    - `define-syntax`, `let-syntax`, `letrec-syntax` (form declaration)
    - `syntax-rules` (pattern-template extension), `syntax-case` (precedural extension)
    - Sequencing + side effects
        - `begin`
    - Iterator + side effects
        - `do`
    - Delayed (lazy) evaluation + promises
        - `delay` e -> promise e, `force` promise e -> e, `lazy` promise e -> promise e,
          `eager`
    - First-class objects
        - `call/cc` (continuation),
        - `with-input-from-file`, `with-output-to-file` (port)
        - (stream)
- Literals `+inf.0`, `-inf.0`, `+nan.0`, `#t`, `#f`, `#\char`, `"string"`, `(car . cdr)`
  (pair), `(list)`, `#(vector)`, `#[char-set]`, `#/reg-exp/`
- String interpolation `#"Value ~expr ~(expr) ~|var|"`
- Parameters `make-parameter`, `parametrize`, `dynamic-wind` (dynamic environment
  management, context switch)
- Combinators (return procedure)
    - `cut` compact parameter spacialization without currying macro
    - `.$`, `compose` procedure composition
    - `identity`, `constantly`, `complement`, `any-pred`, `every-pred`
    - `$`, `$*` procedure application chaining macro

## Equality / equivalence + comparison / ordering + hashing

- Equality (equivalence)
    - `eq?` symbol, keyword, boolean, object reference
    - `eqv?` number `=`, character `char=?`
    - `equal?` aggregate (collection, object), recursive, string `string=?`
        - `object-equal? obj1 obj2` generic function for user-defined data types (UDDT)
          uses `equal?`  recursively
- Comparison (ordering)
    - `compare` -1, 0, 1, number `<`, character `char<?`, string `string<?`
        - `object-compare obj1 obj2` generic function for UDDT
- Hasing
    - `eq-hash` for `eq?`
    - `eqv-hash` for `eqv?`
    - `default-hash` for `equal?`
        - `object-hash obj rec-hash` generic function for UDDT + `combine-hash-value`
- Comparator R7RS = record with equality + comparison/ordering + hasing abstraction
    - `comparator-test-type` type predicate `type?`
    - `=?` equality predicate `equal?`
    - `<?`, `<=?`, `>?`, `>=?` ordering predicate `<` or `compare`
    - `comparator-compare` comparison `compare` or `<`
    - `comparator-hash` hasing `default-hash`
    - `default-comparator`=`equal? compare default-hash` automatically extended for UDDT
        via `object-equal?`, `object-compare`, and `object-hash` that must be defined
    - treemap `compare` keys, hashtable `default-hash` keys
    - `sort` and `merge` `compare` elements

## Exceptions and conditions

- Conditions `define-condition-type` condition, parent, predicate, slots
- Signaling `error`, `errorf` simple error, `raise`, `condition` compound condition
- Handling `guard`
- Cleaning up `unwind-protect` only calls `cleanup` on normal exit or exception, ignors
  continuation control escapes
- Control flow `dynamic-wind` always calls `before` and `after` on any control flow
  transition. Low-level management of exceptions, parameters, continuations and ports

```scheme
(define-condition-type <app-error> <error> app-error? [reason reason])

(guard
 (e
  [(<app-error> e)
   (format #t "ERROR: <app-error> ~a ~a" (reason e) (condition-message e))]
  [else (format #t "ERROR: ~a" e)])
 (error "Message")
 (error <app-error> :reason "Reason" "Message")
 (raise (condition [<app-error> (reason "Reason") (message "Message")])))
```

## Module system

- Open-ended `<module>` creates a namespce, maps symbols -> bindings and controls
  visibility of bindings
- Modules and files are orthogonal concepts
    - A single file may contain multiple modules
    - A module may be defined in multiple files
- Definition `define-module`, `select-module`
- Export`export (rename)`
- Import `import :only :except :rename :prefix` not transitive (module composition)
- Extend `extend` multiple modules (module precedence list) (module inheritance)
- Usage `(use a.module)` = `(require "a/module")` + `(import a.module)` (file = module)
    - `load` file each time at run-time
    - `require` file once due to auto `provide` memoization at compile time
    - `autoload` delayed file / module loading on first reference of a variable / macro

```scheme
(define-module a (export pi) (define pi 3.14))
(define-module b (export e) (define e 2.71))
;; (define-module c (import a b)) ;; composition
(define-module c (extend a b)) ;; inheritance
(select-module c)
(print pi " " e)
```

## Object system

- Gauche type system = classes are used to describe types
    - Implicit (not explicitly expressed)
    - Structural (type name does not matter)
    - Nominal (type name matters)
    - Dynamic (latent) = every value (not variable) knows its type at run-time
    - Strong (value determines what operations can be applied)
    - `<toc>` supertype of all types
    - `<bottom>` subtype of all types, does not have an instance
    - `<object>` supertype of all user-defined classes
    - `class-of obj`, `class-name cls`, `is-a? obj cls`, `subtype? sub sup`
- Class `define-class` (class precedence list)
    - Allocation `:allocation` `:instance`, `:class`,
        - `:virtual` `:slot-ref` computed value, `:slot-set!` data validation
    - Initialization `:init-keyword` in constructor, `:init-value` evaluated once,
      `:init-form` evaluated each time
    - Access GFs: `:accessor` read-write, `:getter` read-only, `:setter` write-only
    - Standard accessors: `slot-ref`, `slot-set!`, `class-slot-ref`, `class-slot-set!`,
      `~` universal accessor
- Instance `make`, `initialize obj args` post initialization
- Method `define-method`, `next-method`, `write-object obj port` serialization
- Property `name` with pseudo private slot `.name`

```scheme
(define-class <person> ()
  ([.name :init-value "nobody"]
   [name :allocation :virtual
         :slot-ref (lambda (p) (~ p '.name))
         :slot-set! (lambda (p n) (when [#/^[A-Z]\w+/ n] (set! (~ p '.name) n)))]))
(let ([p (make <person>)]) (set! (~ p 'name) "Vlad") (print (~ p 'name)))
```

- Record type = portable (standard) and efficient class `<record>`
    - `(use gauche.record)`: `define-record-type` (constructor, predicate, accessors,
      [mutators])

## Pattern matching

- `match`, `match-lambda`, `match-lambda*`, `match-let`, `match-let*` `match-define`
- Pattern syntax
    - Anything `symbol` matches anything and binds the matched value
    - Placeholder `_` matches anything without binding
    - Repetition `...` applies last pattern repeatedly
    - Literals `#t`, `#f`, `#\a`, `1`, `"s"` match themselves with `equal?`
    - Quotes `''symbol`, `''(s expression)`, `':keyword` match themselves with `equal?`
    - List `()`, `(a b)`, `(a b ...)`, `(a b ..k)` k min, `(a b . c)` + pair
    - Vector `#()`, `#(a b)`, `#(a b ...)`, `#(a b ..k)` k min
    - Object `(@ class (slot pat) ...)` matches class / record instance and all slot
      names
    - Procedure `(= procedure pat)` applies `procedure` then matches on the result
    - Predicate `(? predicate pat ...)` matches on the `predicate` and all patterns
    - Pattern combinators `(and pat ...)`, `(or pat ...)`, `(not pat ...)`
    - Quasipattern `'quasipattern` matches itself as leterals except the unquoted
      patterns that are pattern variables
        - Quasiquote turns off the evaluation except the unquoted subtree
        - Quasipattern turns off the pattern syntax except the unquoted subtree

## Collections and sequences

- **Collection** Gauche = unordered set of objects. Collection provides generic
  traversing over list, vector, string, hash table, user-defined class using the method
  dispatch of the object system (CLOS)
    - `(use gauche.collection)`: `fold`, `fold2/3` `map`, `for-each`, `find`,
      `find-min`, `find-max`, `filter`, `remove`, `group-collection`, `size-of`,
      constructive methods: `map-to`, `filter-to`, `remove-to`, `coerce-to`
    - Iterator interface `end?` of collection, `next` element
    - Builder interface `add` element, `get` collection
- **Sequence** Gauche = ordered set of objects built on top of collection. Sequence
  provies index-based access and order-aware operations on top of collection
    - `(use gauche.sequence)`: `ref`=`~` and `subseq` with `set!`, `fold-right`,
      `fold-with-index` `map-with-index`, `for-each-with-index`, `find-with-index`,
      `sequence-contains`, `group-sequence` adjacent

## Eager comprehensions

- Eager comprehensions SRFI-42 `(use srfi-42)`: `(comprehension qualifiers body)`
    - (3) Comprehension-ec (collect, aggregate)
        - Collect `do-ec`, `list-ec`, `vector-ec`, `string-ec`
        - Side effects `(do-ec (:list i '(1 2 3)) (display i))`
        - List `(list-ec (:list i '(1 2 3)) i)`
        - String `(string-ec (:list i '(#\a #\b #\c)) i)`
        - Aggregate `any?-ec`, `every?-ec`, `first-ec`, `last-ec`, `fold-ec`
        - Any (short circuit on first #t) `(any?-ec (:list i '(1 2 3)) (even? i))`
        - Every (short circuit on first #f) `(every?-ec (:list i '(1 2 3)) (even? i))`
        - First (short circuit) `(first-ec 0 (:list i '(1 2 3)) i)`
        - Last `(last-ec 0 (:list i '(1 2 3)) i)`
        - Fold `(fold-ec 0 (:list i '(1 2 3)) i +)`
    - (1) Qualifiers (generate, filter)
        - Generational (typed) qualifiers = :generators `:list`, `:vector`, `:string`,
          `:integers` (infinite), `:rage`, `:real-range`, `:char-range`, `:port`,
          `:parallel` (zip, default is nested), `:while` (stop early), `:let` (introduce
          intermediary variables depending on the outer scope)
        - List append `(list-ec (:list i '(1 2 3) '(4 5)) i)`
        - Index `(list-ec (:list i (index j) '(a b c)) (cons i j))`
        - Range `(list-ec (:range i 1 7 2) i)`
        - Port `(with-input-from-string "Vlad"
          (lambda () (list-ec (:port i (current-input-port) read-char) i)))`
        - Nested = cartesian product (the rightmost generator spins fastest)
          `(list-ec (:list i '(1 2)) (:list j '(a b)) (cons i j))`
        - Zip `(:parallel (:list i '(1 2)) (:list j '(a b)))`
        - While from infinite `(list-ec (:while (:integers i) (< i 5)) i)`
        - Let intermediary variable between generators
          `(list-ec (:list i '(1 2 3)) (:let j (* i 10)) (cons i j))`
        - Control qualifiers `if`, `not`, `and`, `or`, `begin` (side effects)
        - `(list-ec (:list i '(1 2 3 4 5 6)) (if [even? i]) i)`
        - `(list-ec (:list i '(1 2 3 4 5 6)) (not [even? i]) i)`
        - `(list-ec (:list i '(1 2 3 4 5 6)) (and [even? i] [< i 5]) i)`
        - `(list-ec (:list i '(1 2 3 4 5 6)) (or [even? i] [odd? i]) i)`
        - Begin side effects between generators
          `(list-ec (:list i '(1 2 3 4 5 6)) (begin (display i)) i)`
    - (2) Body (evaluate, transform)

## Dictionaries = hash-table + tree-map

- Hash table `<hash-table>` Gauche = unordered key-value mapping with O(1) insertion and
  lookup
    - Construct `make-hash-table`, `alist->hash-table`, `hash-table-r7`
    - Lookup `hash-table-get`, `hash-table-exists?`, `hash-table-num-entries`,
      `hash-table-find`
    - Mutate `hash-table-put!`, `hash-table-delete!`, `hash-table-push!`,
      `hash-table-pop!`, `hash-table-update!`
    - Traverse `hash-table-for-each`, `hash-table-map`, `hash-table-fold`,
      `hash-table-keys`, `hash-table-values`
- Tree map `<tree-map>` Gauche = ordered key-value mapping with O(log n) insersion and
  lookup
    - Construct `make-tree-map`, `alist->tree-map` MISSING `tree-map-r7`
    - Lookup `tree-map-get`, `tree-map-exists?`, `tree-map-num-entries`, `tree-map-min`,
      `tree-map-max`, MISSING `tree-map-find`
    - Mutate `tree-map-put!`, `tree-map-delete!`, `tree-map-push!`, `tree-map-pop!`,
      `tree-map-update!`, `tree-map-pop-min!`, `tree-map-pop-max!`
    - Traverse `tree-map-for-each`, `tree-map-map`, `tree-map-fold`,
      `tree-map-fold-right`, `tree-map-keys`, `tree-map-values`
- **Dictionary** Gauche = key-value mapping. `(use gauche.dictionary)` provides generic
  functions for `<dictionary>` and `<ordered-dictuionary>` lookup, mutation and
  traversal
    - Lookup `dict-get` with `set!`, `dict-exists?`
    - Mutate `dict-put!`, `dict-delete!`, `dict-push!`, `dict-pop!`, `dict-update!`
    - Traverse `dict-for-each`, `dict-map`, `dict-fold`, `dict-fold-right` ordered,
      `dict-keys`, `dict-values`
- `(use scheme.mapping)` R7RS immutable, functional tree map (ordered keys)
    - Construct `mapping`, `mapping-unfold`, `alist->mapping`
    - Lookup `mapping-ref`, `mapping-ref/default` `mapping-empty?`, `mapping-contains?`,
      `mapping-size`, `mapping-min/max-key/value/entry`, `mapping-find`
    - Mutate `mapping-set`, `mapping-delete`, `mapping-update`, `mapping-update/default`
    - Traverse `mapping-for-each`, `mapping-map`, `mapping-filter`, `mapping-remove`,
      `mapping-fold`, `mapping-fold/reverse`, `mapping-keys`, `mapping-values`,
      `mapping-entries`, `mapping-any?`, `mapping-every?`
- `(use scheme.mapping.hash)` R7RS immutable, functional hash table (unordered keys)
    - `mapping` -> `hashmap`
- `(use scheme.set)` unordered set (no duplicates) and bags (with duplicates)
    - Construct `set/bag`, `set/bag-unfold`, `list->set/bag`
    - Lookup `set/bag-empty?`, `set/bag-contains?`, `set/bag-member`, `set/bag-size`,
      `set/bag-find`
    - Mutate `set/bag-adjoin`, `set/bag-delete`
    - Traverse `set/bag-for-each`, `set/bag-map`, `set/bag-filter`, `set/bag-remove`,
      `set/bag-fold`, `set/bag-any?`, `set/bag-every?`
    - Set `set/bag-union`, `set/bag-intersection`, `set/bag-difference`, `set/bag-xor`

## Delayed (lazy) evaluation + promises (thunk + memoization)

- `delay` e -> promise e, creates a promise, requires unbound memory for tail-recursive
  algorithms (R5RS)
- `lazy` promise e -> promise e, creates a promie for space-efficient tail-recursive
  lazy algorithms. Generally `lazy` surrounds the entire body of a function expressing
  lazy algorithm (SRFI-45)
- `force` promise e -> e + memoization
- `eager` e -> promise e, eagerly evaluated type converter to a promise
- **Generators** Gauche `(use gauche.generator)` = a procedure with no arguments that
  yields a series of values with possible side effects ending with the EOF (very
  lightweight implementation of on-demand calculations). Generators work in a pipeline
  (DAG) of generators representing a lazy value-propagation network
    - Construct (create generator)
        - `list->generator`, `vector->generator`, `string->generator`,
          `x->generator coll`
        - `giota`, `grange`, `circular-generator arg ...`, `gunfold`,
          `generate coroutine`

```scheme
(let ([g (generate (lambda (yield) (do ([i 0 (+ i 1)]) ([> i 10]) (yield i))))])
  #?=(generator->list g))
```

    - Combine (return generator)
        - `gcons* item ... gen`, `gappend gen ...` generators, `gconcatenate gen`
          generator of generators, `gflatten gen`
        - `gmap proc gen ...`, `gfilter pred gen`, `gremove pred gen` `gmerge < gen ...`
        - `gtake gen n`, `gdrop gen n`, `gtake-while pred gen`, `gdrop-while pred gen`
    - Consume (return non generator)
        - `generator->list gen`, `generator->vector gen`, `generator->string gen`,
          `generator-unfold gen unfold arg ...`
        - `generator-for-each proc gen ...`, `generator-fold proc seed gen ...`,
          `generator-find pred gen`, `generator-any pred gen`, `generator-every pred gen`
- **Lazy sequence** Gauche `(use gauche.lazy)` = indistinguishable from ordinary list
  structure (all list procedures can be used on a lazy sequence) with a lazy pair, whose
  `car` is immediately / eagerly evaluated and whose `cdr` is implicitly / automatically
  forced on demand. Lazy sequence is not strictly lazy and always evaluates one item
  ahead. Lazy sequences are built on top of generators
    - Construct `generator->lseq`, `lrange`, `liota`,`lunfold p f g s`, `x->lseq obj`
    - Combine `lcons` thunk for each item, `lappend seq ...` sequences, `lconcatenate
      seqs` sequence of sequences, `lmap proc seq ...`, `lfilter pred seq`, `ltake seq n`,
      `ltake-while pred seq`
- **Streams** Gauche = strictly lazy (both `car` and `cdr` are lazily evaluated when
  aboslutely needed) data structure with spacial procedures
    - `(use util.stream)`

## Input and output

- Current ports `current-input-port`, `current-output-port`, `current-error-port`
- Standard ports `standard-input-port`, `standard-output-port`, `standard-error-port`
- File I/O `with-input-from-file file thunk`, `with-output-to-file file thunk`
- String I/O `with-input-from-string str thunk`, `with-output-to-string thunk`,
  `with-string-io str thunk`
- Input `read` s-expression, `read-char`, `read-line`, `port->string port`, `eof-object?
  obj`, `define-reader-ctor` reader-time constructor
- Output `write obj`, `write-char char`, `write-string str`, `write-object obj port`
  machine, `display obj`, `print obj ...`, `pprint obj` human, `format dest fstr args
  ...`, `flush`
- Automatically read and write user-defined class

```scheme
(define-class <data> ()
  ([a :init-keyword :a :init-value 'a]
   [b :init-keyword :b :init-value 'b]))
(define-method write-object ([d <data>] p)
  (format p "#,(<data> ~a ~b)" (~ d 'a) (~ d 'b)))
(define-reader-ctor '<data>
  (lambda (a b) (make <data> :a a :b b)))
(let ([d #,(<data> 'A 'B)])
  (with-input-from-string (format #f "~a" d) (cut print (read))))
```

## Regular expressions

- Regular expressions syntax
    - `.` any char, `[char-set]` char set, `\w`, `\d`, `\s`, `\W`, `\D`, `\S`
    - `^` begin, `$` end
    - `\b` word boundary, `\B` non-word boundary
    - `re*` zero or more, `re+` one or more, `re?` zero or one (greedy matching)
    - `re{n}`, `re{n,m}` from n to m, either can be omitted (greedy matching)
    - `re*?`, `re+?`, `re??`, `re{n,m}?` non-greedy, lazy matching
    - `(re...)` capturing group (positional), `\1` positional backreference
    - `(?:re...)` clustering group (without capturing)
    - `(?<name>re...)` named capturing group, `\k<name>` named backreference
    - `(?i:re...)` case-insensitive clustering group
    - `(?-i:re...)` case-sensitive clustering group
    - `re1|re2` alternatives
    - `(?=re...)` positive lookahead, `(?!re...)` negative lookahead
    - `(?<=re...)` positive lookbehind, `(?<!re...)` negative lookbehind
    - `(?>re...)` atomic clustering = once matched to backtracking even if the following
      pattern fails
    - `re*+` = `(?>re*)`, `re++` = `(?>re+)`, `re?+` = `(?>re?)`
    - `(?test-re then-re)`, `(?test-re then-re|else-re)` conditional matching
- Construct `#/reg-exp/` compiled only once, `string->regexp str` compiled every time
- Match `re str` first match, `grxmatch re str` all matches
- Access `match index` group position, `match name` group name, `rxmatch-cond`,
  `rxmatch-case`
- Replace `regexp-replace re str subs`, `regexp-replace-all re str subs`

```scheme
(let* ([s "one 1 two 2 ten 10"] [r #/(?<wd>\w+) (?<nm>\d+)/] [m (r s)])
  #?=(list (m 0) (m 1) (m 2) (m 'wd) (m 'nm))
  #?=($ generator->list $ gmap (lambda (m) (list (m 0) (m 'wd) (m 'nm))) $ grxmatch r s)
  #?=(regexp-replace-all r s "\\2 \\k<wd>")
  #?=(regexp-replace-all r s (lambda (m) (format #f "~a ~a" (m 'nm) (m 1))))
  #?=(rxmatch-cond [(r s) (all wd nm) (list all wd nm)] [else #f])
  #?=(rxmatch-case s [r (all wd nm) (list all wd nm)] [else #f]))
```

## Random values

- `(use srfi-27)` SRFI-27: `(random-source-randomize! default-random-source)`
    - `random-integer`, `random-real`
- `(use data.random)` Gauche: `(set! (random-data-seed) (random-integer (expt 2 32)))`
    - Uniform distribution `integers$`, `integers-between$`, `booleans`, `chars$`,
      `reals$`, `reals-between$`, `samples$`, `regular-strings$`
    - Non-uniform distribution `reals-normal$`, `reals-exponential$`,
      `integers-geometric$`, `integers-poisson$`
    - Combinators `samples-from`, `weighted-samples-from`, `pairs-of`, `tuples-of`,
      `permulations-of`, `combinations-of`, `list-of`, `vector-of`, `string-of`,
      `sequence-of`
