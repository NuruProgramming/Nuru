# Seti Katika Nuru

**Seti** ni mkusaniko wa vipengele visivyorudia. Aina zinazoweza kukokotolewa tu (namba, neno, booleani, namba kubwa) zinaweza kuwa vipengele vya seti. Seti inasaidia ukaguzi wa uwanachama kwa `ktk`, uzunguko kwa `kwa ... ktk`, na mbinu.

## Kutengeneza seti

Tumia kiendeshi **seta**:

```s
seta()           // seti tupu
seta(1, 2, 3)    // seti kutoka hoja
seta([1, 2, 2])  // seti kutoka safu (rudupo huondolewa)
```

## Uwanachama

Tumia `ktk` kuangalia kama kipengele kiko ndani:

```s
2 ktk seta(1, 2, 3)   // kweli
5 ktk seta(1, 2, 3)   // sikweli
```

## Mbinu

| Mbinu | Hoja | Rudishi | Maelezo |
|--------|------|---------|---------|
| **idadi** | — | namba | Idadi ya vipengele. |
| **ona**(kipengele) | 1 | kweli/sikweli | Kweli kama seti ina kipengele hicho. |
| **ongeza**(...) | 1+ | seti | Ongeza kipengele au zaidi; rudisha seti. |
| **ondoa**(kipengele) | 1 | seti | Ondoa kipengele; rudisha seti. |
| **kitanzi** | — | kitanzi | Kitanzi juu ya seti. |

## Kuzunguka

Mpangilio wa uzunguko unafuata umbo la neno la kipengele. Unaweza kutumia seti moja kwa moja kwa `kwa ... ktk` au kupiga `.kitanzi()`:

```s
kwa _, v ktk seta("a", "b", "c") {
    andika(v)
}
```
