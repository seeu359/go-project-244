package differ

import (
	"maps"
	"slices"
)

type Status int

const (
	StatusAdded Status = iota
	StatusRemoved
	StatusChanged
	StatusUnchanged
	StatusNested
)

type Node struct {
	Key      string
	Status   Status
	OldValue any
	NewValue any
	Children []Node
}

func Diff(map1, map2 map[string]any) []Node {
	keys := append(slices.Collect(maps.Keys(map1)), slices.Collect(maps.Keys(map2))...)
	slices.Sort(keys)
	keys = slices.Compact(keys)

	nodes := make([]Node, 0, len(keys))
	for _, key := range keys {
		value1, in1 := map1[key]
		value2, in2 := map2[key]

		switch {
		case !in2:
			nodes = append(nodes, Node{Key: key, Status: StatusRemoved, OldValue: value1})
		case !in1:
			nodes = append(nodes, Node{Key: key, Status: StatusAdded, NewValue: value2})
		default:
			nodes = append(nodes, compareValues(key, value1, value2))
		}
	}

	return nodes
}

func compareValues(key string, value1, value2 any) Node {
	nested1, isMap1 := value1.(map[string]any)
	nested2, isMap2 := value2.(map[string]any)

	switch {
	case isMap1 && isMap2:
		return Node{Key: key, Status: StatusNested, Children: Diff(nested1, nested2)}
	case value1 == value2:
		return Node{Key: key, Status: StatusUnchanged, OldValue: value1}
	default:
		return Node{Key: key, Status: StatusChanged, OldValue: value1, NewValue: value2}
	}
}
