package main

import (
	"brainfuck-compiler-go/pkg/bfcompiler"
	"brainfuck-compiler-go/pkg/bfcompiler/astbuilder"
	"brainfuck-compiler-go/pkg/toplang/bftranslator"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	logger := slog.Default()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: compiler <file.bf|file.tl>")
		os.Exit(1)
	}
	filePath := os.Args[1]

	builder := astbuilder.NewBasicAstBuilder(logger)
	compiler := bfcompiler.NewBasicCompiler(logger, builder, os.Stdin)

	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".bf":
		output, err := compiler.CompileFromFile(filePath)
		if err != nil {
			logger.Error("brainfuck compilation error", "err", err)
			os.Exit(1)
		}
		fmt.Print(output)

	case ".tl":
		source, err := os.ReadFile(filePath)
		if err != nil {
			logger.Error("failed to read file", "err", err)
			os.Exit(1)
		}

		translator := bftranslator.NewTranslator()

		//astStr, err := translator.VisualizeAST(string(source))
		//if err != nil {
		//	logger.Error("parse error", "err", err)
		//	os.Exit(1)
		//}
		//fmt.Fprintln(os.Stderr, "=== ANTLR Parse Tree (AST) ===")
		//fmt.Fprintln(os.Stderr, astStr)
		//fmt.Fprintln(os.Stderr, "==============================")

		bfCode, err := translator.Translate(string(source))
		if err != nil {
			logger.Error("translation error", "err", err)
			os.Exit(1)
		}

		logger.Warn(bfCode)

		logger.Debug("generated brainfuck", "length", len(bfCode))

		output, err := compiler.Compile(bfCode)
		if err != nil {
			logger.Error("brainfuck execution error", "err", err)
			os.Exit(1)
		}
		fmt.Print(output)

	default:
		fmt.Fprintf(os.Stderr, "unsupported file extension: %s (use .bf or .tl)\n", ext)
		os.Exit(1)
	}
}
