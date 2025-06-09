// Base64 with Try-Catch Minimal Example
andika("Base64 with Try-Catch Demo")

// Valid Base64 operation
validString = "Hello, Nuru!"
andika("Original string:", validString)

// Encode to Base64
encoded = kodeBase64(validString)
andika("Base64 encoded:", encoded)

// Successful decoding with try-catch
andika("\nSuccessful decoding with try-catch:")
jaribu {
    decoded = katuaBase64(encoded)
    andika("Decoded string:", decoded)
} bila error {
    andika("Error:", error)
}

// Invalid Base64 decoding with try-catch
andika("\nInvalid Base64 with try-catch:")
invalidBase64 = "This is not valid Base64!"
jaribu {
    result = katuaBase64(invalidBase64)
    andika("Should not reach here!")
} bila error {
    andika("Caught error:", error)
} 