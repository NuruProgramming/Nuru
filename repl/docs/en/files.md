# Files in Nuru

Nuru's ability to deal with files is primitive, and as for now it only allows you to read contents of a file.

Use the **faili** module: `tumia faili`.

## Opening a File

You open a file with `faili.fungua(path)` or `faili.fungua(path, "r")`. This returns an object of type `FAILI`:
```
tumia faili
fileYangu = faili.fungua("file.txt", "r")

aina(fileYangu) // FAILI
```

## Reading a File

Once you have a file object you can read its contents with the `soma()` method. This will return the contents of the file as a string:
```
tumia faili
fileYangu = faili.fungua("file.txt", "r")

fileYangu.soma()
```
