// Comprehensive example: Demonstrate all core functionalities of the 'njia' (path) module
andika("=== Njia Module Full Example ===")

tumia njia

// 1. jina: Get the last part of a path
andika("jina: " + njia.jina("/home/user/file.txt")) // file.txt
andika("jina (with ext): " + njia.jina("/home/user/file.txt", ".txt")) // file

// 2. sarafa: Get the directory part
andika("sarafa: " + njia.sarafa("/home/user/file.txt")) // /home/user

// 3. ext: Get the file extension
andika("ext: " + njia.ext("/home/user/file.txt")) // .txt

// 4. umbiza: Format a path from an object
parts = {"root": "/", "dir": "/home/user", "base": "file.txt", "ext": ".txt", "name": "file"}
andika("umbiza: " + njia.umbiza(parts))

// 5. niKamili: Check if a path is absolute
andika("niKamili (absolute): " + tungo(njia.niKamili("/home/user/file.txt"))) // true
andika("niKamili (relative): " + tungo(njia.niKamili("file.txt"))) // false

// 6. unganisha: Join path segments
andika("unganisha: " + njia.unganisha("/home", "user", "docs", "file.txt")) // /home/user/docs/file.txt

// 7. sawazisha: Normalize a path
andika("sawazisha: " + njia.sawazisha("/home/user/../user/./file.txt")) // /home/user/file.txt

// 8. changanua: Parse a path into components
parsed = njia.changanua("/home/user/file.txt")
andika("changanua: root=" + parsed["root"] + ", dir=" + parsed["dir"] + ", base=" + parsed["base"] + ", ext=" + parsed["ext"] + ", name=" + parsed["name"])

// 9. husika: Get the relative path from one to another
andika("husika: " + njia.husika("/home/user/docs", "/home/user/images/file.jpg")) // ../images/file.jpg

// 10. tatua: Resolve paths into an absolute path
andika("tatua: " + njia.tatua("docs", "../images", "file.jpg"))

// 11. kitenga: Platform-specific path separator
andika("kitenga: " + njia.kitenga())

// 12. kigawaji: Platform-specific path delimiter
andika("kigawaji: " + njia.kigawaji())

// 13. glob: Globbing (wildcard matching)
andika("glob (all .nr files): " + tungo(njia.glob("examples/*.nr")))

// 14. posix/win32: Use POSIX or Windows path rules
posix_path = njia.posix()
andika("posix unganisha: " + tungo(posix_path.unganisha("/home", "user", "file.txt")))
win32_path = njia.win32()
andika("win32 unganisha: " + tungo(win32_path.unganisha("C:\\", "Users", "file.txt")))
