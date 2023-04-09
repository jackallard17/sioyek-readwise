package main

import (
	"fmt"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var highlights = getHighlights()
	testHighlight := highlights[0]

	fileName := filepath.Base(getDocumentPath(testHighlight.document_path))

	fmt.Println(fileName) // Output: file.txt
	fmt.Println(testHighlight.desc)
}
