# Base64 Implementation in Nuru

This document summarizes the implementation of Base64 encoding and decoding functionality in the Nuru programming language.

## Overview

Base64 is a binary-to-text encoding scheme that represents binary data in an ASCII string format by translating it into a radix-64 representation. This implementation adds native support for Base64 encoding and decoding in Nuru, making it easier to work with binary data, especially for web applications, file transfers, and data storage.

## Implementation Details

The Base64 implementation consists of:

1. **Object Type Definition**:
   - Added `BASE64_OBJ` type constant to `object/object.go`
   - Created `Base64` struct in `object/base64.go` with necessary methods

2. **Built-in Functions**:
   - `kodeBase64(data)`: Encodes string or byte data to Base64
   - `katuaBase64(encodedString)`: Decodes Base64 string back to original data

3. **Documentation**:
   - Added Base64 module documentation to core modules README
   - Created example files demonstrating usage

## Usage

### Basic Encoding and Decoding

```nuru
// Encode a string to Base64
mesiji = "Hello, Nuru with Base64!"
encoded = kodeBase64(mesiji)
andika("Base64 encoded:", encoded)
// Output: SGVsbG8sIE51cnUgd2l0aCBCYXNlNjQh

// Decode from Base64
decoded = katuaBase64(encoded)
andika("Decoded message:", decoded)
// Output: Hello, Nuru with Base64!
```

### File Operations with Base64

Base64 can be used with file operations to encode and decode file contents:

```nuru
ita fs kutoka "faili"

// Read file and encode to Base64
yaliyomo = fs.soma("myfile.txt")
base64Data = kodeBase64(yaliyomo)

// Save encoded data
fs.andika("encoded.txt", base64Data)

// Read and decode
encodedContent = fs.soma("encoded.txt")
originalContent = katuaBase64(encodedContent)
```

## Integration with Other Modules

Base64 works seamlessly with other Nuru modules:

- **HTTP Module**: For sending and receiving encoded data in web requests
- **JSON Module**: For including binary data in JSON objects
- **File System Module**: For storing and retrieving encoded data

## Benefits of the Implementation

1. **Simplicity**: Simple API with just two main functions
2. **Efficiency**: Direct implementation using Go's standard library
3. **Integration**: Works with existing Nuru modules
4. **Practicality**: Useful for real-world applications like web development

## Future Enhancements

Potential enhancements for the Base64 implementation:

1. Support for URL-safe Base64 encoding
2. Stream-based encoding/decoding for large files
3. Support for custom alphabets

## Conclusion

The Base64 implementation enhances Nuru's capabilities for handling binary data and integrates well with the existing module ecosystem. This makes Nuru more versatile for developing applications that need to process, store, or transmit binary data in text form. 