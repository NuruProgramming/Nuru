<h1 align="center"> NURU✨PROGRAMMING✨LANGUAGE</h1>
<p align="center">
    <a href="https://github.com/AvicennaJr/Nuru"><img alt="Nuru Programming Language" src="https://img.shields.io/badge/Nuru-Programming%20Language-yellow"></a>
    <a href="https://github.com/AvicennaJr/Nuru"><img alt="Nuru Programming Language" src="https://img.shields.io/badge/platform-Linux | Windows | Android-green.svg"></a>
    <a href="https://github.com/AvicennaJr/Nuru"><img alt="Nuru Programming Language" src="https://img.shields.io/github/last-commit/AvicennaJr/Nuru"></a>
    <a href="https://github.com/AvicennaJr/Nuru"><img alt="Nuru Programming Language" src="https://img.shields.io/github/v/release/AvicennaJr/Nuru"></a>
<br>
    <a href="https://github.com/AvicennaJr/Nuru"><img alt="Nuru Programming Language" src="https://img.shields.io/github/stars/AvicennaJr/Nuru?style=social"></a>
</p>
A Swahili Programming Language of its kind built from the ground up.

## Installation

To get started download the executables from the release page or follow the
instructions for your device below:

### Linux

 - Download the binary:

```
curl -O -L https://github.com/AvicennaJr/Nuru/releases/download/v0.1.0/nuru_linux_amd64_v0.1.0.tar.gz
```

 - Extract the file to `$HOME/bin`:

```
sudo tar -C $HOME/bin -xzvf nuru_linux_amd64_v0.1.0.tar.gz
```
 - Add $HOME/bin to path:
```
echo 'export PATH="${HOME}/bin:${PATH}"' >> ~/.bashrc
```

 - Confirm installation with:

```
nuru -v
```

### Android (Termux)

 - Make sure you have [Termux](https://f-droid.org/repo/com.termux_118.apk) installed.
 - Download the binary with this command:

```
curl -O -L https://github.com/AvicennaJr/Nuru/releases/download/v0.1.0/nuru_android_arm64_v0.1.0.tar.gz
```
 - Extract the file:

```
tar -xzvf nuru_android_arm64_v0.1.0.tar.gz
```
 - Add it to path:

```
echo "alias nuru='~/nuru'" >> .bashrc
```
 - Confirm installation with:

```
nuru -v 
```

### Windows

 - Make a bin directory if it doesn't exist:

```
mkdir C:\bin
```
 - Download the Nuru Program [Here](https://github.com/AvicennaJr/Nuru/releases/download/v0.1.0/nuru_windows_amd64_v0.1.0.exe)
 - Rename the downloaded program from `nuru_windows_amd64_v0.1.0.exe` to `nuru.exe`
 - Move the file `nuru.exe` to the folder `C:\bin`
 - Add the bin folder to Path with this command:

```
setx PATH "C:\bin;%PATH%"
```
 - Confirm installation with:

```
nuru -v
```

### Building From Source

 - Make sure you have golang installed
 - Run the following command:

```
go build -o nuru main.go
```
 - You can optionally add the binary to $PATH as shown above
 - Confirm installtion with:

```
nuru -v
```

## Syntax

Nuru, although still in its early stage, intends to be a fully functional programming language, and thus it has been baked with many features.

### Defining A Variable

To initiliaze a variable use the `acha` keyword:

```
acha x = 2;
acha y = 3;

andika(x*y) // output is 6
```
You can reassign values to the variable after it has been initiliazed:

```
x = 10

andika(x*y) // output is 30
```
**Note that Semicolons ";" are OPTIONAL**

### Comments

Nuru supports both single line and multiple line comments as shown below:

```
// Single line comment

/*

Multiple
Line
Comment 
*/ 
```

### Arithmetic Operations

For now Nuru supports `+`, `-`, `/` and `*`. More will be added. The `/` operation will truncate (round to a whole number) as Floating points are not supported yet.

Nuru also provides precedence of operations using the BODMAS rule:

```
2 + 2 * 3 // output = 8

2 * (2 + 3) // output = 10
```

### Types

Nuru has the following types:

Type      | Syntax                                    | Comments
--------- | ----------------------------------------- | -----------------------
BOOL      | `kweli sikweli`                           | kweli == true, sikweli == false
INT       | `1, 100, 342, -4`                         | These are signed 64 bit integers
STRING    | `"" "mambo" "habari yako"`                | They MUST be in DOUBLE QUOTES `"`
ARRAY     | `[] [1, 2, 3] [1, "moja", kweli]`       | Arrays can hold any types
DICT      | `{} {"a": 3, 1: "moja", kweli: 2}`        | Keys can be int, string or bool. Values can be anything

### Functions

This is how you define a function in Nuru:

```
acha jumlisha = fn(x, y) {
        rudisha x + y
    }

andika(jumlisha(3,4))
```

Nuru also supports recursion:

```
acha fibo = fn(x) {
	kama (x == 0) {
		rudisha 0;
	} au kama (x == 1) {
			rudisha 1;
	} sivyo {
			rudisha fibo(x - 1) + fibo(x - 2);
	}
}
```

### If Statements

Nuru supports if, elif and else statements with keywords `kama`, `au kama` and `sivyo` respectively:

```
kama (2<1) {
    andika("Mbili ni ndogo kuliko moja")
} au kama (3 < 1) {
    andika ("Tatu ni ndogo kuliko moja")
} sivyo {
    andika("Moja ni ndogo")
}
```

### While Loops

Nuru's while loop syntax is as follows:

```
acha i = 10

wakati (i > 0) {
	andika(i)
	i = i - 1
}
```

### Arrays

This is how you initiliaze and perform other array operations in Nuru:
```
acha arr = []

// To add elements

sukuma(arr, 2)
andika(arr) // output = [2]
// Add two Arrays

acha arr2 = [1,2,3,4]

acha arr3 = arr1 + arr2

andika(arr3) // output = [2,1,2,3,4]

// reassign value

arr3[0] = 0

andika[arr3] // output = [0,1,2,3,4]

// get specific item

andika(arr[3]) // output = 3
```

### Dictionaries

Nuru also supports dictionaris and you can do a lot with them as follows:
```
acha mtu = {"jina": "Mojo", "kabila": "Mnyakusa"}

// get value from key 
andika(mtu["jina"]) // output = Mojo

andika(mtu["kabila"]); // output = Mnyakusa

// You can reassign values

mtu["jina"] = "Avicenna"

andika(mtu["jina"]) // output = Avicenna

// You can also add new values like this:

mtu["anapoishi"] = "Dar Es Salaam"

andika(mtu) // output = {"jina": "Avicenna", "kabila": "Mnyakusa", "anapoishi": "Dar Es Salaam"}

// You can also add two Dictionaries

acha kazi = {"kazi": "jambazi"}

mtu = mtu + kazi

andika(mtu) // output = {"jina": "Avicenna", "kabila": "Mnyakusa", "anapoishi": "Dar Es Salaam", "kazi": "jambazi"}
```

### Getting Input From User

In Nuru you can get input from users using the `jaza()` keyword as follows:
```
acha jina = jaza("Unaitwa nani? ") // will prompt for input

andika("Habari yako " + jina)
```

## How To Run

### Using The Intepreter:

You can enter the intepreter by simply running the `nuru` command:
```
nuru
>>> andika("karibu")
karibu
>>> 2 + 2
4
```
Kindly Note that everything should be placed in a single line. Here's an example:
```
>>> kama (x > y) {andika("X ni kubwa")} sivyo {andika("Y ni kubwa")}
```
### Running From File

To run a Nuru script, write the `nuru` command followed by the name of the file with a `.nr` extension:

```
nuru myFile.nr
```

## Issues

Kindly open an [Issue](https://github.com/AvicennaJr/Nuru/issues) to make suggestions and anything else.

## Contributions

All contributions are welcomed. Clone the repo, hack it, make sure all tests are passing then submit a pull request.

## License

[MIT](http://opensource.org/licenses/MIT)

## Authors

Nuru Programming Language has been authored and being actively maintained by [Avicenna](https://github.com/AvicennaJr)
