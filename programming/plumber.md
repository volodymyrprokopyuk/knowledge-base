# Plumber web API

## Routing and input

- Multiple `@filter`s define a pipeline for handling incoming request
    - `req` and `res` handler parameters are based on R environments and honor
      reference semantics
    - `forward()` the request to another filter or final endpoint after mutating the
      request or producing other side effects
    - `return(list(...))` a response without forwarding the request to the next handler
- Single, final endpoint generates a response for a request
    - `@get | @post | @put | @delete | @head /path/<path_param[:type]>` designates an
      endpoint (single endpoint can support multiple HTTP verbs)
    - `@param qs_param` defines a query string parameter
    - `@assets host_path public_path list()` static file server
- Request
    - `pr("plumber.R")` Plumber router
    - `req$REQUEST_METHOD, PATH_INFO, QUERY_STRING`
    - `req$argsPath, argsQuery, argsBody, args`
    - `req$body, bodyRaw`
    - `req$HTTP_<HEADER>`
    - `req$cookies`
    - `req$REMOTE_ADDR, REMOTE_PORT`, `req$SERVER_NAME, SERVER_PORT`
- Subrouters, by default have their own environments, however an explicitly created
  shared environment can be passed to multiple subrouters

## Rendering and output

- Response
    - `res$headers, setHeader()` list of HTTP headers
    - `res$body` object to be serialized
    - `res$status` HTTP status code
    - `res$setCookie(), removeCookie()`
- `@serializer json list(auto-unbox = T)`
    - `json` -> `unbox(...)` -> scalar
    - `unboxedJSON` -> `I(...)` -> vector
    - Bypass serialization by returning `res` directly from the endpoint
    - `@serializer contentType list(type = "image/svg+xml")` does no serialization
