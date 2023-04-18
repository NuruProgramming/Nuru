# Functions in Nuru

Functions are a fundamental part of Nuru programming, allowing you to define reusable blocks of code. This page covers the syntax and usage of functions in Nuru, including parameters, default parameters, return statements, recursion, and closures.

## Basic Syntax

A function block starts with the unda keyword, followed by parameters enclosed in parentheses () and the body enclosed in curly braces {}. Functions must be assigned to a variable:

```s
jum = unda(x, y) {
    rudisha x + y
}

jum(2, 3) // 5
```

## Parameters

Functions can have zero or any number of arguments. Arguments can be of any type, even other functions:

```s
salamu = unda() {
    andika("Habari yako")
}

salamu()

salamu = unda(jina) {
    andika("Habari yako", jina)
}

salamu("asha") // Habari yako asha
```

## Default Parameters

Functions can be provided with default parameters:

```s
salimu = unda(salamu="Habari") {
    andika(salamu)
}

salimu() // Habari
salimu("Mambo") // Mambo
```

## Return (rudisha)

You can return values with the rudisha keyword. The rudisha keyword will terminate the block and return the value:

```s
mfano = unda(x) {
    rudisha "nimerudi"
    andika(x)
}

mfano("x") // nimerudi
```

## Recursion

Nuru also supports recursion. Here's an example of a recursive Fibonacci function:

```s

fib = unda(n) {
    kama (n <= 1) {
        rudisha n
    } sivyo {
        rudisha fib(n-1) + fib(n-2)
    }
}

andika(fib(10)) // 55
```

The fib function calculates the nth Fibonacci number by recursively calling itself with n-1 and n-2 as arguments until n is less than or equal to 1.

## Closures

Closures are anonymous functions that can capture and store references to variables from their surrounding context. In Nuru, you can create closures using the unda keyword without assigning them to a variable. Here's an example:

```s
fanya jum = unda(x) {
    rudisha unda(y) {
        rudisha x + y
    }
}

fanya jum_x = jum(5)
andika(jum_x(3)) // 8
```

In the example above, the jum function returns another function that takes a single parameter y. The returned function has access to the x variable from its surrounding context.

Now that you understand the basics of functions in Nuru, including recursion and closures, you can create reusable blocks of code to simplify your programs and improve code organization.