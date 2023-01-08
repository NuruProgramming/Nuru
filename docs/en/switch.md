## SWITCH (BADILI)

### Definition

You initialize a switch statement with `badili`, the expression inside parenthesis `()` and all cases will be enclosed inside a bracket `{}`.

A case statement has the keyword `ikiwa` followed by a value to check. Multiple values can be in a single case separated by commas `,`. The consequence to execute if a condition is fulfiled must be inside a bracket `{}`. Here's an example:
```
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

### Multiple Values in a Case

Multiple possibilites can be assigned to a single case (`ikiwa`) statement:
```
badili (a) {
	ikiwa 1,2,3 {
		andika("a ni kati ya 1, 2 au 3")
	}
	ikiwa 4 {
		andika("a ni 4")
	}
}
```

### Default (kawaida)

The default statement will be executed when no condition is satisfied. The default statement is represented by `kawaida`:
```
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