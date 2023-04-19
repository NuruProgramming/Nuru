// basics

jina = "Nuru"
andika(jina) // Nuru

// lists

orodha = [1, "pili", kweli]

namba = [10, 20, 30]
jina = namba[1]  // jina is 20

namba[1] = 25

a = [1, 2, 3]
b = [4, 5, 6]
c = a + b // c is now [1, 2, 3, 4, 5, 6]

namba = [10, 20, 30]
andika(20 ktk namba)  // will print kweli

namba = [1, 2, 3, 4, 5]

kwa thamani ktk namba {
    andika(thamani)
}

majina = ["Juma", "Asha", "Haruna"]

kwa idx, jina ktk majina {
    andika(idx, "-", jina)
}

a = [1, 2, 3]
urefu = a.idadi()
andika(urefu)  // will print 3

a = [1, 2, 3]
a.sukuma("s", "g")
andika(a)  // will print [1, 2, 3, "s", "g"]

a = [1, 2, 3]
mwisho = a.yamwisho()
andika(mwisho)  // will print 3

b = []
mwisho = b.yamwisho()
andika(mwisho)  // will print tupu

andika(1 > 2) // Output: `sikweli`

andika(1 + 3 < 10) // Output: `kweli`

a = 5
b = 10
c = 15

result = (a < b) && (b < c)

kama (result) {
    andika("Both conditions are true")
} sivyo {
    andika("At least one condition is false")
}

// Output: "Both conditions are true"

andika(kweli && kweli) // Output: `kweli`

andika(kweli && sikweli) // Output: `sikweli`

andika(kweli || sikweli) // Output: `kweli`

andika(sikweli || sikweli) // Output: `sikweli`

andika(!kweli) // Output: `sikweli`

andika(!sikweli) // Output: `kweli`

namba = [1, 2, 3, 4, 5]

kwa thamani ktk namba {
    kama (thamani % 2 == 0) {
        andika(thamani, "is even")
    } sivyo {
        andika(thamani, "is odd")
    }
}
// Output:
// 1 is odd
// 2 is even
// 3 is odd
// 4 is even
// 5 is odd

salamu = unda() {
    jina = jaza("Unaitwa nani? ")
    andika("Mambo vipi", jina)
}

salamu()

aina(2) // Output: "NAMBA"
aina("Nuru") // Output: "NENO"


orodha = {"jina": "Juma", "umri": 25}

k = {
    "jina": "Juma",
    "umri": 25,
    kweli: "kweli",
    "salimu": unda(x) { andika("habari", x) },
    "sina value": tupu
}


andika(k[kweli]) // kweli
andika(k["salimu"]("Juma")) // habari Juma

k['umri'] = 30
andika(k['umri']) // 30

k["lugha"] = "Kiswahili"
andika(k["lugha"]) // Kiswahili

matunda = {"a": "apple", "b": "banana"}
mboga = {"c": "carrot", "d": "daikon"}
vyakula = matunda + mboga
andika(vyakula) // {"a": "apple", "b": "banana", "c": "carrot", "d": "daikon"}


"umri" ktk k // kweli
"urefu" ktk k // sikweli


hobby = {"a": "asili", "b": "baiskeli", "c": "chakula"}
kwa i, v ktk hobby {
    andika(i, "=>", v)
}
/* a => asili
   b => baiskeli
   c => chakula */


kwa v ktk hobby {
    andika(v)
}

/*
asili
baiskeli
chakula
*/

jina = "lugano"

kwa i ktk jina {
    andika(i)
}

kamusi = {"a": "andaa", "b": "baba"}

kwa v ktk kamusi {
    andika(v)
}


kwa k, v ktk kamusi {
    andika(k + " ni " + v)
}

kwa v ktk "mojo" {
    andika(v)
}

kwa i, v ktk "mojo" {
    andika(i, "->", v)
}

majina = ["juma", "asha", "haruna"]

kwa v ktk majina {
    andika(v)
}

kwa i, v ktk majina {
    andika(i, "-", v)
}


kwa i, v ktk "mojo" {
    kama (i == 2) {
        andika("nimevunja")
        vunja
    }
    andika(v)
}

kwa i, v ktk "mojo" {
    kama (i == 2) {
        andika("nimeruka")
        endelea
    }
    andika(v)
}

jum = unda(x, y) {
    rudisha x + y
}

jum(2, 3) // 5

salamu = unda() {
    andika("Habari yako")
}

salamu()

salamu = unda(jina) {
    andika("Habari yako", jina)
}

salamu("asha") // Habari yako asha

salimu = unda(salamu="Habari") {
    andika(salamu)
}

salimu() // Habari
salimu("Mambo") // Mambo

mfano = unda(x) {
    rudisha "nimerudi"
    andika(x)
}

mfano("x") // nimerudi


fib = unda(n) {
    kama (n <= 1) {
        rudisha n
    } sivyo {
        rudisha fib(n-1) + fib(n-2)
    }
}

andika(fib(10)) // 55

jum = unda(x) {
    rudisha unda(y) {
        rudisha x + y
    }
}

jum_x = jum(5)
andika(jum_x(3)) // 8

2 + 3 * 5 // 17

a = 2.5
b = 3/5

a + b // 2.8

i = 2.4

i++ // 3.4

i = 2

i *= 3 // 6
i /= 2 // 3
i += 100 // 103
i -= 10 // 93
i %= 90 // 3

i = -10

wakati (i < 0) {
    andika(i)
    i++
}


andika("mambo") // mambo

a = 'niaje'

andika("mambo", a) // mambo niaje

a = "habari" + " " + "yako"

andika(a) // habari yako

b = "habari"

b += " yako"

// habari yako

andika("mambo " * 4)

// mambo mambo mambo mambo

a = "habari"

a *= 4

// habarihabarihabarihabari

jina = "avicenna"

kwa i ktk jina {andika(i)}

a = "nuru"

andika(a == "nuru") // kweli

andika(a == "mambo") // sikweli

a = "mambo"
a.idadi() // 5

a = "NURU"
a.herufindogo() // nuru

a = "nuru mambo habari"
b = a.gawa()
andika(b) // ["nuru", "mambo", "habari"]

a = "nuru,mambo,habari"
b = a.gawa(",")
andika(b) // ["nuru", "mambo", "habari"]

a = 2

badili (a){
	ikiwa 3 {
		andika("a ni tatu")
	}
	ikiwa 2 {
		andika ("a ni mbili")
	}
}

badili (a) {
	ikiwa 1,2,3 {
		andika("a ni kati ya 1, 2 au 3")
	}
	ikiwa 4 {
		andika("a ni 4")
	}
}

z = 20

badili(z) {
	ikiwa 10 {
		andika("kumi")
	}
	ikiwa 30 {
		andika("thelathini")
	}
	kawaida {
		andika("ishirini")
	}
}

i = 1

wakati (i <= 5) {
	andika(i)
	i++
}

i = 1

wakati (i < 5) {
	kama (i == 3) {
		andika("nimevunja")
		vunja
	}
	andika(i)
	i++
}

i = 0

wakati (i < 5) {
	i++
	kama (i == 3) {
		andika("nimeruka")
		endelea
	}
	andika(i)
}

// using time:

tumia muda 

andika(muda.hasahivi())
andika("Tunalala")
muda.lala(2)
andika("Tumeamka")

s = muda.hasahivi()
muda.lala(2)
andika(muda.tangu(s))

andika(s.ongeza(sekunde=3, miaka=10, dakika=4))

// using file:

fail = fungua("sarufi.nr")

andika(fail.soma())

// using net and json:

tumia mtandao, jsoni
url = "https://v2.jokeapi.dev/joke/Any?type=single"

resp = mtandao.peruzi(url)

resp = jsoni.dikodi(resp)

andika(resp["joke"])


// os module

tumia os

andika(os.kimbiza("pwd"))

os.toka(0)
