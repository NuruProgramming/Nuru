# Sets in Nuru

A **set** is a collection of unique elements. Only hashable types (numbers, strings, booleans, big integers) can be set members. Sets support membership tests with `ktk`, iteration with `kwa ... ktk`, and methods.

## Creating a set

Use the **seta** builtin:

```s
seta()           // empty set
seta(1, 2, 3)    // set from arguments
seta([1, 2, 2])  // set from array (duplicates removed)
```

## Membership

Use `ktk` to test membership:

```s
2 ktk seta(1, 2, 3)   // kweli
5 ktk seta(1, 2, 3)   // sikweli
```

## Methods

| Method | Arguments | Returns | Description |
|--------|-----------|---------|-------------|
| **idadi** | — | namba | Number of elements. |
| **ona**(kipengele) | 1 | kweli/sikweli | True if the set contains the element. |
| **ongeza**(...) | 1+ | seti | Add one or more elements; returns the set. |
| **ondoa**(kipengele) | 1 | seti | Remove the element; returns the set. |
| **kitanzi** | — | kitanzi | Iterator over the set. |

## Iteration

Iteration order is by element string form (Inspect). You can use the set directly in `kwa ... ktk` or call `.kitanzi()`:

```s
kwa _, v ktk seta("a", "b", "c") {
    andika(v)
}
```
