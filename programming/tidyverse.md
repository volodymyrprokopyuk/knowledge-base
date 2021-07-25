# tidyverse

- Data analysis lifecycle
    - Import `readr`
    - Tidy `tidyr`
    - while (!done)
        - Transform `dplyr`
        - Visualize `ggplot2` (knowledge generation)
        - Model `modelr` (knowledge generation)
    - Communicate `knitr`
- Data analysis
    - Hypothesis generation (visualization + modeling, each observation used multiple
      times)
    - Hypothesis confirmation (each observation is only used once)

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

- Plot template = the grammar of graphics = formal system for building plots as a set of
  layers
    - 1. `data` set of data to plot
    - 1. `geom` representation of insignts from the dataset
    - 1. `aes`thetic mapping of data values to graphic properties
    - 1. `stat` transformation of the dataset
    - 1. `position` adjustment within the coordinate system
    - 1. `coord` system to represent data values
    - 1. `facet` shceme to partition the dataset into subplots

  ```r
  ggplot(<data>) +
    <geom_function>(aes(<mapping>), <stat>, <position>) +
    <coord_function> +
    <facet_function> +
    ... <layer 2>
  ```

- Plot `ggplot(data, mapping = aes())` maps data properties to graphic properties
- Data (tibble data frame, numerical + categorical) = each row represents one
  observation of potentially multiple variables (stored in columns)
- Geoms (points, lines, bars, boxes)
- Mapping
    - Aesthetic mapping `aes(x, y, color, fill, size, shape, alpha, linetype, group)`
      associates data values with graphic properties (levels)
    - Aesthetic setting `geom_point(color)` without any relation to data values
- Scales
    - Palette `scale_fill_brewer(palette)`, `scale_fill_manual(values)`,
    `scale_color_manual(values)`
    - Percentage scale `scale_x|y_continuous(limits, labels)`
- Guides
    - Labels `labs(title, subtitle, caption, x, y)`, `xlab(label)`, `ylab(label)`
    - Limits `lims(x, y)`, `xlim(lower, upper)`, `ylim(lower, upper)`
- Coordinates
    - `coord_flip()` swaps x with y axes
    - `coord_polar()` turns a bar plot into pie plot
    - `coord_quickmap()`, `coord_map()` sets correct aspect ratio for maps
- Theme `theme(axis, legend, panel, plot, aspect.ratio)`

### Scatter plot (raw data)

- `geom_point(data, mapping, position = "identity | jitter", show.legend)`
    - `position = identity` values are rounded so the points appear on a grid
    - `position = jitter` values have small amount of random nose to avoid overplotting
- `geom_point(position = "jitter")` = `geom_jitter()`

### Line plot (raw data)

- `geom_line(data, mapping)`
- `geom_smooth(data, mapping)` `stat_smooth()` fits a model to data, plots predictions
  from the model

### Bar plot (stat counts or values of discrete categories)

- `stat` = statistical transformation = algorithm to caclulate new values from the
  dataset. `stat` returns special variables `..count..`, `..prop..`, `..density..`
- `stat` and `geom` usually can be used interchangeably as every `geom` has a default
  `stat` and viceversa
- `aes(y = ..count..)` = `aes(stat(count))`
- Count `geom_bar(data, mapping, stat = "count")` = `stat_count(geom = "bar")`
- Value `geom_col(fill, color, position = "dodge | stack | fill", width)` =
  `stat_idenitty(geom = "col")` >= `geom_bar(stat = "identity")`
- Labels `geom_text(aes(label, x, y), vjust, hjust)`
- Cleveland dot plot (cleaner alternative to a bar plot) `geom_point()`

### Histogram (stat counts of bins)

- Count within bins `geom_histogram(bins, binwidth, stat = "bin")` =
  `stat_bin(geom = "histogram")`

### Box plot (stat distribution summary)

- `geom_boxplot(data, mapping)`
- `geom_pointrange(data, mapping)` = `stat_summary(data, mapping, fun.min, fun.max, fun)`

### Function curve

```r
ggplot(tibble(x = -4:4), aes(x = x)) + stat_function(fun = \(x) x^2, geom = "line")
ggplot(tibble(x = -4:4), aes(x = x)) + stat_function(fun = ~ .x^2, geom = "line")
```
### Facet plot

- Partition a plot of categorical variables into subplots
    - One variable `facet_wrap(~ categorical, nrow, ncol)`
    - Two variables `facet_grid(cat1 ~ cat2)`

### Area plot

- `geom_area(mapping, data)`
