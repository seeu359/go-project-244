package formatters

import (
	"encoding/json"

	"code/differ"
)

func JSON(nodes []differ.Node) string {
	result, _ := json.MarshalIndent(toJSONValue(nodes), "", "  ")
	return string(result)
}

func toJSONValue(nodes []differ.Node) map[string]any {
	result := make(map[string]any, len(nodes))

	for _, node := range nodes {
		switch node.Status {
		case differ.StatusNested:
			result[node.Key] = toJSONValue(node.Children)
		case differ.StatusAdded:
			result["+ "+node.Key] = node.NewValue
		case differ.StatusRemoved:
			result["- "+node.Key] = node.OldValue
		case differ.StatusChanged:
			result["- "+node.Key] = node.OldValue
			result["+ "+node.Key] = node.NewValue
		case differ.StatusUnchanged:
			result[node.Key] = node.OldValue
		}
	}

	return result
}
