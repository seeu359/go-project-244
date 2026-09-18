package code

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"code/parsers"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	data1, err := parsers.ParseFile(filepath1)
	if err != nil {
		return "", err
	}

	data2, err := parsers.ParseFile(filepath2)
	if err != nil {
		return "", err
	}

	return buildDiff(data1, data2), nil
}

func buildDiff(data1, data2 map[string]any) string {
	keys := append(slices.Collect(maps.Keys(data1)), slices.Collect(maps.Keys(data2))...)
	slices.Sort(keys)
	keys = slices.Compact(keys)

	var diff strings.Builder
	diff.WriteString("{\n")
	for _, key := range keys {
		value1, in1 := data1[key]
		value2, in2 := data2[key]

		switch {
		case !in2:
			fmt.Fprintf(&diff, "  - %s: %v\n", key, value1)
		case !in1:
			fmt.Fprintf(&diff, "  + %s: %v\n", key, value2)
		case value1 == value2:
			fmt.Fprintf(&diff, "    %s: %v\n", key, value1)
		default:
			fmt.Fprintf(&diff, "  - %s: %v\n", key, value1)
			fmt.Fprintf(&diff, "  + %s: %v\n", key, value2)
		}
	}
	diff.WriteString("}")

	return diff.String()
}
