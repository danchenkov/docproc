package main

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"
)

// A table keeps inputs and expected results together. Each row becomes a named
// subtest, so a failing scenario can be rerun with go test -run TestRun/name.
func TestRun(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		want      string
		wantError string
	}{
		{name: "default fixture", args: []string{"index"}, want: "Index complete: 1 read.\ndata/sample.txt:"},
		{name: "explicit fixture", args: []string{"index", "-file", "internal/import/testdata/hello.txt"}, want: "Index complete: 1 read.\ninternal/import/testdata/hello.txt: 2 words\n"},
		{name: "unicode", args: []string{"index", "-file", "internal/import/testdata/unicode.txt"}, want: "internal/import/testdata/unicode.txt: 4 words\n"},
		{name: "help", args: []string{"--help"}, want: "usage: docproc index"},
		{name: "command help", args: []string{"index", "-h"}},
		{name: "no command", wantError: "usage:"},
		{name: "unknown command", args: []string{"other"}, wantError: "unknown command"},
		{name: "bad flag", args: []string{"index", "-unknown"}, wantError: "flag provided but not defined"},
		{name: "missing flag value", args: []string{"index", "-file"}, wantError: "flag needs an argument"},
		{name: "extra argument", args: []string{"index", "extra"}, wantError: "unexpected arguments"},
		{name: "missing file", args: []string{"index", "-file", "missing.txt"}, wantError: "read"},
		{name: "unsupported format", args: []string{"index", "-file", "sample.json"}, wantError: ".txt extension"},
		{name: "empty file", args: []string{"index", "-file", "internal/import/testdata/empty.txt"}, wantError: "contains no text"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			err := run(context.Background(), tt.args, &stdout, &stderr)
			if tt.wantError != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantError) {
					t.Fatalf("error = %v; want %q", err, tt.wantError)
				}
				if stdout.Len() != 0 {
					t.Fatalf("failed command printed success: %q", stdout.String())
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(stdout.String(), tt.want) {
				t.Fatalf("output = %q; want %q", stdout.String(), tt.want)
			}
		})
	}
}

func TestRunCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var out bytes.Buffer
	if err := run(ctx, []string{"index"}, &out, &out); !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v", err)
	}
	if out.Len() != 0 {
		t.Fatalf("cancelled command wrote %q", out.String())
	}
}
