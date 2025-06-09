// njia_glob_example.nr
// Demonstrate path.glob (wildcard matching)

tumia njia

andika("=== Njia Glob Example ===")

andika("All .go files in module/: ")
files = njia.glob("module/*.go")
andika(files)

andika("All .md files in root: ")
mds = njia.glob("*.md")
andika(mds)

andika("=== Njia Glob Example Complete ===")
