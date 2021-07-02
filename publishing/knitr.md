# Knitr dynamic documents using R

- Global R options
    - Set `options(option = value)`
    - Table `knitr.table.format`, `knitr.kable.NA`
- Global knitr chunk options
    - Set `opts_chunk$set(option = value)`
    - Evaluation
        - Source `eval`, `error`, `collapse` source + output
        - Output `include` all, `echo` source code, `results` text output, `message`,
          `warning`
    - Plot
        - Device `dev`, `dev.args`
        - Figure `fig.cap`, `fig.width`, `fig.height`, `fig.dim`, `fig.show`,
          `fig.process`
    - Code
        - Format `tidy`, `tidy.opts`
    - Output
        - Raw markdwon `results = 'asis'`
        - Uncommented `comment = ""`
        - Prompt `prompt = T` prefix source with `>` and `+`
    - HTML attribute`[class|attr].[source|output|message|warning|error]`
- Code
    - Inline code `r ...`
    - Code block ```{r label, options ...}\newline ... \newline```
- Table
    - `kable(x, "pipe|html", caption, col.names, row.names, align, digits, format.args)`
    - TODO `gt`, `reactable`
