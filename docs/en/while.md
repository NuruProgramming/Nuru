# WHILE (WAKATI)

While loops in Nuru are used to execute a block of code repeatedly, as long as a given condition is true. This page covers the basics of while loops, including how to use the break and continue keywords within them.

## Basic Syntax

A while loop is executed when a specified condition is true. You initiliaze a while loop with the `wakati` keyword followed by the condition in paranthesis  `()`. The consequence of the loop should be enclosed in brackets `{}`:
```s
fanya i = 1

wakati (i <= 5) {
	andika(i)
	i++
}
```
Output
```s
1
2
3
4
5
```

## Break (vunja) and Continue (endelea)
### Break (Vunja)

Use the vunja keyword to terminate a loop:

```s
fanya i = 1

wakati (i < 5) {
	kama (i == 3) {
		andika("nimevunja")
		vunja
	}
	andika(i)
	i++
}
```
Output
```s
1
2
nimevunja
```

### Continue (Endelea)

Use the endelea keyword to skip a specific iteration:

```s
fanya i = 0

wakati (i < 5) {
	i++
	kama (i == 3) {
		andika("nimeruka")
		endelea
	}
	andika(i)
}
```
Output
```s
1
2
nimeruka
4
5
```

By understanding while loops in Nuru, you can create code that repeats a specific action or checks for certain conditions, offering more flexibility and control over your code execution.