# Cascading Style Sheets (CSS)

## General aspects

- CSS inclusion
    - `<div style="inline: style;"/>`
    - `<style>sel { style: block; }</style>`
    - `<link rel="stylesheet" href="style.css"/>`
    - `@import "style.css";`
- CSS variables
    - Variable inheritance: variables cascade down to descendant elements
    - `:root { --global-variable: value; }` <- `var(--global-variable, [default])`
    - Counter `:root { counter-reset: h1c; }`, `h1:before { counter-increment: h1c;
        content: counter(h1c) ". "; }`
- CSS units
    - `px` usually used only for root `font-size`
    - `em` variable, relative to inherited `font-size`
    - `rem` constant, relative to root `font-size`
    - `%` percentage of inherited `font-size` or `width` (responsive design)
    - `wv`, `wh` relative to viewport (responsive design)
    - `fr` CSS grid fraction
    - `cm`, `mm`, `pt` absolute units (print styling)
    - `calc(...)` unit calculation with CSS variables
- CSS colors
    - `#1a2b3c`
    - `rgb(255, 255, 255)`, `rgb(255 255 255)`
    - `rgba(255, 255, 255, 0.5)`, `rgba(255 255 255 / 0.5)`

## CSS selectors

- Cascade = CSS rules importance by source
    - 1. User `important!`
    - 2. Author `important!`
    - 3. Author CSS
    - 4. User CSS
    - 5. Browser CSS
    - Overwrite cascade and specificity `property: value !important;` (prefer more
      specific rules)
- Specificity = CSS rules importance by selector (same specificity last rule wins =
  order of CSS files and CSS selectors matter)
    1. Inline `<div style="inline: style;"/>`
    2. Id `#id`
    3. Class `.class`, `[attribute]`, `:pseudo-class`
    4. Element `element`, `:pseudo-element`
    - Set general style for common context-free elements, then override style for more
      specific elements. Avoid tying CSS to `#id` and document-spacific context, use
      context-free `.classes` instead
- Common selectors `*` universal selector, `element`, `.class`, `#id`
    - Attribute selectors `[attribute]`
    - Whole word `[attribute="exact"]`, `[attribute~="space-separated"]`
    - Substring `[attribute*="substring"]`, `[attribute^="start"]`, `[attribute$="end"]`
    - Compound selector (AND) `element#id.class[attribute="exact"]`
    - Independent selectors (OR) `element, #id, .class, [attribute="exact"]`
- CSS combinators
    - [In]direct descendant combinator `element descendant`
    - Direct child combinator `element > child`
    - General sibling combinator `element ~ sibling`
    - Adjacent sibling combinator `element + sibling`
- Pseudo-class
    - UI state pseudo-class (state-based implicit classification)
        - Link state `:link` unvisited, `:visited`,`:hover`, `:focus`, `:active` (order
          matters: lord vader hates furry animals)
        - Form validation `:required`, `:optional`, `:[in]valid`
    - Document structure pseudo-class (position-based implicit classification)
        - Root `:root` = `html`
        - First / last child `:first-child`, `:last-chaild`
        - Children `:nth-child(odd | even)`, `:nth-child(n)`, `:nth-child(3+2n)`,
          `nth-last-child(3-n)`
    - Negate selector `:not(...)`
    - Target selector `:target` element `#id` contained in the URL hash `/content#id`
- Pseudo-element = part of an element
    - Content `::before`, `::after` + `content`
    - Letter / line `::first-letter`, `::first-line`

## Box model

- Box model `margin`, `border`, `padding` and `content`
- `box-sizing: content-box | border-box` -> `width`, `height` (`margin` is never
    considered) + `[min|max]-[width|height]`
- `display: block | inline | inline-block`
- `block` respects `width` and `height`, is palced on its own line, takes up the
    full width of the container and has just enough height to fit the content
- `inline` ignores `width` and `height`, is palced inline in the text flow, takes up
    enough width and height to fit the content
- `inline-block` respects `width` and `height` and is palced inline in the text flow
- `display: none` removes element from the flow
- `visibility: hidden` preserves element space in the flow

## Borders and backgrounds

- Border `border`, `border-width`, `border-style`, `border-color`, `border-radius`,
  `box-shadow`, `opacity`
- Background `background`, `background-color`, `background-image`, `background-repeat`,
  `background-position`, `background-size`, `background-clip`
- Gradient `background-image: linear-gradient() | radial-gradient()`

## CSS typography

- Text styling `font-family`, `font-size`, `color`, `font-weight`, `font-style`,
  `text-decoration`, `text-transform`, `letter-spacing`, `font-variant`, `text-shadow`
- Text layout `white-space`, `line-height`, `text-indent`, `text-overflow`,
  `text-align`, `vertical-align`
- Web font
    - Import `@font-face { font-family: "FF"; src: url("ff.woff2") format("woff2");
      font-weight: ...; font-style: ...; }`
    - Use `body { font-family: "FF", sans; }`

## Positioning and floats

- CSS overflow `overflow-[x|y]: visible | hiddent | scroll | auto`
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
- `position: sticky` element is `relative` when scrolling up to a `top`, `left`,
    `right`, `botton` point, after which element becomes `fixed`
- `z-index` stacking
- `float: left | right` element is removed from the flow, element flows to the left
    or to the right of the parent, other inline elements flow around the floated
    element + `clear: left | right | both` to resume content below the floated element

## Flexbox layout

- Display `display: flex | inline-flex` flex container parent + flex items children,
  flex containers can be nested forming flexbox layout hierarchy with absolutely
  centered flex items
- Direction `flex-direction: [row|column][-reverse]` one-dimensional either horizontal
  (row) or vertical (column) layout, block flex items are laid out one after another
  along the main axis,
- Sizing
    - Initial flex item size `flext-basis`
    - Grow `flex-grow: 1` flex items grow proportionally to fit the container
    - Shrink `flext-shrink: 1` flex items shrink proportionally to fit the container,
      otherwise flex items overflow, unless container `flex-wrap: wrap`
- Alignment
    - Main axis `justify-content: flex-[start|end] | center |
      spece-[between|around|evenly]`
    - Cross axis `align-items: stretch | flex-[start|end] | center | baseline`
      (`align-self` item-level override)
    - Container cross axis `align-content: stretch | flex-[start|end] | center |
      space-[between|around|evenly]`

## CSS grid layout

- `display: grid | inline-grid` grid container + grid items (container's immediate
  children) laid out in two-dimensional grid layout with rows and columns
- Sizing / naming
    - Row `grid-template-rows: [row-name] repeat(n | auto-fill | auto-fit, minmax(n, m |
      auto))`, `grid-auto-rows`-
    - Column `grid-template-columns: [column-name] ...`
    - Area `grid-template-areas`
    - Gap `gap`
- Positioning
    - Explicit positioning `grid-[row|column]: n | name`
    - Row / column spanning `grid-[row|column]-[start|end]`, `grid-[row|column]: n |
       name / m | name [span n]`
- Alignment
    - Row axis `justify-items: stretch | start | end | center` (`justify-self`
      item-level override)
    - Column axis `align-items: stretch | start | end | center` (`align-self` item-level
      override)
    - Container row axis `justify-content: stretch | start | end | center |
      space-[between|around|evenly]`
    - Container column axis `align-content: stretch | start | end | center |
      space-[between|around|evenly]`

## Responsive design

- Responsive design = use `@media` queries with breakpoints to apply different
  stylesheets or CSS rules depending on the viewport size
- `<meta name="viewport" content="width=device-width, initial-scale=1.0,
  user-scalable=yes">`
- `<link rel="stylesheet" media="print" href="print.css">`
- `@media screen | print | all [and | or | not ([min|max]-width: breakpoint)] {...}`
- Responsive layout with flexbox `flex-direction: row | column` + `flex-wrap: wrap`
- Fluid typography = `font-size: clamp(min, preferred, max)` automatically scales font
  size depending on the viewport size without `@media` queries
- Responsive images `img { max-width: 100%; height: auto; }`

## TODO

- `transform: perspective | rotate | translate | scale | skew`
- `transition` between two states + timing
- `animation` + `@keyframes` more than two states + timing
