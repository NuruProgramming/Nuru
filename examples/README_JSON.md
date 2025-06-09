# Moduli ya JSON kwa Nuru (JSON Module for Nuru)

Moduli ya JSON hutumika kwa kubadilisha data kati ya umbizo la Nuru na umbizo la JSON. Moduli hii inaruhusu kusoma na kuandika data JSON, ambayo ni muhimu kwa kazi za mtandao na hifadhi data.

## Usajili wa Moduli (Module Registration)

```nuru
ita jsoni = pata("jsoni")
```

## Vipengee vya Moduli (Module Components)

| Jina (Name) | Maelezo (Description) |
|-------------|------------------------|
| `dikodi`    | Kubadilisha JSON kuwa object ya Nuru (mfano wa `json.loads` katika Python) |
| `enkodi`    | Kubadilisha object ya Nuru kuwa JSON (mfano wa `json.dumps` katika Python) |
| `soma`      | Kusoma JSON kutoka faili (mfano wa `json.load` katika Python) |
| `hifadhi`   | Kuhifadhi JSON kwenye faili (mfano wa `json.dump` katika Python) |
| `pendeza`   | Kuandika JSON kwa mpangilio mzuri na nafasi (pretty printing) |
| `msailiaji` | Kutengeneza msailiaji wa JSON (mfano wa `JSONEncoder` katika Python) |
| `msailiaji_bora` | Msailiaji wa JSON ulio na vipengele zaidi |

## Mifano ya Matumizi (Usage Examples)

### 1. Kubadilisha Object kuwa JSON (Object to JSON)

```nuru
ita jsoni = pata("jsoni")

// Tengeneza object ya Nuru
fanya mtu = {
    "jina": "Juma",
    "umri": 30,
    "mjini": kweli,
    "watoto": ["Asha", "Bakari"],
    "simu": {
        "nyumbani": "0222334455",
        "kazini": "0777889900"
    }
}

// Badilisha kuwa JSON
fanya neno_json = jsoni.enkodi(mtu)
andika(neno_json)

// Output: {"jina":"Juma","umri":30,"mjini":true,"watoto":["Asha","Bakari"],"simu":{"nyumbani":"0222334455","kazini":"0777889900"}}
```

### 2. Kubadilisha String ya JSON kuwa Object (JSON to Object)

```nuru
ita jsoni = pata("jsoni")

// JSON string
fanya data_json = '{"jina":"Fatuma","umri":25,"mjini":true,"nambari":[1,2,3]}'

// Badilisha kuwa object ya Nuru
fanya mtu = jsoni.dikodi(data_json)

// Tumia object hiyo
andika(mtu.jina)       // "Fatuma"
andika(mtu.umri)       // 25
andika(mtu.mjini)      // kweli
andika(mtu.nambari[1]) // 2
```

### 3. Kusoma na Kuhifadhi JSON kwenye Faili (File Operations)

```nuru
ita jsoni = pata("jsoni")

// Tengeneza object
fanya wanafunzi = [
    {"jina": "Ali", "alama": 85},
    {"jina": "Maimuna", "alama": 92},
    {"jina": "John", "alama": 78}
]

// Hifadhi kwenye faili
jsoni.hifadhi(wanafunzi, "wanafunzi.json")

// Soma kutoka faili
fanya data = jsoni.soma("wanafunzi.json")

// Tumia data hiyo
kwa kila mwanafunzi katika data {
    andika(mwanafunzi.jina + ": " + Neno(mwanafunzi.alama))
}
```

### 4. Kutumia Pretty Printing

```nuru
ita jsoni = pata("jsoni")

fanya config = {
    "jina": "Mfumo wa Wanafunzi",
    "toleo": "1.0.0",
    "mpangilio": {
        "rangi": "bluu",
        "ukubwa": "kati",
        "lugha": "Swahili"
    }
}

// Andika JSON kwa mpangilio mzuri
fanya json_nzuri = jsoni.pendeza(config)
andika(json_nzuri)

/* Output yaweza kuwa:
{
    "jina": "Mfumo wa Wanafunzi",
    "toleo": "1.0.0",
    "mpangilio": {
        "rangi": "bluu",
        "ukubwa": "kati",
        "lugha": "Swahili"
    }
}
*/
```

## Chaguo za Dikodi (Decode Options)

Inawezekana kutoa chaguo zaidi kwa `dikodi`:

- `parse_float`: Function ya kutafsiri namba za desimali
- `parse_int`: Function ya kutafsiri namba za kawaida
- `object_hook`: Function ya kubadilisha objects baada ya kutafsiriwa

## Chaguo za Enkodi (Encode Options)

Inawezekana kutoa chaguo zaidi kwa `enkodi`:

- `skipkeys`: Kama `kweli`, funguo ambazo si strings zitapuuzwa
- `ensure_ascii`: Kama `kweli`, herufi zote zitaandikwa kwa ASCII
- `indent`: Idadi ya nafasi za kutumia kwa indentation
- `sort_keys`: Kama `kweli`, funguo za dictionary zitapangwa kwa alfabeti
- `separators`: Viweka tofauti kwa items na key-value pairs (mfano: [",", ":"])

## JSONEncoder Maalum (Custom JSONEncoder)

Unaweza kutengeneza JSONEncoder maalum:

```nuru
ita jsoni = pata("jsoni")

// Tengeneza encoder
fanya encoder = jsoni.msailiaji({
    "indent": 2,
    "sort_keys": kweli
})

// Tumia encoder hiyo
fanya obj = {"c": 3, "a": 1, "b": 2}
fanya json_nzuri = encoder.enkodi(obj)
andika(json_nzuri)

/* Output:
{
  "a": 1,
  "b": 2,
  "c": 3
}
*/
```

## Maelezo (Notes)

1. Moduli ya JSON inahitaji moduli ya `faili` kwa functions za `soma` na `hifadhi`.
2. Aina za Nuru zinabadilishwa hivi kwenda JSON:
   - `Dict` → JSON object
   - `Array` → JSON array
   - `String` → JSON string
   - `Integer/Float` → JSON number
   - `Boolean` → JSON boolean
   - `Null` → JSON null

3. Aina za JSON zinabadilishwa hivi kwenda Nuru:
   - JSON object → `Dict`
   - JSON array → `Array`
   - JSON string → `String`
   - JSON number → `Integer` au `Float`
   - JSON boolean → `Boolean`
   - JSON null → `Null` 