# Moduli ya Faili (File System) kwa Nuru

Hati hii inaeleza jinsi ya kutumia moduli ya Faili (File System) katika lugha ya programu ya Nuru. Moduli hii inasaidia kushughulikia faili na sarafa kwenye diski, sawa na maktaba ya Node.js [fs module](https://nodejs.org/docs/v22.9.0/api/fs.html).

## Vipengee Vikuu vya Moduli ya Faili

| Jina | Maelezo |
|------|---------|
| `soma(njia)` | Inasoma faili na kurudisha yaliyomo kama neno |
| `andika(njia, data)` | Inaandika data kwenye faili, ikifuta yaliyomo kama faili ipo tayari |
| `ongeza(njia, data)` | Inaongeza data mwishoni mwa faili, ikitengeneza faili kama haipo |
| `fungua(njia, mode?)` | Inafungua faili na kurudisha kifaa cha faili (file descriptor) |
| `funga(fd)` | Inafunga kifaa cha faili |
| `futa(njia)` | Inafuta faili au sarafa |
| `fanya(njia)` | Inatengeneza faili mpya, ikirudisha kifaa cha faili |
| `ipo(njia)` | Inaangalia kama faili au sarafa ipo |
| `orodha(njia)` | Inasoma yaliyomo kwenye sarafa |
| `tengenezaSarafa(njia, ruhusu?)` | Inatengeneza sarafa |
| `futaSarafa(njia)` | Inafuta sarafa tupu |
| `hali(njia)` | Inapata taarifa za faili au sarafa |
| `niSarafa(njia)` | Inaangalia kama njia ni sarafa |
| `niFaili(njia)` | Inaangalia kama njia ni faili ya kawaida |
| `ruhusu(njia, mode)` | Inabadilisha ruhusu za faili |
| `mmiliki(njia, uid, gid)` | Inabadilisha mmiliki wa faili |
| `badilisha(njiaZamani, njiaMpya)` | Inabadilisha jina la faili au sarafa (kuhamisha) |
| `kiungo(lengo, kiungoJina)` | Inatengeneza kiungo-ishara (symbolic link) |
| `somaKiungo(njia)` | Inasoma thamani ya kiungo-ishara |

## Matumizi ya Msingi

### Kusoma na Kuandika Faili

```nuru
ita faili kutoka "faili"

// Kusoma faili
yaliyomo = faili.soma("maandiko.txt")
andika(yaliyomo)

// Kuandika kwenye faili
faili.andika("maandiko.txt", "Haya ni maandiko mapya")

// Kuongeza kwenye faili
faili.ongeza("maandiko.txt", "\nHii ni laini mpya")
```

### Kufanya Kazi na Sarafa

```nuru
ita faili kutoka "faili"

// Kutengeneza sarafa
faili.tengenezaSarafa("sarafa_mpya")

// Sarafa nyingi (recursive)
faili.tengenezaSarafa("sarafa/ndani/zaidi", rekesia: kweli)

// Kusoma yaliyomo kwenye sarafa
vitu = faili.orodha("sarafa_mpya")
kwa kitu ktk vitu {
    andika(kitu.jina)
    andika("  Ni sarafa? " + kitu.ni_sarafa)
    andika("  Ukubwa: " + kitu.ukubwa + " bytes")
}

// Kufuta sarafa
faili.futaSarafa("sarafa_tupu")

// Kufuta sarafa na yaliyomo
faili.futaSarafa("sarafa/ndani", rekesia: kweli)
```

### Kupata Taarifa za Faili

```nuru
ita faili kutoka "faili"

// Kupata taarifa za faili
taarifa = faili.hali("maandiko.txt")
andika("Jina: " + taarifa.jina)
andika("Ukubwa: " + taarifa.ukubwa + " bytes")
andika("Ni sarafa? " + taarifa.ni_sarafa)
andika("Ruhusu: " + taarifa.ruhusu)
andika("Muda: " + taarifa.muda)

// Kuangalia aina
andika("Ni sarafa? " + faili.niSarafa("maandiko.txt"))
andika("Ni faili? " + faili.niFaili("maandiko.txt"))

// Kuangalia kama faili ipo
andika("Faili ipo? " + faili.ipo("hayupo.txt"))
```

## Kifaa cha Faili (File Descriptor)

Kifaa cha faili ni kitu kinachowakilisha faili iliyofunguliwa.

```nuru
ita faili kutoka "faili"

// Fungua faili kwa kusoma
fd = faili.fungua("maandiko.txt", "r")
andika(fd.Content)

// Fungua faili kwa kuandika
fd_andika = faili.fungua("mpya.txt", "w")
// Andika kwenye faili
fd_andika.andika("Haya ni maandiko kwenye faili mpya")

// Funga faili
faili.funga(fd)
faili.funga(fd_andika)
```

## Modi za Ufunguzi wa Faili

Modi zifuatazo zinaweza kutumika wakati wa kufungua faili:

| Modi | Maelezo |
|------|---------|
| `"r"` | Fungua kwa kusoma (default) |
| `"r+"` | Fungua kwa kusoma na kuandika |
| `"w"` | Fungua kwa kuandika, ikitengeneza faili au kuifuta kama ipo |
| `"w+"` | Fungua kwa kusoma na kuandika, ikitengeneza faili au kuifuta kama ipo |
| `"a"` | Fungua kwa kuongeza, ikitengeneza faili kama haipo |
| `"a+"` | Fungua kwa kusoma na kuongeza, ikitengeneza faili kama haipo |

## Viungo-ishara (Symbolic Links)

Viungo-ishara ni aina maalum ya faili inayoelekeza kwa faili nyingine au sarafa.

```nuru
ita faili kutoka "faili"

// Tengeneza kiungo-ishara
faili.kiungo("maandiko.txt", "kiashiria.txt")

// Soma kiungo-ishara
lengo = faili.somaKiungo("kiashiria.txt")
andika("Kiungo kinaelekeza kwa: " + lengo)

// Soma taarifa za kiungo lenyewe
taarifa = faili.hali("kiashiria.txt")
andika("Ukubwa wa kiungo: " + taarifa.ukubwa)
```

## Kushughulikia Makosa

Njia bora ya kushughulikia makosa ni kutumia kifungu cha `jaribu...shika`:

```nuru
ita faili kutoka "faili"

jaribu {
    faili.soma("faili_isiyopo.txt")
} shika kosa {
    andika("Kosa: " + kosa)
}
```

## Mifano Bora ya Matumizi

### Kunakili Faili

```nuru
ita faili kutoka "faili"

function kunakiliFaili(chanzo, lengo) {
    data = faili.soma(chanzo)
    faili.andika(lengo, data)
    andika("Faili imenakiliwa kutoka " + chanzo + " hadi " + lengo)
}

kunakiliFaili("asili.txt", "nakala.txt")
```

### Kupata Faili na Mifumo Aina Maalum

```nuru
ita faili kutoka "faili"
ita njia kutoka "njia"

function pataFailiNaAina(sarafa, aina) {
    vitu = faili.orodha(sarafa)
    matokeo = []
    
    kwa kitu ktk vitu {
        kama (!kitu.ni_sarafa && njia.ext(kitu.jina) === aina) {
            matokeo.sukuma(njia.unganisha(sarafa, kitu.jina))
        }
    }
    
    rudisha matokeo
}

// Pata faili zote za .txt katika sarafa
faili_za_txt = pataFailiNaAina(".", ".txt")
andika(faili_za_txt)
```

### Kutengeneza Trazi la Sarafa (Directory Tree)

```nuru
ita faili kutoka "faili"
ita njia kutoka "njia"

function andikaSarafa(sarafa, nafasi = "") {
    vitu = faili.orodha(sarafa)
    
    kwa kitu ktk vitu {
        andika(nafasi + "- " + kitu.jina)
        
        kama (kitu.ni_sarafa) {
            njia_mpya = njia.unganisha(sarafa, kitu.jina)
            andikaSarafa(njia_mpya, nafasi + "  ")
        }
    }
}

// Andika trazi la sarafa ya sasa
andikaSarafa(".")
```

## Hitimisho

Moduli ya Faili kwa Nuru inafuata mwelekeo wa Node.js, lakini kwa kutumia sintaksia na majina ya Kiswahili. Moduli hii inasaidia kufanya kazi na mfumo wa faili kwa njia rahisi na ya kimatumizi. 