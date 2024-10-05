# Kufanya Kazi na Buliani Katika Nuru

Vitu vyote katika Nuru ni kweli, yaani thamani yoyote ni kweli isipokua tupu and sikweli. Hutumika kutathmini semi ambazo zinarudisha kweli au sikweli.

## Kutathmini Semi za Buliani

### Kutathmini Semi Rahisi

Katika Nuru, unaweza kutathmini semi rahisi zinazorudisha thamani ya buliani:

```go
andika(1 > 2) // Matokeo: `sikweli`

andika(1 + 3 < 10) // Matokeo: `kweli`
```

### Kutathmini Semi Tata

Katika Nuru, unaweza kutumia viendeshaji vya buliani kutathmini semi tata:

```go
a = 5
b = 10
c = 15

tokeo = (a < b) && (b < c)

kama (tokeo) {
    andika("Hali zote mbili ni kweli")
} sivyo {
    andika("Angalau hali moja ni sikweli")
}
// Tokeo: "Hali zote mbili ni kweli"
```

Hapa tumetengeneza vibadilika vitatu a,b,c. Kisha tukatathmini semi (a < b) && (b < c). Kwa sababu semi zote mbili ni kweli, tokeo litakua "Hali zote mbili ni kweli".

## Vitendakazi vya Buliani

Nuru ina vitendakazi vya buliani kadhaa ambavyo unaweza ukatumia kutathmini semi:

### Kitendakazi `&&`

Kitendakazi `&&` hutathmini kwenda kweli kama tu vitu vyote vinavyohusika ni kweli. Kwa mfano:

```go
andika(kweli && kweli) // Tokeo: `kweli`

andika(kweli && sikweli) // Tokeo: `sikweli`
```

### Kitendakazi `||`

Kitendakazi || hutathmini kwenda kweli kama angalau kitu kimoja kati ya vyote vinavyohusika ni kweli. Kwa mfano:

```go
andika(kweli || sikweli) // Tokeo: `kweli`

andika(sikweli || sikweli) // Tokeo: `sikweli`
```

### Kitendakazi `!`

Kitendakazi `!` hukanusha thamani ya kitu. Kwa mfano:

```go
andika(!kweli) // Tokeo: `sikweli`

andika(!sikweli) // Tokeo: `kweli`
```

## Kufanya Kazi na Thamani za Buliani Katika Vitanzi

Katika Nuru, unaweza ukatumia semi za buliani katika vitanzi kuendesha tabia zake. Kwa mfano:

```go
namba = [1, 2, 3, 4, 5]

kwa thamani ktk namba {
    kama (thamani % 2 == 0) {
        andika(thamani, " ni namba shufwa")
    } sivyo {
        andika(thamani, " ni namba witiri")
    }
}

// Output:
// 1 ni namba witiri
// 2 ni namba shufwa
// 3 ni namba witiri
// 4 ni namba shufwa
// 5 ni namba witiri
```

Hapa , tumetengeneza safu yenye namba 1 hadi 5 kisha tukazunguka ndani ya safu hiyo na kwa kila namba tukatumia kitendakazi `%` ilikubaini kama namba ni shufwa au witiri. Matokeo yatakua ni "ni namba shufwa" kwa namba shufwa na "ni namba witiri" kwa namba witiri.

Vitu buliani katika Nuru vinaweza kutumika kutathmini semi ambazo zinarudisha thamani ya kweli au sikweli. Unaweza kutumia vitendakazi vya buliani kutathmini semi tata na kuendesha tabia ya vitanzi. Kuelewa namna ya kufanya kazi na thamani za buliani ni ujuzi wamsingi kwa mtengenezaji programu yeyote wa Nuru.
