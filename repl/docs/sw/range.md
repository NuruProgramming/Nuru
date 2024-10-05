## Kitendakazi cha Mfululizo

Kitendakazi cha `mfululizo` hutoa mfululizo wa nambari, sawa na kitendakazi cha `range()` cha Python. Kinaweza kutumika katika vitanzi au kuunda safu za nambari zinazofuatana.

### Muundo

```go
mfululizo(mwisho)
mfululizo(mwanzo, mwisho)
mfululizo(mwanzo, mwisho, hatua)
```

### Vipengele

- `mwisho`: Kikomo cha juu cha mfululizo (haijumuishwi).
- `mwanzo` (si lazima): Thamani ya kuanzia ya mfululizo. Chaguo-msingi ni 0.
- `hatua` (si lazima): Ongezeko kati ya kila nambari katika mfululizo. Chaguo-msingi ni 1.

### Thamani Inayorudishwa

Hurudisha safu ya nambari kamili.

### Mifano

```go
// Toa nambari kutoka 0 hadi 4
kwa i katika mfululizo(5) {
    andika(i)
}
// Tokeo: 0 1 2 3 4

// Toa nambari kutoka 1 hadi 9
kwa i katika mfululizo(1, 10) {
    andika(i)
}
// Tokeo: 1 2 3 4 5 6 7 8 9

// Toa nambari shufwa kutoka 0 hadi 8
kwa i katika mfululizo(0, 10, 2) {
    andika(i)
}
// Tokeo: 0 2 4 6 8

// Toa nambari kwa mpangilio wa kurudi nyuma
kwa i katika mfululizo(10, 0, -1) {
    andika(i)
}
// Tokeo: 10 9 8 7 6 5 4 3 2 1
```

### Vidokezo

- Thamani ya `mwisho` haijumuishwi, ikimaanisha mfululizo utasimama kabla ya kufikia thamani hii.
- Ikiwa `hatua` hasi imetolewa, `mwanzo` inapaswa kuwa kubwa kuliko `mwisho`.
- Thamani ya `hatua` haiwezi kuwa sifuri.
