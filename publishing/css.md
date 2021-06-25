# CSS

- CSS inclusion
    - `<div style="inline: style;"/>`
    - `<style>sel { style: block; }</style>`
    - `<link rel="stylesheet" href="style.css"/>`
    - `@import "style.css";`
- CSS selectors
    - Specificity (form higher to lower, same specificity last rule wins)
        1. Inline `<div style="inline: style;"/>`
        2. Id `#id`
        3. Class `.class [attribute] :pseudo-class`
        4. Element `element ::pseudo-element`
    - Overwrite specificity `property: value !important;` (prefer more specific rules)
    - Universal selector `*`, `element`, `.class`, `#id`, `[attribute]`,
      `[attribute="exact"]`, `[attribute~="whitespace"]`, `[attribute*="substring"]`,
      `[attribute^="start"]`, `[attribute$="end"]`
    - Compound selector `element.class#id[attribute="exact"]`
    - Independent selectors `element, .class, #id, [attribute="exact"]`
    - [In]direct descendant combinator `element descendant`
    - Direct child combinator `element > child`
    - General sibling combinator `element ~ sibling`
    - Adjacent sibling combinator `element + sibling`
    - UI state pseudo-class `:active`, `:checked`, `:focus`, `:hover`, `:[in]valid`
    - Doc structure pseudo-class `:first-child`, `:last-chaild`, `:nth-child(n)`,
      `:nth-child(2n)`, `:nth-child(odd|even)`, `:root` = `html`
    - Negate selector `:not(...)`
    - Pseudo-element `::first-line`, `::first-letter`, `::before`, `::after` + `content`
- Box model `margin`, `border`, `padding` and `content`
    - `box-sizing: content-box | border-box;` -> `width`, `height` (`margin` is never
      considered) + `[min|max]-[width|height]`
    - `display: block | inline | inline-block;`
    - `block` respects `width` and `height`, is palced on its own line, takes up the
      full width of the container and has just enough height to fit the content
    - `inline` ignores `width` and `height`, is palced inline in the text flow, takes up
      enough width and height to fit the content
    - `inline-block` respects `width` and `height` and is palced inline in the text flow
    - `display: none;` removes element from the flow
    - `visibility: hidden` preserves element space in the flow
- CSS units
    - `px` usually used only for root `font-size`
    - `em` variable, relative to inherited `font-size`
    - `rem` constant, relative to root `font-size`
    - `wv`, `wh` relative to viewport (responsive design)
    - `%` percentage of inherited `font-size` or `width` (responsive design)
    - `cm`, `mm`, `pt` absolute units (print styling)
    - `calc(...)` unit calculation with CSS variables
- CSS colors
    - `#1a2b3c`
    - `rgb(255, 255, 255)`, `rgb(255 255 255)`
    - `rgba(255, 255, 255, 0.5)`, `rgba(255 255 255 / 0.5)`
- CSS overflow `overflow-[x|y]: visible | hiddent | scroll | auto;`
- CSS variables
    - Variable inheritance: variables cascade down to descendant elements
    - `:root { --global-variable: value; }` <- `var(--global-variable, [default])`
    - Counter `:root { counter-reset: h1c; }`, `h1:before { counter-increment: h1c;
       content: counter(h1c) ". "; }`
- Border `border`, `border-width`, `border-style`, `border-color`, `border-radius`,
  `box-shadow`, `opacity`
- Background `background`, `background-color`, `background-image`, `background-repeat`,
  `background-position`, `background-size`, `background-clip`
- Gradient `background-image: linear-gradient() | radial-gradient()`
- Text styling `font-family`, `font-size`, `color`, `font-weight`, `font-style`,
  `text-decoration`, `text-transform`, `letter-spacing`, `font-variant`, `text-shadow`
- Text layout `white-space`, `line-height`, `text-indent`, `text-overflow`,
  `text-align`, `vertical-align`
- Web font
    - Import `@font-face { font-family: "FF"; src: url("ff.woff2") format("woff2");
      font-weight: ...; font-style: ...; }`
    - Use `body { font-family: "FF", sans; }`
- Positioning
    - Vertical margin collaps with the larger margin
    - `maring: auto` center horizontally
    - `position: static` (defualt) element in the flow, elements takes up the full width
      of its parent
    - `position: relative` element's original position remains in the flow, other
      elements are not affected + `top`, `left`, `right`, `bottom` relative to element's
      original position, elements takes up the full width of its parent
    - `position: absolute` element is removed from the flow and floats abouve the flow,
      other elements take spece of the original element + `top`, `left`, `right`,
      `bottom` relative to the closest explicitly positioned ancestor element, element
      takes only the necessary width
    - `position: fixed` element is removed from the flow + `top`, `left`, `right`,
      `bottom` always relative to the viewport (even on scroll, element remains in the
      same position), element takes only the necessary width
    - `position.sticky` element is `relative` when scrolling up to a `top`, `left`,
      `right`, `botton` point, after which element becomes `fixed`

- Boxes, shadows and opacity
- Backgrounds and gradients
- Web fonts and typography
- Layout positioning and stacking
- CSS transforms, transitions and animations
- Flexbox
- Responsive design, media queries and fluid typography
- CSS grid
