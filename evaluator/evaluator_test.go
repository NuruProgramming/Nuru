package evaluator

import (
	"fmt"
	"testing"
	"time"

	"github.com/NuruProgramming/Nuru/lexer"
	"github.com/NuruProgramming/Nuru/object"
	"github.com/NuruProgramming/Nuru/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2", 16},
		{"2 / 2 + 1", 2},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"2**3", 8.0},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"kweli", true},
		{"sikweli", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 > 1", false},
		{"1 < 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"kweli == kweli", true},
		{"sikweli == sikweli", true},
		{"kweli == sikweli", false},
		{"kweli != sikweli", true},
		{"sikweli != kweli", true},
		{"(1 < 2) == kweli", true},
		{"!kweli", false},
		{"!sikweli", true},
		{"!tupu", true},
		{"!'kitu'", false},
		{"2 > 1 && 1 < 4", true},
		{"2 > 1 && 1 > 4", false},
		{"2 < 1 && 1 < 4", false},
		{"2 < 1 && 1 > 4", false},
		{"5 < 2 || 3 > 2", true},
		{"5 == 5 || 4 == 4", true},
		{"5 > 2 || 3 < 2", true},
		{"5 < 2 || 3 < 2", false},
		{"5 >= 2", true},
		{"5 <= 2", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!kweli", false},
		{"!sikweli", true},
		{"!5", false},
		{"!!kweli", true},
		{"!!sikweli", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)

	if !ok {
		t.Errorf("Object is not Integer, got=%T(%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)

	if !ok {
		t.Errorf("Object is not Float, got=%T(%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f", result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean, got=%T(%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value, got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"kama (kweli) {10}", 10},
		{"kama (sikweli) {10}", nil},
		{"kama (1) {10}", 10},
		{"kama (1 < 2) {10}", 10},
		{"kama (1 > 2) {10}", nil},
		{"kama (1 > 2) {10} sivyo {20}", 20},
		{"kama (1 < 2) {10} sivyo {20}", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not null, got=%T(+%v)", obj, obj)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"rudisha 10", 10},
		{"rudisha 10; 9;", 10},
		{"rudisha 2 * 5; 9;", 10},
		{"9; rudisha 2 * 5; 9;", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + kweli",
			"Mstari 1: Aina Hazilingani: NAMBA + BOOLEAN",
		},
		{
			"5 + kweli; 5;",
			"Mstari 1: Aina Hazilingani: NAMBA + BOOLEAN",
		},
		{
			"-kweli",
			"Mstari 1: Operesheni Haieleweki: -BOOLEAN",
		},
		{
			"kweli + sikweli",
			"Mstari 1: Operesheni Haieleweki: BOOLEAN + BOOLEAN",
		},
		{
			"5; kweli + sikweli; 5",
			"Mstari 1: Operesheni Haieleweki: BOOLEAN + BOOLEAN",
		},
		{
			"kama (10 > 1) { kweli + sikweli;}",
			"Mstari 1: Operesheni Haieleweki: BOOLEAN + BOOLEAN",
		},
		{
			`
kama (10 > 1) {
	kama (10 > 1) {
		rudisha kweli + kweli;
	}

	rudisha 1;
}
			`,
			"Mstari 4: Operesheni Haieleweki: BOOLEAN + BOOLEAN",
		},
		{
			"bangi",
			"Mstari 1: Neno Halifahamiki: bangi",
		},
		{
			`"Habari" - "Habari"`,
			"Mstari 1: Operesheni Haieleweki: NENO - NENO",
		},
		{
			`{"jina": "Avi"}[unda(x) {x}];`,
			"Mstari 1: Samahani, UNDO (FUNCTION) haitumiki kama ufunguo",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object return, got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != fmt.Sprintf(tt.expectedMessage) {
			t.Errorf("wrong error message, expected=%q, got=%q", fmt.Sprintf(tt.expectedMessage), errObj.Message)
		}
	}
}

func TestLetStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"fanya a = 5; a;", 5},
		{"fanya a = 5 * 5; a;", 25},
		{"fanya a = 5; fanya b = a; b;", 5},
		{"fanya a = 5; fanya b = a; fanya c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "unda(x) { x + 2 ;};"

	evaluated := testEval(input)
	unda, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not a Function, got=%T(%+v)", evaluated, evaluated)
	}

	if len(unda.Parameters) != 1 {
		t.Fatalf("function has wrong parameters,Parameters=%+v", unda.Parameters)
	}

	if unda.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not x, got=%q", unda.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if unda.Body.String() != expectedBody {
		t.Fatalf("body is not %q, got=%q", expectedBody, unda.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"fanya mfano = unda(x) {x;}; mfano(5);", 5},
		{"fanya mfano = unda(x) {rudisha x;}; mfano(5);", 5},
		{"fanya double = unda(x) { x * 2;}; double(5);", 10},
		{"fanya add = unda(x, y) {x + y;}; add(5,5);", 10},
		{"fanya add = unda(x, y) {x + y;}; add(5 + 5, add(5, 5));", 20},
		{"unda(x) {x;}(5)", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
fanya newAdder = unda(x) {
	unda(y) { x + y};
};

fanya addTwo = newAdder(2);
addTwo(2);
`
	testIntegerObject(t, testEval(input), 4)
}

func TestStringLiteral(t *testing.T) {
	input := `"Habari yako!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("Object is not string, got=%T(%+v)", evaluated, evaluated)
	}

	if str.Value != "Habari yako!" {
		t.Errorf("String has wrong value, got=%q", str.Value)
	}
}

func TestStringconcatenation(t *testing.T) {
	input := `"Mambo" + " " + "Vipi" + "?"`

	evaluated := testEval(input)

	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not a string, got=%T(%+v)", evaluated, evaluated)
	}

	if str.Value != "Mambo Vipi?" {
		t.Errorf("String has wrong value, got=%q", str.Value)
	}
}

func TestStringMultiplyInteger(t *testing.T) {
	input := `"Mambo" * 4`

	evaluated := testEval(input)

	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not a string, got=%T(%+v)", evaluated, evaluated)
	}

	if str.Value != "MamboMamboMamboMambo" {
		t.Errorf("String has wrong value, got=%q", str.Value)
	}
}

// func TestBuiltinFunctions(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected interface{}
// 	}{
// 		{`jumla()`, "Hoja hazilingani, tunahitaji=1, tumepewa=0"},
// 		{`jumla("")`, "Samahani, hii function haitumiki na NENO"},
// 		{`jumla(1)`, "Samahani, hii function haitumiki na NAMBA"},
// 		{`jumla([1,2,3])`, 6},
// 		{`jumla([1,2,3.4])`, 6.4},
// 		{`jumla([1.1,2.5,3.4])`, 7},
// 		{`jumla([1.1,2.5,"q"])`, "Samahani namba tu zinahitajika"},
// 	}

// 	for _, tt := range tests {
// 		evaluated := testEval(tt.input)

// 		switch expected := tt.expected.(type) {
// 		case int:
// 			testIntegerObject(t, evaluated, int64(expected))
// 		case float64:
// 			testFloatObject(t, evaluated, float64(expected))

// 		case string:
// 			errObj, ok := evaluated.(*object.Error)
// 			if !ok {
// 				t.Errorf("Object is not Error, got=%T(%+v)", evaluated, evaluated)
// 				continue
// 			}
// 			if errObj.Message != fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, expected) {
// 				t.Errorf("Wrong eror message, expected=%q, got=%q", expected, errObj.Message)
// 			}
// 		}
// 	}
// }

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("Object is not an Array, got=%T(%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("Array has wrong number of elements, got=%d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"fanya i = 0; [1][i];",
			1,
		},
		{
			"fanya myArr = [1, 2, 3]; myArr[2];",
			3,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestDictLiterals(t *testing.T) {
	input := `fanya two = "two";
{
	"one": 10 - 9,
	two: 1 +1,
	"thr" + "ee": 6 / 2,
	4: 4,
	kweli: 5,
	sikweli: 6
}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Dict)
	if !ok {
		t.Fatalf("Eval didn't return a dict, got=%T(%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		TRUE.HashKey():                             5,
		FALSE.HashKey():                            6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Dict has wrong number of pairs, got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("No pair for give key")
		}

		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestDictIndexExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`fanya key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{kweli: 5}[kweli]`,
			5,
		},
		{
			`{sikweli: 5}[sikweli]`,
			5,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestPrefixInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"-4",
			-4,
		},
		{
			"+5",
			5,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if !ok {
			t.Errorf("Object is not an integer")
		}
		testIntegerObject(t, evaluated, int64(integer))
	}
}

func TestPrefixFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"-4.4",
			-4.4,
		},
		{
			"+5.5",
			5.5,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		float, ok := tt.expected.(float64)
		if !ok {
			t.Errorf("Object is not a float")
		}
		testFloatObject(t, evaluated, float)
	}
}

func TestInExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			"'a' ktk 'habari'",
			true,
		},
		{
			"'c' ktk 'habari'",
			false,
		},
		{
			"1 ktk [1, 2, 3]",
			true,
		},
		{
			"4 ktk [1, 2, 3]",
			false,
		},
		{
			"'a' ktk {'a': 'apple', 'b': 'banana'}",
			true,
		},
		{
			"'apple' ktk {'a': 'apple', 'b': 'banana'}",
			false,
		},
		{
			"'c' ktk {'a': 'apple', 'b': 'banana'}",
			false,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestArrayConcatenation(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"['a', 'b', 'c'] + [1, 2, 3]",
			"[a, b, c, 1, 2, 3]",
		},
		{
			"[1, 2, 3] * 4",
			"[1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3]",
		},
		{
			"4 * [1, 2, 3]",
			"[1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3]",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		arr, ok := evaluated.(*object.Array)
		if !ok {
			t.Fatalf("Object is not an array, got=%T(%+v)", evaluated, evaluated)
		}

		if arr.Inspect() != tt.expected {
			t.Errorf("Array has wrong values, got=%s want=%s", arr.Inspect(), tt.expected)
		}
	}
}

func TestDictConcatenation(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]string
	}{
		{
			input:    "{'a': 'apple', 'b': 'banana'} + {'c': 'cat'}",
			expected: map[string]string{"a": "apple", "b": "banana", "c": "cat"},
		},
		{
			input:    "{'a':'bbb'} + {'a':'ccc'}",
			expected: map[string]string{"a": "ccc"},
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		dict, ok := evaluated.(*object.Dict)
		if !ok {
			t.Fatalf("Object is not an dict, got=%T(%+v)", evaluated, evaluated)
		}

		if len(dict.Pairs) != len(tt.expected) {
			t.Errorf("Dictionary has wrong number of pairs, got=%d want=%d", len(dict.Pairs), len(tt.expected))
		}
	}
}

func TestPostfixExpression(t *testing.T) {
	inttests := []struct {
		input    string
		expected int64
	}{
		{
			"a=5; a++",
			6,
		},
		{
			"a=5; a--",
			4,
		},
	}

	for _, tt := range inttests {
		evaluated := testEval(tt.input)
		integer, ok := evaluated.(*object.Integer)
		if !ok {
			t.Fatalf("Object is not an integer, got=%T(%+v)", evaluated, evaluated)
		}
		testIntegerObject(t, integer, tt.expected)
	}
	floattests := []struct {
		input    string
		expected float64
	}{
		{
			"a=5.5; a++",
			6.5,
		},
		{
			"a=5.5; a--",
			4.5,
		},
	}

	for _, tt := range floattests {
		evaluated := testEval(tt.input)
		float, ok := evaluated.(*object.Float)
		if !ok {
			t.Fatalf("Object is not an float, got=%T(%+v)", evaluated, evaluated)
		}
		testFloatObject(t, float, tt.expected)
	}
}

func TestWhileLoop(t *testing.T) {
	input := `
	i = 10
	wakati (i > 0){
		i--
	}
	i
	`

	evaluated := testEval(input)
	i, ok := evaluated.(*object.Integer)
	if !ok {
		t.Fatalf("Object is not an integer, got=%T(+%v)", evaluated, evaluated)
	}

	if i.Value != 0 {
		t.Errorf("Incorrect value, want=0 got=%d", i.Value)
	}
}

func TestForLoop(t *testing.T) {
	input := `
	output = ""
	kwa i ktk "mojo" {
		output += i
	}
	output
	`
	evaluated := testEval(input)
	i, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("Object is not a string, got=%T(+%v)", evaluated, evaluated)
	}

	if i.Value != "mojo" {
		t.Errorf("Wrong value: want=%s got=%s", "mojo", i.Value)
	}
}

func TestBreakLoop(t *testing.T) {
	input := `
	i = 0
	wakati (i < 10) {
		kama (i == 5) {
			vunja
		}
		i++
	}
	i
	`
	evaluated := testEval(input)
	i, ok := evaluated.(*object.Integer)
	if !ok {
		t.Fatalf("Object is not an integer, got=%T(+%v)", evaluated, evaluated)
	}

	if i.Value != 5 {
		t.Errorf("Wrong value: want=5, got=%d", i.Value)
	}

	input = `
	output = ""
	kwa i ktk "mojo" {
		output += i
		kama (i == 'o') {
			vunja
		}
	}
	output
	`

	evaluatedFor := testEval(input)
	j, ok := evaluatedFor.(*object.String)
	if !ok {
		t.Fatalf("Object is not a string, got=%T", evaluated)
	}

	if j.Value != "mo" {
		t.Errorf("Wrong value: want=%s, got=%s", "mo", j.Value)
	}
}

func TestContinueLoop(t *testing.T) {
	input := `
	i = 0
	wakati (i < 10) {
		i++
		kama (i == 5) {
			endelea
		}
		i++
	}
	i
	`
	evaluated := testEval(input)
	i, ok := evaluated.(*object.Integer)
	if !ok {
		t.Fatalf("Object is not an integer, got=%T(+%v)", evaluated, evaluated)
	}

	if i.Value != 11 {
		t.Errorf("Wrong value: want=11, got=%d", i.Value)
	}

	input = `
	output = ""
	kwa i ktk "mojo" {
		kama (i == 'o') {
			endelea
		}
		output += i
	}
	output
	`

	evaluatedFor := testEval(input)
	j, ok := evaluatedFor.(*object.String)
	if !ok {
		t.Fatalf("Object is not a string, got=%T", evaluated)
	}

	if j.Value != "mj" {
		t.Errorf("Wrong value: want=%s, got=%s", "mj", j.Value)
	}
}

func TestSwitchStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`
			i = 5
			badili (i) {
				ikiwa 2 {
					output = 2
				}
				ikiwa 5 {
					output = 5
				}
				kawaida {
					output = "haijulikani"
				}
			}
			output
			`,
			5,
		},
		{
			`
			i = 5
			badili (i) {
				ikiwa 2 {
					output = 2
				}
				kawaida {
					output = "haijulikani"
				}
			}
			output
			`,
			"haijulikani",
		},
		{
			`
			i = 5
			badili (i) {
				ikiwa 5 {
					output = 5
				}
				ikiwa 2 {
					output = 2
				}
				kawaida {
					output = "haijulikani"
				}
			}
			output
			`,
			5,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			s, ok := evaluated.(*object.String)
			if !ok {
				t.Fatalf("Object is not a string, got=%T", evaluated)
			}

			if s.Value != tt.expected {
				t.Errorf("Wrong Value, want='haijulikani', got=%s", s.Value)
			}

		}
	}
}

func TestAssignEqual(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"a = 5; a += 5",
			10,
		},
		{
			"a = 5; a -= 5",
			0,
		},
		{
			"a = 5; a *= 10",
			50,
		},
		{
			"a = 100; a /= 4",
			25,
		},
		{
			`
		a = [1, 2, 3]
		a[0] += 500
		a[0]
		`,
			501,
		},
		{
			`
		a = "mambo"
		a += " vipi"
		`,
			"mambo vipi",
		},
		{
			"a = 5.5; a += 4.5",
			10.0,
		},
		{
			"a = 11.3; a -= 0.8",
			10.5,
		},
		{
			"a = 0.4; a /= 2",
			0.2,
		},
		{
			"a = 0.1; a *= 10",
			1.0,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case float64:
			testFloatObject(t, evaluated, float64(expected))
		case string:
			s, ok := evaluated.(*object.String)
			if !ok {
				t.Fatalf("Object not a string, got=%T", evaluated)
			}

			if s.Value != tt.expected {
				t.Errorf("Wrong value, want=%s, got=%s", tt.expected, s.Value)
			}
		}
	}
}

func TestStringMethods(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"'mambo'.idadi()",
			5,
		},
		{
			"'mambo'.herufikubwa()",
			"MAMBO",
		},
		{
			"'MaMbO'.herufindogo()",
			"mambo",
		},
		{
			"'habari'.gawa('a')",
			"[h, b, ri]",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			switch eval := evaluated.(type) {
			case *object.String:
				s, ok := evaluated.(*object.String)
				if !ok {
					t.Fatalf("Object not of type string, got=%T", eval)
				}
				if s.Value != tt.expected {
					t.Errorf("Wrong value: want=%s, got=%s", tt.expected, s.Value)
				}
			case *object.Array:
				arr, ok := evaluated.(*object.Array)
				if !ok {
					t.Fatalf("Object not of type array, got=%T", eval)
				}

				if arr.Inspect() != tt.expected {
					t.Errorf("Wrong value: want=%s, got=%s", tt.expected, arr.Inspect())
				}
			}
		}
	}
}

func TestTimeModule(t *testing.T) {
	input := `
	tumia muda
	muda.hasahivi()
	`

	evaluated := testEval(input)
	muda, ok := evaluated.(*object.Time)
	if !ok {
		t.Fatalf("Object is not a time object, got=%T", evaluated)
	}

	_, err := time.Parse("15:04:05 02-01-2006", muda.TimeValue)
	if err != nil {
		t.Errorf("Wrong time value: got=%v", err)
	}
}
