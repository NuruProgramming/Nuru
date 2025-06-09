// jaribu_bila_example.nr
//
// Mfano wa matumizi ya jaribu...bila (try-catch) kwa kushughulikia makosa
// (Example of using try-catch for error handling)

andika("=== Mfano wa Jaribu-Bila (Try-Catch) ===\n")

// Mfano 1: Kushughulikia kosa la faili isiyopo
// Example 1: Handling non-existent file error
andika("Mfano 1: Kushughulikia kosa la faili isiyopo")

ita faili = pata("faili")

jaribu {
    andika("Kujaribu kusoma faili isiyopo...")
    yaliyomo = faili.soma("faili_isiyopo.txt")
    andika("Yaliyomo:", yaliyomo)  // Hii haitafika
} bila kosa {
    andika("Kosa limekamatwa: " + kosa)
}

andika("")

// Mfano 2: Kushughulikia kosa la gawanya kwa sifuri
// Example 2: Handling division by zero error
andika("Mfano 2: Kushughulikia kosa la gawanya kwa sifuri")

gawanya = unda(a, b) {
    kama (b == 0) {
        rudisha "Kosa: Hauwezi kugawanya kwa sifuri"
    }
    rudisha a / b
}

hesabu = unda(a, b) {
    jaribu {
        andika("Kujaribu kugawanya " + a + " kwa " + b)
        matokeo = gawanya(a, b)
        andika("Matokeo: " + matokeo)
        rudisha matokeo
    } bila {
        andika("Kosa la kihesabu lilitokea!")
        rudisha 0
    }
}

hesabu(10, 2)  // Inapaswa kufanya kazi
hesabu(10, 0)  // Inapaswa kukamata kosa

andika("")

// Mfano 3: Kutumia jaribu katika function
// Example 3: Using try in a function
andika("Mfano 3: Kutumia jaribu katika function")

pataData = unda(jina) {
    jaribu {
        data = faili.soma(jina)
        rudisha data
    } bila tatizo {
        andika("Tatizo la faili: " + tatizo)
        rudisha ""
    }
}

data1 = pataData("jaribu_bila_example.nr")  // Faili iliopo
andika("Urefu wa data 1: " + data1.idadi())

data2 = pataData("faili_isiyopo.txt")  // Faili isiyopo
andika("Urefu wa data 2: " + data2.idadi())

andika("\n=== Mwisho wa Mfano ===") 