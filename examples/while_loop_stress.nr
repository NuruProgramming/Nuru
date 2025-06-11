// Stress test for while loops

andika("=== WHILE LOOP STRESS TESTS ===")

// Test 1: Large iteration count
andika("Test 1: Large iteration count")
i = 0
hakika = 100000
wakati (i < hakika) {
    i = i + 1
}
andika("Completed " + aina(hakika) + " iterations: " + aina(i))

// Test 2: Nested while loops
andika("Test 2: Nested while loops")
i = 0
jumla = 0
wakati (i < 1000) {
    j = 0
    wakati (j < 100) {
        jumla = jumla + 1
        j = j + 1
    }
    i = i + 1
}
andika("Nested loops total: " + aina(jumla))

// Test 3: While loop with complex conditions
andika("Test 3: Complex conditions")
a = 1
b = 2
c = 3
wakati (a < 1000 && (b * c) < 10000 && (a + b + c) % 7 != 0) {
    a = a + 1
    b = b + 2
    c = c + 3
}
andika("Complex condition result: a=" + aina(a) + ", b=" + aina(b) + ", c=" + aina(c))

// Test 4: While loop with array operations
andika("Test 4: Array operations in while")
arr = []
i = 0
wakati (i < 5000) {
    arr.sukuma(i * 2)
    i = i + 1
}
andika("Array length after while: " + aina(arr.idadi()))

// Test 5: While loop with dictionary operations
andika("Test 5: Dictionary operations in while")
dict = {}
i = 0
wakati (i < 1000) {
    dict["key_" + aina(i)] = i * i
    i = i + 1
}
andika("Dictionary size after while: " + aina(dict.idadi()))

// Test 6: Infinite loop with break condition (simulate break with flag)
andika("Test 6: Long-running loop with exit condition")
counter = 0
flag = kweli
wakati (flag) {
    counter = counter + 1
    kama (counter > 50000) {
        flag = sikweli
    }
}
andika("Break simulation counter: " + aina(counter))
