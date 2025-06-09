// Base64 with Try-Catch Complete Example
andika("===== Base64 with Try-Catch Complete Demo =====")

// Example 1: Basic encoding and decoding
andika("\n--- Example 1: Basic Base64 Operations ---")
validString = "Hello, Nuru with Try-Catch!"
andika("Original string:", validString)

// Encode and decode with try-catch
jaribu {
    // Encode to Base64
    encoded = kodeBase64(validString)
    andika("Base64 encoded:", encoded)
    
    // Decode from Base64
    decoded = katuaBase64(encoded)
    andika("Decoded string:", decoded)
    
    // Verify success
    kama (validString == decoded) {
        andika("Verification: Original and decoded strings match!")
    } 
} bila error {
    andika("Error occurred:", error)
}

// Example 2: Invalid Base64 string
andika("\n--- Example 2: Invalid Base64 String ---")
invalidBase64 = "This is not valid Base64!"
andika("Invalid Base64 string:", invalidBase64)

jaribu {
    result = katuaBase64(invalidBase64)
    andika("This line should not be reached!")
} bila error {
    andika("Successfully caught decoding error:")
    andika("  Error details:", error)
}

// Example 3: Type error - using a number instead of a string
andika("\n--- Example 3: Type Error Handling ---")

jaribu {
    // Try to use a number - this should cause a type error
    badInput = 12345
    andika("Attempting to use a number:", badInput)
    
    result = katuaBase64(badInput)
    andika("This line should not be reached!")
} bila error {
    andika("Successfully caught type error:")
    andika("  Error details:", error)
}

// Example 4: Using try-catch with a function
andika("\n--- Example 4: Function with Try-Catch ---")

// Define a function that handles Base64 operations safely
decodeAndProcessBase64 = unda(input) {
    jaribu {
        decoded = katuaBase64(input)
        andika("Successfully decoded input")
        rudisha decoded
    } bila error {
        andika("Function caught an error:", error)
        rudisha "ERROR: " + error
    }
}

// Test the function with valid input
andika("\nTesting function with valid input:")
result1 = decodeAndProcessBase64(encoded)
andika("Function result:", result1)

// Test the function with invalid input
andika("\nTesting function with invalid input:")
result2 = decodeAndProcessBase64(invalidBase64)
andika("Function result:", result2)

andika("\n===== Test Complete =====") 