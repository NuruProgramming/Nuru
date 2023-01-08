## FOR (KWA)

### Definition

For is used to iterate over an iterable object. An iterable object is a `string`, `array` or `dictionaries`. You use the `kwa` keyword followed by an identifier such as `k` or `v` followed by an iterable. The iterable block must be enclosed in a bracket `{}`. Here's an example:
```
fanya jina = "lugano"

kwa i ktk jina {
	andika(i)
}

/*
l
u
g
a
n
o
*/
```

### Key Value Pairs

Nuru allows you to get both the value or the key/value pair of an iterable. To get only the value, use one temporary identifier as such:
```
fanya kamusi = {"a": "andaa", "b": "baba"}

kwa v ktk kamusi {
	andika(v)
}

/*
andaa
baba
*/
```
To get both the key and the value, use two temporary identifiers:
```
kwa k, v ktk kamusi {
	andika(k + " ni + " v)
}

/*
a ni andaa
b ni baba
*/
```
- Note that key-value pair iteration also works for `strings` and `lists`:
```
kwa i, v ktk "mojo" {
	andika(i, "->", v)
}
/*
0 -> m
1 -> o
2 -> j
3 -> o
*/
fanya majina = ["juma", "asha", "haruna"]

kwa i, v ktk majina {
	andika(i, "-", v)
}

/*
0 - juma
1 - asha
2 - haruna
*/
```

### Break (Vunja) and Continue (Endelea)

- A loop can be terminated using the `vunja` keyword:
```
kwa i, v ktk "mojo" {
        kama (i == 2) {
                andika("nimevunja")
                vunja
        }
        andika(v)
}
/*
m
o
nimevunja
*/
```

- A specific iteration can be skipped using the `endelea` keyword:
```
kwa i, v ktk "mojo" {
        kama (i == 2) {
                andika("nimeruka")
                endelea
        }
        andika(v)
}

/*
m
o
nimeruka
o
*/
```

**CAUTION**
> In nested loops, the `vunja` and `endelea` keyword MIGHT misbehave