// Base64 with Try-Catch Simple Example
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
    andika("Error details:", kosa)
}

andika("\n===== Test Complete =====") 