package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Highlight struct {
	document_path string
	desc          string
}

func main() {
	var highlights = getHighlights()
	fmt.Println(highlights[0].document_path)
	fmt.Println(highlights[0].desc)
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
