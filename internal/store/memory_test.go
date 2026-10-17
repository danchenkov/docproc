package store

import (
	"context"
	"errors"
	"testing"

	"github.com/danchenkov/docproc/internal/domain"
)

func TestMemory(t *testing.T) {
	ctx := context.Background()
	index := NewMemory()
	// Save out of order, then replace one path. This tests observable behavior,
	// rather than relying on how the map is implemented.
	for _, document := range []domain.Document{
		{Path: "b.txt", WordCount: 1},
		{Path: "a.txt", WordCount: 2},
		{Path: "b.txt", WordCount: 3},
	} {
		if err := index.Save(ctx, document); err != nil {
			t.Fatal(err)
		}
	}
	got, err := index.All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].Path != "a.txt" || got[1].WordCount != 3 {
		t.Fatalf("index = %+v", got)
	}
	got[0].WordCount = 99
	again, err := index.All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if again[0].WordCount != 2 {
		t.Fatal("caller modified stored data")
	}
}

func TestMemoryCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	index := NewMemory()
	if err := index.Save(ctx, domain.Document{Path: "a.txt"}); !errors.Is(err, context.Canceled) {
		t.Fatalf("Save error = %v", err)
	}
	if _, err := index.All(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("All error = %v", err)
	}
	documents, err := index.All(context.Background())
	if err != nil || len(documents) != 0 {
		t.Fatalf("cancelled save changed index: %+v, %v", documents, err)
	}
}
