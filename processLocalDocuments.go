package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Highlight struct {
	document_path string
	desc          string
}

type Document struct {
	path       string
	highlights []string
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

func processLocalDocuments() []Document {
	var documents []Document
	highlights := getHighlights()
	for _, highlight := range highlights {
		fileName := getDocumentPath(highlight.document_path)
		//check if the document is already in the document list
		var documentExists bool
		for _, document := range documents {
			if document.path == fileName {
				documentExists = true
			}
		}
		if documentExists {
			for index, document := range documents {
				if document.path == fileName {
					documents[index].highlights = append(documents[index].highlights, highlight.desc)
				}
			}
		} else {
			var document Document
			document.path = fileName
			document.highlights = append(document.highlights, highlight.desc)
			documents = append(documents, document)
		}
	}
	return documents
}
