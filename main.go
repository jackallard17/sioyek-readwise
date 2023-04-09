package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	highlights := getHighlights()

	printHighlights(highlights)
}

func printHighlights(highlights []Highlight) {
	//print each highlight on a new line
	for _, highlight := range highlights {
		fmt.Println(highlight)
		fmt.Println("--------------------------------------------------")
	}

	fmt.Println("Total highlights: ", len(highlights))
}
