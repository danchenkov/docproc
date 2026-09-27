package app

import (
	"github.com/danchenkov/docproc/internal/process"
)

func Index(sourceFile string) (result int, err error) {
	return process.Index(sourceFile)
}
