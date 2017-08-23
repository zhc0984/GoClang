package evaluator

import (
	"GoClang/lexer"
	"GoClang/object"
	"GoClang/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T)  {
	tests := []struct{
		input string
		expected int64
	}{
		{"5",5},
		{"10", 10},
	}



	for _, tt := range tests {
		evaluator := testEval(tt.input)
		testIntegerObject(t, evaluator, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParserProgram()
	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool{
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}
	return true
}


