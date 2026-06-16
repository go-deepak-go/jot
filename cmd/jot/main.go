package main

import (
	"fmt"
	"os"

	"jot/notes"
)

type Command struct {
	Name string
	Args []string
}

func parseArgs(args []string) (Command, error) {
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: jot <command> [args]")
	}
	return Command{
		Name: args[1],
		Args: args[2:],
	}, nil
}

func listNotes(store notes.Store) {
	for _, note := range store.GetAllNotes() {
		fmt.Println(note.Summary())
	}
}

func viewNote(store notes.Store, args []string) {
	if len(args) < 1 {
		fmt.Println("usage: jot view <id>")
		return
	}
	note, ok := store.GetNote(args[0])
	if !ok {
		fmt.Println("no note with ID", args[0])
		return
	}
	fmt.Println(note.Summary())
	fmt.Println("---")
	fmt.Println(note.Body)
}

func searchNotes(store notes.Store, args []string) {
	if len(args) < 1 {
		fmt.Println("usage: jot search <query>")
		return
	}
	matches := store.Search(args[0])
	if len(matches) == 0 {
		fmt.Println("no matches for:", args[0])
		return
	}
	for _, note := range matches {
		fmt.Println(note.Summary())
	}
}

func notebookNotes(store notes.Store, args []string) {
	if len(args) < 1 {
		fmt.Println("usage: jot notebook <name>")
		return
	}
	result := store.InNotebook(args[0])
	if len(result) == 0 {
		fmt.Println("no notes in notebook:", args[0])
		return
	}
	for _, note := range result {
		fmt.Println(note.Summary())
	}
}

func pinNote(store notes.Store, args []string) {
	if len(args) < 1 {
		fmt.Println("usage: jot pin <id>")
		return
	}
	err := store.Pin(args[0])
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("pinned note", args[0])
	fmt.Println(store[args[0]].Summary())
}

func tagNote(store notes.Store, args []string) {
	if len(args) < 2 {
		fmt.Println("usage: jot tag <id> <tag>")
		return
	}
	err := store.AddTag(args[0], args[1])
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("tagged note", args[0], "with", args[1])
	fmt.Println(store[args[0]].Summary())
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
			viewNote(store, command.Args)
		case "search":
			searchNotes(store, command.Args)
		case "notebook":
			notebookNotes(store, command.Args)
		case "pin":
			pinNote(store, command.Args)
		case "tag":
			tagNote(store, command.Args)
		default:
			fmt.Println("unknown command:", command.Name)
			fmt.Println("commands: list, view, search, notebook, pin, tag")
	}
}