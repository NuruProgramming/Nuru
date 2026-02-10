package module

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/NuruProgramming/Nuru/object"
	"golang.org/x/crypto/pbkdf2"
)

// CryptoFunctions: Inajumuisha kazi zote zinazohusiana na cryptografia
var CryptoFunctions = map[string]object.ModuleFunction{}

func init() {
	CryptoFunctions["md5"] = md5Hash
	CryptoFunctions["sha1"] = sha1Hash
	CryptoFunctions["sha256"] = sha256Hash
	CryptoFunctions["sha512"] = sha512Hash
	CryptoFunctions["hmac_sha256"] = hmacSha256
	CryptoFunctions["hmac_sha512"] = hmacSha512
	CryptoFunctions["bahatiNasibu_baiti"] = randomBytes
	CryptoFunctions["bahatiNasibu_neno"] = randomString
	CryptoFunctions["base64_encode"] = base64Encode
	CryptoFunctions["base64_decode"] = base64Decode
	CryptoFunctions["kodeBase64"] = cryptoKodeBase64
	CryptoFunctions["katuaBase64"] = cryptoKatuaBase64
	CryptoFunctions["hex_encode"] = hexEncode
	CryptoFunctions["hex_decode"] = hexDecode
	CryptoFunctions["sha256_faili"] = sha256Faili
	CryptoFunctions["sha512_faili"] = sha512Faili
	CryptoFunctions["md5_faili"] = md5Faili
	CryptoFunctions["sha1_faili"] = sha1Faili
	CryptoFunctions["pbkdf2_sha256"] = pbkdf2Sha256
}

// md5Hash: Rudisha hash ya MD5 ya neno
func md5Hash(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("md5: inahitaji hoja 1 (neno)")
	}
	str, ok := args[0].(*object.String)
	if !ok {
		return newError("md5: hoja lazima iwe neno (string)")
	}
	h := md5.Sum([]byte(str.Value))
	hash := hex.EncodeToString(h[:])
	return &object.String{Value: hash}
}

// sha1Hash: Rudisha hash ya SHA1 ya neno
func sha1Hash(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("sha1: inahitaji hoja 1 (neno)")
	}
	str, ok := args[0].(*object.String)
	if !ok {
		return newError("sha1: hoja lazima iwe neno (string)")
	}
	h := sha1.Sum([]byte(str.Value))
	hash := hex.EncodeToString(h[:])
	return &object.String{Value: hash}
}

// sha256Hash: Rudisha hash ya SHA256 ya neno
func sha256Hash(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("sha256: inahitaji hoja 1 (neno)")
	}
	str, ok := args[0].(*object.String)
	if !ok {
		return newError("sha256: hoja lazima iwe neno (string)")
	}
	h := sha256.Sum256([]byte(str.Value))
	hash := hex.EncodeToString(h[:])
	return &object.String{Value: hash}
}

// sha512Hash: Rudisha hash ya SHA512 ya neno
func sha512Hash(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("sha512: inahitaji hoja 1 (neno)")
	}
	str, ok := args[0].(*object.String)
	if !ok {
		return newError("sha512: hoja lazima iwe neno (string)")
	}
	h := sha512.Sum512([]byte(str.Value))
	hash := hex.EncodeToString(h[:])
	return &object.String{Value: hash}
}

// hmacSha256: Rudisha HMAC-SHA256 ya neno na funguo
func hmacSha256(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return newError("hmac_sha256: inahitaji hoja 2 (ujumbe, funguo)")
	}
	msg, ok := args[0].(*object.String)
	if !ok {
		return newError("hmac_sha256: ujumbe lazima uwe neno (string)")
	}
	key, ok := args[1].(*object.String)
	if !ok {
		return newError("hmac_sha256: funguo lazima iwe neno (string)")
	}
	h := hmac.New(sha256.New, []byte(key.Value))
	h.Write([]byte(msg.Value))
	hash := hex.EncodeToString(h.Sum(nil))
	return &object.String{Value: hash}
}

// hmacSha512: Rudisha HMAC-SHA512 ya neno na funguo
func hmacSha512(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return newError("hmac_sha512: inahitaji hoja 2 (ujumbe, funguo)")
	}
	msg, ok := args[0].(*object.String)
	if !ok {
		return newError("hmac_sha512: ujumbe lazima uwe neno (string)")
	}
	key, ok := args[1].(*object.String)
	if !ok {
		return newError("hmac_sha512: funguo lazima iwe neno (string)")
	}
	h := hmac.New(sha512.New, []byte(key.Value))
	h.Write([]byte(msg.Value))
	hash := hex.EncodeToString(h.Sum(nil))
	return &object.String{Value: hash}
}

// randomBytes: Rudisha safu ya baiti za bahati nasibu
func randomBytes(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("bahatiNasibu_baiti: inahitaji hoja 1 (idadi ya baiti)")
	}
	n, ok := args[0].(*object.Integer)
	if !ok || n.Value <= 0 {
		return newError("bahatiNasibu_baiti: hoja lazima iwe namba chanya (integer > 0)")
	}
	buf := make([]byte, n.Value)
	_, err := rand.Read(buf)
	if err != nil {
		return newError("bahatiNasibu_baiti: %s", err.Error())
	}
	elements := make([]object.Object, n.Value)
	for i, b := range buf {
		elements[i] = &object.Integer{Value: int64(b)}
	}
	return &object.Array{Elements: elements}
}

// randomString: Rudisha neno la bahati nasibu lenye urefu maalum
func randomString(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("bahatiNasibu_neno: inahitaji hoja 1 (urefu)")
	}
	n, ok := args[0].(*object.Integer)
	if !ok || n.Value <= 0 {
		return newError("bahatiNasibu_neno: hoja lazima iwe namba chanya (integer > 0)")
	}
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n.Value)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return newError("bahatiNasibu_neno: %s", err.Error())
		}
		result[i] = letters[num.Int64()]
	}
	return &object.String{Value: string(result)}
}

// base64Encode: Fanya usimbaji wa base64 kwa neno
func base64Encode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("base64_encode: inahitaji hoja 1 (neno)")
	}
	str, ok := args[0].(*object.String)
	if !ok {
		return newError("base64_encode: hoja lazima iwe neno (string)")
	}
	return &object.String{Value: base64.StdEncoding.EncodeToString([]byte(str.Value))}
}

// base64Decode: Fanya ufichuaji wa base64 kwa neno
func base64Decode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("base64_decode: inahitaji hoja 1 (neno)")
	}
	str, ok := args[0].(*object.String)
	if !ok {
		return newError("base64_decode: hoja lazima iwe neno (string)")
	}
	data, err := base64.StdEncoding.DecodeString(str.Value)
	if err != nil {
		return newError("base64_decode: %s", err.Error())
	}
	return &object.String{Value: string(data)}
}

// cryptoKodeBase64: Encode string, Byte, or array of bytes to Base64 (full API matching former builtin)
func cryptoKodeBase64(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return newError("Samahani, kodeBase64 inahitaji hoja 1 au 2, wewe umeweka %d", len(args))
	}
	var data []byte
	var urlSafe bool
	if len(args) == 2 {
		if s, ok := args[1].(*object.String); ok && s.Value == "urlsafe" {
			urlSafe = true
		} else {
			return newError("Chaguo la pili lazima liwe 'urlsafe' kama unataka URL-safe encoding")
		}
	}
	switch arg := args[0].(type) {
	case *object.String:
		data = []byte(arg.Value)
	case *object.Byte:
		data = arg.Value
	case *object.Array:
		arr := arg.Elements
		data = make([]byte, len(arr))
		for i, v := range arr {
			switch vv := v.(type) {
			case *object.Integer:
				if vv.Value < 0 || vv.Value > 255 {
					return newError("Thamani ya orodha lazima iwe kati ya 0 na 255 kwa byte encoding")
				}
				data[i] = byte(vv.Value)
			case *object.Byte:
				if len(vv.Value) != 1 {
					return newError("Kila Byte kwenye orodha lazima iwe na urefu wa 1")
				}
				data[i] = vv.Value[0]
			default:
				return newError("Orodha lazima iwe na namba au Byte pekee, imepata %s", v.Type())
			}
		}
	default:
		return newError("Samahani, kodeBase64 inatumika na Neno, Byte, au Orodha ya namba/Byte tu, umeweka %s", args[0].Type())
	}
	var encoded string
	if urlSafe {
		encoded = base64.URLEncoding.EncodeToString(data)
	} else {
		encoded = base64.StdEncoding.EncodeToString(data)
	}
	return &object.String{Value: encoded}
}

// cryptoKatuaBase64: Decode Base64 string with options urlsafe, byte, array (full API matching former builtin)
func cryptoKatuaBase64(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 || len(args) > 3 {
		return newError("Samahani, katuaBase64 inahitaji hoja 1 hadi 3, wewe umeweka %d", len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return newError("Samahani, katuaBase64 inatumika na Neno pekee, umeweka %s", args[0].Type())
	}
	encoded := args[0].(*object.String).Value
	var urlSafe bool
	var asByte bool
	if len(args) > 1 {
		for i := 1; i < len(args); i++ {
			if s, ok := args[i].(*object.String); ok {
				switch s.Value {
				case "urlsafe":
					urlSafe = true
				case "byte":
					asByte = true
				case "array":
					// handled below
				default:
					return newError("Chaguo lisilojulikana: %s. Tumia 'urlsafe', 'byte', au 'array' tu.", s.Value)
				}
			} else {
				return newError("Chaguo la ziada lazima liwe Neno ('urlsafe', 'byte', 'array')")
			}
		}
	}
	var data []byte
	var err error
	if urlSafe {
		data, err = base64.URLEncoding.DecodeString(encoded)
	} else {
		data, err = base64.StdEncoding.DecodeString(encoded)
	}
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Shida wakati wa kuondoa usimbaji Base64: %s", err.Error())}
	}
	if asByte {
		return &object.Byte{Value: data, String: string(data)}
	}
	if len(args) > 1 {
		for i := 1; i < len(args); i++ {
			if s, ok := args[i].(*object.String); ok && s.Value == "array" {
				elements := make([]object.Object, len(data))
				for j, b := range data {
					elements[j] = &object.Integer{Value: int64(b)}
				}
				return &object.Array{Elements: elements}
			}
		}
	}
	return &object.String{Value: string(data)}
}

// hexEncode: Fanya usimbaji wa hex kwa neno
func hexEncode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("hex_encode: inahitaji hoja 1 (neno)")
	}
	str, ok := args[0].(*object.String)
	if !ok {
		return newError("hex_encode: hoja lazima iwe neno (string)")
	}
	return &object.String{Value: hex.EncodeToString([]byte(str.Value))}
}

// hexDecode: Fanya ufichuaji wa hex kwa neno
func hexDecode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("hex_decode: inahitaji hoja 1 (neno)")
	}
	str, ok := args[0].(*object.String)
	if !ok {
		return newError("hex_decode: hoja lazima iwe neno (string)")
	}
	data, err := hex.DecodeString(str.Value)
	if err != nil {
		return newError("hex_decode: %s", err.Error())
	}
	return &object.String{Value: string(data)}
}

// sha256Faili: Rudisha hash ya SHA256 ya faili
func sha256Faili(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("sha256_faili: inahitaji hoja 1 (jina la faili)")
	}
	filename, ok := args[0].(*object.String)
	if !ok {
		return newError("sha256_faili: hoja lazima iwe neno (string)")
	}
	f, err := os.Open(filename.Value)
	if err != nil {
		return newError("sha256_faili: %s", err.Error())
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return newError("sha256_faili: %s", err.Error())
	}
	hash := hex.EncodeToString(h.Sum(nil))
	return &object.String{Value: hash}
}

// sha512Faili: Rudisha hash ya SHA512 ya faili
func sha512Faili(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("sha512_faili: inahitaji hoja 1 (jina la faili)")
	}
	filename, ok := args[0].(*object.String)
	if !ok {
		return newError("sha512_faili: hoja lazima iwe neno (string)")
	}
	f, err := os.Open(filename.Value)
	if err != nil {
		return newError("sha512_faili: %s", err.Error())
	}
	defer f.Close()
	h := sha512.New()
	if _, err := io.Copy(h, f); err != nil {
		return newError("sha512_faili: %s", err.Error())
	}
	hash := hex.EncodeToString(h.Sum(nil))
	return &object.String{Value: hash}
}

// md5Faili: Rudisha hash ya MD5 ya faili
func md5Faili(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("md5_faili: inahitaji hoja 1 (jina la faili)")
	}
	filename, ok := args[0].(*object.String)
	if !ok {
		return newError("md5_faili: hoja lazima iwe neno (string)")
	}
	f, err := os.Open(filename.Value)
	if err != nil {
		return newError("md5_faili: %s", err.Error())
	}
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return newError("md5_faili: %s", err.Error())
	}
	hash := hex.EncodeToString(h.Sum(nil))
	return &object.String{Value: hash}
}

// sha1Faili: Rudisha hash ya SHA1 ya faili
func sha1Faili(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("sha1_faili: inahitaji hoja 1 (jina la faili)")
	}
	filename, ok := args[0].(*object.String)
	if !ok {
		return newError("sha1_faili: hoja lazima iwe neno (string)")
	}
	f, err := os.Open(filename.Value)
	if err != nil {
		return newError("sha1_faili: %s", err.Error())
	}
	defer f.Close()
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return newError("sha1_faili: %s", err.Error())
	}
	hash := hex.EncodeToString(h.Sum(nil))
	return &object.String{Value: hash}
}

// pbkdf2Sha256: Rudisha funguo iliyotokana na PBKDF2-SHA256
func pbkdf2Sha256(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 4 {
		return newError("pbkdf2_sha256: inahitaji hoja 4 (neno, chumvi, marudio, urefu)")
	}
	password, ok := args[0].(*object.String)
	if !ok {
		return newError("pbkdf2_sha256: neno lazima iwe neno (string)")
	}
	salt, ok := args[1].(*object.String)
	if !ok {
		return newError("pbkdf2_sha256: chumvi lazima iwe neno (string)")
	}
	iter, ok := args[2].(*object.Integer)
	if !ok || iter.Value <= 0 {
		return newError("pbkdf2_sha256: marudio lazima iwe namba chanya (integer > 0)")
	}
	keyLen, ok := args[3].(*object.Integer)
	if !ok || keyLen.Value <= 0 {
		return newError("pbkdf2_sha256: urefu lazima iwe namba chanya (integer > 0)")
	}
	key := pbkdf2.Key([]byte(password.Value), []byte(salt.Value), int(iter.Value), int(keyLen.Value), sha256.New)
	return &object.String{Value: hex.EncodeToString(key)}
}
