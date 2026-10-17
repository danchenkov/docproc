# docproc

A beginner Go project following [GO_README.md](GO_README.md), through **Initial scope and setup: days 1–3**. The starter reads one UTF-8 TXT document, rejects empty text, counts whitespace-separated words, saves the document in memory, and prints a report. Application and test code use only the Go standard library.

## Run it

The existing `go.mod` declares Go **1.27.1**; use that version or newer. Check your installation with `go version`. The module already exists, so you do not need to run `go mod init` again or download dependencies.

Run these commands from the repository root:

```sh
go run . index
go run . index -file internal/import/testdata/hello.txt
go run . index -h
```

The second command prints:

```text
Index complete: 1 read.
internal/import/testdata/hello.txt: 2 words
```

`go run .` compiles and runs the main package. `index` selects the command, and `-file` changes the input path; the default is `data/sample.txt`. Relative file paths are resolved from your current working directory. A failed command reports an error on standard error and exits with a nonzero status. Ctrl+C requests cancellation.

To build a reusable executable:

```sh
go build -o bin/docproc .
./bin/docproc index -file data/sample.txt
```

The index exists only for the lifetime of the command. Each run starts empty. JSON import, disk persistence, classification, and the pattern assignments are future exercises. The TXT importer reads the whole file into memory, so this baseline targets small documents. Cancellation is cooperative: it is checked between steps, but cannot interrupt `os.ReadFile` itself. The store is for sequential use.

## Find your way through the code

| File or directory | Responsibility |
| --- | --- |
| `go.mod` | Module import path and required Go version. A module contains packages; each directory here is a package. |
| `main.go` | Entry point, flags, Ctrl+C handling, dependency construction, and output. |
| `internal/app` | Coordinates the import → process → store flow. Defines the small interfaces its collaborators satisfy. |
| `internal/domain` | Shared `Document` value: source path, text, and word count. |
| `internal/import` | Reads and validates UTF-8 TXT files. Its Go package name is `importer`, because `import` is a keyword. |
| `internal/process` | Rejects missing paths or blank text and counts words using `strings.Fields`. |
| `internal/store` | Map-backed index; saving the same cleaned path replaces its entry. Reports use sorted paths. |
| `internal/export` | Prints a text report to an `io.Writer`, such as a terminal or test buffer. |
| `data/sample.txt` | Original sample input. |
| `internal/import/testdata` | Small, repeatable test fixtures. Go tooling treats `testdata` as fixture data rather than a package. |

Go's `internal` directory rule prevents unrelated external projects from importing these implementation packages. Imports beginning with `github.com/danchenkov/docproc` refer to packages in this module; building this checkout does not fetch its own code from GitHub.

Read `internal/domain/document.go` first, then the importer, processor, store, app, exporter, and finally `main.go`. Follow that with `main_test.go` to see the pieces exercised together. Comments explain pointers, exported names, constructors, implicit interfaces, error wrapping, and context propagation where they appear.

The app receives its importer and store through `app.New`. This makes dependencies visible and permits a test to substitute a failing store. The processing step stays a plain function while there is one algorithm. `io.Writer` is already a suitable standard interface for reporting. No extra pattern framework is needed for this baseline.

Paths identify documents; they are cleaned but not made absolute or resolved through symlinks. Word counts measure whitespace-separated fields: punctuation stays attached, and languages without spaces do not receive linguistic word segmentation.

## Check your work

```sh
go test ./...
go vet ./...
```

`./...` means the current package and all packages below it. Tests live beside their code in files ending in `_test.go`; functions named `Test...` receive `*testing.T`. Most scenarios use a table of inputs and expectations and `t.Run` for named subtests.

To inspect individual scenarios:

```sh
go test -v ./...
go test -run TestRun/explicit_fixture .
```

The suite covers valid TXT input, Unicode and whitespace, empty and missing files, unsupported formats, invalid UTF-8, CLI argument errors, cancellation, storage replacement, and error propagation. Test fixtures are local; temporary files are cleaned up by `t.TempDir`.

## Written reading resources

Start with the fundamentals, then read the references alongside the corresponding package. All links below lead to written material; no videos are required.

1. [A Tour of Go](https://go.dev/tour/) — interactive text lessons covering syntax, structs, slices, maps, pointers, methods, and interfaces.
2. [Create a Go module](https://go.dev/doc/tutorial/create-module) — explains modules, packages, import paths, and the `go.mod` file; the tutorial series continues through errors, tests, and building.
3. [Organizing a Go module](https://go.dev/doc/modules/layout) — explains package layout and the purpose of `internal` directories.
4. [Effective Go](https://go.dev/doc/effective_go) — read names, functions, methods, interfaces, and errors as you encounter them. This is a language idiom guide, not a complete guide to modern tooling.
5. [Table-driven tests](https://go.dev/wiki/TableDrivenTests) — the table and subtest structure used in this project.
6. [The testing package](https://pkg.go.dev/testing) — test function conventions, `T.Run`, failure reporting, and temporary directories.
7. [The context package](https://pkg.go.dev/context) — cancellation, deadlines, and passing context explicitly to operations.
8. [The flag package](https://pkg.go.dev/flag) — CLI flag parsing and `FlagSet`.
9. [The io.Writer interface](https://pkg.go.dev/io#Writer) — the small output contract shared by terminals, files, and buffers.

Keep [the full study plan](GO_README.md) for the later assignments and its pattern reading list. For this initial block, try tracing an error from a missing file back to the CLI, adding a table row for a new whitespace case, and explaining why re-indexing a path replaces one entry.
