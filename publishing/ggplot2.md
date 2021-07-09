# ggplot2 grammar for graphics

## General aspects

- Plot `ggplot(data, mapping = aes())`
- Mapping `aes(x, y, color, fill)`
- Labels `labs(title, subtitle, caption, x, y)`, `xlab(label)`, `ylab(label)`
- Theme `theme(axis, legend, panel, plot)`
- Palette `scale_color_manual(values)`, `scale_fill_manual(values)`

## Scatter plot

- `geom_point(mapping, data)`

## Line plot

- `geom_line(mapping, data)`

## Bar plot

- Count `geom_bar()` <- `stat_count()`
- Value `geom_col()` = `geom_bar(stat = "identity")` <- `stat_identity()`

## Area plot

- `geom_area(mapping, data)`
