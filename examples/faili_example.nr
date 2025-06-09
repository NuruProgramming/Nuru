// faili_example.nr
// Mfano wa kutumia moduli ya faili (File System) katika Nuru

ita faili kutoka "faili"
ita njia kutoka "njia"

// Sehemu 1: Kazi za msingi za faili
andika("=== Kazi za Msingi za Faili ===")

// Kuandika faili
andika("\n- Kuandika faili mpya:")
faili.andika("mfano_faili.txt", "Huu ni mfano wa faili.\nIna laini mbili.")
andika("  Faili imeandikwa!")

// Kusoma faili
andika("\n- Kusoma faili:")
andika("  " + faili.soma("mfano_faili.txt"))

// Kuongeza kwenye faili
andika("\n- Kuongeza kwa faili:")
faili.ongeza("mfano_faili.txt", "\nHii ni laini ya tatu iliyoongezwa.")
andika("  " + faili.soma("mfano_faili.txt"))

// Kuangalia kama faili ipo
andika("\n- Kuangalia kama faili ipo:")
andika("  mfano_faili.txt ipo: " + faili.ipo("mfano_faili.txt"))
andika("  faili_isiyopo.txt ipo: " + faili.ipo("faili_isiyopo.txt"))

// Kupata taarifa za faili
andika("\n- Kupata taarifa za faili (stat):")
taarifa = faili.hali("mfano_faili.txt")
andika("  Jina: " + taarifa.jina)
andika("  Ukubwa: " + taarifa.ukubwa + " bytes")
andika("  Ni sarafa: " + taarifa.ni_sarafa)

// Sehemu 2: Kazi za sarafa
andika("\n\n=== Kazi za Sarafa ===")

// Kutengeneza sarafa
andika("\n- Kutengeneza sarafa mpya:")
faili.tengenezaSarafa("mfano_sarafa")
andika("  Sarafa imetengenezwa!")

// Kuorodhesha yaliyomo kwenye sarafa
andika("\n- Orodha ya yaliyomo kwenye sarafa ya sasa:")
vitu = faili.orodha(".")
andika("  Idadi ya vitu: " + vitu.idadi())

// Kuonyesha vitu 5 vya kwanza tu kwa ufupi
kwa i, kitu ktk vitu {
    kama (i >= 5) {
        vunja
    }
    aina = kitu.ni_sarafa ? "Sarafa" : "Faili"
    andika("  " + (i+1) + ". " + kitu.jina + " (" + aina + ")")
}

// Kutengeneza faili ndani ya sarafa
andika("\n- Kutengeneza faili ndani ya sarafa mpya:")
faili.andika("mfano_sarafa/faili_ndani.txt", "Faili ndani ya sarafa mpya")
andika("  Faili ndani ya sarafa imetengenezwa!")

// Sehemu 3: Kubadilisha na kuhamisha
andika("\n\n=== Kubadilisha na Kuhamisha ===")

// Kubadilisha jina la faili
andika("\n- Kubadilisha jina la faili:")
faili.badilisha("mfano_faili.txt", "faili_jina_jipya.txt")
andika("  Jina la faili limebadilishwa!")
andika("  Faili ya zamani ipo: " + faili.ipo("mfano_faili.txt"))
andika("  Faili mpya ipo: " + faili.ipo("faili_jina_jipya.txt"))

// Kuhamisha faili kwenye sarafa
andika("\n- Kuhamisha faili kwenye sarafa:")
faili.badilisha("faili_jina_jipya.txt", "mfano_sarafa/faili_iliyohamishwa.txt")
andika("  Faili imehamishwa kwenye sarafa!")

// Sehemu 4: Kutumia fungua na kufunga
andika("\n\n=== Kutumia Fungua na Funga ===")

// Kufungua faili kwa kuandika
andika("\n- Kufungua faili kwa kuandika:")
fd = faili.fungua("faili_mpya_kwa_fungua.txt", "w")
andika("  Faili imefunguliwa!")

// Kuandika kwenye faili iliyofunguliwa
fd.andika("Hii ni faili iliyotengenezwa na kufunguliwa kwa kutumia fungua().")
andika("  Maandishi yameandikwa kwenye faili!")

// Kufunga faili
faili.funga(fd)
andika("  Faili imefungwa!")

// Kusoma faili iliyoandikwa
andika("\n- Kusoma faili iliyotengenezwa na fungua:")
andika("  " + faili.soma("faili_mpya_kwa_fungua.txt"))

// Kusafisha - Kufuta faili zilizotengenezwa na mfano huu
andika("\n\n=== Kusafisha ===")
andika("Kufuta faili zilizotengenezwa...")

faili.futa("faili_mpya_kwa_fungua.txt")
faili.futa("mfano_sarafa/faili_ndani.txt")
faili.futa("mfano_sarafa/faili_iliyohamishwa.txt")
faili.futaSarafa("mfano_sarafa")

andika("Usafishaji umekamilika!") 