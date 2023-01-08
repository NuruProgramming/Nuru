## ARRAYS (ORODHA)

### Definition

Arrays are enclosed in square brackets `[]` and they can hold any type, even function definitions:
```go
fanya arr = [1, "mambo", kweli, unda(x, y){rudisha x + y}, 2 * 3 + 20]

andika(arr)

/*
[1, mambo, kweli, unda(x, y) {rudisha (x + y);}, 26]
*/
```

### Accessing Elements

You can access individual elements through indexes starting from zero:
```go
fanya herufi = ["a", "b", "c"]

andika(herufi[0]) // a
```

### Reassigning Elements

You can also reassign values in elements:
```go
fanya herufi = ["a", "b", "c"]

herufi[1] = "z"

andika(herufi) // ["a", "z", "c"]
```

### Looping over an Array

- You can also iterate through an array:
```go
fanya herufi = ["a", "b", "c"]

kwa i ktk herufi {
	andika(i)
}
/* a
   b
   c  */
```

- And for a key, value pair:
```go
kwa i, v ktk herufi {
	andika(i, "=>", v)
}

/* 0 => a
   1 => b
   2 => c */
```

### Check if an Element exists

You can also check if elements exist in an array:
```go
andika("d" ktk herufi) // sikweli
andika("a" ktk herufi) // kweli
```

### Concatenating Arrays

- You can also add two arrays as follows:
```
fanya h1 = ["a", "b", "c"]
fanya h2 = [1, 2, 3]
fanya h3 = h1 + h2

andika(h3) // ["a", "b", "c", 1, 2, 3]

h2 += h3

andika(h2) // [1, 2, 3, "a", "b", "c", 1, 2, 3]
```

- You can also multiply an array as follows:
```
fanya a = [1, 2, 3]

andika(a * 2) // [1, 2, 3, 1, 2, 3]
```

### Length of an Array

You can get the length of an array with `idadi`:
```
fanya a = ["a", "b", "c"]

andika(idadi(a)) // 3
```

### Adding Elements to an Array

You can add new elements to an array with `sukuma`:
```go
fanya a = [1, 2, 3]

// you must reassign for the new value to be saved
a = sukuma(a, "mambo")

andika(a) // [1, 2, 3, "mambo"]
```

### Getting the Last Element in an Array

You can get the last element of an array with `yamwisho`:
```
fanya a = [1, 2, 3]

andika(yamwisho(a)) // 3
```
**Please Note**
> A lot more array methods will be added in the future