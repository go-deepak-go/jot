package main

import (
	"fmt"
	"os"

	notes "github.com/go-deepak-go/jot"
)

type Command struct {
	Name string
	ID string
	Query string
	Tag string
	Notebook string
	Title string
	Body string
}
 
func main() {
	store, err := notes.Load("notes.json")
	if err != nil {
		store = notes.New()
	}
 
	command, err := parseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
 
	switch command.Name {
	case "list":
		listNotes(store)
	case "view":
		viewNote(store, command.ID)
	case "search":
		searchNotes(store, command.Query)
	case "notebook":
		notebookNotes(store, command.Notebook)
	case "pin":
		pinNote(store, command.ID)
		saveOrPrint(store)
	case "tag":
		tagNote(store, command.ID, command.Tag)
		saveOrPrint(store)
	case "add":
		addNote(&store, command.Title, command.Body)
		saveOrPrint(store)
	case "delete":
		deleteNote(store, command.ID)
		saveOrPrint(store)
	}
}
