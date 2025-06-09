// Base64 with Try-Catch Example
// This example demonstrates how to use try-catch with Base64 operations

andika("===== Base64 with Try-Catch =====")

// Example 1: Successful encoding and decoding
andika("\n--- Example 1: Successful Base64 Operations ---")

neno = "Hello, Nuru with Base64!"
andika("Original message:", neno)

jaribu {
    // Encode to Base64
    encoded = kodeBase64(neno)
    andika("Base64 encoded:", encoded)
    
    // Decode from Base64
    decoded = katuaBase64(encoded)
    andika("Decoded message:", decoded)
    
    // Verify
    kama (neno == decoded) {
        andika("SUCCESS: Original and decoded messages match!")
    } sivyo {
        andika("ERROR: Messages don't match")
    }
} bila kosa {
    andika("Error occurred:", kosa)
}

// Example 2: Handling invalid Base64 input
andika("\n--- Example 2: Invalid Base64 Input ---")

invalidBase64 = "This is not valid Base64!"
andika("Invalid Base64 string:", invalidBase64)

jaribu {
    result = katuaBase64(invalidBase64)
    andika("This should not be reached!")
} bila kosa {
    andika("Successfully caught error:")
    andika("  ", kosa)
}

// Example 3: Handling incorrect input type
andika("\n--- Example 3: Incorrect Input Type ---")

jaribu {
    // Try to decode a number instead of a string
    result = katuaBase64(42)
    andika("This should not be reached!")
} bila kosa {
    andika("Successfully caught type error:")
    andika("  ", kosa)
}

// Example 4: More complex scenario with file operations
andika("\n--- Example 4: File Operations with Base64 and Error Handling ---")

// Let's create a file handler that safely encodes file content to Base64
encodeFileToBase64 = unda(filename) {
    jaribu {
        // Attempt to read file
        ita faili = pata("faili")
        content = faili.soma(filename)
        
        // Attempt to encode content
        encoded = kodeBase64(content)
        
        rudisha encoded
    } bila kosa {
        // Something went wrong - let the caller handle it
        rudisha kosa  // Pass the error up to the caller
    }
}

// Now let's use our function with a file that doesn't exist
jaribu {
    result = encodeFileToBase64("nonexistent_file.txt")
    
    // Check if result is an error
    kama (aina(result) == "KOSA") {
        // Re-throw the error
        rudisha result
    }
    
    andika("Encoded file content:", result)
} bila kosa {
    andika("File operation failed as expected:")
    andika("  ", kosa)
}

andika("\n===== Test Complete =====") 