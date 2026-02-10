// smoke_no_io.nr – smoke test without I/O (re, iterator, C-style for, big integer).
// Run: nuru examples/smoke_no_io.nr  (expect exit 0)

tumia re
re.linganisha("[0-9]+", "sala 123")

orodha = [1, 2, 3]
jumla = 0
kwa _, v ktk orodha.kitanzi() {
	jumla = jumla + v
}

n = 0
kwa i = 0; i < 3; i = i + 1 { n = n + 1 }

kubwa = badilisha("100", "NAMBA_KUBWA")
badilisha(kubwa, "NENO")

andika("tayari")
