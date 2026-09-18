package formatters

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"code/differ"
)

func Stylish(nodes []differ.Node) string {
	var diff strings.Builder
	diff.WriteString("{\n")
	renderNodes(&diff, nodes, 1)
	diff.WriteString("}")

	return diff.String()
}

func renderNodes(diff *strings.Builder, nodes []differ.Node, depth int) {
	for _, node := range nodes {
		renderNode(diff, node, depth)
	}
}

func renderNode(diff *strings.Builder, node differ.Node, depth int) {
	indent := strings.Repeat(" ", depth*4-2)

	switch node.Status {
	case differ.StatusNested:
		fmt.Fprintf(diff, "%s  %s: {\n", indent, node.Key)
		renderNodes(diff, node.Children, depth+1)
		fmt.Fprintf(diff, "%s}\n", strings.Repeat(" ", depth*4))
	case differ.StatusAdded:
		renderValueLine(diff, indent, "+", node.Key, node.NewValue, depth)
	case differ.StatusRemoved:
		renderValueLine(diff, indent, "-", node.Key, node.OldValue, depth)
	case differ.StatusChanged:
		renderValueLine(diff, indent, "-", node.Key, node.OldValue, depth)
		renderValueLine(diff, indent, "+", node.Key, node.NewValue, depth)
	case differ.StatusUnchanged:
		renderValueLine(diff, indent, " ", node.Key, node.OldValue, depth)
	}
}

func renderValueLine(diff *strings.Builder, indent, marker, key string, value any, depth int) {
	nested, isMap := value.(map[string]any)
	if !isMap {
		fmt.Fprintf(diff, "%s%s %s: %s\n", indent, marker, key, stringifyScalar(value))
		return
	}

	fmt.Fprintf(diff, "%s%s %s: {\n", indent, marker, key)
	dumpObject(diff, nested, depth+1)
	fmt.Fprintf(diff, "%s}\n", strings.Repeat(" ", depth*4))
}

func dumpObject(diff *strings.Builder, object map[string]any, depth int) {
	indent := strings.Repeat(" ", depth*4)

	for _, key := range slices.Sorted(maps.Keys(object)) {
		value := object[key]

		nested, isMap := value.(map[string]any)
		if !isMap {
			fmt.Fprintf(diff, "%s%s: %s\n", indent, key, stringifyScalar(value))
			continue
		}

		fmt.Fprintf(diff, "%s%s: {\n", indent, key)
		dumpObject(diff, nested, depth+1)
		fmt.Fprintf(diff, "%s}\n", indent)
	}
}

func stringifyScalar(value any) string {
	if value == nil {
		return "null"
	}
	return fmt.Sprintf("%v", value)
}
