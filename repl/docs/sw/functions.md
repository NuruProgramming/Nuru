# Undo (Functions)

Vitendakazi ni sehemu ya msingi ya Nuru inayokuwezesha kuainisha mapande ya msimbo yanayoweza kutumika tena. Ukurasa huu unaainisha sintaksia na matumizi ya vitendakazi katika nuru ikiwemo vipengele, vipengele vya msingi, matamshi ya kurudisha, kujirudia, na vifungizi.

## Sintaksia

Pande la kitendakazi huanza na neno msingi `unda` likifuatiwa na vipengele vinavyowekwa ndani ya mabano `()` na mwili unaowekwa ndani ya mabano singasinga `{}`. Vitendakazi lazima viwekwe kwenye kibadiliki:

```go
jumla = unda(x, y) {
    rudisha x + y
}

jumla(2, 3) // 5
```

## Vipengele

Vitendakazi vinawezakuwa nazifuri au idadi yoyote ya vipengele. Vipengele vinawezakua vya aina yoyote hata vitendakazi vingine:

```go
salamu = unda(jina) {
    andika("Habari yako", jina)
}

salamu("asha") // Habari yako asha
```

## Vipengele Vya Msingi

Vitendakazi vinawezakupewa vipengele vya msingi:

```go
salimu = unda(salamu="Habari") {
    andika(salamu)
}

salimu() // Habari
salimu("Mambo") // Mambo
```

## Rudisha

Unaweza pia ukarudisha thamani kwa kutumia neno msingi `rudisha`. Neno msingi `rudisha` husitisha pande la msimbo na kurudisha thamani:

```go
mfano = unda(x) {
    rudisha "nimerudi"
    andika(x)
}

mfano("x") // nimerudi
```

## Kujirudia

Nuru pia inahimili kujirudia. Mfano wa kujirudia kwa kitendakazi cha Fibonacci:

```go

fib = unda(n) {
    kama (n <= 1) {
        rudisha n
    } sivyo {
        rudisha fib(n-1) + fib(n-2)
    }
}

andika(fib(10)) // 55
```

Kitendakazi fib kinakokotoa namba ya Fibonacci ya n kwa kujiita yenyewe ikiwa na n-1 na n-2 kama vipengele mpaka ambapo n ni ndogo kuliko au sawa na moja.

## Vifungizi

Vifungizi ni vitendakazi visivyo na jina ambayo vinaweza kudaka na kuhifadhi marejeo ya vibadilika kutoka katika muktadha unaovizunguka. Katika Nuru, unaweza kutengeneza vifungizi kwa kutumia neno msingin `unda` bila kuiweka kwenye kibadiliki. Mfano:

```go
fanya jum = unda(x) {
    rudisha unda(y) {
        rudisha x + y
    }
}

fanya jum_x = jum(5)
andika(jum_x(3)) // 8
```

Katika mfano hapo juu, kitendakazi `jum` kinarudisha kitendakazi kingine ambacho kinabeba kipengele kimoja tu `y`. Kitendakazi kinachorudisha kinawezakupata kibadiliki x kutoka katika muktadha unaokizunguka.

Sasa umeshaelewa misingi ya vitendakazi katika Nuru, ikiwemo kujirudia na vifungizi, unaweza ukatengeneza mapande ya msimbo yanayoweza kutumika tena na tena na kurahisisha programu zako na kuboresha mpangilio wa msimbo wako.
