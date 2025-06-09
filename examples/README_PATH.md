# Moduli ya Njia (Path) kwa Nuru

Hati hii inaeleza jinsi ya kutumia moduli ya Njia (Path) katika lugha ya programu ya Nuru. Moduli hii inasaidia kushughulikia njia za faili na sarafa, sawa na maktaba ya Node.js [path module](https://nodejs.org/docs/v10.0.0/api/path.html).

## Vipengee Vikuu vya Moduli ya Njia

| Jina | Maelezo |
|------|---------|
| `jina(njia, kitanzi?)` | Inapata jina la mwisho la njia (kama Unix basename) |
| `kigawaji` | Kigawaji maalum cha mfumo wa uendeshaji (`:` kwa POSIX, `;` kwa Windows) |
| `sarafa(njia)` | Inapata sehemu ya sarafa ya njia |
| `ext(njia)` | Inapata kitanzi cha faili kutoka njia |
| `umbiza(njiaObj)` | Inatengeneza neno la njia kutoka kwa kamusi |
| `niKamili(njia)` | Inaangalia kama njia ni kamili |
| `unganisha(njia...)` | Inaunganisha vipande vya njia |
| `sawazisha(njia)` | Inasawazisha njia, kutatua `.` na `..` |
| `changanua(njia)` | Inagawa njia kuwa kamusi ya vijenzi |
| `husika(kutoka, mpaka)` | Inapata njia ya husika kati ya njia mbili |
| `tatua(njia...)` | Inatatua njia kuwa njia kamili |
| `kitenga` | Kitenga maalum cha mfumo wa uendeshaji (`/` kwa POSIX, `\` kwa Windows) |
| `posix` | Kitu cha moduli ya njia inayofuata sheria za POSIX |
| `win32` | Kitu cha moduli ya njia inayofuata sheria za Windows |

## Windows na POSIX

Utendaji wa kawaida wa moduli ya Njia unatofautiana kulingana na mfumo wa uendeshaji ambapo programu ya Nuru inaendeshwa. Hasa, wakati wa kuendesha kwenye mfumo wa uendeshaji wa Windows, moduli ya Njia itachukulia kuwa njia za mtindo wa Windows zinatumika.

## Mfano wa Matumizi ya Njia

```nuru
ita njia kutoka "njia"

// Jina la mwisho la njia
jina = njia.jina("/home/user/faili.txt")    // "faili.txt"
jina_bila_kitanzi = njia.jina("/home/user/faili.txt", ".txt")  // "faili"

// Sarafa ya njia
sarafa = njia.sarafa("/home/user/faili.txt")  // "/home/user"

// Kitanzi cha faili
kitanzi = njia.ext("/home/user/faili.txt")   // ".txt"

// Kuunganisha njia
kamili = njia.unganisha("/home", "user", "faili.txt")  // "/home/user/faili.txt"

// Kutatua njia kuwa kamili
kamili = njia.tatua("docs", "../picha")  // Njia kamili kutegemea sarafa ya sasa
```

## Kigawaji (Delimiter)

Inatoa kigawaji cha njia maalum kwa mfumo:

* `;` kwa Windows
* `:` kwa POSIX

Mfano:

```nuru
andika(njia.kigawaji)  // ":" kwa POSIX, ";" kwa Windows

// Kugawa PATH
njia_za_mfumo = PATH.split(njia.kigawaji)
```

## Kitenga (Separator)

Inatoa kitenga cha njia maalum kwa mfumo:

* `\` kwa Windows
* `/` kwa POSIX

Mfano:

```nuru
andika(njia.kitenga)  // "/" kwa POSIX, "\" kwa Windows

// Kugawa njia
vipande = "faili/moja/mbili".split(njia.kitenga)  // ["faili", "moja", "mbili"]
```

## Changanua (Parse)

Njia ya `changanua()` inagawa njia kuwa kamusi yenye sehemu muhimu za njia:

```nuru
sehemu = njia.changanua("/home/user/faili.txt")

// Inajenga kamusi:
// {
//   root: "/",
//   dir: "/home/user",
//   base: "faili.txt",
//   ext: ".txt",
//   name: "faili"
// }
```

Muundo wa matokeo:

```
┌─────────────────────┬────────────┐
│          dir        │    base    │
├──────┬              ├──────┬─────┤
│ root │              │ name │ ext │
"  /    home/user      / faili .txt "
└──────┴──────────────┴──────┴─────┘
```

## Umbiza (Format)

Njia ya `umbiza()` inatengeneza neno la njia kutoka kwa kamusi, kinyume cha `changanua()`:

```nuru
njia_mpya = njia.umbiza({
    root: "/",
    dir: "/home/user",
    base: "faili.txt",
    name: "faili",
    ext: ".txt"
})

// Inajenga: "/home/user/faili.txt"
```

Wakati wa kutoa vijenzi kwa `umbiza`, kumbuka:

* `root` hupuuzwa kama `dir` imetolewa
* `ext` na `name` hupuuzwa kama `base` ipo

## Mifano ya POSIX na Windows

Kwa matokeo thabiti hata wakati wa kushughulikia njia za mifumo tofauti, moduli ya Njia inatoa vijenzi maalum:

```nuru
// Kupata tabia za POSIX kwenye Windows au POSIX
posix = njia.posix
njia_posix = posix.unganisha("/home", "user", "faili.txt")
// Kila wakati: "/home/user/faili.txt"

// Kupata tabia za Windows kwenye Windows au POSIX
win = njia.win32
njia_win = win.unganisha("C:", "Users", "faili.txt")
// Kila wakati: "C:\Users\faili.txt"
```

## Mifano Kamili

Mfano wa kuunganisha na kusawazisha njia:

```nuru
// Kuunganisha njia
njia1 = njia.unganisha("/home", "user", "./docs", "../picha")
// Matokeo: "/home/user/picha"

// Kusawazisha njia
njia2 = njia.sawazisha("/home/./user/../user/picha/../picha/.")
// Matokeo: "/home/user/picha"
```

Mfano wa kutatua njia ya husika:

```nuru
husika = njia.husika("/home/user/docs", "/home/user/picha/sunset.jpg")
// Matokeo: "../picha/sunset.jpg"
```

Mfano wa kutatua njia kuwa kamili:

```nuru
kamili = njia.tatua("docs", "../picha", "./sunset.jpg")
// Kama sarafa ya sasa ni "/home/user", matokeo: "/home/user/picha/sunset.jpg"
```

## Hitimisho

Moduli ya Njia kwa Nuru inafuata mwelekeo wa Node.js, lakini kwa kutumia sintaksia na majina ya Kiswahili. Moduli hii inasaidia kutekeleza ushughulikiaji wa njia wa kifaili na kisarafa kwa urahisi kwa mazingira mbalimbali ya uendeshaji. 