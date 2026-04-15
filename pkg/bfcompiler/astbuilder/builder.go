package astbuilder

type logger interface {
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type AstBuilder interface {
	Build(source string) (*Program, error)
}
