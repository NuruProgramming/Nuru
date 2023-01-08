package evaluator

import (
	"fmt"
	"testing"

	"github.com/AvicennaJr/Nuru/lexer"
	"github.com/AvicennaJr/Nuru/object"
	"github.com/AvicennaJr/Nuru/parser"
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
			"Mstari 0: Aina Hazilingani: NAMBA + BOOLEAN",
		},
		{
			"5 + kweli; 5;",
			"Mstari 0: Aina Hazilingani: NAMBA + BOOLEAN",
		},
		{
			"-kweli",
			"Mstari 0: Operesheni Haielweki: -BOOLEAN",
		},
		{
			"kweli + sikweli",
			"Mstari 0: Operesheni Haielweki: BOOLEAN + BOOLEAN",
		},
		{
			"5; kweli + sikweli; 5",
			"Mstari 0: Operesheni Haielweki: BOOLEAN + BOOLEAN",
		},
		{
			"kama (10 > 1) { kweli + sikweli;}",
			"Mstari 0: Operesheni Haielweki: BOOLEAN + BOOLEAN",
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
			"Mstari 3: Operesheni Haielweki: BOOLEAN + BOOLEAN",
		},
		{
			"bangi",
			"Mstari 0: Neno Halifahamiki: bangi",
		},
		{
			`"Habari" - "Habari"`,
			"Mstari 0: Operesheni Haielweki: NENO - NENO",
		},
		{
			`{"jina": "Avi"}[unda(x) {x}];`,
			"Mstari 0: Samahani, UNDO (FUNCTION) haitumiki kama key",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object return, got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, tt.expectedMessage) {
			t.Errorf("wrong error message, expected=%q, got=%q", fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, tt.expectedMessage), errObj.Message)
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
		t.Fatalf("function haas wrong paramters,Parameters=%+v", unda.Parameters)
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

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`idadi("")`, 0},
		{`idadi("four")`, 4},
		{`idadi("hello world")`, 11},
		{`idadi(1)`, "Samahani, hii function haitumiki na NAMBA"},
		{`idadi("one", "two")`, "Hoja hazilingani, tunahitaji=1, tumepewa=2"},
		{`jumla()`, "Hoja hazilingani, tunahitaji=1, tumepewa=0"},
		{`jumla("")`, "Samahani, hii function haitumiki na NENO"},
		{`jumla(1)`, "Samahani, hii function haitumiki na NAMBA"},
		{`jumla([1,2,3])`, 6},
		{`jumla([1,2,3.4])`, 6.4},
		{`jumla([1.1,2.5,3.4])`, 7},
		{`jumla([1.1,2.5,"q"])`, "Samahani namba tu zinahitajika"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
        
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case float64:
             testFloatObject(t, evaluated, float64(expected))

		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("Object is not Error, got=%T(%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, expected) {
				t.Errorf("Wrong eror message, expected=%q, got=%q", expected, errObj.Message)
			}
		}
	}
}

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
