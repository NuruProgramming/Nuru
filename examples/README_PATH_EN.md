# Path Module for Nuru

This document explains how to use the Path module in the Nuru programming language. This module helps with handling file and directory paths, similar to the Node.js [path module](https://nodejs.org/docs/v10.0.0/api/path.html).

## Main Components of the Path Module

| Name | Description |
|------|-------------|
| `jina(path, ext?)` | Gets the last part of a path (like Unix basename) |
| `kigawaji` | Platform-specific path delimiter (`:` for POSIX, `;` for Windows) |
| `sarafa(path)` | Gets the directory part of a path |
| `ext(path)` | Gets the file extension from a path |
| `umbiza(pathObj)` | Creates a path string from an object |
| `niKamili(path)` | Checks if a path is absolute |
| `unganisha(path...)` | Joins path segments together |
| `sawazisha(path)` | Normalizes a path, resolving `.` and `..` |
| `changanua(path)` | Splits a path into an object of components |
| `husika(from, to)` | Gets the relative path between two paths |
| `tatua(path...)` | Resolves paths into an absolute path |
| `kitenga` | Platform-specific path separator (`/` for POSIX, `\` for Windows) |
| `posix` | Path module object that follows POSIX rules |
| `win32` | Path module object that follows Windows rules |

## Windows and POSIX

The default behavior of the Path module varies depending on the operating system where the Nuru program is running. Specifically, when running on a Windows operating system, the Path module will assume Windows-style paths are being used.

## Example Path Usage

```nuru
ita njia kutoka "njia"

// Last part of a path
name = njia.jina("/home/user/file.txt")    // "file.txt"
name_without_ext = njia.jina("/home/user/file.txt", ".txt")  // "file"

// Directory of a path
dir = njia.sarafa("/home/user/file.txt")  // "/home/user"

// File extension
extension = njia.ext("/home/user/file.txt")   // ".txt"

// Joining paths
full = njia.unganisha("/home", "user", "file.txt")  // "/home/user/file.txt"

// Resolving paths to absolute
absolute = njia.tatua("docs", "../images")  // Absolute path depending on current directory
```

## Delimiter

Provides the path delimiter specific to the platform:

* `;` for Windows
* `:` for POSIX

Example:

```nuru
andika(njia.kigawaji)  // ":" for POSIX, ";" for Windows

// Splitting PATH
system_paths = PATH.split(njia.kigawaji)
```

## Separator

Provides the path separator specific to the platform:

* `\` for Windows
* `/` for POSIX

Example:

```nuru
andika(njia.kitenga)  // "/" for POSIX, "\" for Windows

// Splitting a path
segments = "file/one/two".split(njia.kitenga)  // ["file", "one", "two"]
```

## Parse

The `changanua()` method splits a path into an object with important path components:

```nuru
parts = njia.changanua("/home/user/file.txt")

// Builds an object:
// {
//   root: "/",
//   dir: "/home/user",
//   base: "file.txt",
//   ext: ".txt",
//   name: "file"
// }
```

Result structure:

```
┌─────────────────────┬────────────┐
│          dir        │    base    │
├──────┬              ├──────┬─────┤
│ root │              │ name │ ext │
"  /    home/user      / file  .txt "
└──────┴──────────────┴──────┴─────┘
```

## Format

The `umbiza()` method creates a path string from an object, the opposite of `changanua()`:

```nuru
new_path = njia.umbiza({
    root: "/",
    dir: "/home/user",
    base: "file.txt",
    name: "file",
    ext: ".txt"
})

// Builds: "/home/user/file.txt"
```

When providing components to `umbiza`, note:

* `root` is ignored if `dir` is provided
* `ext` and `name` are ignored if `base` exists

## POSIX and Windows Examples

For consistent results even when dealing with paths from different systems, the Path module provides specific components:

```nuru
// Get POSIX behavior on either Windows or POSIX
posix = njia.posix
posix_path = posix.unganisha("/home", "user", "file.txt")
// Always: "/home/user/file.txt"

// Get Windows behavior on either Windows or POSIX
win = njia.win32
win_path = win.unganisha("C:", "Users", "file.txt")
// Always: "C:\Users\file.txt"
```

## Complete Examples

Example of joining and normalizing paths:

```nuru
// Joining paths
path1 = njia.unganisha("/home", "user", "./docs", "../images")
// Result: "/home/user/images"

// Normalizing paths
path2 = njia.sawazisha("/home/./user/../user/images/../images/.")
// Result: "/home/user/images"
```

Example of resolving relative paths:

```nuru
relative = njia.husika("/home/user/docs", "/home/user/images/sunset.jpg")
// Result: "../images/sunset.jpg"
```

Example of resolving paths to absolute:

```nuru
absolute = njia.tatua("docs", "../images", "./sunset.jpg")
// If current directory is "/home/user", result: "/home/user/images/sunset.jpg"
```

## Conclusion

The Path module for Nuru follows the Node.js approach, but using Swahili syntax and naming. This module helps implement file and directory path handling easily across different operating environments. 