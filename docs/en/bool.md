## BOOLEANS

Boolean objects are `truthy`, meaning any value is true. A value is false only when its null (ie. `tupu`) or false (ie `sikweli`):
### Example 1
```
fanya x = 0

kama (x) {
	andika("I am true")
} sivyo {
	andika("I am not true")
}

// it will print "I am true"
```

### Example 2
```
kama (tupu) {
andika("I am true")
} sivyo {
	andika("I am not true")
}

// will print "I am not true"
```

Expressions can also be evaluated to true or false:
```
andika(1 > 2) // sikweli

andika(1 + 3 < 10) // kweli