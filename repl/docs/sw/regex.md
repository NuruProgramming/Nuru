# Usawiri wa Pattern (regex)

Moduli **re** inatoa ulinganizi wa pattern na ubadilishaji wa maandishi kwa kutumia usawiri wa kawaida (regular expressions). Nuru inatumia injini ya RE2 ya Go, hivyo ulinganizi unaenda kwa wakati wa mstari na ni salama kutoka ReDoS.

## Tumia moduli

```s
tumia re
```

## Kazi

| Kazi | Hoja | Inarudisha | Maelezo |
|------|------|------------|---------|
| **linganisha** | (pattern, neno) | kweli/sikweli | Kweli ikiwa neno linalingana na pattern. |
| **tafuta** | (pattern, neno) | neno au tupu | Lingano la kwanza, au tupu ikiwa hakuna. |
| **tafutaZote** | (pattern, neno) | orodha | Matches yote yasiyoshikana. |
| **vikundi** | (pattern, neno) | orodha au tupu | Match ya kwanza pamoja na vikundi (mabano). Index 0 ni match nzima, 1 ni kikundi cha kwanza, n.k. |
| **badilisha** | (pattern, neno, badiliko) | neno | Badilisha matches zote na neno la badiliko. |
| **gawa** | (pattern, neno) | orodha | Gawa neno kwa pattern. |
| **tayari** | (pattern) | re_iliyotayarishwa | Tayarisha pattern mara moja; rudisha kitu chenye mbinu zinazopokea neno tu (tazama chini). |

Kama pattern si sahihi, kila kazi inarudisha kosa (KOSA) na haivunji programu.

### Re iliyotayarishwa (tayari)

Tumia **re.tayari(pattern)** unapotumia pattern ile ile mara nyingi. Inarudisha kitu chenye mbinu: linganisha(neno), tafuta(neno), tafutaZote(neno), vikundi(neno), badilisha(neno, badiliko), gawa(neno).

## Mifano

```s
tumia re
re.linganisha("[0-9]+", "sala 123")   // kweli
re.tafuta("[0-9]+", "sala 123 zaidi") // "123"
re.tafutaZote("[0-9]+", "1 na 2 na 3") // ["1", "2", "3"]
re.vikundi("(\\w+)\\s+(\\w+)", "jina asha")  // ["jina asha", "jina", "asha"]
re.badilisha("[0-9]", "a1b2c3", "X")  // "aXbXcX"
re.gawa("\\s+", "a   b   c")          // ["a", "b", "c"]
```

## Vidokezo

- Pattern na maandishi ni maneno. Kwenye maneno ya Nuru, backslash ni `\\` kwa backslash halisi kwenye pattern.
- Pattern zisizo sahihi zinazrudisha kosa; hazivunji programu.
- Ulinganizi haubadilishi neno asilia; kazi zote zinazrudisha thamani mpya.
