// Combined stress test for all control flow constructs

andika("=== COMBINED CONTROL FLOW STRESS TEST ===")

// Test combining all control structures in complex scenarios
andika("Test 1: All constructs together")

fanya complex_algorithm = unda(data_size) {
    // Initialize data structures
    results = {}
    matrix = []
    
    // Create matrix using nested while loops
    i = 0
    wakati (i < data_size) {
        row = []
        j = 0
        wakati (j < data_size) {
            // Complex conditional logic
            kama (i == j) {
                row.sukuma(i * j + 1)
            } au kama (i > j) {
                row.sukuma((i - j) * 2)
            } sivyo {
                row.sukuma((j - i) * 3)
            }
            j = j + 1
        }
        matrix.sukuma(row)
        i = i + 1
    }
    
    // Process matrix with for...in loops and conditionals
    total_sum = 0
    even_count = 0
    odd_count = 0
    
    kwa row_idx, row ktk matrix {
        row_sum = 0
        kwa col_idx, value ktk row {
            row_sum = row_sum + value
            
            kama (value % 2 == 0) {
                even_count = even_count + 1
            } sivyo {
                odd_count = odd_count + 1
            }
            
            // Nested conditional with while loop
            kama (value > 10) {
                temp = value
                counter = 0
                wakati (temp > 1 && counter < 5) {
                    temp = temp / 2
                    counter = counter + 1
                }
                total_sum = total_sum + temp
            } sivyo {
                total_sum = total_sum + value
            }
        }
        
        // Classify row based on sum
        kama (row_sum < 50) {
            kama (results["small_rows"] == tupu) {
                results["small_rows"] = 1
            } sivyo {
                results["small_rows"] = results["small_rows"] + 1
            }
        } au kama (row_sum < 100) {
            kama (results["medium_rows"] == tupu) {
                results["medium_rows"] = 1
            } sivyo {
                results["medium_rows"] = results["medium_rows"] + 1
            }
        } sivyo {
            kama (results["large_rows"] == tupu) {
                results["large_rows"] = 1
            } sivyo {
                results["large_rows"] = results["large_rows"] + 1
            }
        }
    }
    
    results["total_sum"] = total_sum
    results["even_count"] = even_count
    results["odd_count"] = odd_count
    results["matrix_size"] = data_size * data_size
    
    rudisha results
}

// Run the complex algorithm
result = complex_algorithm(50)

andika("Complex algorithm results:")
kwa key, value ktk result {
    andika(key + ": " + aina(value))
}

andika("\nTest 2: Stress testing with recursive functions and loops")

fanya recursive_with_loops = unda(n, depth) {
    kama (depth <= 0 || n <= 0) {
        rudisha 0
    }
    
    sum = 0
    i = 1
    wakati (i <= n) {
        kama (i % 2 == 0) {
            // Even numbers: use recursion
            sum = sum + recursive_with_loops(i / 2, depth - 1)
        } au kama (i % 3 == 0) {
            // Multiples of 3: nested loop
            j = 1
            inner_sum = 0
            wakati (j <= 3) {
                inner_sum = inner_sum + (i * j)
                j = j + 1
            }
            sum = sum + inner_sum
        } sivyo {
            // Other numbers: simple addition
            sum = sum + i
        }
        i = i + 1
    }
    
    rudisha sum
}

recursive_result = recursive_with_loops(20, 5)
andika("Recursive with loops result: " + aina(recursive_result))

andika("\nTest 3: Error handling with control structures")

fanya error_prone_function = unda(data) {
    results = []
    errors = 0
    
    kwa item ktk data {
        kama (aina(item) == "INT") {
            kama (item > 0) {
                // Simulate division
                temp = 100
                divisor = item
                wakati (divisor > 1) {
                    temp = temp / divisor
                    divisor = divisor - 1
                    kama (divisor == 0) {
                        errors = errors + 1
                        temp = 0
                        divisor = 1  // Break condition
                    }
                }
                results.sukuma(temp)
            } sivyo {
                errors = errors + 1
                results.sukuma(0)
            }
        } au kama (aina(item) == "STRING") {
            // Process string
            char_sum = 0
            kwa char ktk item {
                // Simulate character processing
                i = 0
                wakati (i < 3) {
                    char_sum = char_sum + 1
                    i = i + 1
                }                }
                results.sukuma(char_sum)
            } sivyo {
                errors = errors + 1
                results.sukuma(-1)
            }
    }
    
    rudisha {"results": results, "errors": errors}
}

test_data = [5, -2, "hello", 0, "world", 3.14, 10, ""]
error_result = error_prone_function(test_data)
andika("Error handling results:")
andika("Processed items: " + aina(error_result["results"].idadi()))
andika("Errors encountered: " + aina(error_result["errors"]))

andika("\n=== COMBINED STRESS TEST COMPLETED ===")
