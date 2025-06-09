# File System Module for Nuru

This document explains how to use the File System module in the Nuru programming language. This module helps with handling files and directories on disk, similar to the Node.js [fs module](https://nodejs.org/docs/v22.9.0/api/fs.html).

## Main Components of the File System Module

| Name | Description |
|------|-------------|
| `soma(path)` | Reads a file and returns the contents as a string |
| `andika(path, data)` | Writes data to a file, replacing the file if it already exists |
| `ongeza(path, data)` | Appends data to the end of a file, creating the file if it doesn't exist |
| `fungua(path, mode?)` | Opens a file and returns a file descriptor |
| `funga(fd)` | Closes a file descriptor |
| `futa(path)` | Deletes a file or directory |
| `fanya(path)` | Creates a new file, returning a file descriptor |
| `ipo(path)` | Checks if a file or directory exists |
| `orodha(path)` | Reads the contents of a directory |
| `tengenezaSarafa(path, mode?)` | Creates a directory |
| `futaSarafa(path)` | Deletes an empty directory |
| `hali(path)` | Gets file or directory information |
| `niSarafa(path)` | Checks if a path is a directory |
| `niFaili(path)` | Checks if a path is a regular file |
| `ruhusu(path, mode)` | Changes file permissions |
| `mmiliki(path, uid, gid)` | Changes file owner |
| `badilisha(oldPath, newPath)` | Renames a file or directory (moves it) |
| `kiungo(target, linkname)` | Creates a symbolic link |
| `somaKiungo(path)` | Reads the value of a symbolic link |

## Basic Usage

### Reading and Writing Files

```nuru
ita faili kutoka "faili"

// Reading a file
content = faili.soma("text.txt")
andika(content)

// Writing to a file
faili.andika("text.txt", "This is new text")

// Appending to a file
faili.ongeza("text.txt", "\nThis is a new line")
```

### Working with Directories

```nuru
ita faili kutoka "faili"

// Creating a directory
faili.tengenezaSarafa("new_directory")

// Creating nested directories (recursive)
faili.tengenezaSarafa("directory/inner/more", rekesia: kweli)

// Reading directory contents
items = faili.orodha("new_directory")
kwa item ktk items {
    andika(item.jina)
    andika("  Is directory? " + item.ni_sarafa)
    andika("  Size: " + item.ukubwa + " bytes")
}

// Removing a directory
faili.futaSarafa("empty_directory")

// Removing a directory and its contents
faili.futaSarafa("directory/inner", rekesia: kweli)
```

### Getting File Information

```nuru
ita faili kutoka "faili"

// Getting file information
stats = faili.hali("text.txt")
andika("Name: " + stats.jina)
andika("Size: " + stats.ukubwa + " bytes")
andika("Is directory? " + stats.ni_sarafa)
andika("Permissions: " + stats.ruhusu)
andika("Time: " + stats.muda)

// Checking types
andika("Is directory? " + faili.niSarafa("text.txt"))
andika("Is file? " + faili.niFaili("text.txt"))

// Checking if a file exists
andika("File exists? " + faili.ipo("doesnotexist.txt"))
```

## File Descriptors

A file descriptor is an object that represents an opened file.

```nuru
ita faili kutoka "faili"

// Open file for reading
fd = faili.fungua("text.txt", "r")
andika(fd.Content)

// Open file for writing
fd_write = faili.fungua("new.txt", "w")
// Write to the file
fd_write.andika("This is text in a new file")

// Close files
faili.funga(fd)
faili.funga(fd_write)
```

## File Opening Modes

The following modes can be used when opening a file:

| Mode | Description |
|------|-------------|
| `"r"` | Open for reading (default) |
| `"r+"` | Open for reading and writing |
| `"w"` | Open for writing, creating the file or truncating if it exists |
| `"w+"` | Open for reading and writing, creating the file or truncating if it exists |
| `"a"` | Open for appending, creating the file if it doesn't exist |
| `"a+"` | Open for reading and appending, creating the file if it doesn't exist |

## Symbolic Links

Symbolic links are a special type of file that points to another file or directory.

```nuru
ita faili kutoka "faili"

// Create a symbolic link
faili.kiungo("text.txt", "pointer.txt")

// Read a symbolic link
target = faili.somaKiungo("pointer.txt")
andika("Link points to: " + target)

// Get information about the link itself
stats = faili.hali("pointer.txt")
andika("Link size: " + stats.ukubwa)
```

## Error Handling

The best way to handle errors is using a `try...catch` block:

```nuru
ita faili kutoka "faili"

jaribu {
    faili.soma("nonexistent_file.txt")
} shika error {
    andika("Error: " + error)
}
```

## Best Practice Examples

### Copying a File

```nuru
ita faili kutoka "faili"

function copyFile(source, destination) {
    data = faili.soma(source)
    faili.andika(destination, data)
    andika("File copied from " + source + " to " + destination)
}

copyFile("original.txt", "copy.txt")
```

### Finding Files with Specific Extensions

```nuru
ita faili kutoka "faili"
ita njia kutoka "njia"

function findFilesWithExtension(directory, extension) {
    items = faili.orodha(directory)
    results = []
    
    kwa item ktk items {
        kama (!item.ni_sarafa && njia.ext(item.jina) === extension) {
            results.sukuma(njia.unganisha(directory, item.jina))
        }
    }
    
    rudisha results
}

// Find all .txt files in the directory
txt_files = findFilesWithExtension(".", ".txt")
andika(txt_files)
```

### Creating a Directory Tree

```nuru
ita faili kutoka "faili"
ita njia kutoka "njia"

function printDirectory(directory, indent = "") {
    items = faili.orodha(directory)
    
    kwa item ktk items {
        andika(indent + "- " + item.jina)
        
        kama (item.ni_sarafa) {
            new_path = njia.unganisha(directory, item.jina)
            printDirectory(new_path, indent + "  ")
        }
    }
}

// Print the directory tree of the current directory
printDirectory(".")
```

## Conclusion

The File System module for Nuru follows the Node.js approach, but using Swahili syntax and naming. This module helps work with the file system in an easy and practical way. 