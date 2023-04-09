package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Highlight struct {
	document_path string
	desc          string
}

func main() {
	var highlights = getHighlights()
	testHighlight := highlights[0]

	fileName := filepath.Base(getDocumentPath(testHighlight.document_path))

	fmt.Println(fileName) // Output: file.txt
	fmt.Println(testHighlight.desc)
}

func getHighlights() []Highlight {
	var highlights []Highlight
	db, err := sql.Open("sqlite3", os.Getenv("HOME")+"/.local/share/sioyek/shared.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT document_path, desc FROM highlights")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var highlight Highlight
		err := rows.Scan(&highlight.document_path, &highlight.desc)
		if err != nil {
			panic(err)
		}
		highlights = append(highlights, highlight)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return highlights
}

// another local database has a table document_hash. there are two fields, hash and path
// take hash as paramater, return path as string
func getDocumentPath(hash string) string {
	var path string
	db, err := sql.Open("sqlite3", os.Getenv("HOME")+"/.local/share/sioyek/local.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT path FROM document_hash WHERE hash = ?", hash)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&path)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return path
}
