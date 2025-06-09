# Kushughulikia Makosa katika Nuru (Error Handling in Nuru)

Lugha ya Nuru ina mfumo wa kushughulikia makosa kupitia muundo wa `jaribu...bila` (try-catch). Nyenzo hii inasaidia kusimamia makosa kwa njia iliyo na mpangilio badala ya kuacha programu iache kufanya kazi.

## Misingi ya Jaribu...Bila (Try-Catch Basics)

Muundo wa msingi wa kauli ya `jaribu...bila` ni:

```nuru
jaribu {
    // Msimbo ambao unaweza kusababisha kosa
} bila {
    // Msimbo wa kushughulikia kosa
}
```

Ukitaka kupata maelezo ya kosa, unaweza kutumia kitambulishi:

```nuru
jaribu {
    // Msimbo ambao unaweza kusababisha kosa
} bila kosa {
    // Msimbo wa kushughulikia kosa, ukitumia kitambulishi 'kosa' 
    andika("Kosa limetokea: " + kosa)
}
```

## Jinsi Inavyofanya Kazi (How It Works)

1. Msimbo ndani ya sehemu ya `jaribu` unatekelezwa.
2. Ikiwa hakuna kosa, programu inaendelea baada ya bloku ya `jaribu...bila`.
3. Ikiwa kosa limetokea ndani ya bloku ya `jaribu`, utekelezaji unasimama na programu inaingia kwenye bloku ya `bila`.
4. Ikiwa kitambulishi kimetolewa (kama `kosa` katika mfano), kosa lililotokea linawekwa kwenye kitambulishi hicho.
5. Baada ya kutekeleza bloku ya `bila`, programu inaendelea na msimbo unaofuata.

## Mifano (Examples)

### Mfano 1: Kushughulikia kosa la faili isiyopo

```nuru
ita faili = pata("faili")

jaribu {
    yaliyomo = faili.soma("faili_isiyopo.txt")
    andika("Yaliyomo:", yaliyomo)  // Hii haitafika kama faili haipo
} bila kosa {
    andika("Kosa limetokea wakati wa kusoma faili: " + kosa)
}
```

### Mfano 2: Kushughulikia kosa katika function

```nuru
// Function ambayo inaweza kusababisha kosa
gawanya = unda(a, b) {
    kama (b == 0) {
        rudisha newError("Hauwezi kugawanya kwa sifuri")
    }
    rudisha a / b
}

// Kutumia jaribu...bila kusimamia kosa
hesabiGawanya = unda(a, b) {
    jaribu {
        jibu = gawanya(a, b)
        andika("Jibu ni: " + jibu)
    } bila tatizo {
        andika("Tatizo la kihesabu: " + tatizo)
        // Unaweza kurudi na thamani mbadala
        rudisha 0
    }
}
```

### Mfano 3: Kushughulikia makosa tofauti

```nuru
ita faili = pata("faili")

somaNaChanganuaJSON = unda(jina_la_faili) {
    jaribu {
        yaliyomo = faili.soma(jina_la_faili)
        ita json = pata("json")
        
        // Jaribu kuchanganua JSON - inaweza kusababisha kosa lingine
        jaribu {
            data = json.dikodi(yaliyomo)
            rudisha data
        } bila kosa_json {
            andika("Kosa la JSON: " + kosa_json)
            rudisha {}  // Rudisha kamusi tupu kama JSON si sahihi
        }
    } bila kosa_faili {
        andika("Kosa la kusoma faili: " + kosa_faili)
        rudisha {}  // Rudisha kamusi tupu kama faili haiwezi kusomwa
    }
}
```

## Vidokezo na Mbinu Bora (Tips and Best Practices)

1. **Tumia Vizuri:** Tumia `jaribu...bila` kwa operesheni zinazoweza kusababisha makosa, kama kusoma/kuandika faili, shughuli za mtandao, au operesheni zinazotegemea mtumiaji.

2. **Kitambulishi Maana:** Toa kitambulishi chenye maana kwenye bloku ya `bila`, kama `kosa`, `tatizo`, au jina linaloelezea aina ya kosa.

3. **Usihukumu Sana:** Usiweke msimbo mwingi sana kwenye bloku ya `jaribu`. Fanya shughuli moja tu ambayo unajua inaweza kusababisha kosa.

4. **Rejesha Kosa:** Wakati mwingine, badala ya kushughulikia kosa, unaweza kutaka kurudisha kosa kwenye msimbo unaoita:

   ```nuru
   somaSalama = unda(jina) {
       jaribu {
           rudisha faili.soma(jina)
       } bila kosa {
           // Unaweza kurudisha kosa tofauti au kile kile
           rudisha newError("Sikuweza kusoma " + jina + ": " + kosa)
       }
   }
   ```

5. **Kamilisha Rasilimali:** Hakikisha unaachilia rasilimali kama vifaa vya faili katika bloku ya `bila` na `jaribu`:

   ```nuru
   jaribu {
       f = faili.fungua("data.txt", "r")
       maandishi = f.soma()
       f.funga()  // Funga faili kabla ya kurudi
       rudisha maandishi
   } bila kosa {
       // Hakikisha faili imefungwa hata kama kosa limetokea
       andika("Kosa: " + kosa)
       jaribu {
           f.funga()
       } bila {
           // Puuza kosa lolote kwenye kufunga
       }
   }
   ```

Kwa kutumia mfumo wa `jaribu...bila`, unaweza kuandika programu ambazo zinatatua makosa vizuri na kutoa maoni mazuri kwa watumiaji wakati hali haziendi kama ilivyotarajiwa. 