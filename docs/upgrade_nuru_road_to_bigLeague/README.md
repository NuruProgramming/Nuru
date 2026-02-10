# Nuru → "Big league" – upgrade roadmap

This document is the consolidated TODO to move Nuru from homework/educational league to "Python league" across **security**, **web framework support**, **computational performance**, and **async/concurrency**.

---

## Foundation / cross-cutting

| # | Task |
|---|------|
| F1 | Add sandbox/restrict mode (e.g. `--sandbox`) that disables or restricts dangerous modules (os, fs, net, http, mfumo) when running untrusted code. |
| F2 | Document production vs educational use and security assumptions in README/ABOUT. |
| F3 | Remove or gate DEBUG prints in `module/http.go` (behind verbose flag or delete for production). |

---

## Security

| # | Task |
|---|------|
| S1 | Harden or restrict `os.kimbiza` (run): sandbox disables it, or allowlist executables; avoid single user-controlled string as full command. |
| S2 | Path traversal protection: in fs, file, crypto modules, resolve paths with `filepath.Clean` + `filepath.Abs` and restrict to an allowed root; reject or sanitize `..` in restricted mode. |
| S3 | Import path validation: in `evaluator/import.go` `findFile`, reject module names containing `..`, `/`, or `\`; optionally restrict to explicit search roots. |
| S4 | Optional request/response size limits in http/net to mitigate DoS and memory abuse. |
| S5 | Add security note in docs (packages, examples): don't commit secrets; prefer env vars or secret store; don't log tokens. |
| S6 | Document `os.toka` (exit) and process control; in sandbox mode consider disabling or restricting it. |

---

## Web framework

| # | Task |
|---|------|
| W1 | Per-server mux: replace global mux in `module/http.go` with a mux stored on the server object from `undaServer`; register routes on that mux so multiple apps can coexist. |
| W2 | Path parameters: add router support for e.g. `/users/:id`; parse params and expose on request (e.g. `ombi["params"]` or `ombi.params["id"]`). |
| W3 | First-class query string: in `createRequestObject`, parse `r.URL.RawQuery` into a dict and set e.g. `ombi["hoja"]` or `ombi["query"]`. |
| W4 | Middleware pipeline: allow registering middleware (e.g. `(jibu, ombi, next)`) and run before route handler; support ordering (logging, auth, CORS). |
| W5 | Response helpers: e.g. `jibu.json(kamusi)`, `jibu.html(neno)`, `jibu.msimbo(404, ujumbe)` that set Content-Type and status. |
| W6 | Static file serving: expose e.g. `http.serveFaili(mpangilio, njiaYaSarafa)` using `http.FileServer`. |
| W7 | Server lifecycle: implement `seva.funga()` (keep `*http.Server` ref, call `Shutdown`); optionally `seva.sikilizaTLS(cert, key)` for HTTPS. |
| W8 | Cookie/session helpers (optional): read Cookie header, set Set-Cookie; optional simple session store. |

---

## Computational (speed + ML)

| # | Task |
|---|------|
| C1 | Typed float array: new type backed by `[]float64`; constructor/literal; indexing and iteration in evaluator. |
| C2 | Bulk numeric ops in Go: dot product, element-wise add/multiply, scalar multiply on typed float arrays (in hisabati or new module). |
| C3 | Matrix type + core ops: object backed by `[]float64` + shape; matmul, transpose (optionally Gonum); expose from Nuru. |
| C4 | Sigmoid/activations in Go: e.g. `hisabati.sigmoid(safu)` (and optionally relu, tanh) on typed float array. |
| C5 | hisabati integration with typed arrays where applicable. |
| C6 | Bytecode VM (optional): compile AST → bytecode, small VM loop for general speed. |
| C7 | Parallelism in native numeric layer: use Gonum/BLAS multi-core where applicable; document. |
| C8 | FFI or native module API: register Go (or C) functions with typed signatures so Nuru can call external libs. |
| C9 | Update perceptron (and similar) examples to use new matrix/vector and activation APIs. |

---

## Async & concurrency

| # | Task |
|---|------|
| A1 | Document: Nuru is single-threaded; no async/threads today (README or spec). |
| A2 | Parallel loop: e.g. `kwa sambamba i ktk orodha { ... }` or `mfumo.pitiaSambamba(orodha, fn)` using a worker pool; no shared mutable state. |
| A3 | Task pool + futures: e.g. `mfumo.tuma(unda(){ ... })` → future, `future.ngoje()` blocks for result (or timeout); implement via goroutines + channels or pool. |
| A4 | Non-blocking HTTP client (optional): client that doesn't block main thread; future or callback when done. |
| A5 | Async/await + event loop (longer-term): language + runtime support for async I/O. |
| A6 | Document concurrency safety: no shared mutable state between parallel tasks; prefer message passing. |

---

## Phased order

- **Phase 1 – Safety & clarity:** F1, F2, F3, S1, S2, S3, S5, S6, A1
- **Phase 2 – Framework:** W1, W3, W2, W5, W4, W6, W7, W8
- **Phase 3 – Computation:** C1, C2, C3, C4, C5, C9
- **Phase 4 – Scale & polish:** C6, C7, C8, S4, A2, A3, A6
- **Later:** A4, A5

---

## Checklist (copy-paste)

```
# Foundation
[ ] F1  Sandbox/restrict mode; disable or restrict os, fs, net, http, mfumo
[ ] F2  Document production vs educational use and security assumptions
[ ] F3  Remove or gate DEBUG prints in module/http.go

# Security
[ ] S1  Harden or restrict os.kimbiza (run)
[ ] S2  Path traversal protection (fs, file, crypto); restrict to root
[ ] S3  Import path validation (reject .. and path separators)
[ ] S4  Optional request/response size limits (http/net)
[ ] S5  Security note in docs (secrets, env, tokens)
[ ] S6  Document os.toka and sandbox behavior

# Web framework
[ ] W1  Per-server mux (no global mux)
[ ] W2  Path parameters (e.g. /users/:id) on request
[ ] W3  First-class query string on request object
[ ] W4  Middleware pipeline
[ ] W5  Response helpers (json, html, status)
[ ] W6  Static file serving
[ ] W7  Server shutdown (funga) and optional TLS
[ ] W8  Cookie/session helpers (optional)

# Computational
[ ] C1  Typed float array ([]float64) + indexing/iteration
[ ] C2  Bulk numeric ops (dot, element-wise) in Go
[ ] C3  Matrix type + matmul/transpose (optionally Gonum)
[ ] C4  Sigmoid/activations on typed arrays in Go
[ ] C5  hisabati integration with typed arrays
[ ] C6  Bytecode VM (optional)
[ ] C7  Parallelism in native numeric layer (Gonum/BLAS)
[ ] C8  FFI / native module API
[ ] C9  Update perceptron (and similar) to use new numeric APIs

# Async & concurrency
[ ] A1  Document: single-threaded; no async/threads today
[ ] A2  Parallel loop (e.g. kwa sambamba ... ktk ... or pitiaSambamba)
[ ] A3  Task pool + futures (e.g. mfumo.tuma → future.ngoje)
[ ] A4  Non-blocking HTTP client (optional)
[ ] A5  Async/await + event loop (longer-term)
[ ] A6  Document concurrency safety (no shared mutable state)
```
