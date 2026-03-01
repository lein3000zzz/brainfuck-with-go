package main

import (
	"brainfuck-compiler-go/pkg/compiler"
	"brainfuck-compiler-go/pkg/compiler/astbuilder"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	logger := slog.Default()

	builder := astbuilder.NewBasicAstBuilder(logger)
	c := compiler.NewBasicCompiler(logger, builder, os.Stdin)

	output, err := c.CompileFromFile("helloworld.bf")
	if err != nil {
		logger.Error("error: %v\n", err)
		os.Exit(1)
	}

	// Hello World
	program := `>+++++++++[<++++++++>-]<.>+++++++[<++++>-]<+.+++++++..+++.[-]` +
		`>++++++++[<++++>-]<.>+++++++++++[<++++++++>-]<-.--------.+++` +
		`.------.--------.[-]>++++++++[<++++>-]<+.[-]++++++++++.`
	ast, _ := builder.Build(program)

	logger.Info(astbuilder.Visualize(ast))

	fmt.Print(output)
}
