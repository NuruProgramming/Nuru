/*
  Mfano wa kutumia moduli ya JSON katika lugha ya Nuru
  (Example of using the JSON module in Nuru language)
*/

andika("=== Mfano wa JSON ===")

// Get JSON module
ita jsoni = pata("jsoni")

// 1. Encode simple array to JSON
andika("\n1. Array to JSON:")
fanya orodha = [1, 2, 3, 4, 5]
fanya json_array = jsoni.enkodi(orodha)
andika(json_array)

// 2. Pretty print
andika("\n2. Pretty print:")
fanya json_pretty = jsoni.pendeza(orodha)
andika(json_pretty)

// 3. Decode JSON string
andika("\n3. Decode JSON string:")
fanya json_data = '{"jina":"Nuru","namba":42}'
fanya decoded = jsoni.dikodi(json_data)
andika("Jina: " + decoded.jina)
andika("Namba: " + Neno(decoded.namba))

// 4. Try file operations if available
jaribu {
    andika("\n4. File operations:");
    // Create simple object
    fanya obj = [1, 2, 3]
    
    // Save to file
    jsoni.hifadhi(obj, "mfano.json")
    andika("File saved")
    
    // Read from file
    fanya loaded = jsoni.soma("mfano.json")
    andika("Loaded: " + Neno(loaded))
    
    // Clean up
    ita fs = pata("faili")
    fs.futa("mfano.json")
} bila {
    andika("File module not available")
}

andika("\n=== End of Example ===") 