package formatters

import (
	"fmt"
	"strings"

	"code/differ"
)

func Plain(nodes []differ.Node) string {
	return strings.Join(collectLines(nodes, ""), "\n")
}

func collectLines(nodes []differ.Node, path string) []string {
	var lines []string

	for _, node := range nodes {
		keyPath := joinPath(path, node.Key)

		switch node.Status {
		case differ.StatusNested:
			lines = append(lines, collectLines(node.Children, keyPath)...)
		case differ.StatusAdded:
			lines = append(lines, fmt.Sprintf("Property '%s' was added with value: %s", keyPath, stringifyPlain(node.NewValue)))
		case differ.StatusRemoved:
			lines = append(lines, fmt.Sprintf("Property '%s' was removed", keyPath))
		case differ.StatusChanged:
			lines = append(lines, fmt.Sprintf("Property '%s' was updated. From %s to %s", keyPath, stringifyPlain(node.OldValue), stringifyPlain(node.NewValue)))
		case differ.StatusUnchanged:
		}
	}

	return lines
}

func joinPath(parent, key string) string {
	if parent == "" {
		return key
	}
	return parent + "." + key
}

func stringifyPlain(value any) string {
	switch typed := value.(type) {
	case nil:
		return "null"
	case string:
		return "'" + typed + "'"
	case map[string]any:
		return "[complex value]"
	default:
		return fmt.Sprintf("%v", typed)
	}
}
