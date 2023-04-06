# Conditional Statements in Nuru

Conditional statements in Nuru are used to perform different actions based on different conditions. The if/else statement is a fundamental control structure that allows you to execute code based on specific conditions. This page covers the basics of if/else statements in Nuru.

## If Statement (Kama)

An if statement starts with the kama keyword, followed by a condition in parentheses (). If the condition is true, the code inside the curly braces {} will be executed.

```s
kama (2 > 1) {
    andika(kweli) // kweli
}
```

In this example, the condition 2 > 1 is true, so the andika(kweli) statement is executed, and the output is kweli.

## Else If and Else Blocks (Au Kama and Sivyo)

You can use au kama to test multiple conditions and sivyo to specify a default block of code to be executed when none of the conditions are true.

```s

fanya a = 10

kama (a > 100) {
    andika("a imezidi 100")
} au kama (a < 10) {
    andika("a ndogo kuliko 10")
} sivyo {
    andika("Thamani ya a ni", a)
}

// The output will be 'Thamani ya a ni 10'
```

In this example, the first condition a > 100 is false, and the second condition a < 10 is also false. Therefore, the code inside the sivyo block is executed, and the output is 'Thamani ya a ni 10'.

By using if/else statements with the kama, au kama, and sivyo keywords, you can control the flow of your Nuru code based on different conditions.