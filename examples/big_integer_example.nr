// big_integer_example.nr – namba kubwa via badilisha and hisabati.namba_kubwa

// Create big integers with badilisha(x, "NAMBA_KUBWA")
a = badilisha("99999999999999999999", "NAMBA_KUBWA")
andika("a =", a)

// Or use the hisabati module: tumia hisabati; hisabati.namba_kubwa("...")
tumia hisabati
b = hisabati.namba_kubwa("11111111111111111111")
andika("b =", b)

// Arithmetic
andika("a + b =", a + b)
andika("a * 2 =", a * badilisha(2, "NAMBA_KUBWA"))

// Comparison
andika("a == b?", a == b)
andika("a > 5?", a > badilisha(5, "NAMBA_KUBWA"))

// Convert to string or (if in range) to NAMBA
andika("as string:", badilisha(a, "NENO"))
c = badilisha(100, "NAMBA_KUBWA")
andika("100 as integer:", badilisha(c, "NAMBA"))
