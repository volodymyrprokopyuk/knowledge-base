# HTML hypertext markup language

## General aspects

- HTML content structure (not presentation CSS) as a hierarchy of elements / tags that
  provide semantics for the content flow
- CSS style separates modular content presentation and layout from content structure and
  semantics (HTML markup)
- `#id` unique, `.class` classification, `attribute="value"` configuration provide
  further semantic context for HTML elements to apply cross-element CSS styling via
  `.class`, attach id-specific JS actions and interlink document sections via `#id`
- `<!-- HTML comment -->`
- Minimal HTML

  ```html
  <!DOCTYPE html> <!-- HTML5 -->
  <html lang="en">
    <head>
      <meta charset="utf-8">
      <title>Minimal HTML</title>
    </head>
    <body>
      <h1>Title</h1>
      <p>Content</p>
    </body>
  </html>
  ```
- Linking external resources
    - JS `...<script src="script.js" [async | defer]></script></body>`
    - CSS `...<link rel="stylesheet" href="style.css" [media="print"]></head>`
    - Import CSS `@import "style.css";`, `@media print { @import "style.css"; }`
    - Favicon `...<link rel="shortcut icon" href="favicon.png"></head>`

## Headings and paragraphs

- Headings `<h1>` title, `<h2>--<h6>` sections / subsections
- Paragraph `<p>`, `<br>` line break, `<pre>` preserves formatting
- Formatting
    - `<em>` ittalic, `<strong>` bold, `<mark>` highlight
    - `<u>` underline, `<s>` strikethrough
    - `<ins>` inserted text (underline), `<del>` deleted text (strikethrough)
    - `<sub>` subscript, `<sup>` superscript
- Abbreviation `<abbr title="Hypertext Markup Language">HTML</abbr>`

## Hyperlinks

- External link `<a href="https://absolute-url | /relative-url" rel="external">External
  site</a>`
- Open in new tab `<a href="url" target="_blank">Link</a>`
- Internal anchor `<a href="[url]#anchor-id">Document section</a>`
- Download link `<a href="https: | blob: | data:" downlaod="document-name">Downlaod
  document</a>`
- Email `<a href="mailto:user@mail.com">Email</a>`

## Lists

- Ordered list `<ol type="1 | a | A" start="2"> > <li value="4">`
- Unordered list `<ul> > <li>`
- Description list `<dl> > <dt> + <dd>`

## Tables

- Basic `<table> > <tr> > <th | td colspan="2" rowspan="2">`
- Extended `<table> > <caption> + <htead> + <tbody> + <tfoot>`
- Column grouping `<table> [<caption>] > <colgroup> [> <col span="2">] + <tr>`
- Row grouping `<tr class="row-group">`
- Heading scope `<th scope="column|row|colgroup|rowgroup">`

## Document structure

- `section`, `header`, `footer`, `nav`, `article`, `aside`, `main`
