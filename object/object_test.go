package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is Avi"}
	diff2 := &String{Value: "My name is Avi"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("string with the same content have different dict keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("String with the same content have different dict keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("Strings with different content have the same dict keys")
	}
}
