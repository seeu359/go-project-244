package parsers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func ParseFile(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file %q: %w", path, err)
	}

	switch strings.ToLower(filepath.Ext(path)) {
	case ".json":
		return parseJSON(data)
	case ".yaml", ".yml":
		return parseYAML(data)
	default:
		return nil, fmt.Errorf("unsupported file format: %q", path)
	}
}

func parseJSON(data []byte) (map[string]any, error) {
	var parsed map[string]any
	if err := json.Unmarshal(data, &parsed); err != nil {
		return nil, fmt.Errorf("parse JSON: %w", err)
	}
	return parsed, nil
}

func parseYAML(data []byte) (map[string]any, error) {
	var parsed map[string]any
	if err := yaml.Unmarshal(data, &parsed); err != nil {
		return nil, fmt.Errorf("parse YAML: %w", err)
	}
	return parsed, nil
}
