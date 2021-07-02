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
    - JS `...<script src="script.js" [async | defer] [crossorigin="anonymous"]>
      </script></body>`
    - CSS `...<link rel="stylesheet" href="style.css" [media="print"]
      [crossorigin="anonymous"]></head>`
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

## Images

- External image `<img scr="image.png" alt="Description">`
- Embedded image `<img src="data:image/png;base64,..." alt="Description">`
- Responsive image = resolution switching per viewport size

  ```html
  <img sizes="viewport-condition image-size, ..."
       scrset="image-url image-size, ..."
       scr="fallback.png" alt="Description">`
  ```

- Alternative images = different images per viewport size

  ```html
  <picture>
    <source media="viewport-condition" srcset="large.png">
    <source media="viewport-condition" srcset="small.png">
    <img src="fallback.png" alt="Description">`
  </picture>
  ```

- Image map = image with clickable areas that usually act as hyperlinks

  ```html
  <img src="image.png" usemap="#image-map">
  <map name="image-map">
    <area shape="rectangle | circle | polygon" coords="x, y, ..." href="url">
    <area shape="rectangle | circle | polygon" coords="x, y, ..." href="url">
  </map>
  ```

## Input

- Label

  ```html
  <label for="n">Label</label>
  <input type="text" name="n">
  ```

- One-line edit box `<input type="text" name="n" id="a">`
    - Types `text | number | range | email | password | color | file | submit`
- File `<input type="file" name="n" id="a" accept="image/png" [multiple]>`
- Multi-line edit area `<textarea>`
- Independent heck box `<input type="checkbox" name="n" id="a" value="v" checked>`
- Mutually exclusive radio button group

  ```html
  <fieldset>
    <legend>Legend</legend>
    <p>
      <input type="radio" name="radio-group" id="a" value="v1" checked>`
      <label for="a">Label a>
    </p>
    <p>
      <input type="radio" name="radio-group" id="b" value="v2">`
      <label for="b">Label b>
    </p>
  </fieldset>
  ```

- Input validation = is done automatically by the browser on form submission based on
  input spacial attributes
    - Length / range `<input required minlength="m" maxlength="n" min="m" max="n">`
    - Pattern / file `<input pattern=".*" accept="image/png" title="Error">`
    - Disabled `<input disabled readonly>`
    - Focus / hint `<input autofocus placeholder="Example" autocomplete>`

## Document structure

- `section`, `header`, `footer`, `nav`, `article`, `aside`, `main`
