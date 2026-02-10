// regex_example.nr – pattern matching with the re module

tumia re

andika("=== linganisha ===")
andika(re.linganisha("[0-9]+", "sala 123"))
andika(re.linganisha("[0-9]+", "hakuna"))

andika("\n=== tafuta ===")
andika(re.tafuta("[0-9]+", "sala 123 zaidi"))

andika("\n=== tafutaZote ===")
andika(re.tafutaZote("[0-9]+", "1 na 2 na 3"))

andika("\n=== vikundi ===")
andika(re.vikundi("(\\w+)\\s+(\\w+)", "jina asha"))

andika("\n=== badilisha ===")
andika(re.badilisha("[0-9]", "a1b2c3", "X"))

andika("\n=== gawa ===")
andika(re.gawa("\\s+", "a   b   c"))
