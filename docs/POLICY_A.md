# Policy A: Standard Library and API Design

Nuru follows **Policy A** to decide where new functionality lives: **global builtins**, **modules**, or **methods**. This keeps the language consistent and the global namespace small.

---

## The three layers

### 1. Global builtins (small kernel)

**Use for:** Core language primitives that every script might need and that are not “owned” by a single type.

**Keep the set small.** Typical builtins:

| Category | Examples | Rationale |
|----------|----------|-----------|
| **I/O** | `andika`, `jaza` | Print and read input; universal. |
| **Type / meta** | `aina`, `badilisha`, `namba`, `tungo` | Type inspection and conversion; apply to any value. |
| **Sequence** | `mfululizo` | Range/sequence helper; kernel utility. |
| **Internal** | `_andika` | If needed for language or tooling; avoid adding more. |

**Do not add builtins for:**

- File operations → use the **faili** module.
- Encoding (e.g. Base64) → use the **crypto** (or encoding) module.
- Math beyond basic operators → use the **hisabati** module.
- Runtime/memory (GC, stats, weak refs) → use the **mfumo** module.
- Anything that is clearly “operations on a value” → use **methods** (see below).

---

### 2. Modules (domain and system)

**Use for:** Domain-specific or system-level features. Anything that is not a universal primitive and not naturally “one value’s behavior” belongs in a module.

**Examples:**

| Module | Role | Examples |
|--------|------|----------|
| **faili** | Files and filesystem | `fungua`, `soma`, `andika`, `futa`, `orodha`, `tengenezaSarafa` |
| **muda** | Time | `hasahivi`, `lala`, `tangu`, `leo`, `baada_ya` |
| **jsoni** | JSON | `dikodi`, `enkodi`, `soma`, `hifadhi` |
| **hisabati** | Math | `sqrt`, `random`, `max`, `min`, trig, constants |
| **crypto** | Cryptography and encoding | `md5`, `sha256`, `kodeBase64`, `katuaBase64`, `hex_encode` |
| **njia** | Paths | `unganisha`, `jina`, `sarafa`, `glob` |
| **mtandao** | HTTP client | `peruzi`, `tuma` |
| **http** | HTTP server and client | `tengenezaServer`, `pata` |
| **url** | URLs | `changanua`, `tengeneza`, `tatua` |
| **os** | Process / OS | `toka`, `kimbiza` |
| **mfumo** | Runtime / memory | `safishaMemori`, `takwimuMemori`, `kumbukumbaDhaifu`, `takwimuMemoriKwa` |

**Usage:** Scripts import with `tumia <moduli>` and call `moduli.kazi(...)`.

**When adding a new feature:** Prefer a new module (or an existing one) over a new global builtin.

---

### 3. Methods (operations on values)

**Use for:** Behavior that belongs to a specific type. The operation is “on” a value (query, transform, or mutate it).

**Examples:**

| Type | Methods (examples) | Rationale |
|------|--------------------|-----------|
| **String** | `idadi`, `herufikubwa`, `herufindogo`, `gawa`, `panga` | Length, case, split, format are string operations. |
| **Array** | `idadi`, `sukuma`, `yamwisho`, `unga`, `chuja`, `tafuta`, `ramani` | Length, push, last, join, filter, find, map are array operations. |
| **Dict** | `idadi`, `fungua` (get by key) | Size and lookup are dict operations. |
| **File** | `soma`, `andika`, `ongeza`, `funga`, `tafuta`, `hali` | Read, write, append, close, seek, stat are file-handle operations. |
| **Time** | time arithmetic and formatting | Operations on a time value. |

**Benefits:**

- Discoverable via `value.` (e.g. autocomplete).
- No global name pollution.
- Clear ownership: “this is what you do with this type.”

**When adding a new feature:** If it is “operate on this value,” add a **method** on the appropriate type rather than a global builtin or a generic module function that takes that value as first argument.

---

## Decision guide

Use this to decide where a new capability should go:

1. **Does it conceptually “belong to” a single type (e.g. string, array, dict, file)?**  
   → **Method** on that type.

2. **Is it a domain or system concern (files, time, crypto, OS, runtime)?**  
   → **Module** (existing or new).

3. **Is it a universal primitive (print, input, type, convert, range)?**  
   → **Builtin** only if it is truly kernel and used everywhere; otherwise prefer module or method.

4. **Would the global builtin list grow quickly if we kept adding this kind of thing?**  
   → Prefer **module** or **method**.

---

## Current builtin kernel (reference)

After Policy A migration, the **global builtins** are:

- `andika` – print
- `_andika` – print to string
- `jaza` – read input
- `aina` – type of value
- `mfululizo` – range (start, end, step)
- `badilisha` – type conversion
- `namba` – convert to integer
- `tungo` – convert to string

File open, Base64, and runtime/memory are in **faili**, **crypto**, and **mfumo** respectively. Big integers (namba kubwa) are created via **badilisha**(x, **"NAMBA_KUBWA"**) or **hisabati.namba_kubwa**.

---

## Summary

| Layer | When to use | Examples |
|-------|-------------|----------|
| **Builtins** | Small kernel: I/O, type, convert, range | `andika`, `jaza`, `aina`, `mfululizo`, `badilisha`, `namba`, `tungo` |
| **Modules** | Domain or system features | faili, muda, jsoni, hisabati, crypto, njia, mtandao, http, url, os, mfumo |
| **Methods** | Operations on a value | `orodha.idadi()`, `neno.gawa()`, `faili.soma()` |

Keeping this split consistent makes Nuru easier to learn and extend. When in doubt, prefer **method** (if it’s “on” a value) or **module** (if it’s a domain feature); add a **builtin** only when it clearly belongs in the kernel.
