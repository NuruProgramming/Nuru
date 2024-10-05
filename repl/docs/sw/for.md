# Vitanzi Vya Kwa Katika Nuru

Vitanzi vya `kwa` ni muundo msingi wa udhibiti katika Nuru ambavyo hutumika kuzunguka vitu vinavyozungukika kama tungo, safu, na kamusi. Ukurasahuu unaangazia sintaksia na matumizi ya Vitanzi katika Nuru, ikiwemo kuzunguka ndani ya jozi ya funguo-thamani, na matumizi ya matamshi `vunja` na `endelea`.

## Sintaksia

Kutengeneza kitanzi cha `kwa`, tumia neno msingi `kwa` likifwatiwa na kitambulishi cha muda mfupi kama vile `i` au `v` na kitu kinachozungukika. Funga mwili wa kitanzi na mabano singasinga `{}`. Mfano unaotumia tungo:

```go
jina = "lugano"

kwa i ktk jina {
    andika(i)
}

// Tokeo:
l
u
g
a
n
o
```

## Kuzunguka Ndani ya Jozi ya Funguo-Thamani

### Kamusi

Nuru inakuruhusu kuzunguka ndani ya kamusi kupata thamani moja moja au jozi ya funguo na thamani yake. Kupata tu thamani, tumia kitambulisha cha muda mfupi:

```go
kamusi = {"a": "andaa", "b": "baba"}

kwa v ktk kamusi {
    andika(v)
}

// Tokeo:

andaa
baba
```

Kupata thamani ya funguo na thamani zake, tumia vitambulishi vya muda mfupi viwili:

```go

kwa k, v ktk kamusi {
    andika(k + " ni " + v)
}

// Tokeo:

a ni andaa
b ni baba
```

### Tungo

Kuzunguka juu ya thamani za tungo, tumia kitambulishi cha muda mfupi:

```go
kwa v ktk "mojo" {
    andika(v)
}

// Tokeo:

m
o
j
o
```

Kuzunguka juu ya funguo na thamani zake, tumia vitambulishi vya muda mfupi viwili:

```go
kwa i, v ktk "mojo" {
    andika(i, "->", v)
}

// Tokeo:

0 -> m
1 -> o
2 -> j
3 -> o
```

### Safu

Kuzunguka juu ya thamani za safu, tumia kitambulishi cha muda mfupi:

```go
majina = ["juma", "asha", "haruna"]

kwa v ktk majina {
    andika(v)
}

// Tokeo:

juma
asha
haruna
```

Kuzunguka juu ya funguo na thamani katika safy, tumia vitambulishi vya muda mfupi viwili:

```go
kwa i, v ktk majina {
    andika(i, "-", v)
}

// Tokeo:

0 - juma
1 - asha
2 - haruna
```

## Vunja na Endelea

### Vunja

Tumia neno msingi `vunja` kisitisha kitanzi:

```go

kwa i, v ktk "mojo" {
    kama (i == 2) {
        andika("nimevunja")
        vunja
    }
    andika(v)
}

// Tokeo:

m
o
j
nimevunja

```

### Endelea

Tumia neno msingi `endelea` kuruka mzunguko maalum:

```go
kwa i, v ktk "mojo" {
    kama (i == 2) {
        andika("nimeruka")
        endelea
    }
    andika(v)
}

// Tokeo:

m
o
nimeruka
o
```
