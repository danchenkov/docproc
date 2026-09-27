package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danchenkov/docproc/internal/app"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: docproc index")
	}
	cmd := os.Args[1]
	if cmd != "index" {
		log.Fatal("command must be index")
	}
	flags := flag.NewFlagSet(cmd, flag.ExitOnError)
	flags.Parse(os.Args[2:])

	sourceFile := "data/sample.txt"

	switch cmd {
	case "index":
		result, err := app.Index(sourceFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Index complete: %d read.\n", result)
	}
}
