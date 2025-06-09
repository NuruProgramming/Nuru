// Example: Combining faili (fs), njia (path), and jsoni (JSON) modules in Nuru
// This script demonstrates reading a JSON file, updating its data, and saving it to a new location using path utilities.

tumia faili
tumia njia
tumia jsoni

// Compose a path for the input JSON file relative to script directory
input_path = njia.unganisha(__DIR__, "data.json")
andika("Input JSON path:", input_path)

// Write a sample JSON file if it doesn't exist
fanya failiIpo = faili.ipo(input_path)
andika("File exists:", failiIpo)

kama (!failiIpo) {
    sample = {"jina": "Asha", "umri": 25, "vipaji": ["kuimba", "kupika"]}
    faili.andika(input_path, jsoni.enkodi(sample))
    andika("Sample data.json created.")
}

// Read and decode the JSON file
data = faili.soma(input_path)
obj = jsoni.dikodi(data)
andika("Original object:", obj)

// Update the object
obj["umri"] = obj["umri"] + 1
obj["vipaji"].sukuma("kusoma")

// Save to a new file in a sibling directory
output_dir = njia.unganisha(__DIR__, "output")
fanya niSarafi = faili.niSarafa(output_dir)

andika("Checking if output directory exists before creation:", niSarafi)

kama (!niSarafi) {
    andika("Attempting to create output directory...")
    faili.tengenezaSarafa(output_dir, {rekesia: kweli})
    andika("Directory created?", faili.niSarafa(output_dir))
}

output_path = njia.unganisha(output_dir, "updated_data.json")
faili.andika(output_path, jsoni.pendeza(obj))
andika("Updated JSON saved to:", output_path)

// Read back and print
andika("Updated file contents:", faili.soma(output_path))
