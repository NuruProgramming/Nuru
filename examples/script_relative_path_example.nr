// Example: Using __FILE__ and __DIR__ for script-relative paths
andika("Script absolute path: " + __FILE__)
andika("Script directory: " + __DIR__)

tumia njia
// Construct a path to a sibling file
sibling = njia.unganisha(__DIR__, "sibling.txt")
andika("Sibling file path: " + sibling)

// Optionally, read a file relative to the script
tumia faili
content = faili.soma(njia.unganisha(__DIR__, "data.txt"))
andika("Data file content: " + content)
