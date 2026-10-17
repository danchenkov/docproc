// Package store supplies the initial in-memory document index.
package store

import (
	"context"
	"sort"

	"github.com/danchenkov/docproc/internal/domain"
)

// Memory owns a map keyed by source path. The CLI uses it sequentially;
// concurrent access would require synchronization in a future assignment.
type Memory struct{ documents map[string]domain.Document }

// NewMemory allocates the map. Reading a nil map is allowed, but writing to one
// panics, which is why callers construct Memory through this function.
func NewMemory() *Memory {
	return &Memory{documents: make(map[string]domain.Document)}
}

// Save adds or replaces the document at this path, preventing duplicate entries.
func (m *Memory) Save(ctx context.Context, document domain.Document) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	m.documents[document.Path] = document
	return nil
}

// All returns a separate slice in path order. Map iteration order is unspecified,
// so sorting makes reports repeatable. Document contains only value fields,
// so callers cannot mutate the map by changing an entry in the returned slice.
func (m *Memory) All(ctx context.Context) ([]domain.Document, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	documents := make([]domain.Document, 0, len(m.documents))
	for _, document := range m.documents {
		documents = append(documents, document)
	}
	sort.Slice(documents, func(i, j int) bool { return documents[i].Path < documents[j].Path })
	return documents, nil
}
