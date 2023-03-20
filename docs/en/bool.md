# Working with Booleans in Nuru

Boolean objects in Nuru are truthy, meaning that any value is true, except tupu and sikweli. They are used to evaluate expressions that return true or false values.

## Evaluating Boolean Expressions

### Evaluating Simple Expressions

In Nuru, you can evaluate simple expressions that return a boolean value:

```s
andika(1 > 2) // Output: `sikweli`

andika(1 + 3 < 10) // Output: `kweli`
```

### Evaluating Complex Expressions

In Nuru, you can use boolean operators to evaluate complex expressions:

```s
a = 5
b = 10
c = 15

result = (a < b) && (b < c)

kama (result) {
    andika("Both conditions are true")
} sivyo {
    andika("At least one condition is false")
}
// Output: "Both conditions are true"
```

Here, we create three variables a, b, and c. We then evaluate the expression (a < b) && (b < c). Since both conditions are true, the output will be "Both conditions are true".

## Boolean Operators

Nuru has several boolean operators that you can use to evaluate expressions:

### The && Operator

The && operator evaluates to true only if both operands are true. Here's an example:

```s
andika(kweli && kweli) // Output: `kweli`

andika(kweli && sikweli) // Output: `sikweli`
```

### The || Operator

The || operator evaluates to true if at least one of the operands is true. Here's an example:

```s
andika(kweli || sikweli) // Output: `kweli`

andika(sikweli || sikweli) // Output: `sikweli`
```

### The ! Operator

The ! operator negates the value of the operand. Here's an example:

```s
andika(!kweli) // Output: `sikweli`

andika(!sikweli) // Output: `kweli`
```

## Working with Boolean Values in Loops

In Nuru, you can use boolean expressions in loops to control their behavior. Here's an example:

```s
namba = [1, 2, 3, 4, 5]

kwa thamani ktk namba {
    kama (thamani % 2 == 0) {
        andika(thamani, "is even")
    } sivyo {
        andika(thamani, "is odd")
    }
}
// Output:
// 1 is odd
// 2 is even
// 3 is odd
// 4 is even
// 5 is odd
```

Here, we create an array namba with the values 1 through 5. We then loop over each value in the array and use the % operator to determine if it is even or odd. The output will be "is even" for even numbers and "is odd" for odd numbers.


Boolean objects in Nuru can be used to evaluate expressions that return true or false values. You can use boolean operators to evaluate complex expressions and control the behavior of loops. Understanding how to work with boolean values is an essential skill for any Nuru programmer.