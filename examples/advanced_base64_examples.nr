// Advanced Base64 Built-in Functionality Demo
andika("===== Advanced Base64 Test =====")

// 1. URL-safe encoding/decoding
data = "Nuru: Base64/URL-safe!"
encoded_url = kodeBase64(data, "urlsafe")
andika("URL-safe Base64:", encoded_url)
decoded_url = katuaBase64(encoded_url, "urlsafe")
andika("Decoded URL-safe:", decoded_url)

// 2. Encode array of bytes (integers)
arr = [78, 117, 114, 117] // ASCII for 'Nuru'
encoded_arr = kodeBase64(arr)
andika("Encoded array:", encoded_arr)
decoded_arr = katuaBase64(encoded_arr, "array")
andika("Decoded array:", decoded_arr)

// 3. Decode to Byte object
encoded = kodeBase64("byte test")
byte_obj = katuaBase64(encoded, "byte")
andika("Decoded as Byte:", byte_obj)

andika("===== Advanced Base64 Test Complete =====")
