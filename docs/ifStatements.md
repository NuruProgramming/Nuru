## IF/ELSE (KAMA/SIVYO)

### Definition

You initiliaze an if block with `kama`, the condition must be inside a paranthesis `()` and the consequence inside a `{}`:
```
kama (2>1) {
	andika(kweli) // kweli
}
```

### Else Block

- For multiple conditions, you can use `kama` ,  `au kama` and `sivyo`:
```
fanya a = 10

kama (a > 100) {
	andika("a imezidi 100")
} au kama (a < 10) {
	andika("a ndogo kuliko 10")
} sivyo {
	andika("Thamani ya a ni", a)
}

// it will print 'Thamani ya a ni 10'
```