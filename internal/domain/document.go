// Package domain defines shared data without knowing about CLI flags or storage.
package domain

// Document represents one text file and its initial derived metadata.
// Exported names start with a capital letter so other packages can use them.
type Document struct {
	Path      string // Cleaned source path; also the in-memory index key.
	Content   string // Original UTF-8 text, including whitespace.
	WordCount int    // Whitespace-separated fields, not linguistic words.
}
