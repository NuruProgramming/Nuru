# Switch Statements in Nuru

Switch statements in Nuru allow you to execute different code blocks based on the value of a given expression. This page covers the basics of switch statements and their usage.

## Basic Syntax

You initialize a switch statement with the badili keyword, the expression inside parentheses (), and all cases enclosed within curly braces {}.

A case statement has the keyword ikiwa followed by a value to check. Multiple values can be in a single case separated by commas ,. The consequence to execute if a condition is fulfilled must be inside curly braces {}. Here's an example:

```s
fanya a = 2

badili (a){
	ikiwa 3 {
		andika("a ni tatu")
	}
	ikiwa 2 {
		andika ("a ni mbili")
	}
}
```

## Multiple Values in a Case

Multiple possibilities can be assigned to a single case (ikiwa) statement:

```s
badili (a) {
	ikiwa 1,2,3 {
		andika("a ni kati ya 1, 2 au 3")
	}
	ikiwa 4 {
		andika("a ni 4")
	}
}
```

## Default Case (kawaida)

The default statement will be executed when no condition is satisfied. The default statement is represented by kawaida:

```s
fanya z = 20

badili(z) {
	ikiwa 10 {
		andika("kumi")
	}
	ikiwa 30 {
		andika("thelathini")
	}
	kawaida {
		andika("ishirini")
	}
}
```

By understanding switch statements in Nuru, you can create more efficient and organized code that can handle multiple conditions easily.