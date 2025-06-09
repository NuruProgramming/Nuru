// URL Module Full Example for Nuru
// Demonstrates all Swahili-named URL utilities, including escaping, query encoding/decoding, and error handling

tumia url

andika("--- Path Escaping/Unescaping ---")
original_path = "/foo bar/äö?x=1"
escaped = url.kimbiaNjia(original_path)
andika("Escaped path:", escaped)
unescaped = url.tatuaNjia(escaped)
andika("Unescaped path:", unescaped)

andika("--- Query Escaping/Unescaping ---")
original_query = "jina=Juma & kazi=mpishi"
escaped_query = url.kimbiaHoja(original_query)
andika("Escaped query:", escaped_query)
unescaped_query = url.tatuaHoja(escaped_query)
andika("Unescaped query:", unescaped_query)

andika("--- Dict <-> Query String ---")
kamusi = { "jina": "Asha", "kazi": "mhandisi", "rangi": ["nyekundu", "bluu"] }
query_str = url.kamusiKwaHoja(kamusi)
andika("Kamusi -> Query string:", query_str)
kamusi2 = url.hojaKwaKamusi(query_str)
andika("Query string -> Kamusi:", kamusi2)

andika("--- Error Handling ---")


andika("--- URL Parsing and Object Creation ---")
url_str = "https://user:pass@example.com:8080/path/to/page?foo=bar&baz=qux#frag"
parsed = url.changanua(url_str)
andika("Parsed URL:", parsed)

// Test URL object creation (constructor)
url_obj = url.URL("/other/path?x=1", "https://user:pass@example.com:8080/path/to/page?foo=bar&baz=qux#frag")
andika("URL object (relative to base):", url_obj)

// Test URL formatting
dict_url = {
    "protocol": "https:",
    "username": "ali",
    "password": "siri",
    "host": "nuru.dev:443",
    "pathname": "/swahili",
    "search": "?kitabu=ndio",
    "hash": "#sehemu"
}
formatted = url.tengeneza(dict_url)
andika("Formatted URL:", formatted)

// Test URL resolving
tatua = url.tatua("https://nuru.dev/kitabu/", "../maktaba?x=2")
andika("Tatua (resolve):", tatua)

andika("--- URLSearchParams Object ---")
params = url.URLSearchParams("a=1&b=2&b=3")
params.ongeza("c", "4")
params.weka("a", "10")
andika("Params after ongeza/weka:", params.kuNeno())
andika("pata('a'):", params.pata("a"))
andika("pataZote('b'):", params.pataZote("b"))
andika("ina('c'):", params.ina("c"))
params.futa("b")
andika("After futa('b'):", params.kuNeno())
funguo = params.funguo()
andika("Funguo:", funguo)
thamani = params.thamani()
andika("Thamani:", thamani)
viingilio = params.viingilio()
andika("Viingilio:", viingilio)
params.panga()
andika("After panga:", params.kuNeno())
// Invalid unescape
kosa = url.tatuaNjia("%ZZinvalid")
andika("TatuaNjia error:", kosa)
kosa2 = url.tatuaHoja("%ZZinvalid")
andika("TatuaHoja error:", kosa2)