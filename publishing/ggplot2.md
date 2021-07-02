# ggplot2 grammar for graphics

## General aspects

- Plot `ggplot(data, mapping = aes())`
- Mapping `aes(x, y, color, fill)`
- Labels `labs(title, subtitle, caption, x, y)`, `xlab(label)`, `ylab(label)`
- Theme `theme(axis, legend, panel, plot)`
- Palette `scale_color_manual(values)`, `scale_fill_manual(values)`

## Line plot

- `geom_line(mapping, data)`

## Area plot

- `geom_area(mapping, data)`
