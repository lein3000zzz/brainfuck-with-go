package astbuilder

import (
	"strings"
)

const (
	connectorStr       = "├── "
	connectorLastStr   = "└── "
	childPrefixBaseStr = "│   "
	spaceStr           = "    "
)

func Visualize(program *Program) string {
	if program == nil {
		return "Program (nil)"
	}

	var sb strings.Builder

	sb.WriteString("Program")
	sb.WriteByte('\n')
	visualizeNodes(&sb, program.Instructions, "")

	return sb.String()
}

func visualizeNodes(sb *strings.Builder, nodes []Node, prefix string) {
	for i, node := range nodes {
		isLast := i == len(nodes)-1

		connector := connectorStr
		if isLast {
			connector = connectorLastStr
		}

		sb.WriteString(prefix)
		sb.WriteString(connector)
		sb.WriteString(node.Type().String())
		sb.WriteByte('\n')

		if loop, ok := node.(*LoopNode); ok {
			childPrefix := prefix + childPrefixBaseStr
			if isLast {
				childPrefix = prefix + spaceStr
			}
			visualizeNodes(sb, loop.Children, childPrefix)
		}
	}
}
