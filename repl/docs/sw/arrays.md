# Orodha Au Safu Katika Nuru

Safu katika nuru ni miundo ya data ambayo inaweza kubeba vitu vingi, ikiwa ni pamoja na aina za data tofauti tofauti kama `namba`, `tungo`, `buliani`, `vitendakazi`, na thamani `tupu`. Ukurasa huu unaangazia vipengele mbalimbali vya safu, ikiwemo namna ya kutengeneza, kuchambua, na kuzunguka ndani yake kwa kutumia vitendakazi vilivyojengwa ndani ya Nuru.

## Kutengeneza Safu

Kutengeneza safu, tumia mabano mraba na tenganisha kila kitu kimoja kwa kutumia mkwaju:

```s
orodha = [1, "pili", kweli]
```

## Kupata na Kubadilisha Vipengele vya Safu

Safu katika Nuru ni zero-indexed; ikimaanisha kipengele cha kwanza katika safu kina kumbukumbu namba 0. Kupata kipengele, unaweza ukatumia kumbukumbu namba yake ndani ya mabano mraba:

```s
namba = [10, 20, 30]
jina = namba[1]  // jina is 20
```

Unaweza ukabadilisha kipengele katika safu kwa kutumia kumbukumbu namba yake:

```s
namba = [10, 20, 30]
namba[1] = 25
andika(namba)  // Tokeo: [10,25,30]
```

## Kuunganisha Safu

Kuunganisha safu mbili au zaidi, tumia kiendeshi `+`:

```s
a = [1, 2, 3]
b = [4, 5, 6]
c = a + b
// c is now [1, 2, 3, 4, 5, 6]
```

## Kuangalia Uanachama Katika Safu

Tumia neno msingi `ktk` kuangalia kama kipengele kipo ndani ya safu:

```s
namba = [10, 20, 30]
andika(20 ktk namba)  // Tokeo: kweli
```

## Kuzunguka Ndani ya Safu

Unaweza kutumia maneno msingi `kwa` na `ktk` kuzunguka ndani ya safu. Kuzunguka ndani ya safu na kupata kipengele peke yake tumia sintaksia ifuatayo:

```s
namba = [1, 2, 3, 4, 5]

kwa thamani ktk namba {
    andika(thamani)
}

//Tokeo:
1
2
3
4
5
```

Kuzunguka ndani ya safu na kupata kumbukumbu namba na kipengele tumia sintaksi aifuatayo:

```s
majina = ["Juma", "Asha", "Haruna"]

kwa idx, jina ktk majina {
    andika(idx, "-", jina)
}

//Tokeo:
0-Juma
1-Asha
2-Haruna
```

## Vitendakazi vya Safu

Nuru ina vitendakazi mbalimbali vilivyojengwa ndani kwa ajili ya Safu:

### idadi()

`idadi()` hurudisha urefu wa safu:

```s
a = [1, 2, 3]
urefu = a.idadi()
andika(urefu)  // Tokeo: 3
```

### sukuma()

`sukuma()` huongeza kipengele kimoja au zaidi mwishoni mwa safu:

```s
a = [1, 2, 3]
a.sukuma("s", "g")
andika(a)  // Tokeo [1, 2, 3, "s", "g"]
```

### yamwisho()

`yamwisho()` hurudisha kipengele cha mwisho katika safu, au `tupu` kama safu haina kitu:

```s
a = [1, 2, 3]
mwisho = a.yamwisho()
andika(mwisho)  // Tokeo: 3

b = []
mwisho = b.yamwisho()
andika(mwisho)  // Tokeo: tupu
```

Kwa kutumia taarifa hii, unaweza ukafanyakazi na safu za Nuru kwa ufanisi, kufanya iwe rahisi kuchambua mikusanyo ya data katika programu zako.
