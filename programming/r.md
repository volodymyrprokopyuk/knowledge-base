# R

- TODO:
    - R tools `littler`, `styler` DONE, `argparser`, `r-optparse`, `docopt`
    - R language `rlang`, `lobstr`
    - OOP `R6` DONE, `proto`, `sloop`, `vctrs`
    - Data structures `hash`, `clock`, `fastmap`
    - IO `fs`
    - Statistics `psych`, `easystats`
    - Math `Matrix`
    - Data `data.table` DONE, `fst`
    - Utils `logger`, `lgr`
    - Profiling `profvis`
    - Benchmarking `bench` DONE
    - Destructuring `zeallot`
    - Database `RPostgres`
    - Web `plumber`, `jsonlite` DONE
    - Parallelism and concurrency `future` interface, `callr` swap / fork, `parallel`
      cluster, `batchtools` batch HPC (Slurm)
    - Plotting `ggally`, `ggstatsplot`

## General aspects

- R evolution
    - Assembler = machine instructions + memory locations
    - Fortran = data types + atomics + arrays + subroutines
    - S = (object = vector + attributes) + side-effect free, functional computation +
      interactive environment + local reference + copy-on-modify of non-local reference
    - R
        - Functional OOP = different objects + same generics (interface) + different
          methods (implementations) + immutable objects created by generics
        - Encapsulated OOP = mutable objects + reference semantics + side effects +
          mutable objects mutataed by methods
- Show package help `package?ggplot2`
- Show object help `?ggplot2::object`
- Search help `??"search"`
- Non-syntactic name `` `_name` ``, `` `%...%` <- function(l, r) {...} ``
- Anonymous body `{ ...; ... }`
- Subsetting + assignment = subassignement `x[i] <- v`, `x[] <- v`
- Immutable objects
    - Copy-on-modify semantics for shared objects `x <- c(1, 2); y <- x; y[[1]] <- 10`
    - Modify-in-place optimization for a single name `y[[2]] <- 20`
    - Environments are always modified in place (reference semantics)
    - Vector deep copy, list shallow copy
    - Character vector uses global string pool
- R type system
    - Vector type
        - Vector (atomic vector)
            - Atomic vector (logical, integer, double, character, complex, raw)
            - Factor (ordered factor)
            - Date and time
        - List (generic vector)
            - List
            - Data frame, tibble
    - Node type
        - Function
        - Environment

## Vectors, factors, datetimes, matrices and arrays

- Missing value `NA`, `is.na(x)`, `na.omit(x)`
- Null object = zero-length vector or absent vector `NULL`, `is.null(x)`
- Atomic vector
    - Logical `TRUE`, `FALSE`, `T`, `F`, `any(cond)`, `all(cond)`, `is.logical(x)`
        - Boolean operations `v[!v & v | xor(v1, v2)]`, `if (T && || F) {...}`
    - Integer `1L`, `0x1aL`, `is.integer(x)`
        - Set operations `union`, `intersect`, `setdiff`, `setdiff(union, intersect)`
    - Double `1.2`, `1.2e3`, `[-]Inf`, `NaN`, `is.double(x)`, `is.[in]finite(x)`,
      `is.nan(x)`
    - Character `"string"`, `'string'`, `is.character(x)`
        - `nchar`, `cat`, `paste`, `sprintf`, `strsplit`, `substr`, `grep[l]`,
          `[g]regexpr`, `regexec`, `[g]sub`
    - Complex `1+2i`
    - Raw (binary)
- Vector (homogeneous, fixed, flat, linear, element)
    - Creation `c(1, ...)`, `c(a = 1, ...)`, `append(x, values, after)` combine /
      append, `seq(start, end, step)`, `rep(obj, times, length.out)`,
      `vector(mode, length) = mode(length)`, `seq_along(x)`
    - Access `length(x)`, `names(x)`, `is.null(dim(x))`, `typeof(x)`,
      `which(x, arr.ind)` logical -> integer, `order | sort(x, decreasing)`, `rev(x)`,
      `unique(x)`, `match(x, table)`, `sample(n)` random permutation, `split(x, factor)`
    - Subsetting `[...]` (preserves structure, returns multiple elements, 1-based)
        - Identity `v[]` original vector, `v[0]` zero-length vector
        - Postion `v[1]`, `v[c(1, 2)]`
        - Logical `v[c(T, F)]`, `v[v > 0]`, `v[... & | ...]`
        - Match names `v[c("a", "b")]` named vector
        - Lookup table = chracter metching agains named vector
    - Remove `v[-c(...)]`, `v[-which(v < 0)]`
- Matrix = vector with `$dim` (homogeneous, column-first, row x column)
    - Creation `matrix(x, nrow, ncol, byrow)`, `cbind(...)`, `rbind(...)`
    - Access `nrow(x)`, `ncol(x)`, `dim(x)`, `rownames()`, `colnames()`, `is.matrix(x)`,
      `diag(n | x)`, `t(x)` transpose, `solve(x)` inverse, `upper | lower.tri(x, diag)`,
      `%*%`
    - Subsetting (matrix subsetting simplifies dimensionality on `m[1 | "a" | T]`,
      `drop = F` preserves dimensionality)
        - Single vector `m[c(...)]` `m` as column-first 1D vector
        - Vector per dimension `m[c(...), c(...), drop = F]`, `m[c(...), ]`,
          `m[, c(...)]`
        - Matrix `m[matrix(...)]` row = location in `m`, column = dimension of `m`
    - Remove`m[-c(...), -c(...)]`
- Array (homogeneous, row x column x layter x block)
    - Creation `array(x, dim, dimnames)`
    - Access `dim(x)`, `dimnames(x)`, `is.array(x)`
- S3 atomic vectors
    - Factor = integer vector that represents categorical data with a fixed set of
      predefined levels
        - Creation `factor(x, levels, labels)`, `ordered(x, levels, labels)`
        - Access `levels(x)`, `droplevles(x)`, `cut(x, breaks)`, `interaction(...)`,
          `reorder(x)`
    - Date vector = double vector that represents the number of days since 1970-01-01
        - Creation `Sys.Date()`, `as.Date("1970-01-01")`
    - Calendar time = double vector that represents the number of seconds since
      1970-01-01
        - Creation `as.POSIXct("1970-01-01 00:00:00", tzone = "UTC")`
    - Duration = double vector that represents the number of units between two datetimes
        - Creation `as.difftime(x, units)`

## Lists, data frames and tibbles

- List = generic vector (heterogeneous, extensible, recursive, hierarchical, member)
    - Create `list(...)`, `append(x, values, after)`, preserve, `c(...)` flatten
    - Access `length(x)`, `names(x)`, `is.list(x)`
    - Subsetting (preserves dimensionality)
        - Identity `l[]` original list, `l[0]` zero-length list
        - Member reference `l[[1 | "a"]]` single object (simplifies structure),
        - Member reference `l$name = l[["name"]]`
        - List slicing `l[1]` list (preserves structure)
        - Recursive subsetting `l[c(1, 2)] = l[[1]][[2]]`
    - Remove `l[[1]] <- NULL`, assign `NULL` `l[[1]] <- list(NULL)`
- Data frame / tibble (hegerogeneous, rectangular named list of equal-length vectors,
  observations [rows] of a set of variables [columns])
    - Creation `data.frame | tibble(a = c(...), ..., [row.names])`, `rbind(...)`
    - Tibble does not support `row.names` -> `rownames_to_column(df)` converts data
      frame's `row.names` into a tibble's regular column
    - Access `ncol(x)`, `nrow(x)`, `dim(x)`, `colnames(x)`, `rownames(x)`,
      `is.data.frame(x)`, `is_tibble(x)`
    - Subsetting
        - Single vector `df[c(...)]` list subsetting
        - Vector per dimension `df[c(...), c(...)]` matrix subsetting
        - `df$name`, `df[df$name > 0,]`

## Control flow

- Conditionals
    - Atomic `if (T && || F) { ... } [else { ... }]` statement
    - Vector `ifelse(test, yes, no)` function
    - Choice `switch(str.expr, match = value, ..., else)` function
- Loops
    - Vector iteration `for (... in ...) { ... }` statement
        - Preallocate the output container with `vector(mode, length)` when generating
          data in a `for` loop
    - Condition `while (...) { ... }` statement
    - Explicit exit `reapet` + `break` | `next` statements
    - Generic -> specific `repeat` -> `while` -> `for`
    - Implicit loop functions
        - Array `apply(x, margin, fun)`
        - List `lapply(x, fun)` -> list, `sapply(s, fun)` -> array
        - Data frame `tapply(x, factor.index, fun)`

## Functions

- Function object = `formals()` positional, named, variadic `...` + `body()` +
  `environment()` lexical scoping
- Lexical scoping = parse-time structures (vs dynamic scoping = runtime structures)
    - Name masking = names defined inside a function mask names defined outside a
      function
    - Namespace sharing = functions and variables share the namespace, but non-function
      objects are not considered when resolving function name
    - Dynamic lookup = name resolution and evaluation happens at runtime, not at
      definition time (function's behavior depends on objects defined outside function's
      environment)
    - Independent function execution = every time a function is called a new environment
      is created to host its execution
- R function can have `attributes()`, but primitive C function has all set to `NULL`
- `f <- function(...) {on.exit(..., add = T, after = F) ... [return(...) |
  invisible(...) | stop(...)] }` call by name = evaluate argument on first use (by need
  if memoized)
- Explicit `return(...)`, `invisible(...)` return prevents automatic printing,
  `(f(...))` prints invisible return. `<-` returns invisibly and allows chaining.
  Side-effecting functions should return their input invisibly to allow explicit capture
  of the input in a pipe
- Function singature `args(func)`, `missing(arg)`
- Order of arguments matching
    - 1. Exact (named arguments take precedence over positional)
    - 2. Partial (avoid)
    - 3. Positional (first or two most commonly used arguments)
- Ellipsis = variadic arguments `list(...)`, `..n` to pass arguments to inner functions
- Function call
    - Prefix `f(x, ...)` any call can be written in prefix form, knowing the name of
      non-prefix funciton allows to override its behavior
    - Infix `x + | %...% y` is left-associative (prefix `` `+`(x, y) ``)
    - Replacement `names(df) <- c(...)` in-place (copy) argument modification, (prefix
      `` `names<-`(df, c(...)) ``)
        - `` `f<-` <- \(x, value) { x[...] <- value; x } `` call `f(x) <- value`
        - `` `f<-` <- \(x, extra, value) { ...; x }`` call `f(x, extra) <- value`
    - Spacial forms (R language features = primitive C functions)
        - Parentheses and braces `` (x) -> `(`(x)  ``, `` {x} -> `{`(x) ``
        - Subsetting `` x[i] -> `[`(x, i) ``, `` x[[i]] -> `[[`(x, i) ``
        - Control flow
            - `` `if`(condition, consequent, alternative) ``
            - `` `for`(i, sequence, action) ``
            - `` `while`(condition, action)  ``
            - `` `repeat`(action) ``
            - `` `next`() ``
            - `` `break`() ``
        - Function `` `function`(arguments, body, environment) ``
    - Apply `do.call(func, list)`
- Lazy evaluation of function arguments = promise / thunk = expression + environment +
  memoized value (et most once evaluation)
- Default arguments are evaluated inside the function, hence can be defined in terms of
  other arguments or even variables defined inside the function. `missing(arg) == T` if
  default value is used
- Explicitly passed function arguments are evaluated in outside (global) environment
- The enclosing environment of the manufactured function is the an execution environment
  of the function factory (closure, force function factory arguments to avoid lazy
  evaluation bugs)

## Environments

- Environment binds names to values and implements reference semantics (in-place
  modification, not copying)
    - Avoid copies of large data via R6 encapsulated OOP built on top of environments
    - Manage state within a package across function calls via explicit environments
      `set_a(x) { o <- e$a; e$a <- x; invisible(o) }`
    - Environment is a hashmap via `hash` package
- Every environment has a parent environment that is used to implement lexical scoping
- The empty environment is the root of the environment hierarchy and does not have a
  parent
- Assignment `<- ` creates a binding in the current environment
- Super assignment `<<-` rebinds an existing name in the parent of a current environment
- Active binding = access and assignment operations on objects in environment are
  programmed in R
- Environment types
    - Global environment = user-defined funcitons and objects
    - Package environment = package external interface that exposes functions to a user
        - Package attached last to the search path by `library(package)`or
          `require(package)` becomes an immediate parent of the global environment
        - Package is loaded automatically when one of its functions is accessed via
          `package::function`
        - Search path `.global -> library(b) -> library(a) -> Autoloads -> package:base`
        - Parent environment of a package varies based on order of other attached
          packages
    - Function environment = environment where a function is defined (closure) used to
      implement lexical scoping
    - Namespace = package internal interface that controls function variables resolution
        - Namespace hides package internal implementation details from a user
        - Every function in a package is associated with a pair of environments: package
          environment + namespace environment
        - Every binding in the package environment is also in the namespace, so every
          function can use every other function in a package
        - Imports environment = bindings to functions used by the package (NAMESPACE
          file) `namespace -> imports -> base -> global`
    - Execution environment = ephemeral environment (child of function environment)
      created every time the function is called
    - Caller environment = environment from which a function was called
        - Call stack is made up of frames = evaluation contexts
        - Frame = expression (function call) + execution environment + previous frame +
          `on.exit()` handlers + `return()` context + condition system handlers
        - Dynamic scoping = lookup variables in a call stack rather than in the
          execution environment
- The parent of the global environment is the last loaded package
- Ancestors of global environment inclue every attached package
- Attach functions to an environment `with(funcs, ...)` -> `attach(funcs)` ->
  `rlang::env_bind(e, !!!funcs)`

## Conditions

- Signal condition (with default behavior)
    - Unrecoverable `stop(err)` = `rlang::abort(err)` abort function execution
    - Recoverable `warning(warn)` = `rlang::warn(warn)` retained till function exit
    - Informational `message(msg)` reported immediately
- Handle condition
    - Ignore `try({...})` error, `suppressWarnings({...})` warnings,
      `suppressMessages({...})` messages (all implemented in terms of `tryCatch()`)
    - Condition handlers temporarily override or supplement the default condition
      behavior
    - Handle error `tryCatch({...}, error | warning | message, finally)` registers
      exiting handler (usually for error conditions), terminates wrapped code and return
      control to the context where `tryCatch()` was called
        - Handler function is passed a condition object
        - The return value of an exiting handler is returned to the caller
        - `finally` is a block of code, not a function (`on.exit()` is implemented using
          `finally`)
        - Existing handler is called in the context of the call to `tryCatch()`
    - Handle warning and message `withCallingHandlers({...}, error | warning | message)`
      registers calling handlers (usually for warning and message conditions), after the
      condition is handled control returns to the context where the condition was
      signaled and the wrapped code is resumed
        - The return value of a calling handler is ignored as the wrapped code resumes,
          so calling handlers are only useful for their side effects
        - By default the condition propagates to parent calling handlers after being
          processed by the current calling handler (`rlang::cnd_muffle(cnd)`)
        - Calling handler is called in the context that signaled the condition
- Custom condition

  ```r
  tryCatch(
    custom_condition = \(cnd) list(message = cnd$message, detail = cnd$detail),
    abort | signal(
      message = "Custom message",
      class = "custom_condition",
      detail = "Extra meta"))
  ```

## Functional programming (FP)

## Object-oriented programming (OOP)

- OOP fundamentals
    - Polymprphism
        - Single function = uniform interface
        - Multiple objects = different implementations
    - Encapsulation
        - Informaiton hiding = representation
        - Abstraction = interpretation
    - Inheritance
        - Structure inheritance = object fileds (single vs multiple inheritance)
        - Behavior inheritance = method dispatch (single vs multiple dispatch)
- OOP systems
    - Functional OOP
        - Methods belong to generic functions (behavior)
        - Method call looks like function call
        - Classes defines only object fields (structure)
    - Encapsulated OOP
        - Methods belong to objects
        - Object encapsulats both structure (fields) and behavior (methods)
    - Prototype-based (classless) OOP
        - Inherit from a chain of objects (not classes) that are dynamically mutated at
          runtime
        - Clone and extend objects that become prototypes for more specialized objects
- `typeof(x)` base objects, `class(x)` S3 and S4 classes
- Base types
    - Vector `NULL`, `logical`, `integer`, `double`, `character`, `complex`, `raw`,
      `list`
    - Function `closure` regular R function, `special` internal R function,
      `builtin` primitive C function
    - Environment `environment`
    - S4 type `S4`
    - Language `symbol` name, `language` expression, `pairlist` formals
- Attributes = are ephemeral, lost by most operations, but `names` and `dim` are
  preserved (to preserve other attributes, custom S3 class has to be created)
    - Get / set indivitual attribute `attr(obj, attr)`
    - Get all attributes `attributes(obj)`
    - Set multiple attributes `structure(obj, attr = value ...)`
- S3
    - Functional OOP = informal structure (attributes) + GF + single inheritance +
      single dispatch, simple and concise
    - S3 object is a base type (class-less vector, list, data frame) with at least a
      `class` attribute (`unclass(x)` returns the base type)
    - Creation `structure(x, class = "a_class")`, `class(x) <- "a_class"`
        - User helper `a_class(base, attrs ...)` provides user interface to a S3 object
          creation, coerces the input to acceptable by the constructor
        - Efficient low-level constructor `new_a_class(base, attrs ...)` enforces
          consistent structure of objects and checks types of the base object and
          attributes
        - Expensive optional validator `validate_a_class(x)` returns a valid object or
          thrown a validation exception
    - Generic function (GF) `a_generic <- \(x, args ...) UseMethod("a_generic", x)`
      performs instance-based method dispatch to a concrete method implementation based
      on the `class` attribute of the first argument (single dispatch)
    - Method `a_generic.a_class(x, args ...)` must implement the generic interface
      defiined by the GF
    - Inheritance
        - `class` vector `c("subclass", "superclass")`
        - Delegetion to a superclass `NextMethod()`
        - To allow subclassing the parent constructor needs `...` and the `class`,
          argument
- S4
    - Functional OOP = formal structure (slots) + GF + multiple inheritance + multiple
      dispatch, strict and suitable for large projects
    - Class definition `setClass(classname, slots, prototype)` prototype = fields
      default values
    - Object instantiation
        - Low-level constructor `new("AClass", slots ...)`
        - User helper `AClass(...)` provides user interface to a S3 object creation,
          coerces the input to acceptable by the constructor
        - Validator `setValidity(classname, \(object) "error" | T)` called automatically
          by `new()`. Explicit validation `validObject(x)`
    - Subsetting = low-level slot access `@`, `slot(object, name)`
    - High-level slot accessor functions
        - Getter
            - Generic `setGeneric("aSlot", \(x) standardGeneric("aSlot"))`
            - Method `setMethod("aSlot", "AClass", \(x) x@aSlot)`
        - Setter
            - Generic `setGeneric("aSlot<-", \(x, value) standardGeneric("aSlot<-"))`
            - Method `setMethod("aSlot<-", "AClass",
              \(x, value) { x@aSlot <- value; validObject(x); x })`
    - Inheritance `setClass(contains)`
    - Generic `setGeneric(signature)` define interface
        - `signature` explicitly defines arguments used in method dispatch, otherwise
          all arguments are used in method dispatch
    - Method dispatch
        - `ANY` pseudo-class matches any class, is always the farther method (equivalent
          to S3 `default`)
        - `MISSING` pseudo-class matches whenever the argument is missing
        - Multiple inheritance -> dispatches on the closest method (on equal distance
          alphabetic order is used = kind of random)
- R6
    - Encapsulated OOP with non-idiomatic to R reference semantics built on top of
      environments (stateful simulation)
    - Encapsulated object = fileds + methods in a local namespace (R6) vs generic
      functions in the global namespace (S3, S4)
    - Methods belong to an in-place mutable object (not GFs) and allow side-effects +
      return value via `invisible(self)` in the same method for method chaining
    - Reference semantics (in-place modification vs copy-on-modify) or `$clone(deep)`
    - Definition `AClass <- R6Class("AClass",
      public = list(a_filed = NULL, a_method = \() ...))`
        - Public `$initialize(...)` is called by `$new(...)`
        - Private `$finalize()` automatic on GC cleanup of resources acquired by
          `$new(...)`
        - Field containing an R6 object
            - Shared across all instances if created in the class definition
            - Per instance if created in the `initialize(...)`
    - Access control `R5Class(public = list(), active = list(), private = list())`
        - `active` = property = dymanic field with an accessor and a mutator implemented
          as a single function using `if (missing(value)) read else write`
    - Creation `instance <- AClass$new(...)`
    - Access `instance$a_field`, `instance$a_field <- value`, `instance$a_method(...)`
    - Internal access to public / private members `self|private$a_member`,
      `self|private$a_member <- value`
    - Inheritance `SubClass <- R6Class("SubClass", inherit = SuperClass)`
        - Delegetion to a super class `super$a_member`
    - Extend an exiting class with `AClass$set("public", "name", a_field | a_method)`
    - R6 example

        ```r
        Counter <-
          R6Class(
            "Counter",
            public = list(
              initialize = \(init = 0) private$counter = init,
              # Public method
              increment = \(val = 1) private$counter <- private$counter + val),
            active = list(
              # Read-only property
              state = \() private$counter,
              # Read-write property
              value = \(val) if (missing(val)) private$counter else private$counter <- val),
            private = list(
              # private field
              counter = 0,
              finalize = \() cat("<Counter>: cleaning up...")))
        ```

## Metaprogramming (MP)

- System functions `getwd`, `setwd`
- Formatting `cat`, `paste`, `print`, `format`, `sprintf`
- Math funcitons `[cum]sum`, `[cum]prod`, `[cum]min|max`, `round`

## Package

- Minimal R package `devtools::create(path)` -> `R/` + `DESCRIPTION` + `NAMESPACE` +
  `[README.md]`
- Package statges
    - Source = local source code repository <- `devtools::create(path)`
    - Bundle = platform-agnostic `.tar.gz` compressed source with built vignettes <-
      `devtools::build()`
    - Binary = platform-specific `.tgz` compressed compiled R bytecode, C code, package
      metadata and documentation <- `devtools::build(binary = T)`
    - Installed = binary package decompressed into a package library in `.libPaths()` <-
      `install.packages(pkgs)` remote packages, `devtools::install(pkg)` local package
    - Loaded = loads and attaches an installed package <- `library(package)`,
      `devtools::load_all()` (`package::function(...)` only loads, but does not attach)
- Library = directory containing installed packages
    - Library search path `.libPaths()` -> system library, user library, package private
      library
- `R/*.R` code
    - Use common prefix in file names as directories are not allowed in `R/`
    - Use `styler` for code formatting
      `Rscript -e "styler::style_file('$SOURCE.R', strict = T)"`
    - Script = code is run when the script is loaded via `source("script.R")` ->
      immediate execution results
    - Package = code is run when the package is built with `devtools::build(package)` ->
      cached build-time results when the package is loaded
    - The top-level R code in a package is only executed when the package is built, not
      when the package is loaded -> never run code in the top-level of a package ->
      package should only create objects, mostly functions
    - Do not use `library(package)` in a package (modifies the global environment) ->
      use `DESCRIPTION` to specify package dependencies at install and load time
    - Do not use `source(file)` in a package (modifies the current environment) -> use
      `devtools::load_all()` which automatically `source()`s all files in `R/`
    - Place package side effects in `.onLoad|.onAttach|.onUnload(libname, pkgname)` that
      are called automatically by R (e. g. to set package `options()`)
- `DESCRIPTION` package metadata + package dependencies on other packages
    - `Title` concise (one line), `Description` extended (one paragraph)
    - `Version` major.minor.patch.devel
    - `Authors@R: c(person(...), ...)`
    - `URL` package repository
    - `License`
        - MIT / BSD (license must always be distirbuted with the code)
        - GPL (derivative work must also be GPL)
        - CC (public domain: anyone can use the code for any purpose)
    - `Imports` (just load) mandatory package deps, `Suggests` optional package deps <-
      `devtools::use_package(package, type)`
    - `Depends: R (>= 4.0.0)` system library, extenal executable (load + attach)
- Documentation
    - Reference documentation `roxygen2 #' @tag Markdown -> man/*.Rd -> TEXT|HTML|PDF`
      <- `devtools::document()` (document functions, S3 / S4 generics and methods, R6
      classes, datasets, packages)
    - Hight-level documentation vignettes `vignettes/*.Rmd` <-
      `devloots::build_vignettes()` -> `browseVignettes("package")`, `vignette("x")`
- Automated testing `tests/testthat/test-*.R`, `tests/testthat.R` <- `devloots::test()`
    - `expect_*(actual, expected)` action (expectation) = verifies the expected outcome
      via a binary assertion
    - `test_that("Goal", { ... })` = defines the goal for a group of expectations (each
      test is run in its own environment and is self contained)
    - Single `context()` per file = groups related tests together
- `NAMESPACE` package imports of objects from other packages + package exports <-
  `devtools::document()`
    - Package dependencies on objects are looked up in the `NAMESPACE` imports (not the
      global namespace)
    - `NAMESPACE` exports specifies objects provided by the package
    - Object lookup = global environment -> `search()` path (reversed list of attached
      packages)
    - Load an installed package with `::` = load code and data + register S3, S4
      generics and methods = access package objects with `package::object`
    - Attach a loaded package with `library(x)` (laod + attach) = add package objects to
      the `search()` path = access package objects directly with `object`
    - `library(x)` throws an error -> use in a script, but never in a package -> use
      `DESCRIPTION.Imports` (just loads a package) and `DESCRIPTION.Depends` (loads +
      attaches a package)
    - Always use `Imports` as the package should minimize changes to the global
      environment (including the `search()` path)
    - Only use `Depends` if the package directly extends tha base package by providing
      new objects along with the objects from the base package
    - `requireNamespace("x", quietly = T)` returns `F` -> use in a package to check
      availability of `Suggest`ed packages
    - `NAMESPACE` is usually generated by `roxygen2`
        - `#' @export` -> `export()`, `S3method()`, `exportClass()`, `exportMethod()`
          (export as little as possible)
- Dataset
    - `data/*.Rdata` files do not use the `NAMESPACE` mechanism and do not need to be
      exported

## `data.table` data manipulation and analysis vs SQL data storage and access

- data.table = in-memory, ordered columns store
- Concise and consistent (cryptic, but powerful) syntax for in-memory (column store vs
  SQL row store, ordered data.table vs unordered SQL) data manipulation
- In-memory processing, by reference in-place update with no memory allocation for
  intermediate results by default or `shallow()` copy-on-update semantics (referential
  transparency)
  `dt[i, j, by / keyby, ...]` =
  `FROM[WHERE / ORDER BY -> row subset, join / reorder; SELECT -> compute, update;
   GROUP BY -> aggregate]` 1. subset / reorder `i`, 3. compute / update `j` 2. group by
   `by / keyby`
    - `i`: `a == b` condition, `a:b` range, `order(a, -b)` sort
    - `j`: `a` vector, `.(A = a | B = expr(b) | ..columns | -c(...) | year:day |
      invisible(ggplot(...)))` data.table
    - `by`: `.(a | B = expr(b))`, `c("a", "b")`
    - Chaining `dt[...][...]`
    - Append rows / data.tables `rbindlist(list(dt, ...))` (rows cannot be efficiently
      inserted or deleted in the middle of a data.table)
    - Append / update columns `:=`
    - Change column order `setcolorder(col, ...)`
- Special variables
    - `.N` = `length(col)` number of observations in the current group / subset
    - `.SD` (reflexive reference to a subset of data) = data.table (without `:=`)
        with all columns (except the grouping columns) or `.SDcols` (character name,
        interger position, logical mask, `patterns(...)`) for the current group if
        access to all columns of `.SD` is required
        - Identity `dt[, .SD]` == `dt[]`
        - Column subsetting `dt[, ..cols]`, `dt[, .SD, .SDcols = cols]`
        - Grouping `dt[, .SD[1 | sample(.N, 1L) | .N], by = ...]`
- Create `data.table(..., key)`, `fread(file|https|stdin)`
- Coerce `setDT(df|list)` in-place, `as.data.table(other)` copy
- Reference semantics = add new columns, update or delete existing columns in-place
    - `shallow(dt)` copy = copies a pointer to an object, leaving a single copy of data
    - Deep `copy(dt)` = copies data of an object, creating two copies of data
    - `dt[, A := a]`, `dt[, c("A", "B") := .(a, b)]`, `dt[, ``:=``(A = a, B = b)]`
    - `:=` returns the result invisibly, use additional `[]` to print the result
    - Delete column `dt[, A := NULL]`, `dt[, c("A", "B") := NULL]`,
      `dt[, ``:=``(A = NULL, B = NULL)]`
    - Dereferencing `dt[, (cols) := .(...)]`
    - Only `:=`, and all `set*()` funcitons have reference sumantics on a data.table
- Keys = fast binary search subsetting by equality (vs vector scan) with shorter syntax
  suitable for repeated subsetting on the same column
    - At most one key can be set on multiple columns, which physically reorders the rows
      by key columns by reference always in increasing order and marks the columns as
      key columns in subsequent subsetting by setting the `sorted` attribute on the
      data.table (key does not use memory)
    - Duplicatge keys are allowed (uniqueness is not enforced)
    - Set a key `dt |> setkey(a, b)`, `dt |> setkeyv(c("a", "b"))`
    - Get a key `key(dt)`
    - Remove the key `dt |> setkey(NULL)`
    - Key subset (only the first column) `dt[.(c("a", "b"))]` == `dt[c("a", "b")]`
      column a matches a OR b
    - Key subset (multiple columns) `dt[.(c("a", "b"), c("c", "d"))]` columna a matches
      a OR b AND column b matches c OR d
    - Key subset (only the second column) `dt[.(unique(a), c("c", "d"))]`
    - Choose matching rows `dt[..., mult = "first|all|last", nomatch = NULL|NA]`
- Indices
    - Compute and cache multiple indices temporarily on the fly without
      physical reordering (but consuming memory) of a data.table
    - The `on` argument is mandatory for indices (cleaner syntax) and is recommended for
      the key
    - Set an index `dt |> setindex(a, b)`, `dt |> setindexv(c("a", "b"))` create and
      cache the index
    - Get indices `dt |> indices()`
    - Remove all indices `dt |> setindex(NULL)`
    - Indices subset `dt[.(c("a", "b")), on = c("a", "b")]` creates, but does not cache
      the index
    - Auto indexing = automatically computes and caches an index on the first use of
      `==` or `%in%`
- Reshaping
    - Wide-to-long `melt(id.vars, measure.vars, variable.name, value.name)`
    - Long-to-wide `dcast(id.var + ... ~ measure.var + ..., value.name)`
- Programming on data.table
    - data.table provides a robust mechanism built upon `data.table::substitute2()` for
      parameterizing expressions passed to `i`, `j` and `by / keyby` arguments of
      `[.data.table(..., env = list(...))`
- Join as subsetting via keys or indices
    - Right (outer) equi join (default) `x |> setkey(id); y |> setkey(id); x[y, j]`,
      `x[y, j, on = .(id | a = b), nomatch = NA]` (all rows from `y` + matching rows
      from `x` with duplicates)
    - (Inner) equi join `x[y, j, on = .(id), nomatch = NULL]` (only matching rows from
      `x` and `y` with duplicates)
    - Full (outer) equi join `x |> setkey(id); y |> setkey(id)` (key on both tables is
      the implicit joining variable) `x[y[c(x$id, y$id) |> unique()]]` (left join +
      right join)
    - Anti-join `x[!y, j, on = .(id)]` (only `x` rows with no match in `y`)
    - Semi-join `x[na.omit(x[y, on = .(id), which = T])]` (only `x` rows with match in
      `y` without duplicates)
    - Non-equi join `x[y, j, on = .(x.col <= y.col, ...)]`
    - Rolling join `x[y, j, on = .(id), roll = T | -Inf]` (join on a rolling window of
      values, not the exact match)
    - Overlapping range join `x |> setkey(start, end)`, `x |> foverlaps(y)`
    - Grouping by each `i` `x[y, j, by = .EACHI]` evaluates `j` on matched subset of `x`
      corresponding to each row in `y`
    - Update on join (recode factor)

      ```r
      f <- factor(c("small", "large", "large", "small", "medium"))
      data.table(f)[.(f = levels(f), to = c("L", "M", "S")), f := i.to, on = .(f)]$f
      levels(f) <- list(S = "small", M = "medium", L = "large")
      ```
