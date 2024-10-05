## Range Function (mfululizo)

The `mfululizo` function generates a sequence of numbers. It can be used in loops or to create arrays of sequential numbers.

### Syntax

```go
mfululizo(end)
mfululizo(start, end)
mfululizo(start, end, step)
```

### Parameters

- `end`: The upper limit of the sequence (exclusive).
- `start` (optional): The starting value of the sequence. Default is 0.
- `step` (optional): The increment between each number in the sequence. Default is 1.

### Return Value

Returns an array of integers.

### Examples

```go
// Generate numbers from 0 to 4
kwa i katika mfululizo(5) {
    andika(i)
}
// Output: 0 1 2 3 4

// Generate numbers from 1 to 9
kwa i katika mfululizo(1, 10) {
    andika(i)
}
// Output: 1 2 3 4 5 6 7 8 9

// Generate even numbers from 0 to 8
kwa i katika mfululizo(0, 10, 2) {
    andika(i)
}
// Output: 0 2 4 6 8

// Generate numbers in reverse order
kwa i katika mfululizo(10, 0, -1) {
    andika(i)
}
// Output: 10 9 8 7 6 5 4 3 2 1
```

### Notes

- The `end` value is exclusive, meaning the sequence will stop before reaching this value.
- If a negative `step` is provided, `start` should be greater than `end`.
- The `step` value cannot be zero.
