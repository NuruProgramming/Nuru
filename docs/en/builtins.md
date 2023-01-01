## BUILTINS

Nuru has a few builtin functions and more will be added in the future

### andika()

This function will print out whatever is placed inside the parenthesis `()`. It can take zero or multiple number of arguments. Arguments will be printed out with a space in between them:
```
andika(1,2,3) // 1 2 3
```
`andika()` also supports some basic formatting such as:
- `/n` for a new line
- `/t` for a tab space
- `\\` for a backslash

### jaza()

This is a function to get input from a user. It can have zero or one argument. The only acceptable argument is a string:
```
fanya salamu = unda(){
	fanya jina = jaza("Unaitwa nani? ")
	andika("mambo vipi", jina)
}

salamu()
```

### aina()

`Aina()` is a function to help identify the type of an object. It only accepts one argument:
```
aina(2) // NAMBA
```

### idadi()

`idadi` is a function to know a length of an object. It accepts only one argument which can be a `string`, `list` or `dictionary`:
```
idadi("mambo") // 5
```

### sukuma()

`sukuma()` is a function that adds a new element to an array. The function accepts two arguments, the first must be a list and the second is the element to be added/appended:
```
fanya majina = ["juma", "asha"]

majina = sukuma(majina, "mojo")
```
**Notice that the list is reassigned for the change to take effect**

### yamwisho()

This is a function to get the last element in an array. It only accepts one argument which must be an array:
```
fanya namba = [1,2,3,4,5]

yamwisho(namba) // 5
```

**MORE BUILTIN FUNCTIONS WILL BE ADDED WITH TIME**