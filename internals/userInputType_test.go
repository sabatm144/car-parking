package internals

import (
	"testing"
)

func TestReadFromFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
	}{
		{"Invalid Type - 1", ".txt"},
		{"Invalid Type - 2", "examples/file_input.tx"},
		{"Valid", "examples/file_input.txt"},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadFromFile(tt.filePath)
		})
	}
}

func TestReadFromConsole(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"read from console"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadFromConsole()
		})
	}
}
