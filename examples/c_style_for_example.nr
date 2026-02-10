// c_style_for_example.nr – C-style for: kwa i = start; condition; update { }

andika("Namba 0 mpaka 4:")
kwa i = 0; i < 5; i = i + 1 {
    andika(i)
}

andika("\nJumla 1+2+...+10:")
jumla = 0
kwa n = 1; n <= 10; n = n + 1 {
    jumla = jumla + n
}
andika(jumla)

andika("\nVunja (stop at 2):")
kwa i = 0; i < 5; i = i + 1 {
    andika(i)
    kama (i == 2) {
        vunja
    }
}

andika("\nEndelea (skip 2):")
kwa i = 0; i < 5; i = i + 1 {
    kama (i == 2) {
        endelea
    }
    andika(i)
}
