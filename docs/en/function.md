## FUNCTIONS (UNDA)

### Definition

A function block starts with the `unda` keyword, parameters are surrounded by `()` and the body by `{}`. They must also be assigned to a variable as follows:
```
fanya jum = unda(x, y) {
	rudisha x + y
}

jum(2, 3) // 5
```

### Parameters

Functions can have zero or any number of arguments. Arguments can be of any type, even other functions:
```
fanya salamu = unda() {
	andika("Habari yako")
}

salamu()

salamu = unda(jina) {
	andika("Habari yako", jina)
}

salamu(asha) // Habari yako asha
```

### Return (rudisha)

You can return items with the `rudisha` keyword. The `rudisha` keyword will terminate the block and return the value:
```
fanya mfano = unda(x) {
	rudisha "nimerudi"
	andika(x)
}

mfano(x) // nimerudi
```

### Recursion

Nuru also supports recursion. Here's an example:
```
fanya fib = unda(n) {
    kama (n < 3) {
        rudisha 1
    } sivyo {
        rudisha fib(n-1) + fib(n-2)
    }
}

andika(fib(10)) // 55