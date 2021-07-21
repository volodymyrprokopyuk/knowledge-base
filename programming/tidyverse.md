# tidyverse

## purrr functional programming in R

### General aspects

- All `purrr` functions translate the `~` formula into an anonymous `function` where
  `.x` = `..1` and `.y` = `..2`
- `safely(.f, otherwise, quite)` returns `list(result, error)` applied as
  `x |> map(safely(f)) |> transpose()`
    - `quietly(.f)` captures messages and warnings

### Structure and environment indexing

- `pluck(.x, ..., .default)` generalized form of `[[...]]` to index deeply nested
  structures by name or by position, returns a single element or NULL when not exist
    - `pluck(.x, ...) <- value`
    - `chuck(.x, ...)` throws an error when not exist

### Map family (one-dimensional vector)

- `map(.x, .f, ...)`returns a list (generic vector) of the same length
    - `map_lgl()`, `map_int()`, `map_dbl()`, `map_chr()` returns an atomic vector
    - `map(.x, .selector, .default)` extracts by a `pluck()` selector
    - `map_if(.x, .p, .f, ..., .else)` maps based on condition
    - `map_at(.x, .at, .f, ...)` maps by name or position
- `modify(.x, .f, ...)` returns a modified copy preserving structure
    - `modify_if(.x, .p, .f, ..., .else)` modifies column-wise based on condition
    - `modify_at(.x, TODO)` modifies column-wise by name or position
    - `modify_in(.x, TODO)` modifies a single element by a `pluck()` selector
- `walk(.x, .f, ...)` executes the function for side effects and returns invisibly the
  input
- `map2(.x, .y, .f, ...)`
- `pmap(.l, .f, ...)` applies a function to each row of a data frame
- `imap` maps with index or name

### Reduce family

- `reduce(.x, .f, ..., .init)` reduces a list to a single value by iterativel applying a
  binary function, generalizes a binary function over a vector
    - `reduce2(.x, .y, .f, ..., .init)` function of three arguments
- `accumulate(.x, .f, ..., .init)` accumulates intermediate results of a vector
  reduction
    - `accumulate2(.x, .y, .f, ..., .init)` function of three arguments

### Predicates

- Any / all matches `some(.x, .p)`, `every(.x, .p)`
- First match value / index `detect(.x, .p)`, `detect_index(.x, .p)`
- Keep discard `keep(.x, .p)`, `discard(.x, .p)`

## ggplot2 grammar for graphics

### General aspects

- Plot `ggplot(data, mapping = aes())` maps data properties to graphic properties
- Data (tibble data frame, numerical + categorical) = each row represents one
  observation of potentially multiple variables (stored in columns)
- Geoms (points, lines, bars, boxes)
- Mapping
    - Aesthetic mapping `aes(x, y, color, fill, size, shape, group)`
    - Aesthetic setting `geom_point(color)`
- Scales
    - Palette `scale_fill_brewer(palette)`, `scale_fill_manual(values)`,
    `scale_color_manual(values)`
    - Percentage scale `scale_x|y_continuous(limits, labels)`
- Guides
    - Labels `labs(title, subtitle, caption, x, y)`, `xlab(label)`, `ylab(label)`
    - Limits `lims(x, y)`, `xlim(lower, upper)`, `ylim(lower, upper)`
- Theme `theme(axis, legend, panel, plot)`

### Scatter plot

- `geom_point(mapping, data)`

### Line plot

- `geom_line(mapping, data)`

### Bar plot (numeric values of discrete categories)

- Count `geom_bar()` (`stat_count()`)
- Value `geom_col(fill, color, position = "dodge | stack | fill", width)`
  (`stat_idenitty()`) >= `geom_bar(stat = "identity")`
- Labels `geom_text(aes(label, x, y), vjust, hjust)`
- Cleveland dot plot (cleaner alternative to a bar plot) `geom_point()`

### Histogram (continuous bins)

- Count within bins `geom_histogram(bins, binwidth)` (`stat_bin()`)

### Box plot (distribution comparison)

- `geom_boxplot()`

### Function curve

```r
ggplot(tibble(x = -4:4), aes(x = x)) + stat_function(fun = \(x) x^2, geom = "line")
ggplot(tibble(x = -4:4), aes(x = x)) + stat_function(fun = ~ .x^2, geom = "line")
```

### Area plot

- `geom_area(mapping, data)`
