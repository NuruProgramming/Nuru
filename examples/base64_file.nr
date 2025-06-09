// Base64 File Operations Example
ita fs kutoka "faili"

andika("===== Base64 File Operations =====")

// Create a test file
failiJina = "base64_test.txt"
ujumbe = "This is a test file for Base64 encoding and decoding in Nuru."

andika("Writing test content to", failiJina)
fs.andika(failiJina, ujumbe)

// Read the file content
yaliyomo = fs.soma(failiJina)
andika("Original file content:", yaliyomo)

// Encode the file content to Base64
base64Data = kodeBase64(yaliyomo)
andika("Base64 encoded data:", base64Data)

// Save the encoded data to a new file
failiBase64 = "base64_encoded.txt"
fs.andika(failiBase64, base64Data)
andika("Base64 data saved to", failiBase64)

// Read the Base64 file
encodedYaliyomo = fs.soma(failiBase64)
andika("Read encoded data from file:", encodedYaliyomo)

// Decode the Base64 data back to original
decodedData = katuaBase64(encodedYaliyomo)
andika("Decoded data:", decodedData)

// Save the decoded data to a new file to verify
failiDecoded = "base64_decoded.txt"
fs.andika(failiDecoded, decodedData)
andika("Decoded data saved to", failiDecoded)

// Verify the original and decoded data are the same
kama (ujumbe == decodedData) {
    andika("SUCCESS: Original and decoded data match!")
} sivyo {
    andika("ERROR: Original and decoded data do not match!")
    andika("Original:", ujumbe)
    andika("Decoded:", decodedData)
}

// Clean up files
andika("\nCleaning up files...")
fs.futa(failiJina)
fs.futa(failiBase64)
fs.futa(failiDecoded)
andika("Files removed.")

andika("===== Test Complete =====") 