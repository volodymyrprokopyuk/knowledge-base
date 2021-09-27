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
- Exploratory data analysis = develop an understaning of the data
    1. Ask questions about the data that guide the investigation and discovery
       - What type of variation occurse within each variable? Bar plot + histogram, most
         common + rare values, clusters + distributions within clusters
       - What type of covariation occurs between variables?
    1. Search answers by visualizing, transforming and modeling the data
    1. Use discovered insights to refine existing questions or ask new questions



## lubridate + hms

- lubridate always uses UTC that does not have DST (equivalent to its predecessor GMT)
- Date
    - Current date `today(tzone)` = `Sys.Date()`
    - Parse `ymd(chr|int, tz)`
    - Create `make_date(components ...)`
    - Cast `as_date(x)`
    - Accessor / mutator `year(x)`, `yday(x)`, `month(x, label, abbr)`, `mday(x)`,
      `wday(x, label, abbr)`
    - Copy + update `update(x, components ...)`
- Date-time
    - Current date-time `now(tzone)` = `Sys.time()`
    - Parse `ymd_hms(chr|int, tz)`
    - Create `make_datetime(components ...)`
    - Cast `as_datetime(x)`
    - Accessor / mutator `hour(x)`, `minute(x)`, `second(x)`, `tz(x)`
- Time span
    - Duration = always in seconds, takes into account TZ and DST, precise results
        - Creation `as.duration(difftime)`, `dseconds(x)`, `dminites(x)`, `dhours(x)`,
          `ddays(x)`, `dweeks(x)`, `dyears(x)`
    - Period = human intuitive units without a fixed length in second and TZ, DST
      surprises, intuitive results
        - Creation `as.period(difftime)`, `seconds(x)`, `minites(x)`, `hours(x)`,
          `days(x)`, `weeks(x)`, `months(x)`, `years(x)`
    - Interval = specific date of starting point + duration to figure out how long a
      time span is in human units
        - Creation `interval(start, end, tzone)`, `start %--% end`
- Time zone
    - `with_tz(x, tzone)` displays a UTC in a specified time zone
    - `force_tz(x, tzone)` changes the underlaying time zone
- Operations
    - Rounding `floot_date(x)`, `round_date(x)`, `ceiling_date(x)`
    - Date subtraction `date - date = difftime`



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



## dplyr data manipulation

- Verb families
    - Mutating join `inner|left|right|full_join(x, y, by)` = adds new variables to `x`
      from the matching observations in `y` (duplicates observations from `x` or `y` if
      keys are not unique)
    - Filtering join `semy|anti_join(x, y, by)` = keeps | drops all observation in `x`
      that have a match in `y` (never duplicates observations from `x`)
    - Set operations `insersect|union|setdiff(x, y)`
- The language of data manipulation
    - `filter(data, == != < <= > >= ! & | xor() %in% between())` restricts rows
    - `arrange(data, cols ... desc())` sorts rows
    - `select(data, cols ... start:stop c(...) -col starts|ends_with() contains()
      matches() everything())` restricts columns or changes order of columns
        - `rename(data, new_col = old_col)` renames columns
    - `mutate(data, new_col = ...)` computes new columns
        - `transmute(data, col = ...)` keeps only new columns
    - `summarize(data, new_col = ... n() n_distinct())` collapses a data frame into a
      single row
        - `group_by(data, cols ...)` delimits the unit of analysis with `summarize()`
          from a whole data set to individual groups
        - When grouping by multiple variables, each summary peels of one level of the
          grouping (progressively roll up a dataset)
        - `ungroup` removes grouping
    - `top_n(n)` limits to `n` observations
    - `count(x, ...)` = `group_by(x, ...) + summarize(n = n())`
    - `slice(x, ...)` subset rows by position



## ggplot2 grammar for graphics

### General aspects

- The layered grammar for graphics (multi-layer plot template) = formal system embedded
  into R for building plots as a set of layers by declaratively describing graphic
  components
    1. `data` = rectangular dataset of observations (rows) of variables (columns)
    1. `geom` graphical representation (points, lines, bars) of data values
    1. `aes`thetic mapping of data values to graphic properties (position, color, size,
       shape, alpha)
    1. `stat` statistical transformation of data (bin, boxplot, contour, density,
       identity, jitter, smooth, summary, unique)
    1. `position` plot adjustment within the coordinate system
    1. `coord` drawing system to represent data values
    1. `scales` = converts data units to aesthetic units
    1. `guides` = axes (positional scales), legends (continuous: size, color, alpha;
       categorical: shape, color), title and annotations
    1. `facet` = data partitioning for subplots visualization

  ```r
  ggplot(<default_data>, <default_mapping>) +
    <geom_function>(aes(<data>, <mapping>), <stat>, <position>) +
    <coord_function> +
    <facet_function> +
    ... <layer 2>
  ```

- `ggplot` is composed by
    - A default `data`set + set of `mapping`s from variables to `aes`thetics
    - One or more layers using a default dataset or a different dataset with each layer
      having
        - One `geom`etric object (with default `stat`istical transformation) that
          displays aesthetics (position, color, shape, size)
        - One `stat`istical transformation (with default `geom`) adds new variables to
          the original `data`set
        - One `position` adjustment of the plot in the `coord`inate system
        - Optionally a different `data`set and set of `aes`thetic `mapping`s
    - One `scale` (axes, legends, title, annotations) for each `aes`thetic `mapping`
      (scales are common across layers)
    - One `coord`inate system
    - One `facet` specification
- Hierarchy of defaults
    - `ggplot(data, mapping = aes())` maps data properties to graphic properties
    - `geom(mapping, data)` specified in a layer overrides the default from
      `ggplot(data, mapping)`
    - Specify either `geom` or `stat` as each `geom` has a default `stat` and vice versa
    - `coord_cartesian()` is the default
    - Default scales are added for continuous and categorical variables
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
    - Limits `lims(x, y)`, `xlim(lower, upper)`, `ylim(lower, upper)`,
      `expand_limits()`
- Coordinates
    - `coord_cartesian()` is the default
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

### Bar plot (stat counts or values of categorical variable)

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

### Histogram (stat counts of bins of continuous variable)

- Count within bins + bars `geom_histogram(bins, binwidth, breaks, stat = "bin")` =
  `stat_bin(geom = "histogram")`
- Count within bins + lines `geom_freqpoly(bins, binwidth, breaks, stat = "bin")`

### Box and whiskers plot (stat distribution summary)

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

### Pie chart, bullseye plot, coxcomb plot



## jsonlite consistent, bidirectional mapping between JSON and R

- `toJSON()`, `fromJSON()` class-based method dispatch (can be extended for user defined
  classes)
- Mapping between JSON and R
    - Atomic vector -> [ always primitives ]
    - Matrix / array -> [ array of always primitives ]
    - Data frame / tibble (columns-first) -> [ object of always primitives ] (row-first)
    - Unnamed list -> [ object | array ] (never primitives) ordered sequence of
      heterogeneous values
    - Named list -> { object | array } (never primitives) unorderd collection of
      key-value pairs
