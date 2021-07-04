# R

- TODO: `littler`, `styler`, `lintr`

## General aspects

- Help in R console `?obj` show help, `??"search"` search help
- Non-syntactic name `` `_name` ``
- Subsetting + assignment = subassignement `x[i] <- v`, `x[] <- v`
- Immutable objects
    - Copy-on-modify semantics for shared objects `x <- c(1, 2); y <- x; y[[1]] <- 10`
    - Modify-in-place optimization for a single name `y[[2]] <- 20`
    - Environments are always modified in place (reference semantics)
    - Vector deep copy, list shallow copy
    - Character vector uses global string pool
- Type system
    - Vector type: atomic vector + factor + date and time, list + data frame + tibble
    - Node type: function, environment

## Vectors, factors, datetimes, matrices and arrays

- Missing value `NA`, `is.na(x)`, `na.omit(x)`
- Null object = zero-length vector or absent vector `NULL`, `is.null(x)`
- Atomic vector
    - Logical `TRUE`, `FALSE`, `T`, `F`, `any(cond)`, `all(cond)`, `T && || F`,
      `v & | v`, `! v`, `is.logical(x)`
    - Integer `1L`, `0x1aL`, `is.integer(x)`
    - Double `1.2`, `1.2e3`, `[-]Inf`, `NaN`, `is.double(x)`, `is.[in]finite(x)`,
      `is.nan(x)`
    - Character `"string"`, `'string'`, `is.character(x)`
    - Complex
    - Raw (binary)
- Vector (homogeneous, fixed, flat, linear, element)
    - Creation `c(1, ...)`, `c(a = 1, ...)` combine / append, `seq(start, end, step)`,
      `rep(obj, times, length.out)`
    - Access `length(x)`, `names(x)`, `is.null(dim(x))`, `typeof(x)`,
      `which(x, arr.ind)`, `order | sort(x, decreasing)`, `rev(x)`, `match(x, table)`
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
    - Atomic `if (... && || ...) { ... } [else { ... }]` statement
    - Vector `ifelse(test, yes, no)` function
    - Choice `switch(str.expr, match = value, ..., else)` function
- Loops
    - Iteration `for (... in ...) { ... }` statement
    - Condition `while (...) { ... }` statement
    - Explicit exit `reapet` + `break` | `next` statements
    - Implicit loop functions
        - Array `apply(x, margin, fun)`
        - List `lapply(x, fun)` -> list, `sapply(s, fun)` -> array
        - Data frame `tapply(x, factor.index, fun)`

## Functions

- `f <- function(...) { ... [return(...) | invisible(...)] }` call by name (by need if
  memoized)
- Function singature `args(func)`, `missing(arg)`
- Arguments matching positional, exact / partial (avoid), mixed, variadic `list(...)`

## Environments

- Global environment user-defined funcitons and objects `ls()`, `rm(obj)`,
  `detach("pakcage:x"|obj)`, `attach(obj)`
- Package environment = build-in functions and objects `ls("package:ggplot2")`
- Local environment = function lexical scope
- Search path `search()` -> `.GlobalEnv, library(...), package:base`
- Environment where object is defined `environment(obj)`

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
