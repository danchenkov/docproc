// Package export writes a human-readable report of the current index.
package export

import (
	"context"
	"fmt"
	"io"

	"github.com/danchenkov/docproc/internal/domain"
)

// WriteReport accepts io.Writer, the standard library's small output interface.
// A terminal, file, and test buffer can all receive the same report.
func WriteReport(ctx context.Context, out io.Writer, documents []domain.Document) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(out, "Index complete: %d read.\n", len(documents)); err != nil {
		return fmt.Errorf("write summary: %w", err)
	}
	for _, document := range documents {
		if err := ctx.Err(); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(out, "%s: %d words\n", document.Path, document.WordCount); err != nil {
			return fmt.Errorf("write document report: %w", err)
		}
	}
	return nil
}
