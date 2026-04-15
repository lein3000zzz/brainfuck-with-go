package bfcompiler

// Compiler compiles and executes a brainfuck program, returning the output.
type Compiler interface {
	Compile(program string) (string, error)
	CompileFromFile(filePath string) (string, error)
}

type logger interface {
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}
