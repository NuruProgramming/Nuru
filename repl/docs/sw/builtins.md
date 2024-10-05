# Vitendakazi Vilivyojengwa Ndani ya Nuru

Nuru ina vitendakazi kadhaa vilivyojengwa ndani vinavyofanya kazi husika.

## Kitendakazi andika()

Kitendakazi `andika()` kinatumika kuchapisha ujumbe kwenye konsoli. Inawezakuchukua hoja sifuri au zaidi, na hoja zitachapishwa na nafasi kati yao. Kwa kuongeza, `andika()` huhimili uundaji wa msingi kama vile `/n` kwa ajili ya mstari mpya, `/t` kwa ajili ya nafasi ya kichupo, na `\\` kwa ajili ya mkwajunyuma. Mfano:

```go
andika(1, 2, 3) // Output: 1 2 3
```

```go
andika("Jina: Asha /n Umri: 20 /n Chuo: IFM")

// Output:
// Jina: Asha
// Umri: 20
// Chuo: IFM
```

## Kitendakazi jaza()

Kitendakazi `jaza()` kinatumika kupata ingizo kutoka kwa mtumiaji. Inawezakuchukua hoja sifuri au moja, ambayo ni utungo utakao tumika kama kimahasishi kwa mtumiaji. Mfano:

```go
fanya salamu = unda() {
    fanya jina = jaza("Unaitwa nani? ")
    andika("Mambo vipi", jina)
}

salamu()
```

Katika mfano huu, tunaainisha kitendakazi `salamu()` ambacho kinamhamasisha mtumiaji kuingiza jina kwa kutumia kitendakazi `jaza()`. Kisha tunatumia kitendakazi `andika()` kuchapisha ujumbe unaobeba jina la mtumiaji aliloingiza.

## Kitendakazi aina()

Kitendakazi `aina()` kinatumika kutambua aina ya kitu. Inakubali hoja moja, na thamani inayorudi hua ni utungo unaoonyesha aina ya kitu. Mfano:

```go
aina(2) // Output: "NAMBA"
aina("Nuru") // Output: "NENO"
```

## Kitendakazi fungua()

Kitendakazi `fungua()` kinatumika kufungua faili. Inakubali hoja moja, ambayo ni njia ya faili unalotaka kufungua. Mfano:

```go
faili = fungua("data.txt")
```

Katika mfano huu, tumetumia kitendakazi `fungua()` kufungua faili linaloitwa "data.txt". Kibadilika `faili` kinabeba kumbukumbu ya faili lililofunguliwa.
