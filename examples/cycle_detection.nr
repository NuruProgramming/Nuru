// Nuru - Cycle Detection Example
// This example demonstrates the reference counting and cycle detection features

// Function to create a circular reference
tengenezaMduara = unda() {
    andika("Kutengeneza muundo wa kamusi na mduara...")
    
    // Create two dictionaries
    a = {"jina": "Kamusi A"}
    b = {"jina": "Kamusi B"}
    
    // Create a circular reference
    a["rejea"] = b
    b["rejea"] = a
    
    andika("Muundo wa mduara umetengenezwa!")
    
    // Print memory statistics
    andika("Takwimu za memori kabla ya kusafisha:")
    takwimuMemori()
    
    // Force garbage collection
    safishaMemori()
    
    // Print memory statistics again
    andika("Takwimu za memori baada ya kusafisha:")
    takwimuMemori()
    
    rudisha [a, b]
}

// Function to create a circular reference with weak references
tengenezaMduaraDhaifu = unda() {
    andika("Kutengeneza muundo wa kamusi na mduara kwa kumbukumbu dhaifu...")
    
    // Create two dictionaries
    a = {"jina": "Kamusi A"}
    b = {"jina": "Kamusi B"}
    
    // Create a circular reference with one weak reference
    a["rejea"] = b
    b["rejea"] = kumbukumbaDhaifu(a)  // Use weak reference
    
    andika("Muundo wa mduara dhaifu umetengenezwa!")
    
    // Print memory statistics
    andika("Takwimu za memori kabla ya kusafisha:")
    takwimuMemori()
    
    // Force garbage collection
    safishaMemori()
    
    // Print memory statistics again
    andika("Takwimu za memori baada ya kusafisha:")
    takwimuMemori()
    
    rudisha [a, b]
}

// Function to test the garbage collection
jaribioGC = unda() {
    andika("Kujaribu garbage collection...")
    
    // Create a lot of objects to trigger garbage collection
    kwa i ktk mfululizo(1, 1000) {
        kamusi = {}
        kwa j ktk mfululizo(1, 10) {
            kamusi[tungo(j)] = j
        }
    }
    
    // Print memory statistics
    andika("Takwimu za memori baada ya kutengeneza objects nyingi:")
    takwimuMemori()
    
    // Force garbage collection
    andika("Kusafisha memori...")
    safishaMemori()
    
    // Print memory statistics again
    andika("Takwimu za memori baada ya kusafisha:")
    takwimuMemori()
}

// Function to show detailed memory statistics
takwimuKamili = unda() {
    stats = {}
    stats = takwimuMemoriKwa(stats)
    
    andika("=== Takwimu Kamili za Memori ===")
    andika("Jumla ya objects: " + tungo(stats["jumla"]))
    
    // Print stats for each type
    kwa aina, idadi ktk stats {
        kama (aina != "jumla") {
            andika(aina + ": " + tungo(idadi))
        }
    }
}

// Main program
andika("=== Majaribio ya Kumbukumbu na Mduara ===")
andika()

// Test 1: Regular circular reference
andika("=== Jaribio 1: Mduara wa Kawaida ===")
mduaraKawaida = tengenezaMduara()
andika()

// Test 2: Weak reference circular reference
andika("=== Jaribio 2: Mduara wa Kumbukumbu Dhaifu ===")
mduaraDhaifu = tengenezaMduaraDhaifu()
andika()

// Test 3: GC stress test
andika("=== Jaribio 3: Kushinikiza GC ===")
jaribioGC()
andika()

// Show detailed memory statistics
andika("=== Jaribio 4: Takwimu Kamili ===")
takwimuKamili()
andika()

andika("Majaribio yamekamilika!") 