# Identifiers in Nuru

Identifiers are used to name variables, functions, and other elements in your Nuru code. This page covers the rules and best practices for creating identifiers in Nuru.

## Syntax Rules

Identifiers can contain letters, numbers, and underscores. However, there are a few rules you must follow when creating identifiers:
- Identifiers cannot start with a number.
- Identifiers are case-sensitive. For example, myVar and myvar are considered distinct identifiers.

Here are some examples of valid identifiers:

```s
fanya birth_year = 2020
andika(birth_year) // 2020

fanya convert_c_to_p = "C to P"
andika(convert_c_to_p) // "C to P"
```

In the examples above, birth_year and convert_c_to_p are both valid identifiers.

## Best Practices

When choosing identifiers, it's important to follow best practices to ensure your code is clear and easy to understand:

- Use descriptive names that clearly indicate the purpose or meaning of the variable or function.
- Follow a consistent naming convention, such as camelCase (myVariableName) or snake_case (my_variable_name).
- Avoid using single-letter variable names, except for commonly accepted cases like loop counters (i, j, k).

By following these best practices when creating identifiers, you will make your Nuru code more readable and maintainable for yourself and others.
