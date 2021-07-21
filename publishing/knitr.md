# knitr dynamic HTML documents with R

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

# gt declarative HTML tables with R

- Table `tibble |> gt(rowname_col, groupname_col)`
- Header `tab_header(title, subtitle)`
- Body
    - Stubhead
        - Spanner column label `tab_spanner(label, columns)`
        - Column label `cols_label(name = label)`
        - Stubhead label `tab_stubhead(label)`
    - Stub
        - Row group label `tab_row_group(label, rows)`, `gt(groupname_col)`,
          `row_group_order(groups)`
        - Row label `gt(rowname_col)` is a row header, not a data column
        - Summary label `summary_rows(groups, columns, fns = list(...))`,
          `grand_summary_rows(columns, fns = list(...))`
        - Cell
- Footer
    - Footnotes `tab_footnote(footnote, locations = cells_body(columns, rows))` attached
      to cells data
    - Source notes `tab_source_note(source_note)` related to the whole table
- Markup `html(str)`, `md(str)`
- Export `as_raw_html(gt, inline_css)`
