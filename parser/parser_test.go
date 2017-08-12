package parser

import (
	"testing"
	"../ast"
	"../lexer"
	"../parser"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foo = 838383;
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements is not len=3, got %d", len(program.Statements))
	}

	tests := []struct{
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		stm := program.Statements[i]

		if !testLetStatement(t, stm, tt.expectedIdentifier) {
			return 
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Literal is not let, got=%q", s.TokenLiteral())
		return false
	}

	letStm, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not ast.LetStatement, got=%T", s)
		return false
	}

	if letStm.Name.Value != name {
		t.Errorf("letStm.Name.Value is not '%s', got '%s'", name, letStm.Name.Value)
		return false
	}

	if letStm.Token.Literal != name {
		t.Errorf("letStm.Token.Literal is not '%s' got '%s'", name, letStm.Token.Literal)
		return false
	}

	return true
}
