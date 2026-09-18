package code

import (
	"fmt"

	"code/differ"
	"code/formatters"
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

	nodes := differ.Diff(data1, data2)

	formatter, err := getFormatter(format)
	if err != nil {
		return "", err
	}

	return formatter(nodes), nil
}

func getFormatter(format string) (func([]differ.Node) string, error) {
	switch format {
	case "", "stylish":
		return formatters.Stylish, nil
	default:
		return nil, fmt.Errorf("unknown format %q", format)
	}
}
