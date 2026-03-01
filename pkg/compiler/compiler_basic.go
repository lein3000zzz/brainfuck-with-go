package compiler

import (
	"brainfuck-compiler-go/pkg/compiler/astbuilder"
	"fmt"
	"io"
	"os"
	"strings"
)

const MemorySizeInBytes = 1024

type BasicCompiler struct {
	logger     logger
	astBuilder astbuilder.AstBuilder
	input      io.Reader
}

func NewBasicCompiler(logger logger, astBuilder astbuilder.AstBuilder, input io.Reader) *BasicCompiler {
	return &BasicCompiler{
		logger:     logger,
		astBuilder: astBuilder,
		input:      input,
	}
}

func (c *BasicCompiler) Compile(program string) (string, error) {
	c.logger.Debug("parsing brainfuck program", "length", len(program))

	ast, err := c.astBuilder.Build(program)
	if err != nil {
		c.logger.Warn("error parsing brainfuck program", "err", err)
		return "", err
	}

	c.logger.Debug("AST built", "instructions", len(ast.Instructions))

	state := &runtimeState{
		memory: [MemorySizeInBytes]byte{},
		output: strings.Builder{},
		input:  c.input,
	}

	if err := c.executeNodes(ast.Instructions, state); err != nil {
		return state.output.String(), err
	}

	return state.output.String(), nil
}

type runtimeState struct {
	memory  [MemorySizeInBytes]byte
	pointer int
	output  strings.Builder
	input   io.Reader
}

// executeNodes walks a slice of AST nodes and executes them.
func (c *BasicCompiler) executeNodes(nodes []astbuilder.Node, state *runtimeState) error {
	for _, node := range nodes {
		if err := c.executeNode(node, state); err != nil {
			c.logger.Warn("error executing node", "err", err)
			return err
		}
	}
	return nil
}

// executeNode executes a single AST node.
func (c *BasicCompiler) executeNode(node astbuilder.Node, state *runtimeState) error {
	switch node.Type() {
	case astbuilder.NodeIncPtr:
		state.pointer++
		if state.pointer >= MemorySizeInBytes {
			c.logger.Error("%s: %d", ErrPointerOutOfBounds.Error(), state.pointer)
			return fmt.Errorf("%w: %d", ErrPointerOutOfBounds, state.pointer)
		}

	case astbuilder.NodeDecPtr:
		state.pointer--
		if state.pointer < 0 {
			c.logger.Error("%s: %d", ErrPointerOutOfBounds.Error(), state.pointer)
			return fmt.Errorf("%w: %d", ErrPointerOutOfBounds, state.pointer)
		}

	case astbuilder.NodeIncVal:
		state.memory[state.pointer]++

	case astbuilder.NodeDecVal:
		state.memory[state.pointer]--

	case astbuilder.NodeOutput:
		state.output.WriteByte(state.memory[state.pointer])

	case astbuilder.NodeInput:
		if state.input == nil {
			return ErrInputUnavailable
		}
		buf := make([]byte, 1)
		_, err := state.input.Read(buf)
		if err != nil {
			if err == io.EOF {
				// On EOF, set cell to 0 (common convention).
				state.memory[state.pointer] = 0
			} else {
				return err
			}
		} else {
			state.memory[state.pointer] = buf[0]
		}

	case astbuilder.NodeLoop:
		loop := node.(*astbuilder.LoopNode)
		for state.memory[state.pointer] != 0 {
			if err := c.executeNodes(loop.Children, state); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *BasicCompiler) CompileFromFile(filePath string) (string, error) {
	program, err := os.ReadFile(filePath)
	if err != nil {
		c.logger.Warn("error compiling program", "err", err)
		return "", err
	}

	return c.Compile(string(program))
}
