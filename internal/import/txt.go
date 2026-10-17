// Package importer reads source documents. Its directory is named import to
// match the study plan; the package name avoids Go's reserved keyword "import".
package importer

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/danchenkov/docproc/internal/domain"
)

// TXT imports small UTF-8 text files. It has no fields because it holds no state.
type TXT struct{}

// NewTXT constructs a TXT importer; no configuration is needed yet.
func NewTXT() *TXT { return &TXT{} }

// Import reads exactly one file. Directory walking and JSON are later exercises.
func (t *TXT) Import(ctx context.Context, path string) (domain.Document, error) {
	if err := ctx.Err(); err != nil {
		return domain.Document{}, err
	}
	if !strings.EqualFold(filepath.Ext(path), ".txt") {
		return domain.Document{}, fmt.Errorf("source %q must have a .txt extension", path)
	}
	// ReadFile reads the whole file and closes it for us. This simple baseline
	// is for small fixtures, not streaming very large documents.
	content, err := os.ReadFile(path)
	if err != nil {
		return domain.Document{}, fmt.Errorf("read %q: %w", path, err)
	}
	// ReadFile cannot be interrupted through context. Checking afterward keeps
	// cancelled work from continuing into processing and storage.
	if err := ctx.Err(); err != nil {
		return domain.Document{}, err
	}
	if !utf8.Valid(content) {
		return domain.Document{}, fmt.Errorf("source %q is not valid UTF-8", path)
	}
	return domain.Document{Path: filepath.Clean(path), Content: string(content)}, nil
}
