package bfcompiler

import (
	"errors"
)

var (
	// ErrPointerOutOfBounds is returned when the memory pointer moves outside [0, MemorySizeInBytes).
	ErrPointerOutOfBounds = errors.New("pointer out of bounds")

	// ErrInputUnavailable is returned when a ',' instruction is executed but no input reader is configured.
	ErrInputUnavailable = errors.New("input instruction executed but no input source is available")
)
