// iterator_example.nr – using kitanzi() with orodha, kamusi, and neno

andika("=== Orodha ===")
orodha = [10, 20, 30]
it = orodha.kitanzi()
kwa i, v ktk it {
    andika("index", i, "->", v)
}

andika("\n=== Neno ===")
neno = "Nuru"
kwa _, herufi ktk neno.kitanzi() {
    andika(herufi)
}

andika("\n=== Kamusi ===")
kamusi = {"a": 1, "b": 2}
kwa k, v ktk kamusi.kitanzi() {
    andika(k, "->", v)
}

andika("\n=== Viitanzi viwili kwa orodha moja ===")
a = [1, 2]
it1 = a.kitanzi()
it2 = a.kitanzi()
andika("Kitanzi 1:")
kwa _, v ktk it1 { andika(v) }
andika("Kitanzi 2:")
kwa _, v ktk it2 { andika(v) }
