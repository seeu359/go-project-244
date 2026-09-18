package code

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const fixturePath = "testdata/fixture"

func TestGenDiffFlat(t *testing.T) {
	tests := []struct {
		name         string
		file1, file2 string
		expected     string
	}{
		{
			name:     "json removed, added and changed keys",
			file1:    "file1.json",
			file2:    "file2.json",
			expected: "expected_flat.txt",
		},
		{
			name:     "json identical files",
			file1:    "file1.json",
			file2:    "file1.json",
			expected: "expected_identical.txt",
		},
		{
			name:     "yaml removed, added and changed keys",
			file1:    "file1.yml",
			file2:    "file2.yml",
			expected: "expected_flat.txt",
		},
		{
			name:     "yaml identical files",
			file1:    "file1.yml",
			file2:    "file1.yml",
			expected: "expected_identical.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := readFixture(t, tt.expected)

			got, err := GenDiff(
				filepath.Join(fixturePath, tt.file1),
				filepath.Join(fixturePath, tt.file2),
				"stylish",
			)
			if err != nil {
				t.Fatalf("GenDiff() unexpected error: %v", err)
			}
			if got != expected {
				t.Errorf("GenDiff() mismatch:\ngot:\n%s\nwant:\n%s", got, expected)
			}
		})
	}
}

func TestGenDiffErrors(t *testing.T) {
	tests := []struct {
		name         string
		file1, file2 string
	}{
		{"missing first file", "missing.json", "file1.json"},
		{"missing second file", "file1.json", "missing.json"},
		{"unsupported extension", "note.txt", "file1.json"},
		{"invalid json", "invalid.json", "file1.json"},
		{"invalid yaml", "invalid.yml", "file1.yml"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GenDiff(
				filepath.Join(fixturePath, tt.file1),
				filepath.Join(fixturePath, tt.file2),
				"stylish",
			); err == nil {
				t.Error("GenDiff() expected error, got nil")
			}
		})
	}
}

func readFixture(t *testing.T, name string) string {
	t.Helper()

	data, err := os.ReadFile(filepath.Join(fixturePath, name))
	if err != nil {
		t.Fatalf("read fixture %q: %v", name, err)
	}
	return strings.TrimRight(string(data), "\n")
}
