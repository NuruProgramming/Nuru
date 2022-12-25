## DICTIONARIES (KAMUSI)

### DEFINITION

Dictionaries are enclosed by curly braces `{}` and have keys and values. You can define a dictionary as follows:
```
fanya k = {"jina": "juma"}
```
- The `keys` can be `string, int, float` and `boolean`
- The `values` can be of any type; `string, int, float, boolean, null` and even a `function`:
```
fanya k = {
	"jina": "juma",
	"umri": 2,
	kweli : "true",
	"mi ni function": unda(x){andika("habari", x)}
	"sina value": tupu
}

andika(k["sina value"]) // tupu
```

### Accessing Elements

You can access individual elements as follows:
```
andika(k[kweli]) // true

andika(k["mi ni function"]("juma")) // habari juma
```

### Updating Elements
You can update the value of an element as follows:
```
k['umri'] = 50

andika(k['umri']) // 50
```

### Adding New Elements
If a key-value pair doesn't exist, you can add one as follows:
```
k["I am new"] = "new element"

andika(k["I am new"]) // new element
```

### Concatenating Dictionaries

You can add two dictionaries as follows:
```
fanya a = {"a": "andazi"}
fanya b = {"b": "bunduki"}
fanya c = a + b

andika(c) // {"a": "andazi", "b": "bunduki"}
```

### Checking If Key Exists In A Dictionary

Use the `ktk` keyword to check if a key exists:
```
"umri" ktk k // kweli
"ubini" ktk k // sikweli
```

### Looping Over A Dictionary

- You can loop over a dictionary as follows:

```go
fanya k = {"a": "afya", "b": "buibui", "c": "chapa"}
kwa i, v ktk k {
	andika(i, "=>", v)
}
/* a => afya
   b => buibui
   c => chapa */
```

- You can also loop over just values as follows:

```
kwa v ktk k {
    andika(v)
}

/*
afya
buibui
chapa
*/
```

**Please Note**
> A lot more dict methods will be added in the future