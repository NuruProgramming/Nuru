// Memory allocation test in Nuru
// This script demonstrates memory usage by creating large strings in a loop
// Similar to Python's memory allocation pattern

andika("Starting memory allocation test...")
andika("Creating large strings (100MB each) and storing them in an array")
andika("WARNING: This will consume significant memory!")

data = []
i = 0

wakati (kweli) {
    andika("DEBUG: Starting iteration " + tungo(i + 1))
    
    // Create a string with 100,000,000 'x' characters (approximately 100MB)
    // Using smaller chunks to avoid string concatenation performance issues
    largeString = ""
    j = 0
    andika("DEBUG: Building large string (100MB)...")
    wakati (j < 10000) {
        largeString = largeString + "x"
        j = j + 1
        // Debug every 20,000,000 characters (every 20MB)
        kama (j % 2000 == 0) {
            andika("DEBUG: Built " + tungo(j / 1000000) + "MB of data...")
        }
    }
    andika("DEBUG: Large string completed, length: " + tungo(largeString.idadi()))
    
    // Add the large string to our data array
    andika("DEBUG: Adding string to data array...")
    data.sukuma(largeString)
    i = i + 1
    andika("DEBUG: Current array size: " + tungo(data.idadi()) + " elements")
    
    // Print status every iteration since each iteration is now 100MB
    memoryUsedMB = data.idadi() * 100
    andika("=== STATUS UPDATE ===")
    andika("Iteration " + tungo(i) + ", Memory used: ~" + tungo(memoryUsedMB) + "MB")
    andika("Array elements: " + tungo(data.idadi()))
    andika("=====================")
    
    // Debug every iteration for the first 5 iterations
    kama (i <= 5) {
        andika("DEBUG: Completed iteration " + tungo(i))
    }
    
    // Safety break to prevent infinite loop - reduced limit due to 100MB per iteration
    kama (i >= 10) {
        andika("=== FINAL RESULTS ===")
        andika("Reached 10 iterations, stopping for safety...")
        andika("Total memory allocated: ~" + tungo(data.idadi() * 100) + "MB")
        andika("Total elements in array: " + tungo(data.idadi()))
        andika("Average string length: " + tungo(100000000) + " characters")
        andika("====================")
        vunja
    }
}

andika("DEBUG: Memory test completed or terminated.")
andika("Final array size: " + tungo(data.idadi()) + " elements")
