## INTEGERS (NAMBA) AND FLOATS (DESIMALI)

### PRECEDENCE

Integers and floats work the way you'd expect them to. They precedence in mathematical operations follow the BODMAS rule:

```go
2 + 3 * 5 // 17

fanya a = 2.5
fanya b = 3/5

a + b // 2.8
```

### UNARY INCREMENTS

You can perform unary increments (++ and --) on both floats and integers. These will add or subtract 1 from the current value. Note that the float or int have to be assigned to a variable for this operation to work. Here's an example:

```go
fanya i = 2.4

i++ // 3.4
```

### SHORTHAND ASSIGNMENT

You can also perform shorthand assignments with `+=`, `-=`, `/=`, `*=` and `%=` as follows:

```go
fanya i = 2

i *= 3 // 6
i /= 2 // 3
i += 100 // 103
i -= 10 // 93
i %= 90 // 3
```

### NEGATIVE NUMBERS

Negative numbers also behave as expected:

```go
fanya i = -10

wakati (i < 0) {
    andika(i)
    i++
}

/*
-10
-9
-8
-7
-6
-5
-4
-3
-2
-1
0
1
2
3
4
5
6
7
8
9 
*/
```
