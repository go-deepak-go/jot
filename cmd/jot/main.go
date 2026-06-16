package main

import (
	"fmt"
	"os"

	"jot/notes"
)

type Command struct {
	Name string
	ID string
	Query string
	Tag string
	Notebook string
}

func parseArgs(args []string) (Command, error) {
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: jot <command> [args]")
	}
	name := args[1]
	rest := args[2:]

	switch name {
	case "list":
		return Command{Name: "list"}, nil

	case "view":
		if len(rest) < 1 {
			return Command{}, fmt.Errorf("usage: jot view <id>")
		}
		return Command{Name: "view", ID: rest[0]}, nil
	case "search":
		if len(rest) < 1 {
			return Command{}, fmt.Errorf("usage: jot search <query>")
		}
		return Command{Name: "search", Query: rest[0]}, nil
	case "notebook":
		if len(rest) < 1 {
			return Command{}, fmt.Errorf("usage: jot notebook <name>")
		}
		return Command{Name: "notebook", Notebook: rest[0]}, nil
	case "pin":
		if len(rest) < 1 {
			return Command{}, fmt.Errorf("usage: jot pin <id>")
		}
		return Command{Name: "pin", ID: rest[0]}, nil
	case "tag":
		if len(rest) < 2 {
			return Command{}, fmt.Errorf("usage: jot tag <id> <tag>")
		}
		return Command{Name: "tag", ID: rest[0], Tag: rest[1]}, nil
	default:
		return Command{}, fmt.Errorf("unknown command: %s", name)
	}
}
 
func listNotes(store notes.Store) {
	for _, note := range store.GetAllNotes() {
		fmt.Println(note.Summary())
	}
}
 
func viewNote(store notes.Store, id string) {
	note, ok := store.GetNote(id)
	if !ok {
		fmt.Println("no note with ID", id)
		return
	}
	fmt.Println(note.Summary())
	fmt.Println("---")
	fmt.Println(note.Body)
}
 
func searchNotes(store notes.Store, query string) {
	matches := store.Search(query)
	if len(matches) == 0 {
		fmt.Println("no matches for:", query)
		return
	}
	for _, note := range matches {
		fmt.Println(note.Summary())
	}
}
 
func notebookNotes(store notes.Store, notebook string) {
	result := store.InNotebook(notebook)
	if len(result) == 0 {
		fmt.Println("no notes in notebook:", notebook)
		return
	}
	for _, note := range result {
		fmt.Println(note.Summary())
	}
}
 
func pinNote(store notes.Store, id string) {
	err := store.Pin(id)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("pinned note", id)
	fmt.Println(store[id].Summary())
}
 
func tagNote(store notes.Store, id string, tag string) {
	err := store.AddTag(id, tag)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("tagged note", id, "with", tag)
	fmt.Println(store[id].Summary())
}
 
func main() {
	store := notes.Store{
		"1": {ID: "1", Title: "Go maps", Body: "Maps are great.", Notebook: "Go", Tags: []string{"go"}, Pinned: false},
		"2": {ID: "2", Title: "Sourdough", Body: "Mix flour and water.", Notebook: "Cooking", Tags: []string{}, Pinned: false},
		"3": {ID: "3", Title: "Go slices", Body: "Slices are fun.", Notebook: "Go", Tags: []string{}, Pinned: false},
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
	case "tag":
		tagNote(store, command.ID, command.Tag)
	}
}