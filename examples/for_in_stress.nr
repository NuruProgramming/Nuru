// Stress test for for...in loops

andika("=== FOR...IN LOOP STRESS TESTS ===")

// Test 1: Large array iteration
andika("Test 1: Large array iteration")
big_array = []
i = 0
wakati (i < 10000) {
    big_array.sukuma(i)
    i = i + 1
}

counter = 0
kwa item ktk big_array {
    counter = counter + 1
}
andika("Iterated through array of size: " + aina(counter))

// Test 2: Large dictionary iteration
andika("Test 2: Large dictionary iteration")
big_dict = {}
i = 0
wakati (i < 5000) {
    big_dict["key_" + aina(i)] = "value_" + aina(i)
    i = i + 1
}

key_count = 0
value_sum = 0
kwa key, value ktk big_dict {
    key_count = key_count + 1
    // Process value (simulate work)
    kama (key_count % 100 == 0) {
        // Every 100th iteration, do something
        value_sum = value_sum + key_count
    }
}
andika("Dictionary iteration - keys: " + aina(key_count) + ", value_sum: " + aina(value_sum))

// Test 3: Nested for...in loops
andika("Test 3: Nested for...in loops")
matrix = []
i = 0
wakati (i < 100) {
    row = []
    j = 0
    wakati (j < 100) {
        row.sukuma(i * j)
        j = j + 1
    }
    matrix.sukuma(row)
    i = i + 1
}

total = 0
kwa row ktk matrix {
    kwa cell ktk row {
        total = total + 1
    }
}
andika("Nested for...in processed: " + aina(total) + " cells")

// Test 4: String iteration stress
andika("Test 4: Large string iteration")
long_string = ""
i = 0
wakati (i < 1000) {
    long_string = long_string + "abcdefghij"
    i = i + 1
}

char_count = 0
kwa char ktk long_string {
    char_count = char_count + 1
}
andika("String character count: " + aina(char_count))

// Test 5: Mixed data type array iteration
andika("Test 5: Mixed data types")
mixed_array = []
i = 0
wakati (i < 1000) {
    kama (i % 4 == 0) {
        mixed_array.sukuma(i)
    } au kama (i % 4 == 1) {
        mixed_array.sukuma("string_" + aina(i))
    } au kama (i % 4 == 2) {
        mixed_array.sukuma(i % 2 == 0)
    } sivyo {
        mixed_array.sukuma({"id": i, "value": i * 2})
    }
    i = i + 1
}

type_counts = {"INT": 0, "STRING": 0, "BOOL": 0, "DICT": 0}
kwa item ktk mixed_array {
    item_type = aina(item)
    kama (item_type == "INT") {
        type_counts["INT"] = type_counts["INT"] + 1
    } au kama (item_type == "STRING") {
        type_counts["STRING"] = type_counts["STRING"] + 1
    } au kama (item_type == "BOOL") {
        type_counts["BOOL"] = type_counts["BOOL"] + 1
    } sivyo {
        type_counts["DICT"] = type_counts["DICT"] + 1
    }
}
andika("Mixed types count: " + aina(type_counts))
