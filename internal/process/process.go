// Package process validates text and derives its first metadata.
package process

import (
	"context"
	"errors"
	"strings"

	"github.com/danchenkov/docproc/internal/domain"
)

// Document returns a processed copy. A plain function is enough for this one
// algorithm; later assignments can introduce variation when it is needed.
func Document(ctx context.Context, document domain.Document) (domain.Document, error) {
	if err := ctx.Err(); err != nil {
		return domain.Document{}, err
	}
	if strings.TrimSpace(document.Path) == "" {
		return domain.Document{}, errors.New("document path is required")
	}
	// Fields splits on Unicode whitespace, handling spaces, tabs and newlines.
	words := strings.Fields(document.Content)
	if len(words) == 0 {
		return domain.Document{}, errors.New("document contains no text")
	}
	document.WordCount = len(words)
	if err := ctx.Err(); err != nil {
		return domain.Document{}, err
	}
	return document, nil
}
