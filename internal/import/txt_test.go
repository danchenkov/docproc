package importer

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestImport(t *testing.T) {
	tests := []struct{ path, content string }{
		{"testdata/hello.txt", "Hello Go!\n"},
		{"testdata/unicode.txt", "Hello 世界\nLearning Go\n"},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got, err := NewTXT().Import(context.Background(), tt.path)
			if err != nil {
				t.Fatal(err)
			}
			if got.Path != tt.path || got.Content != tt.content || got.WordCount != 0 {
				t.Fatalf("document = %+v", got)
			}
		})
	}
}

func TestImportErrors(t *testing.T) {
	// TempDir is isolated per test and automatically removed by the test runner.
	invalid := filepath.Join(t.TempDir(), "invalid.txt")
	if err := os.WriteFile(invalid, []byte{0xff}, 0600); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"missing.txt", "sample.json", invalid, "testdata"} {
		t.Run(path, func(t *testing.T) {
			if _, err := NewTXT().Import(context.Background(), path); err == nil {
				t.Fatal("expected an error")
			}
		})
	}
	_, err := NewTXT().Import(context.Background(), "missing.txt")
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("original filesystem error lost: %v", err)
	}
}
