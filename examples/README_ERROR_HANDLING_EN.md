# Error Handling in Nuru

Nuru language provides a mechanism for handling errors through the `jaribu...bila` (try-catch) construct. This feature helps manage errors in a structured way rather than allowing the program to crash.

## Try-Catch Basics

The basic structure of a `jaribu...bila` (try-catch) statement is:

```nuru
jaribu {
    // Code that might cause an error
} bila {
    // Code to handle the error
}
```

If you want to access the error details, you can use an identifier:

```nuru
jaribu {
    // Code that might cause an error
} bila error {
    // Code to handle the error, using the 'error' identifier
    andika("An error occurred: " + error)
}
```

## How It Works

1. The code inside the `jaribu` (try) block is executed.
2. If no error occurs, the program continues after the `jaribu...bila` block.
3. If an error occurs within the `jaribu` block, execution stops and the program enters the `bila` (catch) block.
4. If an identifier is provided (like `error` in the example), the error that occurred is assigned to that identifier.
5. After executing the `bila` block, the program continues with the code that follows.

## Examples

### Example 1: Handling a non-existent file error

```nuru
ita faili = pata("faili")

jaribu {
    yaliyomo = faili.soma("non_existent_file.txt")
    andika("Contents:", yaliyomo)  // This won't be reached if the file doesn't exist
} bila error {
    andika("An error occurred while reading the file: " + error)
}
```

### Example 2: Handling errors in functions

```nuru
// Function that might cause an error
gawanya = unda(a, b) {
    kama (b == 0) {
        rudisha newError("Cannot divide by zero")
    }
    rudisha a / b
}

// Using try-catch to manage the error
calculateDivision = unda(a, b) {
    jaribu {
        result = gawanya(a, b)
        andika("The result is: " + result)
    } bila problem {
        andika("Math problem: " + problem)
        // You can return an alternative value
        rudisha 0
    }
}
```

### Example 3: Handling different types of errors

```nuru
ita faili = pata("faili")

readAndParseJSON = unda(filename) {
    jaribu {
        content = faili.soma(filename)
        ita json = pata("json")
        
        // Try to parse JSON - this could cause another error
        jaribu {
            data = json.dikodi(content)
            rudisha data
        } bila json_error {
            andika("JSON error: " + json_error)
            rudisha {}  // Return empty object if JSON is invalid
        }
    } bila file_error {
        andika("File reading error: " + file_error)
        rudisha {}  // Return empty object if file can't be read
    }
}
```

## Tips and Best Practices

1. **Use Appropriately:** Use `jaribu...bila` for operations that might cause errors, such as file I/O, network operations, or user-dependent operations.

2. **Meaningful Identifiers:** Provide meaningful identifiers in the `bila` block, such as `error`, `problem`, or a name that describes the type of error.

3. **Don't Overdo It:** Don't place too much code inside a `jaribu` block. Do just one operation that you know might cause an error.

4. **Propagate Errors:** Sometimes, instead of handling an error, you might want to return an error to the calling code:

   ```nuru
   readSafely = unda(name) {
       jaribu {
           rudisha faili.soma(name)
       } bila error {
           // You can return a different error or the same one
           rudisha newError("Couldn't read " + name + ": " + error)
       }
   }
   ```

5. **Clean Up Resources:** Make sure to release resources like file handles in both the `bila` and `jaribu` blocks:

   ```nuru
   jaribu {
       f = faili.fungua("data.txt", "r")
       text = f.soma()
       f.funga()  // Close the file before returning
       rudisha text
   } bila error {
       // Ensure the file is closed even if an error occurred
       andika("Error: " + error)
       jaribu {
           f.funga()
       } bila {
           // Ignore any errors on closing
       }
   }
   ```

By using the `jaribu...bila` system, you can write programs that handle errors gracefully and provide good feedback to users when things don't go as expected. 