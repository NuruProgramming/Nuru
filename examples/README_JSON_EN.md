# JSON Module for Nuru

The JSON module is used to convert data between Nuru objects and JSON format. This module allows reading and writing JSON data, which is essential for network operations and data storage.

## Module Registration

```nuru
ita jsoni = pata("jsoni")  // Get the JSON module
```

## Module Components

| Name | Description |
|------|-------------|
| `dikodi` | Decode JSON string to Nuru object (similar to `json.loads` in Python) |
| `enkodi` | Encode Nuru object to JSON string (similar to `json.dumps` in Python) |
| `soma` | Load JSON from file (similar to `json.load` in Python) |
| `hifadhi` | Save JSON to file (similar to `json.dump` in Python) |
| `pendeza` | Pretty print JSON with proper indentation |
| `msailiaji` | Create a JSON encoder object (similar to `JSONEncoder` in Python) |
| `msailiaji_bora` | Enhanced JSON encoder with additional features |

## Usage Examples

### 1. Converting Objects to JSON

```nuru
ita jsoni = pata("jsoni")

// Create a Nuru object
fanya person = {
    "name": "John",
    "age": 30,
    "isUrban": true,
    "children": ["Asha", "Bakari"],
    "phones": {
        "home": "0222334455",
        "work": "0777889900"
    }
}

// Convert to JSON
fanya json_string = jsoni.enkodi(person)
andika(json_string)

// Output: {"name":"John","age":30,"isUrban":true,"children":["Asha","Bakari"],"phones":{"home":"0222334455","work":"0777889900"}}
```

### 2. Converting JSON Strings to Objects

```nuru
ita jsoni = pata("jsoni")

// JSON string
fanya json_data = '{"name":"Fatima","age":25,"isUrban":true,"numbers":[1,2,3]}'

// Convert to Nuru object
fanya person = jsoni.dikodi(json_data)

// Use the object
andika(person.name)        // "Fatima"
andika(person.age)         // 25
andika(person.isUrban)     // true
andika(person.numbers[1])  // 2
```

### 3. Reading and Writing JSON Files

```nuru
ita jsoni = pata("jsoni")

// Create an object
fanya students = [
    {"name": "Ali", "score": 85},
    {"name": "Maria", "score": 92},
    {"name": "John", "score": 78}
]

// Save to file
jsoni.hifadhi(students, "students.json")

// Read from file
fanya data = jsoni.soma("students.json")

// Use the data
kwa kila student katika data {
    andika(student.name + ": " + Neno(student.score))
}
```

### 4. Using Pretty Printing

```nuru
ita jsoni = pata("jsoni")

fanya config = {
    "name": "Student Management System",
    "version": "1.0.0",
    "settings": {
        "color": "blue",
        "size": "medium",
        "language": "Swahili"
    }
}

// Pretty print the JSON
fanya pretty_json = jsoni.pendeza(config)
andika(pretty_json)

/* Output might be:
{
    "name": "Student Management System",
    "version": "1.0.0",
    "settings": {
        "color": "blue",
        "size": "medium",
        "language": "Swahili"
    }
}
*/
```

## Decode Options

You can provide additional options to `dikodi`:

- `parse_float`: Function to parse float numbers
- `parse_int`: Function to parse integer numbers
- `object_hook`: Function to transform objects after parsing

## Encode Options

You can provide additional options to `enkodi`:

- `skipkeys`: If `true`, keys that are not strings will be skipped
- `ensure_ascii`: If `true`, all characters will be written in ASCII
- `indent`: Number of spaces to use for indentation
- `sort_keys`: If `true`, dictionary keys will be sorted alphabetically
- `separators`: Different separators for items and key-value pairs (e.g., [",", ":"])

## Custom JSONEncoder

You can create a custom JSONEncoder:

```nuru
ita jsoni = pata("jsoni")

// Create encoder
fanya encoder = jsoni.msailiaji({
    "indent": 2,
    "sort_keys": true
})

// Use the encoder
fanya obj = {"c": 3, "a": 1, "b": 2}
fanya pretty_json = encoder.enkodi(obj)
andika(pretty_json)

/* Output:
{
  "a": 1,
  "b": 2,
  "c": 3
}
*/
```

## Notes

1. The JSON module requires the `faili` (file) module for the `soma` and `hifadhi` functions.
2. Nuru types are converted to JSON as follows:
   - `Dict` → JSON object
   - `Array` → JSON array
   - `String` → JSON string
   - `Integer/Float` → JSON number
   - `Boolean` → JSON boolean
   - `Null` → JSON null

3. JSON types are converted to Nuru as follows:
   - JSON object → `Dict`
   - JSON array → `Array`
   - JSON string → `String`
   - JSON number → `Integer` or `Float`
   - JSON boolean → `Boolean`
   - JSON null → `Null` 