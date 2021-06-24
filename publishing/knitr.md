# Knitr dynamic documents using R

- Global R options
    - Set `options(option = value)`
    - Table `knitr.table.format`, `knitr.kable.NA`
- Global knitr chunk options
    - Set `opts_chunk$set(option = value)`
    - Evaluation `eval`, `include` all, `echo` source code, `results` text output,
      `message`, `warning`, `fig.show` plot, `error`, `collapse` source + output
    - Plot `dev`, `dev.args`, `fig.cap`, `fig.width`, `fig.height`, `fig.dim`
    - Code `tidy`, `tidy.opts`. TODO `formatR`, `styler`
    - HTML class `class.source`, `class.output`
- Code
    - Inline code `r ...`
    - Code block ```{r label, options ...}\newline ... \newline```
- Table
    - `kable(x, "pipe|html", caption, col.names, row.names, align, digits, format.args)`
    - TODO `gt`, `reactable`
