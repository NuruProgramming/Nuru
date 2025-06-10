tumia http

andika("\n--- MFANO WA SEVA ---")
http.sajiliNjia("GET", "/salamu", unda(jibu, ombi) {
    jibu.wekaKichwa("Content-Type", "text/plain")
    jibu.mwisho("Habari dunia kutoka kwa seva ya Nuru!")
})
seva = http.undaServer()
seva.sikiliza(":8080")
