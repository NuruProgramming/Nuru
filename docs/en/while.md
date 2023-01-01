## WHILE (WAKATI)

### Definition

A while loop is executed when a specified condition is true. You initiliaze a while loop with the `wakati` keyword followed by the condition in paranthesis  `()`. The consequence of the loop should be enclosed in brackets `{}`:
```
fanya i = 1

wakati (i <= 5) {
	andika(i)
	i++
}
/*
1
2
3
4
5
*/
```

### Break (vunja) and Continue (endelea)

- A loop can be terminated using the `vunja` keyword:
```
fanya i = 1

wakati (i < 5) {
	kama (i == 3) {
		andika("nimevunja")
		vunja
	}
	andika(i)
	i++
}
/*
1
2
nimevunja
*/
```

- A specific iteration can be skipped using the `endelea` keyword:
```
fanya i = 0

wakati (i < 5) {
	i++
	kama (i == 3) {
		andika("nimeruka")
		endelea
	}
	andika(i)
}
/*
1
2
nimeruka
4
5
*/
```

**CAUTION**
> In nested loops, the `vunja` and `endelea` keyword MIGHT misbehave
