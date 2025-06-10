// Mfano kamili wa matumizi ya moduli ya http ya Nuru (API mpya ya Kiswahili, Go-idiomatic)

tumia http

// --- MFANO WA MTEJA (CLIENT) ---
andika("\n--- MFANO WA MTEJA ---")

// Ombi la GET kwa kutumia mteja chaguo-msingi
jibu = http.pata("https://httpbin.org/get")
andika("Jibu la GET:")
andika("Aina ya jibu:")
andika(aina(jibu))
andika("Inspect ya jibu:")
andika(jibu)
andika("Aina ya jibu.mwili:")
andika(aina(jibu.mwili))
andika("Inspect ya jibu.mwili:")
andika(jibu.mwili)
kama (aina(jibu.mwili) == "NENO") {
    andika(jibu.mwili)
} sivyo {
    andika("Hakuna mwili au kosa: ", jibu.mwili)
}

// Unda mteja maalum na muda wa mwisho
mteja = http.undaMteja({"mudaWaMwisho": 5000})
jibu2 = mteja.pata("https://httpbin.org/headers")
andika("Jibu la mteja maalum:")
andika("Aina ya jibu2:")
andika(aina(jibu2))
andika("Inspect ya jibu2:")
andika(jibu2)
andika("Aina ya jibu2.mwili:")
andika(aina(jibu2.mwili))
andika("Inspect ya jibu2.mwili:")
andika(jibu2.mwili)
kama (aina(jibu2.mwili) == "NENO") {
    andika(jibu2.mwili)
} sivyo {
    andika("Hakuna mwili au kosa: ", jibu2.mwili)
}

// Unda ombi maalum (POST)
ombi = http.undaOmbi("POST", "https://httpbin.org/post", "{json: 'ndio'}")
jibu3 = mteja.tekeleza(ombi)
andika("Jibu la POST:")
andika("Aina ya jibu3:")
andika(aina(jibu3))
andika("Inspect ya jibu3:")
andika(jibu3)
andika("Aina ya jibu3.mwili:")
andika(aina(jibu3.mwili))
andika("Inspect ya jibu3.mwili:")
andika(jibu3.mwili)
kama (aina(jibu3.mwili) == "NENO") {
    andika(jibu3.mwili)
} sivyo {
    andika("Hakuna mwili au kosa: ", jibu3.mwili)
}

// --- MFANO WA SEVA (SERVER) ---
andika("\n--- MFANO WA SEVA ---")

// Sajili njia ya GET kwenye /salamu
http.sajiliNjia("GET", "/salamu", unda(jibu, ombi) {
    jibu.wekaKichwa("Content-Type", "text/plain")
    jibu.andika("Habari dunia kutoka kwa seva ya Nuru!")
    jibu.mwisho()
})

// Unda seva na sikiliza kwenye bandari 8080 (itaendelea kukimbia)
seva = http.undaServer()
seva.sikiliza(":8080")
andika("Seva inasikiliza kwenye http://localhost:8080/salamu (jaribu kwa browser au curl)")

// Jaribu server yetu kwa kutumia http.pata
andika("\n--- KUTUMA OMBI KWA SERVER YETU ---")
jibu_local = http.pata("http://localhost:8080/salamu")
andika("Jibu la server yetu:")
andika(jibu_local.mwili)

// Mwisho wa mfano
andika("Mfano wa http umekamilika.")
