// Simple Base64 example
andika("===== Simple Base64 Test =====")

// Create a string 
mesiji = "Hello, Nuru with Base64!"
andika("Original message:", mesiji)

// Encode to Base64
encoded = kodeBase64(mesiji)
andika("Base64 encoded:", encoded)

// Decode from Base64
decoded = katuaBase64(encoded)
andika("Decoded message:", decoded)

// Test complete
andika("===== Test Complete =====") 