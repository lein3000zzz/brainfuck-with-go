package bftranslator

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

var (
	// ErrUndeclaredVariable is returned when a variable is used but not declared.
	ErrUndeclaredVariable = errors.New("undeclared variable")

	// ErrRedeclaredVariable is returned when a variable is declared twice.
	ErrRedeclaredVariable = errors.New("variable already declared")

	// ErrOutOfMemory is returned when there are no more free cells on the tape.
	ErrOutOfMemory = errors.New("out of brainfuck tape memory")

	// ErrParseFailed is returned when the TopLang source fails to parse.
	ErrParseFailed = errors.New("parse error")

	// ErrIntOverflow is returned when an integer literal exceeds 255.
	ErrIntOverflow = errors.New("integer literal exceeds cell range 0-255")

	// ErrUnsupportedOperator is returned for an unknown operator token.
	ErrUnsupportedOperator = errors.New("unsupported operator")

	// ErrInvalidString is returned when a string literal contains invalid escape sequences.
	ErrInvalidString = errors.New("invalid string literal")
)

type errCollector struct {
	antlr.DefaultErrorListener
	errs []string
}

func (e *errCollector) SyntaxError(_ antlr.Recognizer, _ interface{}, line, col int, msg string, _ antlr.RecognitionException) {
	e.errs = append(e.errs, fmt.Sprintf("line %d:%d %s", line, col, msg))
}
