# R

- TODO:
    - R tools `littler`, `styler`, `lintr`
    - R language `rlang`, `lobstr`
    - OOP `R6`
    - Data structures `hash`
    - Statistics `psych`

## General aspects

- Help in R console `?obj`, `` ?`(` `` show help, `??"search"` search help
- Non-syntactic name `` `_name` ``, `` `%...%` <- function(l, r) {...} ``
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
    - Complex
    - Raw (binary)
- Vector (homogeneous, fixed, flat, linear, element)
    - Creation `c(1, ...)`, `c(a = 1, ...)` combine / append, `seq(start, end, step)`,
      `rep(obj, times, length.out)`, `vector(mode, length)`, `seq_along(length)`
    - Access `length(x)`, `names(x)`, `is.null(dim(x))`, `typeof(x)`,
      `which(x, arr.ind)` logical -> integer, `order | sort(x, decreasing)`, `rev(x)`,
      `unique(x)`, `match(x, table)`, `sample(n)` random permutation
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
        - Access `levels(x)`, `cut(x, breaks)`
    - Date vector = double vector that represents the number of days since 1970-01-01
        - Creation `Sys.Date()`, `as.Date("1970-01-01")`
    - Calendar time = double vector that represents the number of seconds since
      1970-01-01
        - Creation `as.POSIXct("1970-01-01 00:00:00", tzone = "UTC")`
    - Duration = double vector that represents the number of units between two datetimes
        - Creation `as.difftime(x, units)`

## Lists, data frames and tibbles

- List = generic vector (heterogeneous, extensible, recursive, hierarchical, member)
    - Create `list(...)` preserve, `c(...)` flatten
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
  Side-effecting functions should return invisibly
- Function singature `args(func)`, `missing(arg)`
- Order of arguments matching
    - 1. Exact
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

## Conditions

- Signal condition (with default behavior)
    - Unrecoverable `stop(err)` = `rlang::abort(err)` abort function execution
    - Recoverable `warning(warn)` = `rlang::warn(warn)` retained till function exit
    - Informational `message(msg)` reported immediately
- Handle condition
    - Ignore `try({...})` error, `suppressWarnings({...})` warnings,
      `suppressMessages({...})` messages
    - Condition handlers temporarily override or supplement the default condition
      behavior
    - Handle error `tryCatch({...}, error | warning | message, finally)` registers
      exiting handler (usually for error conditions), terminates wrapped code and return
      control to the context where `tryCatch()` was called
        - Handler function is passed a condition object
        - `finally` is a block of code, not a function (`on.exit()` is implemented using
          `finally`)
    - Handle warning and message `withCallingHandlers({...}, error | warning | message)`
      registers calling handlers (usually for warning and message conditions), after the
      condition is handled control returns to the context where the condition was
      signaled and the wrapped code is resumed
        - The return value of a calling handler is ignored as the wrapped code resumes,
          so calling handlers are only useful for their side effects
        - By default the condition propagates to parent calling handlers after being
          processed by the current calling handler (`rlang::cnd_muffle(cnd)`)

## Functional programming (FP)

## Object-oriented programming (OOP)

- Attributes = are ephemeral, lost by most operations, but `names` and `dim` are
  preserved (to preserve other attributes, custom S3 class has to be created)
    - Get / set ndivitual attribute `attr(obj, attr)`
    - Get all attributes `attributes(obj)`
    - Set multiple attributes `structure(obj, attr = value ...)`
- Data sets serialization `read.table`, `read.csv`, `write.table`, `read.csv`
- R objects serialization `dput`, `dget`
- S4 `@ = $`, `slot(...) = [[...]]`

## Metaprogramming (MP)

- System functions `getwd`, `setwd`, `format`, `sprintf`
- Math funcitons `sum`, `prod`, `round`
- Evaluation `eval(parse(text = "1 + 2"))`
- Errors: `message`, `warning`, `stop`
