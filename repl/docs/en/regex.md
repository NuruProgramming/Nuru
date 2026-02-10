# Regular expressions (regex)

The **re** module provides pattern matching and replacement on strings using regular expressions. Nuru uses Go’s RE2-based engine, so matching runs in linear time and is safe from ReDoS (catastrophic backtracking).

## Import

```s
tumia re
```

## Functions

| Function | Arguments | Returns | Description |
|----------|-----------|---------|-------------|
| **linganisha** | (pattern, neno) | kweli/sikweli | True if the string matches the pattern. |
| **tafuta** | (pattern, neno) | neno or tupu | First match, or tupu if none. |
| **tafutaZote** | (pattern, neno) | orodha | All non-overlapping matches. |
| **vikundi** | (pattern, neno) | orodha or tupu | First match plus capturing groups (parentheses). Index 0 is full match, 1 is first group, etc. |
| **badilisha** | (pattern, neno, badiliko) | neno | Replace all matches with the replacement string. |
| **gawa** | (pattern, neno) | orodha | Split the string by the pattern. |
| **tayari** | (pattern) | re_iliyotayarishwa | Compile a pattern once; returns an object with methods that take only the string (see below). |

If the pattern is invalid, each function returns an error (KOSA) and does not crash the program.

### Compiled regex (tayari)

Use **re.tayari(pattern)** when you use the same pattern many times. It returns a compiled-regex object with methods that take one (or two) arguments instead of repeating the pattern:

| Method | Arguments | Returns |
|--------|------------|---------|
| **linganisha** | (neno) | kweli/sikweli |
| **tafuta** | (neno) | neno or tupu |
| **tafutaZote** | (neno) | orodha |
| **vikundi** | (neno) | orodha or tupu |
| **badilisha** | (neno, badiliko) | neno |
| **gawa** | (neno) | orodha |

Example:

```s
tumia re
r = re.tayari("[0-9]+")
r.linganisha("sala 123")   // kweli
r.tafuta("sala 123")       // "123"
```

## Examples

### linganisha (match)

```s
tumia re
re.linganisha("[0-9]+", "sala 123")   // kweli
re.linganisha("[0-9]+", "hakuna")     // sikweli
```

### tafuta (first match)

```s
tumia re
re.tafuta("[0-9]+", "sala 123 zaidi")  // "123"
re.tafuta("[a-z]+", "123")             // tupu
```

### tafutaZote (all matches)

```s
tumia re
re.tafutaZote("[0-9]+", "1 na 2 na 3")  // ["1", "2", "3"]
```

### vikundi (capturing groups)

```s
tumia re
re.vikundi("(\\w+)\\s+(\\w+)", "jina asha")  // ["jina asha", "jina", "asha"]
```

### badilisha (replace all)

```s
tumia re
re.badilisha("[0-9]", "a1b2c3", "X")   // "aXbXcX"
```

### gawa (split by pattern)

```s
tumia re
re.gawa("\\s+", "a   b   c")   // ["a", "b", "c"]
```

## Notes

- Pattern and text are strings. In Nuru strings, backslash is `\\` for a literal backslash in the pattern (e.g. `"\\d+"` for digits).
- Invalid patterns return an error; they do not crash the program.
- Matching does not modify the original string; all functions return new values.
