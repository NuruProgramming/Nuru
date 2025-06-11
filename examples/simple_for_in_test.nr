// Simple test for for...in loops

andika("=== SIMPLE FOR...IN TEST ===")

// Test 1: Small array iteration
andika("Test 1: Small array iteration")
small_array = []
i = 0
wakati (i < 10) {
    small_array.sukuma(i)
    i = i + 1
}

counter = 0
kwa item ktk small_array {
    counter = counter + 1
}
andika("Small array size: " + aina(counter))

// Test 2: String iteration
andika("Test 2: String iteration")
test_string = "hello"
char_count = 0
kwa char ktk test_string {
    char_count = char_count + 1
}
andika("String length: " + aina(char_count))

andika("=== SIMPLE TEST COMPLETED ===")
