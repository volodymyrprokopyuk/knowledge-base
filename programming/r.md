# R

- TODO: `littler`, `styler`, `lintr`

## General aspects

- Help in R console `?obj` show help, `??"search"` search help
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
        - Creation `factor(x, levels)`, `ordered(x, levels)`
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

## Strings

- String `nchar`, `cat`, `paste`, `sprintf`, `strsplit`, `substr`, `grep[l]`,
  `[g]regexpr`, `regexec`, `[g]sub`

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
- `f <- function(...) { ... [return(...) | invisible(...)] }` call by name (by need if
  memoized)
- Function singature `args(func)`, `missing(arg)`
- Arguments matching positional, exact / partial (avoid), mixed, variadic `list(...)`
- `do.call(func, list)`
- Lazy evaluation of function arguments = promise / thunk = expression + environment +
  memoized value (et most once evaluation)
- Default arguments are evaluated inside the function, hence can be defined in terms of
  other arguments or even variables defined inside the function. `missing(arg) == T` if
  default value is used
- Explicitly passed function arguments are evaluated in outside (global) environment

## Environments

- Global environment user-defined funcitons and objects `ls()`, `rm(obj)`,
  `detach("pakcage:x"|obj)`, `attach(obj)`
- Package environment = build-in functions and objects `ls("package:ggplot2")`
- Local environment = function lexical scope
- Search path `search()` -> `.GlobalEnv, library(...), package:base`
- Environment where object is defined `environment(obj)`, `exists("name")`

## Conditions

- Exceptions `message(msg)`, `warning(warn)`, `stop(err)`, `try(expr)`,
  `tryCatch(expr, ..., finally)`, `suppressWarnings(expr)`

## Object-oriented programming (OOP)

- Attributes = are ephemeral, lost by most operations, but `names` and `dim` are
  preserved (to preserve other attributes, custom S3 class has to be created)
    - Get / set ndivitual attribute `attr(obj, attr)`
    - Get all attributes `attributes(obj)`
    - Set multiple attributes `structure(obj, attr = value ...)`
- Data sets serialization `read.table`, `read.csv`, `write.table`, `read.csv`
- R objects serialization `dput`, `dget`
- S4 `@ = $`, `slot(...) = [[...]]`

## Functional programming (FP)

## Metaprogramming (MP)

- System functions `getwd`, `setwd`, `format`, `sprintf`
- Math funcitons `sum`, `prod`, `round`
- Evaluation `eval(parse(text = "1 + 2"))`
- Errors: `message`, `warning`, `stop`
