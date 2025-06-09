// jsoni_full_test.nr
// Comprehensive test for the JSON module in Nuru

andika("===== JSON Module Full Test =====")
tumia faili
tumia jsoni

// 1. Encode a dict to JSON
andika("\n1. Encode dict to JSON:")
data = {"jina": "Nuru", "mwaka": 2025, "float": 3.14, "kweli": kweli, "tupu": tupu, "arr": [1, 2, 3]}
json_str = jsoni.enkodi(data)
andika(json_str)

// 2. Decode JSON string to dict
andika("\n2. Decode JSON string:")
decoded = jsoni.dikodi(json_str)
andika(decoded)
andika("Jina:", decoded["jina"])
andika("Mwaka:", decoded["mwaka"])
andika("Float:", decoded["float"])
andika("Kweli:", decoded["kweli"])
andika("Tupu:", decoded["tupu"])
andika("Arr:", decoded["arr"])

// 3. Pretty print
andika("\n3. Pretty print:")
pretty = jsoni.pendeza(data)
andika(pretty)

// 4. File operations (with njia):
andika("\n4. File operations (with njia):")
tumia njia
folder = "."
json_path = njia.unganisha(folder, "json_test.json")
faili.andika(json_path, pretty)
loaded = faili.soma(json_path)
andika("Loaded from file:", loaded)
decoded2 = jsoni.dikodi(loaded)
andika("Decoded from file:", decoded2)
faili.futa(json_path)

// 4A. Demonstrate explicit creation and reading of a .json file in the same folder
andika("\n4A. Create and read a .json file in the same folder:")
datafile = njia.unganisha(".", "explicit_test.json")
obj = {"salamu": "Habari dunia", "nambari": [1,2,3], "kweli": kweli}
jsoni_str = jsoni.enkodi(obj)
faili.andika(datafile, jsoni_str)
andika("File created:", datafile)
read_str = faili.soma(datafile)
andika("Read from file:", read_str)
read_obj = jsoni.dikodi(read_str)
andika("Decoded object:", read_obj)
faili.futa(datafile)

// 5. Test parse_float, parse_int, object_hook
andika("\n5. Test parse_float, parse_int, object_hook:")
parse_float = unda(x) { andika("parse_float called with:", x); rudisha 42 }
parse_int = unda(x) { andika("parse_int called with:", x); rudisha -1 }
object_hook = unda(obj) { andika("object_hook called with:", obj); obj.hooked = kweli; rudisha obj }
json_str2 = '{"a": 1, "b": 2.5, "c": {"d": 3}}'
decoded3 = jsoni.dikodi(json_str2, {"parse_float": parse_float, "parse_int": parse_int, "object_hook": object_hook})
andika(decoded3)
andika("a:", decoded3["a"])
andika("b:", decoded3["b"])
andika("c:", decoded3["c"])
andika("c.hooked:", decoded3["c"]["hooked"])

// 6. Test path module
andika("\n6. Test path module:")
tumia njia
p = "/home/user/docs/file.txt"
andika("Njia: ", p)
andika("Jina: ", njia.jina(p))
andika("Sarafa: ", njia.sarafa(p))
andika("Ext: ", njia.ext(p))
andika("Umbiza: ", njia.umbiza({"dir": "/home/user/docs", "base": "file.txt"}))
andika("NiKamili: ", njia.niKamili(p))
andika("Unganisha: ", njia.unganisha("/home", "user", "docs", "file.txt"))
andika("Sawazisha: ", njia.sawazisha("/home/user/../user/docs/file.txt"))
andika("Changanua: ", njia.changanua(p))
andika("Husika: ", njia.husika("/home/user", "/home/user/docs/file.txt"))
andika("Tatua: ", njia.tatua("docs", "file.txt"))
andika("Kitenga: ", njia.kitenga())
andika("Kigawaji: ", njia.kigawaji())

andika("===== JSON Module Full Test Complete =====")
