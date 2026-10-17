package process

import (
	"context"
	"testing"

	"github.com/danchenkov/docproc/internal/domain"
)

func TestDocument(t *testing.T) {
	tests := []struct {
		name, path, text string
		words            int
		wantError        bool
	}{
		{"plain", "sample.txt", "Hello Go!", 2, false},
		{"unicode whitespace", "sample.txt", "Hello\t世界\r\nGo\u2003text", 4, false},
		{"empty", "sample.txt", "", 0, true},
		{"whitespace only", "sample.txt", " \r\n\t", 0, true},
		{"missing path", "", "Hello", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := domain.Document{Path: tt.path, Content: tt.text}
			got, err := Document(context.Background(), input)
			if (err != nil) != tt.wantError {
				t.Fatalf("error = %v", err)
			}
			if tt.wantError {
				return
			}
			if got.WordCount != tt.words || got.Content != tt.text {
				t.Fatalf("document = %+v", got)
			}
			// Processing a value should leave the caller's original document unchanged.
			if input.WordCount != 0 {
				t.Fatal("input was mutated")
			}
		})
	}
}
