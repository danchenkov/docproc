// Command docproc imports one TXT document into a temporary, in-memory index.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/danchenkov/docproc/internal/app"
	"github.com/danchenkov/docproc/internal/export"
	"github.com/danchenkov/docproc/internal/import"
	"github.com/danchenkov/docproc/internal/store"
)

func main() {
	// Ctrl+C cancels this context. Passing it down lets operations check whether
	// the caller still wants the work. defer releases resources when main returns.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := run(ctx, os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "docproc:", err)
		// os.Exit skips deferred functions, so clean up before exiting.
		stop()
		os.Exit(1)
	}
}

// run receives its inputs instead of reading globals, so tests can exercise
// actual CLI logic with argument slices and buffers without starting a process.
func run(ctx context.Context, args []string, stdout, stderr io.Writer) error {
	const usage = "usage: docproc index [-file data/sample.txt]"
	if len(args) == 0 {
		return errors.New(usage)
	}
	if args[0] == "-h" || args[0] == "--help" {
		_, err := fmt.Fprintln(stdout, usage)
		return err
	}
	if args[0] != "index" {
		return fmt.Errorf("unknown command %q; %s", args[0], usage)
	}
	// ContinueOnError returns parsing errors instead of exiting the process.
	flags := flag.NewFlagSet("index", flag.ContinueOnError)
	flags.SetOutput(stderr)
	source := flags.String("file", "data/sample.txt", "path to a single TXT document")
	if err := flags.Parse(args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("unexpected arguments; %s", usage)
	}

	// Constructors wire concrete collaborators together here. Each invocation
	// gets a fresh index: no data is persisted to disk.
	index := store.NewMemory()
	application := app.New(importer.NewTXT(), index)
	if err := application.Index(ctx, *source); err != nil {
		return err
	}
	documents, err := index.All(ctx)
	if err != nil {
		return err
	}
	return export.WriteReport(ctx, stdout, documents)
}
