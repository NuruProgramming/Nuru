# Null (Tupu) in Nuru

The null data type in Nuru represents the absence of a value or the concept of "nothing" or "empty." This page covers the syntax and usage of the null data type in Nuru, including its definition and evaluation.

## Definition

A null data type is a data type with no value, defined with the tupu keyword:

```s
fanya a = tupu
```
## Evaluation

When evaluating a null data type in a conditional expression, it will evaluate to false:

```s
kama (a) {
    andika("niko tupu")
} sivyo {
    andika("nimevaa nguo")
}

// Output: nimevaa nguo
```

The null data type is useful in Nuru when you need to represent an uninitialized, missing, or undefined value in your programs. By understanding the null data type, you can create more robust and flexible code.