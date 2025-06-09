// Import the HTTP module
http = ita("http")

// Example 1: Simple HTTP Server
andika("Mfano wa 1: HTTP Server rahisi")

server = http.tengenezaServer(unda(ombi, jibu) {
    jibu.wekaKichwa("Content-Type", "text/html")
    jibu.andikaKichwa(200, "OK")
    jibu.mwisho("<h1>Karibu kwenye Nuru HTTP Server!</h1>")
})

andika("Kuanzisha server katika port 3000...")
server.sikiliza(3000)
andika("Server inasikiliza kwenye port 3000")

// Example 2: HTTP GET Request
andika("\nMfano wa 2: HTTP GET Request")

http.pata("https://jsonplaceholder.typicode.com/posts/1", unda(jibu) {
    jibu.kwenyeData(unda(data) {
        andika("Tumepokea data:")
        andika(data)
    })
    
    jibu.kwenyeMwisho(unda() {
        andika("Ombi limekamilika")
    })
})

// Example 3: More Advanced HTTP Request
andika("\nMfano wa 3: HTTP Request yenye chaguzi zaidi")

ombi = http.tengenezaOmbi({
    "njia": "POST",
    "url": "https://jsonplaceholder.typicode.com/posts",
    "vichwa": {
        "Content-Type": "application/json"
    }
})

// Add event handlers
ombi.kwenyeJibu(unda(jibu) {
    andika("Msimbo wa Hali:", jibu.msimboWaHali)
    
    jibu.kwenyeData(unda(data) {
        andika("Mwili wa jibu:", data)
    })
})

ombi.kwenyeHitilafu(unda(kosa) {
    andika("Kosa limetokea:", kosa)
})

// Write data to request
ombi.andika('{"title": "Mtihani wa Nuru", "body": "Huu ni mtihani wa moduli ya HTTP", "userId": 1}')

// Finish request and send
ombi.mwisho()

// Example 4: Using HTTP Status Codes
andika("\nMfano wa 4: Kutumia msimbo wa hali za HTTP")
msimbo = http.STATUS_CODES
andika("Msimbo 200 inamaanisha:", msimbo["200"])
andika("Msimbo 404 inamaanisha:", msimbo["404"])

// Example 5: Using HTTP Methods
andika("\nMfano wa 5: Njia za HTTP zinazokubalika")
andika("Njia za HTTP:", http.METHODS)

// Print message about the Agent
andika("\nAgenti Dunia inayotumika:")
agenti = http.agentiDunia()
andika("Soketi za juu:", agenti.maxSockets) 