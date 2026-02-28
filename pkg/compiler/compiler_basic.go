package compiler

import "log/slog"

type BasicCompiler struct {
	logger *slog.Logger
}

func NewBasicCompiler(logger logger) *BasicCompiler {
	return &BasicCompiler{
		logger: slog.Default(),
	}
}

func (c *BasicCompiler) Compile(program string) (string, error) {

}
