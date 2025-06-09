// faili_json_example.nr
//
// Mfano wa matumizi ya moduli za faili na json pamoja
// (Example of using file and json modules together)

ita faili kutoka "faili"
ita json kutoka "json"

// ============================================================
// Kuandika data kwenye faili JSON
// (Writing data to a JSON file)
// ============================================================

andika("=== Kuandika JSON kwenye Faili ===")

// Kuunda data
andika("\n- Kuunda na kuandika data JSON:")
fanya mwaka = 2023
fanya majina = ["Founder", "Contributors"]
fanya data = {
    "jina": "Nuru Programming Language",
    "lugha": "Swahili",
    "mwaka": mwaka,
    "sifa": {
        "rahisi": kweli,
        "haraka": kweli
    }
}

// Kuandika kwenye faili
faili.andika("data.json", json.pendeza(data))
andika("  Data imeandikwa kwenye data.json!")

// ============================================================
// Kusoma data kutoka kwenye faili JSON
// (Reading data from a JSON file)
// ============================================================

andika("\n\n=== Kusoma JSON kutoka Faili ===")

// Kusoma faili
andika("\n- Kusoma na kuchanganua JSON:")
fanya yaliyomo = faili.soma("data.json")
andika("  Faili imesomwa!")

// Kuchanganua JSON
fanya obj = json.dikodi(yaliyomo)
andika("  Data imechanganuliwa!")

// Kuonyesha data iliyosomwa
andika("\n- Taarifa za programu:")
andika("  Jina: " + obj.jina)
andika("  Lugha: " + obj.lugha)
andika("  Mwaka: " + obj.mwaka)

andika("\n- Sifa:")
andika("  Rahisi: " + obj.sifa.rahisi)
andika("  Haraka: " + obj.sifa.haraka)

// ============================================================
// Kuhariri JSON na Kuokoa Upya
// ============================================================

andika("\n\n=== Kuhariri JSON na Kuokoa Upya ===")

// Kubadilisha data
andika("\n- Kubadilisha data:")
obj.toleo = "1.0.0"
obj.sifa.salama = kweli
andika("  Data imebadilishwa!")

// Kuokoa upya
andika("\n- Kuokoa data iliyobadilishwa:")
faili.andika("data_mpya.json", json.pendeza(obj))
andika("  Data mpya imeandikwa kwenye data_mpya.json!")

// ============================================================
// Kutumia mipito ya faili
// (Using file streams)
// ============================================================

andika("\n\n=== Kutumia Fungua na Funga ===")

andika("\n- Kufungua na kusoma JSON:")
fd = faili.fungua("data.json", "r")
fanya data_kutoka_fd = json.dikodi(fd.soma())
faili.funga(fd)

andika("  Jina kutoka fd: " + data_kutoka_fd.jina)
andika("  Lugha kutoka fd: " + data_kutoka_fd.lugha)

// ============================================================
// Kusafisha (Cleanup)
// ============================================================

andika("\n\n=== Kusafisha ===")
andika("Kufuta faili zilizotengenezwa...")

faili.futa("data.json")
faili.futa("data_mpya.json")

andika("Usafishaji umekamilika!") 