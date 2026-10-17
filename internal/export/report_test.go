package export

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/danchenkov/docproc/internal/domain"
)

func TestWriteReport(t *testing.T) {
	tests := []struct {
		name      string
		documents []domain.Document
		want      string
	}{
		{"empty index", nil, "Index complete: 0 read.\n"},
		{"one document", []domain.Document{{Path: "hello.txt", WordCount: 2}}, "Index complete: 1 read.\nhello.txt: 2 words\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			if err := WriteReport(context.Background(), &out, tt.documents); err != nil {
				t.Fatal(err)
			}
			if out.String() != tt.want {
				t.Fatalf("report = %q; want %q", out.String(), tt.want)
			}
		})
	}
}

type failingWriter struct{ err error }

func (w failingWriter) Write([]byte) (int, error) { return 0, w.err }

func TestWriteReportErrors(t *testing.T) {
	failure := errors.New("output unavailable")
	if err := WriteReport(context.Background(), failingWriter{failure}, nil); !errors.Is(err, failure) {
		t.Fatalf("error = %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var out bytes.Buffer
	if err := WriteReport(ctx, &out, nil); !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v", err)
	}
	if out.Len() != 0 {
		t.Fatal("cancelled report wrote output")
	}
}
