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
			"Aina Hazilingani: NAMBARI + BOOLEAN",
		},
		{
			"5 + kweli; 5;",
			"Aina Hazilingani: NAMBARI + BOOLEAN",
		},
		{
			"-kweli",
			"Operesheni Haielweki: -BOOLEAN",
		},
		{
			"kweli + sikweli",
			"Operesheni Haielweki: BOOLEAN + BOOLEAN",
		},
		{
			"5; kweli + sikweli; 5",
			"Operesheni Haielweki: BOOLEAN + BOOLEAN",
		},
		{
			"kama (10 > 1) { kweli + sikweli;}",
			"Operesheni Haielweki: BOOLEAN + BOOLEAN",
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
			"Operesheni Haielweki: BOOLEAN + BOOLEAN",
		},
		{
			"bangi",
			"Neno Halifahamiki: bangi",
		},
		{
			`"Habari" - "Habari"`,
			"Operesheni Haielweki: NENO - NENO",
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
			t.Errorf("wrong error message, expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestLetStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"acha a = 5; a;", 5},
		{"acha a = 5 * 5; a;", 25},
		{"acha a = 5; acha b = a; b;", 5},
		{"acha a = 5; acha b = a; acha c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2 ;};"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not a Function, got=%T(%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function haas wrong paramters,Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not x, got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q, got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"acha mfano = fn(x) {x;}; mfano(5);", 5},
		{"acha mfano = fn(x) {rudisha x;}; mfano(5);", 5},
		{"acha double = fn(x) { x * 2;}; double(5);", 10},
		{"acha add = fn(x, y) {x + y;}; add(5,5);", 10},
		{"acha add = fn(x, y) {x + y;}; add(5 + 5, add(5, 5));", 20},
		{"fn(x) {x;}(5)", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
acha newAdder = fn(x) {
	fn(y) { x + y};
};

acha addTwo = newAdder(2);
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
