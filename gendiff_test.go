package code

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const fixturePath = "testdata/fixture"

func TestGenDiff(t *testing.T) {
	tests := []struct {
		name     string
		file1    string
		file2    string
		format   string
		expected string
	}{
		{
			name:     "nested json, stylish",
			file1:    "file1.json",
			file2:    "file2.json",
			format:   "stylish",
			expected: "expected_stylish.txt",
		},
		{
			name:     "nested yaml, stylish",
			file1:    "file1.yml",
			file2:    "file2.yml",
			format:   "stylish",
			expected: "expected_stylish.txt",
		},
		{
			name:     "default format is stylish",
			file1:    "file1.json",
			file2:    "file2.json",
			format:   "",
			expected: "expected_stylish.txt",
		},
		{
			name:     "plain format",
			file1:    "file1.json",
			file2:    "file2.json",
			format:   "plain",
			expected: "expected_plain.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := readFixture(t, tt.expected)

			got, err := GenDiff(
				filepath.Join(fixturePath, tt.file1),
				filepath.Join(fixturePath, tt.file2),
				tt.format,
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

func TestGenDiffUnknownFormat(t *testing.T) {
	_, err := GenDiff(
		filepath.Join(fixturePath, "file1.json"),
		filepath.Join(fixturePath, "file2.json"),
		"bogus",
	)
	if err == nil {
		t.Error("GenDiff() expected error for unknown format, got nil")
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
