# Policy: Standard Library and API Design - v2.1

Nuru follows a **Policy** to decide where new functionality lives: **global builtins**, **modules**, or **methods**. This keeps the language consistent and the global namespace small.

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
| **Collections** | `seti`, `jozi` | Create set or immutable tuple from args or array. |
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
| **muda** | Time | `hasahivi`, `lala`, `tangu`, `leo`, `baada_ya`, `siku` |
| **re** | Regex | `linganisha`, `tafuta`, `tafutaZote`, `vikundi`, `badilisha`, `gawa`, `tayari` |
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
| **String** | `idadi`, `herufikubwa`, `herufindogo`, `gawa`, `panga`, `ondoaNafasi`, `anzaNa`, `ishiaNa`, `ina`, `badilishaNeno` | Length, case, split, format, trim, prefix/suffix/contains, replace. |
| **Array** | `idadi`, `sukuma`, `yamwisho`, `unga`, `chuja`, `tafuta`, `ramani`, `geuza`, `panga`, `gawa` | Length, push, last, join, filter, find, map, reverse, sort, chunk. |
| **Dict** | `idadi`, `fungua`, `funguo`, `maana`, `vikundi` | Size, lookup, keys, values, key-value pairs. |
| **Set** | `idadi`, `ona`, `ongeza`, `ondoa`, `kitanzi` | Size, membership, add, remove, iterate. |
| **Tuple** | `idadi`, `kitanzi` | Immutable sequence; read-only index, length, iterate. |
| **File** | `soma`, `andika`, `ongeza`, `funga`, `tafuta`, `hali`, `isFungwa` | Read, write, append, close, seek, stat. |
| **Time** | `ongeza`, `tangu`, `panga` | Add time (named args), since, format. |
| **Date** | `panga` | Date-only; format by layout. |
| **Compiled regex** | `linganisha`, `tafuta`, `tafutaZote`, `vikundi`, `badilisha`, `gawa` | Match/find/replace/split with precompiled pattern (from `re.tayari`). |

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


## Summary

| Layer | When to use | Examples |
|-------|-------------|----------|
| **Builtins** | Small kernel: I/O, type, convert, range, collections | `andika`, `jaza`, `aina`, `mfululizo`, `badilisha`, `namba`, `tungo`, `seta`, `jozi` |
| **Modules** | Domain or system features | faili, muda, re, jsoni, hisabati, crypto, njia, mtandao, http, url, os, mfumo |
| **Methods** | Operations on a value | `orodha.idadi()`, `neno.gawa()`, `faili.soma()` |

Keeping this split consistent makes Nuru easier to learn and extend. When in doubt, prefer **method** (if it’s “on” a value) or **module** (if it’s a domain feature); add a **builtin** only when it clearly belongs in the kernel.

---

## Reference: Builtins, modules, and types

### Global builtins

| Builtin | Description |
|---------|-------------|
| **andika** | Print arguments (Inspect) to stdout, space-separated; 0 args = newline. |
| **_andika** | Same as andika but returns a string instead of printing (for capture). |
| **jaza** | Read a line from stdin; optional prompt string. Returns string. |
| **aina** | Return type name of value (e.g. NENO, NAMBA, ORODHA). |
| **mfululizo** | Range: 1 arg (end), 2 args (start, end), or 3 (start, end, step). Returns array of integers. |
| **badilisha** | Convert value to type: (value, "NAMBA" \| "DESIMALI" \| "NENO" \| "BOOLEAN" \| "NAMBA_KUBWA"). |
| **namba** | Convert to integer. |
| **tungo** | Convert to string. |
| **seta** | Create set: `seta()`, `seta(1,2,3)`, or `seta([...])`. Returns seti. |
| **jozi** | Create immutable tuple: `jozi(1,2,3)` or `jozi([...])`. Returns jozi. |

### Modules and their functions

| Module | Import | Functions |
|--------|--------|-----------|
| **faili** | `tumia faili` | soma, andika, ongeza, fungua, futa, fanya, ipo, orodha, tengenezaSarafa, futaSarafa, hali, niSarafa, niFaili, ruhusu, mmiliki, badilisha, kiungo, somaKiungo, funga |
| **muda** | `tumia muda` | hasahivi, lala, tangu, leo, baada_ya, tofauti, ongeza, siku |
| **re** | `tumia re` | linganisha, tafuta, tafutaZote, vikundi, badilisha, gawa, tayari |
| **jsoni** | `tumia jsoni` | dikodi, enkodi, soma, hifadhi, pendeza, msailiaji, msailiaji_bora |
| **hisabati** | `tumia hisabati` | PI, e, phi, ln10, ln2, log10e, log2e, log2, sqrt1_2, sqrt2, sqrt3, sqrt5, EPSILON, abs, sign, ceil, floor, sqrt, cbrt, root, hypot, random, factorial, round, max, min, exp, expm1, log, log10, log1p, cos, sin, tan, acos, asin, atan, cosh, sinh, tanh, acosh, asinh, atanh, atan2, namba_kubwa |
| **crypto** | `tumia crypto` | md5, sha1, sha256, sha512, hmac_sha256, hmac_sha512, bahatiNasibu_baiti, bahatiNasibu_neno, base64_encode, base64_decode, kodeBase64, katuaBase64, hex_encode, hex_decode, sha256_faili, sha512_faili, md5_faili, sha1_faili, pbkdf2_sha256 |
| **njia** | `tumia njia` | jina, kigawaji, sarafa, ext, umbiza, niKamili, unganisha, sawazisha, changanua, husika, tatua, kitenga, posix, win32, glob |
| **mtandao** | `tumia mtandao` | peruzi, tuma |
| **http** | `tumia http` | undaMteja, undaOmbi, sajiliNjia, undaServer, pata, kichwa, tuma, weka, futa, bandika, chaguzi |
| **url** | `tumia url` | changanua, URL, tengeneza, tatua, URLSearchParams, kimbiaNjia, tatuaNjia, kimbiaHoja, tatuaHoja, kamusiKwaHoja, hojaKwaKamusi |
| **os** | `tumia os` | toka, kimbiza |
| **mfumo** | `tumia mfumo` | safishaMemori, takwimuMemori, takwimuMemoriKwa, kumbukumbaDhaifu |

**Note:** **re.tayari**(pattern) returns a compiled regex object (re_iliyotayarishwa) with methods: linganisha(neno), tafuta(neno), tafutaZote(neno), vikundi(neno), badilisha(neno, badiliko), gawa(neno).

### Primitives and types with methods

| Type | Object type name | Methods |
|------|------------------|---------|
| **String** | NENO | idadi, herufikubwa, herufindogo, gawa, panga, kitanzi, ondoaNafasi, anzaNa, ishiaNa, ina, badilishaNeno |
| **Array** | ORODHA | idadi, sukuma, yamwisho, unga, chuja, tafuta, kitanzi, geuza, panga, gawa, sukumaKamaRef; **map**, **chuja** (evaluator: take function arg) |
| **Dict** | KAMUSI | idadi, kitanzi, fungua, funguo, maana, vikundi |
| **Set** | SETI | idadi, ona, ongeza, ondoa, kitanzi |
| **Tuple** | JOZI | idadi, kitanzi (immutable; no index assignment) |
| **File** | FAILI | soma, andika, ongeza, funga, tafuta, hali, isFungwa |
| **Time** | MUDA | ongeza (named: sekunde, dakika, masaa, siku, wiki, miezi, miaka), tangu, panga |
| **Date** | TAREHE | panga (format layout) |
| **Compiled regex** | RE_ILIYOTAYARISHWA | linganisha, tafuta, tafutaZote, vikundi, badilisha, gawa (each takes neno or (neno, badiliko)) |
| **Base64** | BASE64 | kukata, kutoka, data |

**Other types** (no methods or module/object-specific): NAMBA (Integer), DESIMALI (Float), BOOLEAN, TUPU (Null), NAMBA_KUBWA (BigInteger), KITANZI (Iterator), PAKEJI (Package/Instance), MODULE, KUMBUKUMBU (WeakReference), etc.
