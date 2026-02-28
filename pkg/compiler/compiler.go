package compiler

type Compiler interface {
	Compile(program string) (string, error)
}

type logger interface {
}
