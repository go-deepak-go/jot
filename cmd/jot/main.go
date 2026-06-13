package main

import (
	"fmt"
	"jot/notes"
)

func main() {
	note := notes.Note{
		ID: "1",
		Title: "ABC",
		Body: "qwerty",
		Notebook: "Test",
		Tags: []string{"go", "learning"},
		Pinned: true,
	}
	fmt.Println(note.Summary())
}