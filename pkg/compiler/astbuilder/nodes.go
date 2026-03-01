package astbuilder

import "fmt"

// NodeType represents the type of AST node.
type NodeType int

const (
	NodeIncPtr NodeType = iota // >
	NodeDecPtr                 // <
	NodeIncVal                 // +
	NodeDecVal                 // -
	NodeOutput                 // .
	NodeInput                  // ,
	NodeLoop                   // [...]
)

// nodeTypeNames maps NodeType to a human-readable name.
var nodeTypeNames = map[NodeType]string{
	NodeIncPtr: "IncPtr (>)",
	NodeDecPtr: "DecPtr (<)",
	NodeIncVal: "IncVal (+)",
	NodeDecVal: "DecVal (-)",
	NodeOutput: "Output (.)",
	NodeInput:  "Input (,)",
	NodeLoop:   "Loop ([...])",
}

// String returns a human-readable name for the NodeType.
func (t NodeType) String() string {
	if name, ok := nodeTypeNames[t]; ok {
		return name
	}
	return fmt.Sprintf("Unknown(%d)", int(t))
}

// Node is the interface that all AST nodes implement.
type Node interface {
	Type() NodeType
}

// Program is the root of the AST, containing a sequence of instructions.
type Program struct {
	Instructions []Node
}

// IncPtrNode represents the ">" operator: move pointer right.
type IncPtrNode struct{}

func (n *IncPtrNode) Type() NodeType {
	return NodeIncPtr
}

// DecPtrNode represents the "<" operator: move pointer left.
type DecPtrNode struct{}

func (n *DecPtrNode) Type() NodeType {
	return NodeDecPtr
}

// IncValNode represents the "+" operator: increment the current cell.
type IncValNode struct{}

func (n *IncValNode) Type() NodeType {
	return NodeIncVal
}

// DecValNode represents the "-" operator: decrement the current cell.
type DecValNode struct{}

func (n *DecValNode) Type() NodeType {
	return NodeDecVal
}

// OutputNode represents the "." operator: output the current cell as a character.
type OutputNode struct{}

func (n *OutputNode) Type() NodeType {
	return NodeOutput
}

// InputNode represents the "," operator: read one character into the current cell.
type InputNode struct{}

func (n *InputNode) Type() NodeType {
	return NodeInput
}

// LoopNode represents a "[...]" loop: execute children while the current cell is nonzero.
type LoopNode struct {
	Children []Node
}

func (n *LoopNode) Type() NodeType {
	return NodeLoop
}
