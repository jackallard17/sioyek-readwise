package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	docs := processLocalDocuments()

	fmt.Println(docs[70].path)
	fmt.Println(docs[70].highlights[2])
}
