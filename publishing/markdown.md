# Markdown markup language

- `pandoc` options
    - HTML `pandoc -f markdown -t html5 -s --self-contained --mathml`
    - PDF `pandoc -f markdown -t pdf -s --pdf-engine wkhtmltopdf --mathml`
- Document metadata
    - `<head>` children

      ```yaml
      # -c style.css
      css:
        - style/style.css
      # -H head.html
      header-includes:
        - <link rel="shortcut icon" href="image/favicon.png"/>
      ```

- Header
    - `# Header{#id .class key="a value"}` delimits a section
- Paragraph
    - `*emphasis*`
    - `**strong emphasis**`
    - `***emphasis + strong emphasis***`
    - `~~strikethrough~~`
    - `^superscript^`
    - `~subscript~`
    - `| preserves formatting`
    - `---` horizontal rule
    - `-` hyphen for compound words
    - `--` en-dash for ranges
    - `---` em-dash substitutes coma, colon, seminolon or parentheses to set apart a
      phrase
    - `\*` escape control character
    - `\<newline>` hard line break
    - `\<space>` non-breaking space
    - `<!-- comment -->`
- Link
    - Web
        - `[Inline link](https://organization.org)`
        - `[Link reference]: https://organization.org` <- `[Link reference]`
        - Raw URI inline link `<https://organization.org>`
    - Email
        - `[Inline email](mailto:user@mail.com)`
        - `[Email reference]: user@mail.com` <- `[Email reference]`
        - Raw email inline link `<user@mail.com>`
    - Internal link
        - `<div id="internal-link"/>` <- `[Internal link](#internal-link)`
        - `<div id="internal-link"/>`, `[Internal link reference]: #internal-link`
          <- `[Internal link reference]`
        - `[Header link]`
    - Link attributes `[Inline link](https://organization.org){#id .class key="a value"}`
    - Reference `[reference-label]: URI` <- `[Reference text][reference-label]`
- Footnote
    - `^[Inline footnote]`
    - `[^footnote]: Footnote reference` <- `[^footnote]`
- Image
    - Inline image `![Image description](image.png)`
    - Image reference `[Image reference]: image.png` <- `![Image reference]`
- List
    - `- Bulleted item`
    - `4. Numbered item` starts from 4
    - `a. Lettered item`
    - Definition list

      ```md
      Term
      : Definition
      ```

- Table
    - Pipe table

      ```md
      Table: Caption

      Left | Center  | Right
      :--- | :---: | ---:
      A | B | C
      ```

- Code
    - `inline code`{.language}
    - ``inline ecode with a ` backtick``{.language}
    - ```{.language} code block``` = ```language code block```
- Quote
    - `> Block quote`
    - `> > Nested block quote`
- Math
    - `$inline tex math$`
    - `$$block tex math$$`
- Raw HTML
    - `:::{#id .class key="a value"}\newline Content \newline:::` =
      `<div id="id" class="class" key="a value">Content</div>`
    - `[Content]{#id .class key="a value"}` =
      `<span id="id" class="class" key="a value">Content</span>`
    - `<span>Raw inline code</span>{=html5}` only for html5 output
    - ```{=html5}\newline <div>Raw code block</div>\newline``` only for html5 output

# Inkscape vector graphics

# PlantUML sequence diagrams

# Dot graph visualization
