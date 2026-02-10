# Built-in Functions in Nuru

Nuru has several built-in functions that perform specific tasks.

## The andika() Function

The andika() function is used to print out messages to the console. It can take zero or more arguments, and the arguments will be printed out with a space in between them. Additionally, andika() supports basic formatting such as /n for a new line, /t for a tab space, and \\ for a backslash. Here's an example:

```s
andika(1, 2, 3) // Output: "1 2 3"
```

## The jaza() Function

The jaza() function is used to get input from the user. It can take zero or one argument, which is a string that will be used as a prompt for the user. Here's an example:

```s
fanya salamu = unda() {
    fanya jina = jaza("Unaitwa nani? ")
    andika("Mambo vipi", jina)
}

salamu()
```

In this example, we define a function `salamu()` that prompts the user to enter their name using the `jaza()` function. We then use the `andika()` function to print out a message that includes the user's name.

## The aina() Function

The `aina()` function is used to determine the type of an object. It accepts one argument, and the return value will be a string indicating the type of the object. Here's an example:

```s
aina(2) // Output: "NAMBA"
aina("Nuru") // Output: "NENO"
```

**badilisha**(value, type) supports the type **"NAMBA_KUBWA"** to create or convert to a big integer (for values beyond 64-bit range). See [Numbers](numbers.md#big-integers-namba_kubwa).

To open a file, use the **faili** module: `tumia faili` then `faili.fungua(path)` or `faili.fungua(path, "r")`. See the [Files](files.md) documentation.

## Base64 encoding and decoding

Use the **crypto** module: `tumia crypto` then `crypto.kodeBase64(...)` and `crypto.katuaBase64(...)`. They support strings, Byte objects, arrays of bytes, URL-safe encoding (second argument `"urlsafe"`), and decode options `"byte"` and `"array"`.

## Runtime and memory (mfumo module)

For garbage collection and memory introspection, use the **mfumo** module: `tumia mfumo`. Available functions: `mfumo.safishaMemori()` (run GC), `mfumo.takwimuMemori()` (object count), `mfumo.takwimuMemoriKwa(kamusi)` (stats dict), `mfumo.kumbukumbaDhaifu(kitu)` (weak reference).