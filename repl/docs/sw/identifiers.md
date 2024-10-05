# Vitambulisho katika Nuru

Vitambulisho hutumika kuweka majina kwenye vigezo, vitendakazi na vipengele vingine katika msimbo wako wa Nuru. Ukurasa huu unashughulikia sheria na mbinu bora za kuunda vitambulisho katika Nuru.

## Sheria za Sintaksia

Vitambulisho vinaweza kuwa na herufi, nambari na nistari wa chini `_`. Walakini, kuna sheria chache ambazo unapaswa kufuata wakati wa kuunda vitambulisho:

- Vitambulisho haviwezi kuanza na nambari.
- Vitambulisho huwa na tofauti kulingana na matumizi ya herufi kubwa na ndogo. Kwa mfano, `kibadilikaChangu` na `kibadilikachangu` huchukuliwa kuwa vitambulisho tofauti.

Hapa kuna mifano ya vitambulisho halali:

```go
fanya mwaka_wa_kuzaliwa = 2020
andika(mwaka_wa_kuzaliwa) // 2020

fanya badili_c_kwenda_p = "C kwenda P"
andika(badili_c_kwenda_p) // "C kwenda P"
```

Katika mifano iliyo hapo juu, mwaka_wa_kuzaliwa na badili_c_kwenda_p zote ni vitambulisho halali.

## Mazoea Bora

Wakati wa kuchagua vitambulisho, ni muhimu kufuata mazoea bora ili kuhakikisha kuwa msimbo wako uko wazi na rahisi kueleweka:

- Tumia majina yanayoelezea wazi kusudi au maana ya kigezo au kitendakazi.
- Fuata kanuni thabiti ya kuweka majina, kama vile camelCase (kibadilikaChangu) au snake_case (kibadilika_changu).
- Epuka kutumia majina tofauti ya herufi moja, isipokuwa kwa vijisehemu vinavyokubalika kwa kawaida kama vile vihesabu vitanzi (i, j, k).

Kwa kufuata mbinu bora hizi unapounda vitambulisho, utafanya code yako ya Nuru iwe rahisi kusoma na kutunza kwa wewe na wengine.
