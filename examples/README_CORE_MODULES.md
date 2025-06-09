# Nuru Core Modules: HTTP, URL, Path, File System, and JSON

This document provides comprehensive documentation for the core modules developed for the Nuru programming language. These modules are designed to follow the Node.js architecture while utilizing Swahili naming conventions and syntax.

## Table of Contents

1. [HTTP Module](#http-module)
2. [URL Module](#url-module)
3. [Path Module](#path-module)
4. [File System Module](#file-system-module)
5. [JSON Module](#json-module)
6. [Base64 Module](#base64-module)
7. [Integration Examples](#integration-examples)
8. [Best Practices](#best-practices)

## HTTP Module

The HTTP module (`http`) provides functionality for creating HTTP servers and clients, similar to Node.js's HTTP module. It allows you to handle HTTP requests and responses, configure HTTP agents, and work with HTTP status codes.

### Key Components

| Name | Description |
|------|-------------|
| `tengenezaServer(callback)` | Creates an HTTP server that listens for requests |
| `sikiliza(port, host?, callback?)` | Makes the server listen on a specific port and host |
| `tengenezaAgent(options?)` | Creates an HTTP agent with custom options |
| `ombaMtandao(options, callback?)` | Makes an HTTP request and returns a response |

### Example: Creating an HTTP Server

```nuru
ita http kutoka "http"

// Create a server
seva = http.tengenezaServer(fanya(ombi, jibu) {
    jibu.andika("Hello, World!")
    jibu.mwisho()
})

// Listen on port 8080
seva.sikiliza(8080, fanya() {
    andika("Server running at http://localhost:8080/")
})
```

### Example: Making an HTTP Request

```nuru
ita http kutoka "http"

// Options for the request
chaguo = {
    host: "example.com",
    port: 80,
    path: "/",
    method: "GET"
}

// Make the request
ombi = http.ombaMtandao(chaguo, fanya(jibu) {
    jibu.kwenye("data", fanya(data) {
        andika(data)
    })
})

ombi.mwisho()
```

## URL Module

The URL module (`url`) provides utilities for URL resolution and parsing, similar to Node.js's URL module. It allows you to parse URLs into their components, manipulate query strings, and construct URLs.

### Key Components

| Name | Description |
|------|-------------|
| `changanua(url, parseQueryString?)` | Parses a URL string into its components |
| `tengeneza(url, base?)` | Creates a new URL object |
| `tengenezaURLSearchParams(input)` | Creates a new URLSearchParams object for query string manipulation |
| `format(urlObject)` | Formats a URL object into a URL string |
| `tatua(from, to)` | Resolves a relative URL against a base URL |

### Example: Parsing a URL

```nuru
ita url kutoka "url"

// Parse a URL
parsed = url.changanua("https://user:pass@example.com:8080/path?query=string#hash")

// Result:
// {
//   protocol: "https:",
//   slashes: true,
//   auth: "user:pass",
//   host: "example.com:8080",
//   port: "8080",
//   hostname: "example.com",
//   hash: "#hash",
//   search: "?query=string",
//   query: "query=string", // or an object if parseQueryString is true
//   pathname: "/path",
//   path: "/path?query=string",
//   href: "https://user:pass@example.com:8080/path?query=string#hash"
// }
```

### Example: Working with URLSearchParams

```nuru
ita url kutoka "url"

// Create URLSearchParams object
params = url.tengenezaURLSearchParams("name=John&age=30")

// Add a parameter
params.ongeza("city", "New York")

// Get a parameter
name = params.pata("name")  // "John"

// Convert to string
str = params.toString()  // "name=John&age=30&city=New%20York"
```

## Path Module

The Path module (`njia`) provides utilities for working with file and directory paths. It follows the Node.js Path module architecture while using Swahili naming conventions.

### Key Components

| Name | Description |
|------|-------------|
| `jina(path, ext?)` | Gets the last part of a path (like Unix basename) |
| `kigawaji` | Platform-specific path delimiter (`:` for POSIX, `;` for Windows) |
| `sarafa(path)` | Gets the directory part of a path |
| `ext(path)` | Gets the file extension from a path |
| `umbiza(pathObj)` | Creates a path string from an object |
| `niKamili(path)` | Checks if a path is absolute |
| `unganisha(path...)` | Joins path segments together |
| `sawazisha(path)` | Normalizes a path, resolving `.` and `..` |
| `changanua(path)` | Splits a path into an object of components |
| `husika(from, to)` | Gets the relative path between two paths |
| `tatua(path...)` | Resolves paths into an absolute path |
| `kitenga` | Platform-specific path separator (`/` for POSIX, `\` for Windows) |

### Example: Working with Paths

```nuru
ita njia kutoka "njia"

// Join paths
full_path = njia.unganisha("/home", "user", "documents", "file.txt")
// Result: "/home/user/documents/file.txt"

// Get file name
file = njia.jina(full_path)  // "file.txt"

// Get directory
dir = njia.sarafa(full_path)  // "/home/user/documents"

// Get extension
ext = njia.ext(full_path)  // ".txt"

// Parse path
parts = njia.changanua(full_path)
// Result: {root: "/", dir: "/home/user/documents", base: "file.txt", ext: ".txt", name: "file"}
```

## File System Module

The File System module (`faili`) provides functionality for working with files and directories on disk, similar to Node.js's fs module.

### Key Components

| Name | Description |
|------|-------------|
| `soma(path)` | Reads a file and returns the contents as a string |
| `andika(path, data)` | Writes data to a file, replacing the file if it already exists |
| `ongeza(path, data)` | Appends data to the end of a file, creating the file if it doesn't exist |
| `fungua(path, mode?)` | Opens a file and returns a file descriptor |
| `funga(fd)` | Closes a file descriptor |
| `futa(path)` | Deletes a file or directory |
| `fanya(path)` | Creates a new file, returning a file descriptor |
| `ipo(path)` | Checks if a file or directory exists |
| `orodha(path)` | Reads the contents of a directory |
| `tengenezaSarafa(path, mode?)` | Creates a directory |
| `futaSarafa(path)` | Deletes an empty directory |
| `hali(path)` | Gets file or directory information |
| `niSarafa(path)` | Checks if a path is a directory |
| `niFaili(path)` | Checks if a path is a regular file |

### Example: Reading and Writing Files

```nuru
ita faili kutoka "faili"

// Writing to a file
faili.andika("hello.txt", "Hello, world!")

// Reading a file
yaliyomo = faili.soma("hello.txt")
andika(yaliyomo)  // Outputs: Hello, world!

// Checking if a file exists
kama (faili.ipo("hello.txt")) {
    andika("File exists!")
}

// Getting file information
taarifa = faili.hali("hello.txt")
andika("File size: " + taarifa.ukubwa + " bytes")
```

### Example: Working with Directories

```nuru
ita faili kutoka "faili"

// Creating a directory
faili.tengenezaSarafa("nyaraka")

// Reading directory contents
vitu = faili.orodha("nyaraka")
kwa kitu ktk vitu {
    andika(kitu.jina + " (" + (kitu.ni_sarafa ? "Directory" : "File") + ")")
}

// Creating a file in the directory
faili.andika("nyaraka/mfano.txt", "This is a test file")

// Removing a file
faili.futa("nyaraka/mfano.txt")

// Removing a directory
faili.futaSarafa("nyaraka")
```

## JSON Module

The JSON module (`jsoni`) provides functionality for encoding and decoding JSON data, similar to Python's json module. It allows you to convert between Nuru objects and JSON strings, as well as read and write JSON files.

### Key Components

| Name | Description |
|------|-------------|
| `dikodi(jsonString, options?)` | Decodes a JSON string into a Nuru object (similar to `json.loads` in Python) |
| `enkodi(obj, options?)` | Encodes a Nuru object into a JSON string (similar to `json.dumps` in Python) |
| `soma(path, options?)` | Reads a JSON file and parses its contents (similar to `json.load` in Python) |
| `hifadhi(obj, path, options?)` | Writes a Nuru object to a JSON file (similar to `json.dump` in Python) |
| `pendeza(obj, options?)` | Pretty-prints a Nuru object or JSON string with proper indentation |
| `msailiaji(options?)` | Creates a custom JSON encoder for specific formatting needs |

### Example: Converting Between Objects and JSON

```nuru
ita jsoni = pata("jsoni")

// Create a Nuru object
fanya mtu = {
    "jina": "John",
    "umri": 30,
    "mjini": kweli,
    "nambari": [1, 2, 3]
}

// Encode to JSON
fanya json_str = jsoni.enkodi(mtu)
andika(json_str)  // {"jina":"John","umri":30,"mjini":true,"nambari":[1,2,3]}

// Decode from JSON
fanya json_data = '{"jina":"Fatima","umri":25,"mjini":true}'
fanya obj = jsoni.dikodi(json_data)
andika(obj.jina)  // "Fatima"
andika(obj.umri)  // 25
```

### Example: Pretty-printing JSON

```nuru
ita jsoni = pata("jsoni")

fanya config = {
    "jina": "Student System",
    "toleo": "1.0.0",
    "mpangilio": {
        "rangi": "blue",
        "ukubwa": "medium"
    }
}

// Pretty-print JSON with indentation
fanya json_nzuri = jsoni.pendeza(config)
andika(json_nzuri)

/* Output:
{
    "jina": "Student System",
    "toleo": "1.0.0",
    "mpangilio": {
        "rangi": "blue",
        "ukubwa": "medium"
    }
}
*/
```

### Example: Reading and Writing JSON Files

```nuru
ita jsoni = pata("jsoni")

// Create an object
fanya wanafunzi = [
    {"jina": "Ali", "alama": 85},
    {"jina": "Maria", "alama": 92},
    {"jina": "John", "alama": 78}
]

// Save to a JSON file
jsoni.hifadhi(wanafunzi, "wanafunzi.json")

// Read from a JSON file
fanya data = jsoni.soma("wanafunzi.json")

// Use the data
kwa kila mwanafunzi katika data {
    andika(mwanafunzi.jina + ": " + Neno(mwanafunzi.alama))
}
```

## Base64 Module

The Base64 module provides functionality for encoding and decoding data using the Base64 encoding scheme. Base64 is commonly used to encode binary data (like images or files) into ASCII string format, making it suitable for transmission over text-based protocols.

### Key Components

| Name | Description |
|------|-------------|
| `kodeBase64(data)` | Encodes String or Byte data into a Base64 object |
| `katuaBase64(encodedString)` | Decodes a Base64 string into a Base64 object |

### Base64 Object Methods

| Method | Description |
|------|-------------|
| `kukata()` | Decodes the Base64 content and returns the raw data as a Byte object |
| `kutoka(data)` | Creates a new Base64 object from the provided String or Byte data |
| `data()` | Returns the raw (decoded) data as a Byte object |

### Example: Basic Base64 Encoding and Decoding

```nuru
// Encoding a string to Base64
neno = "Hujambo Dunia! Nuru ni lugha nzuri ya programu."
andika("Neno la asili:", neno)

// Convert string to Base64
base64 = kodeBase64(neno)
andika("Base64:", base64)
// Output: "SHVqYW1ibyBEdW5pYSEgTnVydSBuaSBsdWdoYSBuenVyaSB5YSBwcm9ncmFtdS4="

// Decoding Base64 back to original data
decoded = katuaBase64(base64)
andika("Data iliyotolewa:", decoded)
// Output: "Hujambo Dunia! Nuru ni lugha nzuri ya programu."
```

### Example: Error Handling with Try-Catch

Base64 operations can fail when given invalid input. You can use the `jaribu...bila` (try-catch) construct to handle these errors gracefully:

```nuru
// Example 1: Successful decoding with try-catch
validBase64 = "SGVsbG8sIE51cnUh"  // "Hello, Nuru!" encoded

jaribu {
    decoded = katuaBase64(validBase64)
    andika("Successfully decoded:", decoded)
} bila error {
    andika("Error during decoding:", error)
}

// Example 2: Handling invalid Base64 input
invalidBase64 = "This is not valid Base64!"

jaribu {
    result = katuaBase64(invalidBase64)
    // This line won't execute if an error occurs
    andika("Decoded result:", result)
} bila error {
    andika("Decoding failed:", error)
    // You can take alternative actions here
}
```

### Example: Working with Files and Base64

```nuru
waagiza fs kutoka "mfumo";

// Create a file to encode
failiJina := "base64_test.txt";
data := "Data ya kutest Base64 encoding na decoding.";
fs.andika(failiJina, data);

// Read the file and encode its content to Base64
failiData := fs.soma(failiJina);
fileBase64 := kodeBase64(failiData);
andika("Base64 ya faili:", fileBase64);

// Save the Base64 encoded string to a new file
fs.andika("base64_encoded.txt", fileBase64.Inspect());

// Read the encoded file and decode it
encodedData := fs.soma("base64_encoded.txt");
decodedData := katuaBase64(encodedData);
andika("Data baada ya kutoa usimbaji:", decodedData.data());
```

### Example: Encoding and Decoding Binary Data

```nuru
// Binary data represented as bytes
byteData := [104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100]; // "hello world" in ASCII

// Convert byte array to string
binData := "";
kwa b katika byteData {
    binData = binData + tungo(b);
}

// Encode binary data to Base64
binBase64 := kodeBase64(binData);
andika("Binary Base64:", binBase64);
// Output similar to: "aGVsbG8gd29ybGQ="

// Decode Base64 back to binary data
decodedBin := katuaBase64(binBase64.Inspect());
andika("Decoded binary data:", decodedBin.data());
```

## Integration Examples

### RESTful API with HTTP, URL, Path, File System, and JSON Modules

```nuru
ita http = pata("http")
ita url = pata("url")
ita njia = pata("njia")
ita faili = pata("faili")
ita jsoni = pata("jsoni")

// Create a simple REST API server
seva = http.tengenezaServer(fanya(ombi, jibu) {
    // Parse the URL from the request
    parsed_url = url.changanua(ombi.url)
    endpoint = parsed_url.pathname

    // Set JSON content type
    jibu.wekaKichwa("Content-Type", "application/json")

    // Create router for different endpoints
    kama (endpoint == "/api/users" && ombi.njia == "GET") {
        // Check if users.json exists
        kama (faili.ipo("data/users.json")) {
            // Read users from file
            data = jsoni.soma("data/users.json")
            jibu.mwisho(jsoni.enkodi(data))
        } sivyo {
            // Return empty array if file doesn't exist
            jibu.mwisho(jsoni.enkodi([]))
        }
    } sivyo kama (endpoint == "/api/users" && ombi.njia == "POST") {
        // Get request body
        ombi.kwenyeData(fanya(data) {
            // Parse JSON data
            user = jsoni.dikodi(data)
            
            // Ensure data directory exists
            kama (!faili.ipo("data")) {
                faili.tengenezaSarafa("data")
            }
            
            // Read existing users or create empty array
            users = []
            kama (faili.ipo("data/users.json")) {
                users = jsoni.soma("data/users.json")
            }
            
            // Add new user and save to file
            users.add(user)
            jsoni.hifadhi(users, "data/users.json")
            
            // Return success response
            jibu.andikaKichwa(201)
            jibu.mwisho(jsoni.enkodi({success: true}))
        })
    } sivyo {
        // Handle 404 Not Found
        jibu.andikaKichwa(404)
        jibu.mwisho(jsoni.enkodi({error: "Endpoint not found"}))
    }
})

// Start the server
seva.sikiliza(8080, fanya() {
    andika("REST API server running at http://localhost:8080/")
})
```

### Image Transfer with HTTP, File System, and Base64 Modules

```nuru
ita http = pata("http")
ita fs = pata("faili")

// Create a server that can receive and serve base64-encoded images
seva = http.tengenezaServer(fanya(ombi, jibu) {
    parsed_url = url.changanua(ombi.url)
    endpoint = parsed_url.pathname
    
    // Handle image upload endpoint
    kama (endpoint == "/upload" && ombi.njia == "POST") {
        jaribu {
            // Receive the request body data
            ombi.kwenyeData(fanya(data) {
                // Parse JSON data
                json_data = jsoni.dikodi(data)
                
                // Get the base64 string from the request
                base64_string = json_data.image
                
                // Decode the base64 data
                image_data = katuaBase64(base64_string)
                
                // Make sure the uploads directory exists
                kama (!fs.ipo("uploads")) {
                    fs.tengenezaSarafa("uploads")
                }
                
                // Save the decoded image to a file
                fs.andika("uploads/" + json_data.filename, image_data.data().Inspect())
                
                // Respond with success
                jibu.wekaKichwa("Content-Type", "application/json")
                jibu.mwisho(jsoni.enkodi({
                    success: kweli,
                    message: "Image uploaded successfully",
                    path: "/images/" + json_data.filename
                }))
            })
        } bila kosa {
            // Handle errors
            jibu.andikaKichwa(500)
            jibu.wekaKichwa("Content-Type", "application/json")
            jibu.mwisho(jsoni.enkodi({
                success: sikweli,
                message: "Failed to upload image: " + kosa
            }))
        }
    }
    // Handle image serving endpoint
    sivyo kama (endpoint.startsWith("/images/")) {
        jaribu {
            // Extract filename from path
            filename = endpoint.substring(8)  // Remove '/images/' prefix
            image_path = "uploads/" + filename
            
            // Check if the file exists
            kama (fs.ipo(image_path)) {
                // Read the file content
                image_data = fs.soma(image_path)
                
                // Set appropriate headers
                jibu.wekaKichwa("Content-Type", "image/jpeg")  // Adjust as needed for different formats
                
                // Send the image data
                jibu.mwisho(image_data)
            } sivyo {
                // Image not found
                jibu.andikaKichwa(404)
                jibu.mwisho("Image not found")
            }
        } bila kosa {
            // Handle errors
            jibu.andikaKichwa(500)
            jibu.mwisho("Server error: " + kosa)
        }
    }
    sivyo {
        // Serve a simple HTML form for uploading images
        kama (endpoint == "/" && ombi.njia == "GET") {
            jibu.wekaKichwa("Content-Type", "text/html")
            jibu.mwisho(`
                <html>
                <head><title>Nuru Image Uploader</title></head>
                <body>
                    <h1>Upload an Image</h1>
                    <input type="file" id="imageInput" />
                    <button onclick="uploadImage()">Upload</button>
                    <div id="result"></div>
                    
                    <script>
                    function uploadImage() {
                        const file = document.getElementById('imageInput').files[0];
                        if (!file) {
                            alert('Please select a file first');
                            return;
                        }
                        
                        const reader = new FileReader();
                        reader.onload = function(e) {
                            // Get base64 string (remove data URL prefix)
                            const base64 = e.target.result.split(',')[1];
                            
                            // Send to server
                            fetch('/upload', {
                                method: 'POST',
                                headers: { 'Content-Type': 'application/json' },
                                body: JSON.stringify({
                                    filename: file.name,
                                    image: base64
                                })
                            })
                            .then(response => response.json())
                            .then(data => {
                                if (data.success) {
                                    document.getElementById('result').innerHTML = 
                                        `<p>Upload successful!</p>
                                         <img src="${data.path}" width="300" />`;
                                } else {
                                    document.getElementById('result').innerHTML = 
                                        `<p>Error: ${data.message}</p>`;
                                }
                            });
                        };
                        reader.readAsDataURL(file);
                    }
                    </script>
                </body>
                </html>
            `)
        } sivyo {
            // Handle 404 Not Found
            jibu.andikaKichwa(404)
            jibu.mwisho("Not Found")
        }
    }
})

// Start the server
seva.sikiliza(8080, fanya() {
    andika("Image server running at http://localhost:8080/")
})
```

### Config Parser with File System, JSON, and Base64 Modules

```nuru
ita faili = pata("faili")
ita jsoni = pata("jsoni")

// Function to load and parse a config file
fanya loadConfig = fanya(path) {
    // Check if the file exists
    kama (!faili.ipo(path)) {
        andika("Config file not found at: " + path)
        rudisha {}
    }
    
    jaribu {
        // Read and parse the config file
        config = jsoni.soma(path)
        andika("Loaded configuration from: " + path)
        rudisha config
    } bila {
        andika("Error loading config: " + hii)
        rudisha {}
    }
}

// Function to save a config file
fanya saveConfig = fanya(config, path) {
    jaribu {
        // Create directory if it doesn't exist
        dir = njia.sarafa(path)
        kama (!faili.ipo(dir)) {
            faili.tengenezaSarafa(dir, {"rekesia": kweli})
        }
        
        // Save the config with pretty formatting
        jsoni.hifadhi(config, path, {"indent": 2})
        andika("Saved configuration to: " + path)
        rudisha kweli
    } bila {
        andika("Error saving config: " + hii)
        rudisha sikweli
    }
}

// Example usage
config = loadConfig("config/app.json")

// Add or modify config values
config.name = "My Application"
config.version = "1.0.1"
config.settings = {
    theme: "dark",
    language: "sw",
    notifications: kweli
}

// Save updated config
saveConfig(config, "config/app.json")
```

## Best Practices

### Error Handling

Always handle errors when working with file systems, HTTP requests, and JSON parsing:

```nuru
// File system error handling
jaribu {
    data = faili.soma("config.json")
} bila kosa {
    andika("Kosa katika kusoma faili:", kosa)
    // Provide fallback or exit gracefully
}

// HTTP request error handling
ombi = http.ombaMtandao(chaguo, fanya(jibu) {
    jibu.kwenye("error", fanya(kosa) {
        andika("Kosa la ombi:", kosa)
    })
})

// JSON parsing error handling
jaribu {
    obj = jsoni.dikodi(data)
} bila kosa {
    andika("JSON imekosewa:", kosa)
}
```

### Resource Cleanup

Always close file descriptors and server connections when no longer needed:

```nuru
// Close file descriptors
fd = faili.fungua("file.txt", "r")
// ... use the file
faili.funga(fd)

// Close server connections
seva = http.tengenezaServer(...)
// ... use the server
seva.funga()
```

### URL Parsing

Validate and sanitize URLs before using them:

```nuru
// Validate URL
parsedUrl = url.changanua(userInput)
kama (!parsedUrl.host || !parsedUrl.protocol) {
    // URL is invalid or incomplete
    andika("URL si sahihi")
    rudi
}
```

### Path Manipulation

Normalize and validate paths to prevent directory traversal attacks:

```nuru
// Normalize path
userPath = njia.sawazisha(userInput)

// Validate path is within allowed directory
rootDir = "/allowed/path"
fullPath = njia.unganisha(rootDir, userPath)
kama (!fullPath.startsWith(rootDir)) {
    andika("Samahani, njia sio sahihi")
    rudi
}
```

### File System Security

Restrict file operations to specific directories and validate file types:

```nuru
// Check file extension
filename = userInput
kama (!filename.endsWith(".txt") && !filename.endsWith(".log")) {
    andika("Samahani, aina ya faili haijakubaliwa")
    rudi
}
```

### JSON Data Handling

Validate JSON structure and sanitize data from untrusted sources:

```nuru
// Validate JSON structure
jaribu {
    data = jsoni.dikodi(userJson)
    // Check for required fields
    kama (!data.username || !data.email) {
        andika("Data haijatosheleza")
        rudi
    }
    
    // Sanitize data
    data.username = data.username.replace(/[<>]/g, "")
} bila kosa {
    andika("JSON imekosewa:", kosa)
}

// Use pretty printing for logging
prettyJson = jsoni.pendeza(data)
andika("Data:", prettyJson)

// Use compact encoding for network transmission
compactJson = jsoni.enkodi(data)
http.ombi.mwisho(compactJson)
```

### Base64 Data Handling

Use Base64 encoding properly for binary data and sensitive information:

```nuru
// Encode binary data
imageData = faili.soma("picha.jpg", "binary")
base64Data = kodeBase64(imageData)

// Use Base64 for storing binary data in JSON
jsonObj = {
    jina: "picha.jpg",
    data: base64Data.Inspect(),
    aina: "image/jpeg"
}
jsoni.hifadhi(jsonObj, "picha_info.json")

// Validate Base64 input from untrusted sources
jaribu {
    input = userInput
    // Check if it looks like Base64 (simplified check)
    kama (input.length % 4 != 0 || /[^A-Za-z0-9+/=]/.test(input)) {
        andika("Samahani, Base64 sio sahihi")
        rudi
    }
    decodedData = katuaBase64(input)
} bila kosa {
    andika("Samahani, siwezi kutafsiri Base64:", kosa)
}

// Don't use Base64 for encryption (it's encoding, not encryption)
// This is bad practice:
password = "siri123"
encodedPassword = kodeBase64(password)  // Don't do this for security!

// Instead, use a proper hashing/encryption library
// (This would be part of a dedicated security module)
```

### Combining Modules

Leverage the interoperability between modules for complex operations:

```nuru
// Create a JSON configuration file with paths
configData = {
    dataPath: njia.unganisha(process.cwd(), "data"),
    tempPath: njia.unganisha(process.cwd(), "temp"),
    allowedTypes: ["txt", "json", "csv"]
}

// Save with pretty-printing
jsoni.hifadhi(configData, "config.json", { pendeza: kweli })

// Upload a Base64-encoded image to a server
imageData = faili.soma("picha.jpg")
base64Image = kodeBase64(imageData)

http.ombaMtandao({
    method: "POST",
    host: "example.com",
    path: "/upload",
    headers: {
        "Content-Type": "application/json"
    }
}, fanya(jibu) {
    // Send the Base64 data
    jibu.mwisho(jsoni.enkodi({
        jina: "picha.jpg",
        data: base64Image.Inspect()
    }))
})
```

By following these best practices, you'll create more robust, secure, and maintainable applications with Nuru's core modules.

For more details on each module, refer to the individual module documentation:
- [HTTP Module Documentation](README_HTTP.md)
- [URL Module Documentation](README_URL.md)
- [Path Module Documentation](README_PATH.md)
- [File System Documentation](README_FS.md) 