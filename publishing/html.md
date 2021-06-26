# HTML hypertext markup language

## General aspects

- HTML content structure (not presentation CSS) as a hierarchy of elements / tags that
  provide semantic meanging for the content flow
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
