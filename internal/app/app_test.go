package app

import (
	"context"
	"errors"
	"testing"

	"github.com/danchenkov/docproc/internal/domain"
	"github.com/danchenkov/docproc/internal/import"
	"github.com/danchenkov/docproc/internal/store"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		name, path string
		words      int
		wantError  bool
	}{
		{"plain", "../import/testdata/hello.txt", 2, false},
		{"unicode", "../import/testdata/unicode.txt", 4, false},
		{"empty", "../import/testdata/empty.txt", 0, true},
		{"missing", "missing.txt", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			index := store.NewMemory()
			err := New(importer.NewTXT(), index).Index(ctx, tt.path)
			if (err != nil) != tt.wantError {
				t.Fatalf("error = %v", err)
			}
			documents, err := index.All(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if tt.wantError {
				if len(documents) != 0 {
					t.Fatal("failed import was stored")
				}
				return
			}
			if len(documents) != 1 || documents[0].WordCount != tt.words {
				t.Fatalf("index = %+v", documents)
			}
		})
	}
}

// A tiny test double satisfies Store implicitly and lets us simulate a failure
// that the in-memory implementation normally cannot produce.
type failingStore struct{ err error }

func (s failingStore) Save(context.Context, domain.Document) error { return s.err }

func TestIndexPreservesStoreError(t *testing.T) {
	failure := errors.New("storage unavailable")
	err := New(importer.NewTXT(), failingStore{failure}).Index(context.Background(), "../import/testdata/hello.txt")
	if !errors.Is(err, failure) {
		t.Fatalf("error = %v; want original failure", err)
	}
}
