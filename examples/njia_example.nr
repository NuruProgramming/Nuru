ita njia kutoka "njia"

// Mfano 1: Kupata jina la faili kutoka njia
andika("=== Mfano 1: Jina la faili ===")
jina_faili = njia.jina("/home/user/faili.txt")
andika("Jina la faili: " + jina_faili)

// Ondoa kitanzi
andika("Jina bila kitanzi: " + njia.jina("/home/user/faili.txt", ".txt"))

// Mfano 2: Kupata sarafa (directory)
andika("\n=== Mfano 2: Sarafa ===")
sarafa = njia.sarafa("/home/user/faili.txt")
andika("Sarafa: " + sarafa)

// Mfano 3: Kupata kitanzi (extension)
andika("\n=== Mfano 3: Kitanzi ===")
kitanzi = njia.ext("/home/user/faili.txt")
andika("Kitanzi: " + kitanzi)

// Mfano 4: Kuunganisha njia
andika("\n=== Mfano 4: Unganisha ===")
njia_kamili = njia.unganisha("/home", "user", "faili.txt")
andika("Njia iliyounganishwa: " + njia_kamili)

// Mfano 5: Sawazisha njia
andika("\n=== Mfano 5: Sawazisha ===")
andika("Njia iliyosawazishwa: " + njia.sawazisha("/home/user/../user/./faili.txt"))

// Mfano 6: Angalia kama njia ni kamili
andika("\n=== Mfano 6: Njia kamili ===")
andika("Je, /home/user ni njia kamili? " + njia.niKamili("/home/user"))
andika("Je, faili.txt ni njia kamili? " + njia.niKamili("faili.txt"))

// Mfano 7: Changanua njia
andika("\n=== Mfano 7: Changanua njia ===")
sehemu = njia.changanua("/home/user/faili.txt")
andika("Mizizi (root): " + sehemu.root)
andika("Sarafa (dir): " + sehemu.dir)
andika("Msingi (base): " + sehemu.base)
andika("Jina (name): " + sehemu.name)
andika("Kitanzi (ext): " + sehemu.ext)

// Mfano 8: Umbiza njia
andika("\n=== Mfano 8: Umbiza njia ===")
njia_mpya = njia.umbiza({
    dir: "/home/user",
    base: "faili.txt"
})
andika("Njia iliyoumbizwa: " + njia_mpya)

// Mfano 9: Njia ya husika
andika("\n=== Mfano 9: Njia ya husika ===")
husika = njia.husika("/home/user/docs", "/home/user/picha/picha.jpg")
andika("Njia ya husika: " + husika)

// Mfano 10: Tatua njia
andika("\n=== Mfano 10: Tatua njia ===")
tatuliwa = njia.tatua("docs", "../picha", "picha.jpg")
andika("Njia iliyotatuliwa: " + tatuliwa)

// Mfano 11: Vitenganishi na vigawaji
andika("\n=== Mfano 11: Vitenganishi na vigawaji ===")
andika("Kitenga cha njia: " + njia.kitenga)
andika("Kigawaji cha njia: " + njia.kigawaji)

// Mfano 12: Njia ya POSIX
andika("\n=== Mfano 12: Njia ya POSIX ===")
posix_path = njia.posix
andika("POSIX kitenga: " + posix_path.kitenga)
andika("POSIX unganisha: " + posix_path.unganisha("/home", "user", "faili.txt"))

// Mfano 13: Njia ya Windows
andika("\n=== Mfano 13: Njia ya Windows ===")
win_path = njia.win32
andika("Windows kitenga: " + win_path.kitenga)
andika("Windows unganisha: " + win_path.unganisha("C:", "Users", "faili.txt")) 