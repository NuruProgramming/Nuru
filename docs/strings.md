## STRINGS (NENO)

### Definition
 
Strings can be enclosed in either a single quote `''` or double quotes `""`:

```
andika("mambo") // mambo

fanya a = 'niaje'

andika("mambo", a) // mambo niaje
```

### Concatenation
 
- Strings can also be concatenated as follows:

```
fanya a = "habari" + " " + "yako"

andika(a) // habari yako

fanya b = "habari"

b += " yako" 

// habari yako
```

- You can also multiply a string `n` number of times:

```
andika("mambo " * 4)

// mambo mambo mambo mambo

fanya a = "habari"

a *= 4

// habarihabarihabarihabari
```

### Looping over a String
 
- You can loop through a string as follows

```
fanya jina = "avicenna"

kwa i ktk jina {andika(i)}

/*  
    a
	v
	i
	c
	e
	n
	n
	a  
*/
```

- And for key, value pairs:
```go
kwa i, v ktk jina {
	andika(i, "=>", v)
}
/*
0 => a
1 => v
2 => i
3 => c
4 => e
5 => n
6 => n
7 => a
*/
```

### Comparing Strings

- You can also check if two strings are the same:
```
fanya a = "nuru"

andika(a == "nuru") // kweli

andika(a == "mambo") // sikweli
```

### Length of a String

You can also check the length of a string with the `idadi` function
```
fanya a = "mambo"

idadi(a) // 5
```

**Please Note**
> A lot more string methods will be added in the future