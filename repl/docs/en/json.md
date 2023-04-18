# JSON in Nuru

Nuru also makes it easy to deal with JSON.

## Import JSONI

Use the following to import the json module:
```
tumia jsoni
```

## Decoding JSON with dikodi()
Use this to convert a string to a dictionary:
```
jsonString = '{                                                                                                            
    "error": false,                                                                                          
    "category": "Pun",                                                                                       
    "type": "single",                                                                                        
    "joke": "I was reading a great book about an immortal dog the other day. It was impossible to put down."
}'

// to make it a dict

tumia jsoni

k = jsoni.dikodi(jsonString)

k["joke"] // I was reading a great book about an immortal dog the other day. It was impossible to put down.
```

## Encoding JSON with enkodi()

You can encode JSON with the `enkodi` method, this will turn a dictionary to a string:
```
tumia jsoni

k = {
        "a": "apple",
        "b": "banana"
    }

j = json.enkodi(k)
```
