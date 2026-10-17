// Package app coordinates importing, processing, and storing a document.
package app

import (
	"context"
	"fmt"

	"github.com/danchenkov/docproc/internal/domain"
	"github.com/danchenkov/docproc/internal/process"
)

// Importer describes only the behavior this consumer needs. Go implements
// interfaces implicitly: a type with this method satisfies the interface
// without an "implements" declaration. Define interfaces where they are used.
type Importer interface {
	Import(context.Context, string) (domain.Document, error)
}

// Store is the other consumer boundary. App needs to save a document but does
// not need to know whether storage uses a map or (later) a file.
type Store interface {
	Save(context.Context, domain.Document) error
}

// App holds dependencies as fields rather than package-level global variables.
type App struct {
	importer Importer
	store    Store
}

// New makes required dependencies explicit. Callers supply non-nil collaborators.
// *App is a pointer to the constructed value; & takes the address of a value.
func New(importer Importer, store Store) *App {
	return &App{importer: importer, store: store}
}

// Index saves a document only after importing and processing both succeed.
func (a *App) Index(ctx context.Context, path string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	document, err := a.importer.Import(ctx, path)
	if err != nil {
		// %w adds context while preserving the original error for errors.Is.
		return fmt.Errorf("import document: %w", err)
	}
	document, err = process.Document(ctx, document)
	if err != nil {
		return fmt.Errorf("process document: %w", err)
	}
	if err := a.store.Save(ctx, document); err != nil {
		return fmt.Errorf("store document: %w", err)
	}
	return nil // A nil error means success.
}
