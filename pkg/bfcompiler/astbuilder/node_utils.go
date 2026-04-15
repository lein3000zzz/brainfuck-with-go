package astbuilder

func NewProgram(instructions ...Node) *Program {
	return &Program{Instructions: instructions}
}

func NewLoop(children ...Node) *LoopNode {
	return &LoopNode{Children: children}
}

func NewIncPtr() *IncPtrNode {
	return &IncPtrNode{}
}

func NewDecPtr() *DecPtrNode {
	return &DecPtrNode{}
}

func NewIncVal() *IncValNode {
	return &IncValNode{}
}

func NewDecVal() *DecValNode {
	return &DecValNode{}
}

func NewOutput() *OutputNode {
	return &OutputNode{}
}

func NewInput() *InputNode {
	return &InputNode{}
}
