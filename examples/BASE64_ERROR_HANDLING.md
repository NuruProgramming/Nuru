# Base64 Integration with Try-Catch Error Handling in Nuru

This document explains how the Base64 module in Nuru is integrated with the error handling mechanism to provide robust error recovery during encoding and decoding operations.

## Overview

Base64 operations, especially decoding, can fail for various reasons:
- Invalid Base64 characters in the input string
- Incorrect padding
- Incorrect input type (e.g., attempting to decode a number instead of a string)

The integration between Base64 and the try-catch construct allows programs to gracefully handle these errors instead of crashing.

## Implementation Details

The Base64 functions return proper Error objects when operations fail:

1. **Error Creation**: When Base64 operations fail, they return `object.Error` objects with descriptive error messages.
2. **Error Propagation**: These errors are automatically propagated to the surrounding code.
3. **Error Catching**: The `jaribu...bila` (try-catch) construct can catch these errors and handle them appropriately.

## Usage Examples

### Basic Error Handling

```nuru
// Try to decode invalid Base64 input
invalidInput = "This is not valid Base64!"

jaribu {
    decoded = katuaBase64(invalidInput)
    // This code won't execute if an error occurs
    andika("Decoded:", decoded)
} bila error {
    // This code will execute when an error occurs
    andika("Error:", error)
    // Take appropriate action
}
```

### Handling Different Error Types

Different errors can be handled based on the error message:

```nuru
jaribu {
    result = katuaBase64(input)
    // Process the result
} bila error {
    // Check error type by examining the error message
    kama (error.Inspect().includes("illegal base64 data")) {
        andika("Invalid Base64 format detected")
    } sivyo kama (error.Inspect().includes("inatumika na Neno pekee")) {
        andika("Type error: Input must be a string")
    } sivyo {
        andika("Unknown error:", error)
    }
}
```

### Function with Error Recovery

```nuru
// Function that safely decodes Base64, returning a default value on error
safeDecode = unda(input, defaultValue) {
    jaribu {
        return katuaBase64(input)
    } bila error {
        andika("Warning: Decoding failed:", error)
        return defaultValue
    }
}

// Usage
result = safeDecode(potentiallyBadInput, "Default value if decoding fails")
```

## Best Practices

1. **Always Use Try-Catch**: When working with untrusted input data, always wrap Base64 operations in try-catch blocks.

2. **Provide Clear Error Messages**: When handling errors, provide clear feedback to users about what went wrong.

3. **Fallback Strategies**: Implement appropriate fallback strategies when Base64 operations fail, such as:
   - Using default values
   - Requesting the user to re-enter data
   - Logging the error for debugging

4. **Validate Input Before Decoding**: When possible, validate Base64 input before attempting to decode:
   ```nuru
   // Simple Base64 validation (checks length and allowed characters)
   isValidBase64 = unda(str) {
       // Base64 length should be a multiple of 4 (with possible padding)
       kama (str.length % 4 != 0) {
           rudisha sikweli
       }
       // Check for valid Base64 characters (A-Z, a-z, 0-9, +, /, =)
       rudisha !str.match(/[^A-Za-z0-9+/=]/)
   }
   ```

By properly integrating Base64 operations with try-catch error handling, Nuru programs can handle encoding and decoding failures gracefully, providing a better user experience and more robust applications. 