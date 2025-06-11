// Stress test for if statements

andika("=== CONDITIONAL STATEMENTS STRESS TESTS ===")

// Test 1: Deep nested if statements
andika("Test 1: Deep nested conditionals")
fanya test_deep_nesting = unda(depth, value) {
    kama (depth <= 0) {
        rudisha value
    } au kama (depth == 1) {
        kama (value > 0) {
            rudisha value * 2
        } sivyo {
            rudisha value
        }
    } au kama (depth == 2) {
        kama (value > 10) {
            kama (value % 2 == 0) {
                rudisha value * 3
            } sivyo {
                rudisha value * 2
            }
        } sivyo {
            rudisha value
        }
    } sivyo {
        kama (value > 100) {
            kama (value % 3 == 0) {
                kama (value % 5 == 0) {
                    rudisha test_deep_nesting(depth - 1, value / 15)
                } sivyo {
                    rudisha test_deep_nesting(depth - 1, value / 3)
                }
            } sivyo {
                rudisha test_deep_nesting(depth - 1, value - 10)
            }
        } sivyo {
            rudisha test_deep_nesting(depth - 1, value + 1)
        }
    }
}

result = test_deep_nesting(10, 150)
andika("Deep nesting result: " + aina(result))

// Test 2: Complex boolean expressions
andika("Test 2: Complex boolean expressions")
i = 0
complex_hits = 0
wakati (i < 10000) {
    a = i % 7
    b = i % 11
    c = i % 13
    
    kama ((a > 3 && b < 6) || (c == 0 && a != b) || (a + b + c > 15 && a * b * c < 100)) {
        kama ((i % 2 == 0 && i % 3 != 0) || (i > 5000 && i < 8000)) {
            kama (a * b + c > 20 || (a == 0 && b == 0) || c > 10) {
                complex_hits = complex_hits + 1
            }
        }
    }
    i = i + 1
}
andika("Complex boolean hits: " + aina(complex_hits))

// Test 3: Many elif chains
andika("Test 3: Long elif chains")
fanya classify_number = unda(n) {
    kama (n < 0) {
        rudisha "negative"
    } au kama (n == 0) {
        rudisha "zero"
    } au kama (n == 1) {
        rudisha "one"
    } au kama (n < 10) {
        rudisha "single_digit"
    } au kama (n < 50) {
        rudisha "small"
    } au kama (n < 100) {
        rudisha "medium_small"
    } au kama (n < 200) {
        rudisha "medium"
    } au kama (n < 500) {
        rudisha "medium_large"
    } au kama (n < 1000) {
        rudisha "large"
    } au kama (n < 5000) {
        rudisha "very_large"
    } au kama (n < 10000) {
        rudisha "huge"
    } sivyo {
        rudisha "massive"
    }
}

categories = {}
i = 0
wakati (i < 10000) {
    category = classify_number(i)
    kama (categories[category] == tupu) {
        categories[category] = 1
    } sivyo {
        categories[category] = categories[category] + 1
    }
    i = i + 1
}
andika("Number classifications: " + aina(categories))

// Test 4: Conditionals with function calls
andika("Test 4: Conditionals with function calls")
fanya expensive_calculation = unda(x) {
    result = 0
    i = 0
    wakati (i < x) {
        result = result + (i * i)
        i = i + 1
    }
    rudisha result
}

fanya fibonacci = unda(n) {
    kama (n <= 1) {
        rudisha n
    } sivyo {
        rudisha fibonacci(n - 1) + fibonacci(n - 2)
    }
}

total_operations = 0
i = 0
wakati (i < 100) {
    kama (expensive_calculation(10) > 285) {
        kama (fibonacci(8) > 20) {
            total_operations = total_operations + 1
        }
    } au kama (expensive_calculation(5) < 30) {
        total_operations = total_operations + 2
    } sivyo {
        total_operations = total_operations + 3
    }
    i = i + 1
}
andika("Function call conditionals result: " + aina(total_operations))

// Test 5: Memory stress with conditionals
andika("Test 5: Memory stress test")
big_data = []
i = 0
wakati (i < 5000) {
    item = {"id": i, "data": []}
    j = 0
    wakati (j < 10) {
        item["data"].sukuma(j * i)
        j = j + 1
    }
    
    kama (i % 2 == 0) {
        item["type"] = "even"
        kama (i % 4 == 0) {
            item["special"] = kweli
        } sivyo {
            item["special"] = sikweli
        }
    } au kama (i % 3 == 0) {
        item["type"] = "triple"
        item["factor"] = 3
    } sivyo {
        item["type"] = "other"
        kama (i % 5 == 0) {
            item["divisible_by_5"] = kweli
        }
    }
    
    big_data.sukuma(item)
    i = i + 1
}

// Process the data with conditionals
processed = 0
kwa item ktk big_data {
    kama (item["type"] == "even") {
        kama (item["special"]) {
            processed = processed + 2
        } sivyo {
            processed = processed + 1
        }
    } au kama (item["type"] == "triple") {
        processed = processed + item["factor"]
    } sivyo {
        kama (item["divisible_by_5"]) {
            processed = processed + 5
        } sivyo {
            processed = processed + 1
        }
    }
}
andika("Memory stress processed: " + aina(processed))
