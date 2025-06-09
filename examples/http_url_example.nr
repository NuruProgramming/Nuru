ita http kutoka "http"
ita url kutoka "url"

// Mfano 1: Kuchanganua URL
andika("=== Mfano 1: Kuchanganua URL ===")
url_neno = "https://username:password@example.com:8080/path/to/page?param1=value1&param2=value2#section"
parsed_url = url.changanua(url_neno)

andika("Protocol: " + parsed_url.protocol)
andika("Host: " + parsed_url.host)
andika("Hostname: " + parsed_url.hostname)
andika("Port: " + parsed_url.port)
andika("Path: " + parsed_url.pathname)
andika("Search: " + parsed_url.search)
andika("Hash: " + parsed_url.hash)
andika("Username: " + parsed_url.username)
andika("Password: " + parsed_url.password)
andika("Origin: " + parsed_url.origin)

// Mfano 2: Kutengeneza URL mpya
andika("\n=== Mfano 2: Kutengeneza URL mpya ===")
msingi_url = "https://example.com/base/"
relative_url = "relative/path?query=123"
full_url = url.URL(relative_url, msingi_url)
andika("URL kamili: " + full_url.href)

// Mfano 3: Kutatua URL
andika("\n=== Mfano 3: Kutatua URL ===")
base = "https://example.com/base/"
relative = "../other/path"
resolved = url.tatua(base, relative)
andika("URL iliyotatuliwa: " + resolved)

// Mfano 4: URLSearchParams
andika("\n=== Mfano 4: URLSearchParams ===")
query_string = "name=John&age=30&hobby=coding&hobby=reading"
params = url.URLSearchParams(query_string)

// Kupata thamani
andika("Jina: " + params.pata("name"))
andika("Umri: " + params.pata("age"))
andika("Hobi zote: " + params.pataZote("hobby")[0] + ", " + params.pataZote("hobby")[1])

// Kuangalia kuwepo kwa ufunguo
andika("Ina 'name'? " + params.ina("name"))
andika("Ina 'address'? " + params.ina("address"))

// Kuongeza na kuweka vigezo
params.ongeza("country", "Tanzania")
params.weka("age", "31")
andika("Vigezo vilivyosasishwa: " + params.kuNeno())

// Mfano 5: Kutumia HTTP na URL pamoja
andika("\n=== Mfano 5: HTTP na URL pamoja ===")

// Tengeneza URL
api_url = url.changanua("https://jsonplaceholder.typicode.com/todos/1")
andika("Tunatuma ombi kwa: " + api_url.href)

// Tumia URL katika ombi la HTTP
ombi = http.pata(api_url)

// Sikiliza tukio la jibu
ombi.kwenyeJibu(unda(jibu) {
    andika("Msimbo wa hali: " + jibu.msimboWaHali)
    
    // Kusanya data
    data = ""
    jibu.kwenyeData(unda(kipande) {
        data = data + kipande
    })
    
    // Maliza ombi
    jibu.kwenyeMwisho(unda() {
        andika("Data iliyopokelewa:")
        andika(data)
        
        // Pia angalia URL ya jibu
        andika("URL ya jibu: " + jibu.url.href)
    })
})

// Sikiliza tukio la hitilafu
ombi.kwenyeHitilafu(unda(kosa) {
    andika("Hitilafu imetokea: " + kosa)
})

// Maliza ombi
ombi.mwisho()

// Mfano 6: Kutengeneza server na kutumia URLSearchParams
andika("\n=== Mfano 6: Server na URLSearchParams ===")
andika("Tengeneza server kwenye port 8080...")
andika("Fungua http://localhost:8080/hello?name=Juma&greeting=Habari katika kivinjari")

server = http.tengenezaServer(unda(ombi, jibu) {
    // Pata URL ya ombi na kuchanganua vigezo
    ombi_url = url.changanua(ombi.url)
    vigezo = url.URLSearchParams(ombi_url.search)
    
    // Angalia njia
    kama (ombi_url.pathname == "/hello") {
        // Pata jina na salamu kutoka vigezo
        jina = vigezo.pata("name")
        kama (!jina) {
            jina = "Mgeni"
        }
        
        salamu = vigezo.pata("greeting")
        kama (!salamu) {
            salamu = "Karibu"
        }
        
        // Unda jibu
        jibu.wekaKichwa("Content-Type", "text/html")
        jibu.andikaKichwa(200)
        jibu.andika("<h1>" + salamu + ", " + jina + "!</h1>")
        jibu.andika("<p>Vigezo vya ombi:</p>")
        jibu.andika("<pre>" + ombi_url.search + "</pre>")
        jibu.mwisho()
    } sivyo {
        // Njia haipo
        jibu.wekaKichwa("Content-Type", "text/html")
        jibu.andikaKichwa(404)
        jibu.andika("<h1>Samahani, ukurasa haujapatikana</h1>")
        jibu.mwisho()
    }
})

// Sikiliza kwenye port 8080
server.sikiliza(8080)

andika("Server imeanza, inasikiza kwenye port 8080...") 