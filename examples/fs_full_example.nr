tumia faili

// --- File Operations ---
faili.andika("testfile.txt", "Hello, Nuru!")
andika("Wrote to testfile.txt")

andika("Read testfile.txt:", faili.soma("testfile.txt"))

faili.ongeza("testfile.txt", "\nAppended line.")
andika("Appended to testfile.txt:", faili.soma("testfile.txt"))

// Open and close file
testfile = faili.fungua("testfile.txt", "r")
andika("Opened file:", testfile)
faili.funga(testfile)
andika("Closed file.")

// File exists?
andika("testfile.txt exists:", faili.ipo("testfile.txt"))

// --- Directory Operations ---
faili.tengenezaSarafa("testdir")
andika("Created directory testdir")

andika("Is testdir a directory?", faili.niSarafa("testdir"))
andika("Is testfile.txt a file?", faili.niFaili("testfile.txt"))

faili.andika("testdir/file1.txt", "File 1")
faili.andika("testdir/file2.txt", "File 2")
andika("Directory contents:", faili.orodha("testdir"))

// Remove file
faili.futa("testdir/file2.txt")
andika("Removed testdir/file2.txt")

// Remove directory (non-recursive, should fail if not empty)
andika("Try remove non-empty dir (should error):", faili.futaSarafa("testdir"))

// Remove directory recursively
faili.futaSarafa("testdir", kweli)
andika("Removed testdir recursively")

// --- File Metadata ---
andika("File stat:", faili.hali("testfile.txt"))

// --- Permissions (chmod) ---
faili.ruhusu("testfile.txt", 0o644)
andika("Changed permissions of testfile.txt to 644")

// --- Rename ---
faili.badilisha("testfile.txt", "renamed.txt")
andika("Renamed testfile.txt to renamed.txt")

// --- Symlinks ---
faili.kiungo("renamed.txt", "link_to_renamed.txt")
andika("Created symlink link_to_renamed.txt -> renamed.txt")
andika("Symlink points to:", faili.somaKiungo("link_to_renamed.txt"))

// --- Clean up ---
faili.futa("renamed.txt")
faili.futa("link_to_renamed.txt")
andika("Cleaned up files.")

// --- Error Handling ---
andika("Try reading non-existent file:", faili.soma("no_such_file.txt"))
andika("Try removing non-existent file:", faili.futa("no_such_file.txt"))
