## OPERATORS

### ASSIGNMENT

Assuming `i` and `v` are predefined variables, Nuru supports the following assignment operators:

- `i = v`: which is the regular assign operator
- `i += v`: which is the equivalent of `i = i + v`
- `i -= v`: which is the equivalent of `i = i - v`
- `i *= v`: which is the equivalent of `i = i * v`
- `i /= v`: which is the equivalent of `i = i / v`
- `i += v`: which is the equivalent of `i = i + v`

For `strings`, `arrays` and `dictionaries`, the `+=` sign operator is permissible. Example:
```
list1 += list2 // this is equivalent to list1 = list1 + list2
```

### ARITHMETIC OPERATORS

The following arithmetic operators are supported:

- `+`: Additon
- `-`: Subtraction
- `*`: Multiplication
- `/`: Division
- `%`: Modulo (ie the remainder of a division)
- `**`: Exponential power (eg: `2**3 = 8`)

### COMPARISON OPERATORS

The following comparison operators are supported:

- `==`: Equal to
- `!=`: Not equal to
- `>`: Greater than
- `>=`: Greater than or equal to
- `<`: Less than
- `<=`: Less than or equal to

### MEMBER OPERATOR

The member operator in Nuru is `ktk`. It will check if an object exists in another object:
```go
fanya majina = ['juma', 'asha', 'haruna']

"haruna" ktk majina // kweli
"halima" ktk majina // sikweli
```

### LOGIC OPERATORS

The following logic operators are supported:

- `&&`: Logical `AND`. It will evaluate to true if both are true, otherwise it will evaluate to false.
- `||`: Logical `OR`. It will evaluate to false if both are false, otherwise it will evaluate to true.
- `!`: Logical `NOT`. It will evaluate to the opposite of a given expression.

### PRECEDENCE OF OPERATORS

The following is the precedence of operators, starting from the HIGHEST PRIORITY to LOWEST.

- `()` : Items in paranthesis have the highest priority
- `!`: Negation
- `%`: Modulo
- `**`: Exponential power
- `/, *`: Division and Multiplication
- `+, +=, -, -=`: Addition and Subtraction
- `>, >=, <, <=`: Comparison operators
- `==, !=`: Equal or Not Equal to
- `=`: Assignment Operator
- `ktk`: Member Operator
- `&&, ||`: Logical AND and OR