package main

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Highlight struct {
	text  string //text of the highlight
	title string //title of the source
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
		err := rows.Scan(&highlight.title, &highlight.text)
		if err != nil {
			panic(err)
		}
		highlight.title = getDocumentPath(highlight.title)
		highlights = append(highlights, highlight)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return highlights
}

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

	return filepath.Base(path)
}
