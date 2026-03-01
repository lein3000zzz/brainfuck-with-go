package astbuilder

import (
	"fmt"
)

type BasicAstBuilder struct {
	logger logger
}

type Frame struct {
	Loop     *LoopNode
	Position int // Position of the '[' in source
}

func NewBasicAstBuilder(logger logger) *BasicAstBuilder {
	return &BasicAstBuilder{
		logger: logger,
	}
}

func (b *BasicAstBuilder) Build(source string) (*Program, error) {
	b.logger.Debug("building AST from source", "source_length", len(source))

	program := NewProgram()
	loopStack := make([]Frame, 0, 16)
	current := &program.Instructions

	for i, char := range source {
		var node Node
		switch char {
		case IncreasePointerSign:
			node = NewIncPtr()
		case DecreasePointerSign:
			node = NewDecPtr()
		case IncreaseValueByPointerSign:
			node = NewIncVal()
		case DecreaseValueByPointerSign:
			node = NewDecVal()
		case PrintOneCharSign:
			node = NewOutput()
		case InputOneCharSign:
			node = NewInput()
		case WhileLoopStartSign:
			loop := NewLoop()
			*current = append(*current, loop)
			loopStack = append(loopStack, Frame{
				Loop:     loop,
				Position: i,
			})
			current = &loop.Children
			continue
		case WhileLoopEndSign:
			if len(loopStack) == 0 {
				b.logger.Debug("unmatched close bracket", "position", i)
				return nil, fmt.Errorf("%w: at postition %d", ErrUnmatchedCloseBracket, i)
			}
			loopStack = loopStack[:len(loopStack)-1]
			if len(loopStack) == 0 {
				current = &program.Instructions
			} else {
				current = &loopStack[len(loopStack)-1].Loop.Children
			}
			continue
		default:
			// Все остальное игнорируется
			continue
		}
		*current = append(*current, node)
	}

	if len(loopStack) > 0 {
		b.logger.Debug("unmatched open bracket", "position", loopStack[len(loopStack)-1].Position)
		return nil, fmt.Errorf("%w: at position %d", ErrUnmatchedOpenBracket, loopStack[len(loopStack)-1].Position)
	}

	b.logger.Debug("AST built successfully", "instructions", len(program.Instructions))

	return program, nil
}
