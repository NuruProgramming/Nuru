// crypto_example.nr
// Example usage of the 'crypto' module in Nuru

tumia crypto, njia

andika("MD5:", crypto.md5("Nuru"))
andika("SHA256:", crypto.sha256("Nuru"))
andika("SHA512:", crypto.sha512("Nuru"))
andika("SHA1:", crypto.sha1("Nuru"))

andika("\n-- HMAC --")
andika("HMAC_SHA256:", crypto.hmac_sha256("Nuru", "siri"))
andika("HMAC_SHA512:", crypto.hmac_sha512("Nuru", "siri"))

andika("\n-- Random --")
andika("Bahati Nasibu Baiti:", crypto.bahatiNasibu_baiti(8))
andika("Bahati Nasibu Neno:", crypto.bahatiNasibu_neno(12))

andika("\n-- Base64/Hex --")
base = "NuruCrypto"
andika("Base64 Encode:", crypto.base64_encode(base))
andika("Base64 Decode:", crypto.base64_decode(crypto.base64_encode(base)))
andika("Hex Encode:", crypto.hex_encode(base))
andika("Hex Decode:", crypto.hex_decode(crypto.hex_encode(base)))

andika("\n-- File Hashing --")
faili = njia.unganisha(__DIR__, "data.json")
andika("SHA256 ya faili:", crypto.sha256_faili(faili))
andika("SHA512 ya faili:", crypto.sha512_faili(faili))
andika("MD5 ya faili:", crypto.md5_faili(faili))
andika("SHA1 ya faili:", crypto.sha1_faili(faili))

andika("\n-- PBKDF2 --")
pw = "siri"
salt = "chumvi"
marudio = 10000
urefu = 32
andika("PBKDF2-SHA256:", crypto.pbkdf2_sha256(pw, salt, marudio, urefu))
