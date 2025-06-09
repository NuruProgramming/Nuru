# Built-in Functions in Nuru

Nuru has several built-in functions that perform specific tasks.

## The andika() Function

The andika() function is used to print out messages to the console. It can take zero or more arguments, and the arguments will be printed out with a space in between them. Additionally, andika() supports basic formatting such as /n for a new line, /t for a tab space, and \\ for a backslash. Here's an example:

```s
andika(1, 2, 3) // Output: "1 2 3"
```

## The jaza() Function

The jaza() function is used to get input from the user. It can take zero or one argument, which is a string that will be used as a prompt for the user. Here's an example:

```s
fanya salamu = unda() {
    fanya jina = jaza("Unaitwa nani? ")
    andika("Mambo vipi", jina)
}

salamu()
```

In this example, we define a function `salamu()` that prompts the user to enter their name using the `jaza()` function. We then use the `andika()` function to print out a message that includes the user's name.

## The aina() Function

The `aina()` function is used to determine the type of an object. It accepts one argument, and the return value will be a string indicating the type of the object. Here's an example:

```s
aina(2) // Output: "NAMBA"
aina("Nuru") // Output: "NENO"
```

## The fungua() Function

The `fungua()` function is used to open a file. It accepts one argument, which is the path to the file that you want to open. Here's an example:

```s
faili = fungua("data.txt")
```

In this example, we use the `fungua()` function to open a file named "data.txt". The variable faili will contain a reference to the opened file.

## The kodeBase64() and katuaBase64() Functions

These functions allow you to encode and decode data using Base64, including support for URL-safe encoding, arrays, and byte objects.

### kodeBase64()
Encodes a string, byte object, or array of bytes/integers to Base64. Optionally, you can specify "urlsafe" as a second argument for URL-safe encoding.

**Examples:**
```s
// Encode a string
encoded = kodeBase64("Nuru Base64!")
andika(encoded) // Output: "TnVydSBCYXNlNjQh"

// URL-safe encoding
encoded_url = kodeBase64("Nuru: Base64/URL-safe!", "urlsafe")
andika(encoded_url)

// Encode an array of bytes (integers)
arr = [78, 117, 114, 117] // ASCII for 'Nuru'
andika(kodeBase64(arr)) // Output: "TnVydQ=="
```

### katuaBase64()
Decodes a Base64-encoded string. You can specify optional arguments:
- "urlsafe" for URL-safe decoding
- "byte" to return a Byte object
- "array" to return an array of integers (bytes)

**Examples:**
```s
// Decode to string
andika(katuaBase64("TnVydSBCYXNlNjQh")) // Output: "Nuru Base64!"

// URL-safe decoding
decoded_url = katuaBase64(encoded_url, "urlsafe")
andika(decoded_url)

// Decode to Byte object
byte_obj = katuaBase64(encoded, "byte")
andika(byte_obj)

// Decode to array of bytes
decoded_arr = katuaBase64("TnVydQ==", "array")
andika(decoded_arr) // Output: [78, 117, 114, 117]
```

These functions make it easy to work with Base64 data in Nuru, including binary and URL-safe use cases.