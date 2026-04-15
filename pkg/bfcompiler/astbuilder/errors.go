package astbuilder

import (
	"errors"
)

var (
	// ErrUnmatchedOpenBracket is returned when a '[' has no matching ']'.
	ErrUnmatchedOpenBracket = errors.New("unmatched open bracket")

	// ErrUnmatchedCloseBracket is returned when a ']' has no matching '['.
	ErrUnmatchedCloseBracket = errors.New("unmatched close bracket")
)
