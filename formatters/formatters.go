package formatters

import (
	"fmt"

	"code/differ"
)

type Formatter func([]differ.Node) string

func Get(format string) (Formatter, error) {
	switch format {
	case "", "stylish":
		return Stylish, nil
	case "plain":
		return Plain, nil
	default:
		return nil, fmt.Errorf("unknown format %q", format)
	}
}
