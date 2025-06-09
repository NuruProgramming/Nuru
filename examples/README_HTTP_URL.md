# Moduli za HTTP na URL kwa Nuru

Hati hii inaeleza jinsi ya kutumia moduli za HTTP na URL katika lugha ya programu ya Nuru. Moduli hizi zinasaidia kushughulikia mawasiliano ya mtandao na uchambuzi wa URL, sawa na maktaba za NodeJS.

## Moduli ya URL

Moduli ya URL inasaidia kuchanganua, kuunda, na kushughulikia anwani za URL. Inafuata muundo wa API ya [Node.js URL module](https://nodejs.org/docs/v10.0.0/api/url.html).

### Vipengee Vikuu vya Moduli ya URL

| Jina | Maelezo |
|------|---------|
| `changanua(url_string)` | Huchanganua neno la URL na kurudisha kamusi ya sehemu zake |
| `URL(url, base_url?)` | Hujenga kitu cha URL kutoka kwa neno la URL |
| `tengeneza(url_obj)` | Hugeuza kamusi ya URL kuwa neno la URL |
| `tatua(base, relative)` | Hutatua URL ya relativi dhidi ya URL ya msingi |
| `URLSearchParams(query)` | Hutengeneza kitu cha kushughulikia vigezo vya utafutaji URL |

### Mfano wa Matumizi ya URL

```nuru
ita url kutoka "url"

// Kuchanganua URL
url_neno = "https://example.com:8080/path?param=value#hash"
parsed_url = url.changanua(url_neno)

// Kupata sehemu za URL
andika(parsed_url.protocol)  // "https:"
andika(parsed_url.hostname)  // "example.com"
andika(parsed_url.port)      // "8080"
andika(parsed_url.pathname)  // "/path"
andika(parsed_url.search)    // "?param=value"
andika(parsed_url.hash)      // "#hash"

// Kutatua URL ya relativi
base = "https://example.com/a/b/c"
resolved = url.tatua(base, "../d/e")
andika(resolved)  // "https://example.com/a/d/e"

// Kushughulikia vigezo vya utafutaji
params = url.URLSearchParams("name=John&age=30")
andika(params.pata("name"))  // "John"
params.weka("age", "31")
andika(params.kuNeno())     // "name=John&age=31"
```

### Kamusi ya URL

URL iliyochanganuliwa inarudishwa kama kamusi yenye ufunguo zifuatazo:

| Ufunguo | Maelezo |
|---------|---------|
| `protocol` | Protokali ya URL pamoja na alama ya koloni (mfano: "https:") |
| `host` | Jina la mwenyeji na port ikiwa ipo (mfano: "example.com:8080") |
| `hostname` | Jina la mwenyeji bila port (mfano: "example.com") |
| `port` | Namba ya port ikiwa imetajwa (mfano: "8080") |
| `pathname` | Njia ya URL (mfano: "/path/to/page") |
| `search` | Sehemu ya utafutaji ya URL pamoja na alama ya swali (mfano: "?query=123") |
| `hash` | Sehemu ya kishiri ya URL pamoja na alama ya reli (mfano: "#section") |
| `origin` | Mwanzo wa URL (mfano: "https://example.com") |
| `href` | URL kamili (mfano: "https://example.com/path?query#hash") |
| `username` | Jina la mtumiaji ikiwa lipo (mfano: "user") |
| `password` | Nywila ikiwa ipo (mfano: "pass") |
| `searchParams` | Kitu cha URLSearchParams kwa ajili ya kushughulikia vigezo vya utafutaji |

### Kitu cha URLSearchParams

Kitu cha URLSearchParams kinatoa njia za kushughulikia vigezo vya utafutaji:

| Njia | Maelezo |
|------|---------|
| `pata(key)` | Kupata thamani ya kwanza ya ufunguo uliotajwa |
| `pataZote(key)` | Kupata safu ya thamani zote za ufunguo uliotajwa |
| `ongeza(key, value)` | Kuongeza jozi ya ufunguo/thamani |
| `weka(key, value)` | Kuweka thamani ya ufunguo, ikifuta thamani zingine zilizopo za ufunguo huo |
| `futa(key)` | Kufuta ufunguo uliotajwa na thamani zake zote |
| `ina(key)` | Kuangalia kama ufunguo upo |
| `kuNeno()` | Kugeuza vigezo kuwa neno |
| `panga()` | Kupanga vigezo kwa mtiririko wa ufunguo |
| `funguo()` | Kupata safu ya funguo zote |
| `thamani()` | Kupata safu ya thamani zote |
| `viingilio()` | Kupata safu ya viingilio vya [ufunguo, thamani] |

## Moduli ya HTTP

Moduli ya HTTP inasaidia kutuma na kupokea maombi ya HTTP. Inafuata muundo wa API ya [Node.js HTTP module](https://nodejs.org/docs/latest-v18.x/api/http.html).

### Vipengee Vikuu vya Moduli ya HTTP

| Jina | Maelezo |
|------|---------|
| `tengenezaServer(callback)` | Huunda seva mpya ya HTTP |
| `tengenezaOmbi(options)` | Huunda ombi la HTTP kwa seva ya mbali |
| `pata(url)` | Husaidia kutuma ombi la GET kwa URL iliyotajwa |
| `STATUS_CODES` | Kamusi ya misimbo ya hali ya HTTP |
| `METHODS` | Safu ya njia za HTTP (GET, POST, n.k) |
| `agentiDunia` | Kitu cha wakala wa HTTP kimataifa |

### Mfano wa Matumizi ya HTTP

```nuru
ita http kutoka "http"

// Kutuma ombi la GET
ombi = http.pata("https://jsonplaceholder.typicode.com/todos/1")

// Kushughulikia jibu
ombi.kwenyeJibu(unda(jibu) {
    andika("Msimbo wa hali: " + jibu.msimboWaHali)
    
    jibu.kwenyeData(unda(data) {
        andika("Data: " + data)
    })
})

// Maliza ombi
ombi.mwisho()

// Kuunda seva
server = http.tengenezaServer(unda(ombi, jibu) {
    jibu.wekaKichwa("Content-Type", "text/plain")
    jibu.andikaKichwa(200)
    jibu.andika("Karibu kwenye seva ya Nuru!")
    jibu.mwisho()
})

// Sikiliza ombi kwenye bandari
server.sikiliza(8080)
```

### Kitu cha Ombi

Kitu cha ombi kinawakilisha ombi la HTTP na kina njia zifuatazo:

| Njia | Maelezo |
|------|---------|
| `andika(data)` | Kuandika data kwa mwili wa ombi |
| `mwisho(data?)` | Kumaliza ombi na kutuma data ya mwisho kama ipo |
| `kwenyeJibu(callback)` | Kusajili tukio la kupokea jibu |
| `kwenyeHitilafu(callback)` | Kusajili tukio la hitilafu |

### Kitu cha Jibu

Kitu cha jibu kinawakilisha jibu la HTTP na kina njia zifuatazo:

| Njia | Maelezo |
|------|---------|
| `andika(data)` | Kuandika data kwa jibu |
| `andikaKichwa(statusCode, statusMessage?, headers?)` | Kuandika kichwa cha hali na vichwa |
| `mwisho(data?)` | Kumaliza jibu na kutuma data ya mwisho kama ipo |
| `wekaKichwa(name, value)` | Kuweka kichwa kimoja |
| `kwenyeData(callback)` | Kusajili mshikaji wa tukio la kupokea data |
| `kwenyeMwisho(callback)` | Kusajili mshikaji wa tukio la kumaliza jibu |

## Mifano ya Kutumia HTTP na URL Pamoja

### Kuchanganua URL na Kutuma Ombi

```nuru
ita http kutoka "http"
ita url kutoka "url"

// Tengeneza URL
api_url = url.changanua("https://jsonplaceholder.typicode.com/todos/1")

// Tumia URL katika ombi la HTTP
ombi = http.pata(api_url)
ombi.kwenyeJibu(unda(jibu) {
    andika("Msimbo wa hali: " + jibu.msimboWaHali)
})
ombi.mwisho()
```

### Kutengeneza Seva na Kuchambua Vigezo vya Utafutaji

```nuru
ita http kutoka "http"
ita url kutoka "url"

server = http.tengenezaServer(unda(ombi, jibu) {
    // Pata URL ya ombi na kuchanganua vigezo
    ombi_url = url.changanua(ombi.url)
    vigezo = url.URLSearchParams(ombi_url.search)
    
    // Pata jina kutoka vigezo
    jina = vigezo.pata("name") || "Mgeni"
    
    // Unda jibu
    jibu.wekaKichwa("Content-Type", "text/html")
    jibu.andikaKichwa(200)
    jibu.andika("<h1>Karibu, " + jina + "!</h1>")
    jibu.mwisho()
})

server.sikiliza(8080)
```

## Hatua Zinazofuata za Maendeleo

1. **Msaada wa HTTPS**: Kuongeza msaada wa mawasiliano salama ya HTTP
2. **Usimamizi wa Muda-Mishi**: Kuboresha usimamizi wa muda-mishi wa maombi
3. **Msaada wa WebSockets**: Kutekeleza msaada wa WebSockets kwa mawasiliano ya haraka
4. **Msaada wa API ya Kuahirisha**: Kutekeleza toleo la Promise/Await kwa maombi ya HTTP
5. **Kusanidi Hifadhi ya Majibu**: Kutekeleza hifadhi ya majibu kwa utendaji bora

## Hitimisho

Moduli za HTTP na URL katika Nuru zinafuata mwelekeo wa Node.js, lakini kwa kutumia sintaksia na majina ya Kiswahili. Kwa pamoja, moduli hizi zinasaidia kutekeleza maombi ya mtandao, uchakato wa majibu, na usimamizi wa URL kwa urahisi. 