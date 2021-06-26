# R

- TODO: `littler`, `styler`

## General aspects

- Help `?obj`, show help `??"search"` search help
- Literals `Inf`, `-Inf`, `NaN`, `NA` missing data, `NULL` null object
- Type predicate `is.finite`, `is.infinite`, `is.nan`, `is.na`, `na.omit`, `is.null`
- Copy-on-write semantics `a <- b`
- Logical `TRUE`, `FALSE`, `T`, `F`, `any`, `all`, `T && || F`, `v & | v`, `! v`

## Vectors and factors

- Vector (homogeneous, fixed, element) `c`, `length`, `seq`, `rep`, `sort`, `which`, `v2
  <- v1`
    - Subsetting (1-based) `v[1]` element (preserves structure), `v[c(1, 2)]`, `v[c(T,
      F)]`, `v[v > 0]` sub-vector, `v[... & | ...]` `v[-c(...)]`, v[-which(v < 0)]
      remove
- Matrix (homogeneous, column-first, row x column) `matrix`, `cbind`, `rbind`, `dim`,
  `nrow`, `ncol`, `t`, `%*%`, `solve`
    - Subsetting `m[1]` column-first vector, `m[1, 1]` element, `m[1,]` row, `m[, 1]`
      column, `m[c(...), c(...)]` sub-matrix, `m[-c(...), -c(...)]` remove, `diag`
      diagonal
- Array (homogeneous, row x column x layter x block) `array` + array subsetting
- Factor (string vector with ordering) `factor`, `levels`, `c`, `cut`, `length` + factor
  subsetting

## Lists and data frames

- List (heterogeneous, extensible, nested) `list`, `length`, `names`
    - Subsetting `l[[1]]` member reference -> single object, `l$name`,
      `l[1]` list slicing -> list (preserves structure)
- Data frame (hegerogeneous named list of equal-length vectors, observation records =
  rows of variables = columns) `data.frame`, `nrow`, `ncol`, `dim`, `rbind`, `colnames`,
  `rownames`
    - Subsetting `df[1, 1]`, `df[c(1), c("name")]`, `df$name`, `df[df$name > 0,]`

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
    - Implicit loop
        - Array `apply(x, margin, fun)`
        - List `lapply(x, fun)` -> list, `sapply(s, fun)` -> array
        - Data frame `tapply(x, factor.index, fun)`

## Functions
    - `f <- function(...) { ... [return(...)] }` call by name (by need if memoized)
    - Function singature `args(func)`, `missing(arg)`
    - Arguments matching exact, partial, positional, mixed, variadic `list(...)`

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

- S3 `attributes`, `attr`, `class`
- Data sets serialization `read.table`, `read.csv`, `write.table`, `read.csv`
- R objects serialization `dput`, `dget`

## Functional programming (FP)

## Metaprogramming (MP)

- System functions `getwd`, `setwd`, `format`, `sprintf`
- Math funcitons `sum`, `prod`, `round`
- Evaluation `eval(parse(text = "1 + 2"))`
- Errors: `message`, `warning`, `stop`
