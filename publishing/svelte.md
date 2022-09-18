# Svelte

## Styling

- CSS is unpredictable as everything in CSS is global
- `style` is scoped to the component (through unique classes generated during
  compilation). Use simple selectors without conflicts between components
- Global CSS through an external file vs `:global { ... }` to opt out from style
  scoping from within a Svelte componen
- JS variables cannot be used in the `<style>` block. Use CSS variable (CSS
  custom properties instread)
- Style directives (for dynamic styling)
    - `<div style:css-property={value}></div>`
    - `<div style:--css-custom-property={value}></div>`
- Class directives (for dynamic styling)
    - `<div class:class-name={variable}></div>` (boolean variable)
    - `<div class:var></div>` (variable name = class name)
    - `<div class={variable}></div>` (class name from variable)
    - Export `class` property
        ```svelte
        // Child
        <script>
          let clName
          export { clName as class }
        </script>
        <div class={clName}></div>
        // Parent
        <Component class="class-name"/>
        ```
- Problem: class on the Component does not work (class is exclusive to elements)
    ```svelte
    <Component class="cmp"/>
    <style>
      .cmp { color: blue; }
    </style >
    ```
    - Solution 1. Wrapping `div` (attach CSS to the wrapping `div`)
        ```svelte
        // Parent
        <div class="cmp">
          <Component/>
        </div>
        <style>
          .cmp { color: blue; }
        </style >
        ```
    - Solution 2. CSS variables / custom properites (only exposed variables)
        ```svelte
        // Parent
        <Component --color="blue">
        // Child
        <div class="cmp">Component</div>
        <style>
          .cmp { color: var(--color, default); }
        </style>
        ```
    - Solution 3. Style directives (only exposed variables) does not work on
      Component (exclusive to elements)
        ```svelte
        // Parent
        <Component color="blue"/>
        // Child
        <script>
          export let color
        </script>
        <div style:color>Component</div>
        ```
