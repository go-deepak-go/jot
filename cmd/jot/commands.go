package main

import (
	"fmt"
	"os"

	"github.com/go-deepak-go/jot/notes"
)

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
	note, _ := store.GetNote(id)
	fmt.Println("pinned note", id)
	fmt.Println(note.Summary())
}

func tagNote(store notes.Store, id string, tag string) {
	err := store.AddTag(id, tag)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	note, _ := store.GetNote(id)
	fmt.Println("tagged note", id, "with", tag)
	fmt.Println(note.Summary())
}

func addNote(store *notes.Store, title string, body string) {
	note, err := store.AddNote(notes.Note{Title: title, Body: body})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("added:", note.Summary())
}

func saveOrPrint(store notes.Store) {
    err := store.Save("notes.json")
    if err != nil {
        fmt.Println("Error saving notes:", err)
    }
}

func deleteNote(store notes.Store, id string) {
	err := store.Delete(id)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("deleted note", id)
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
	case "add":
		if len(rest) < 2 {
			return Command{}, fmt.Errorf("usage: jot add <title> <body>")
		}
		return Command{Name: "add", Title: rest[0], Body: rest[1]}, nil
	case "delete":
		if len(rest) < 1 {
			return Command{}, fmt.Errorf("usage: jot delete <id>")
		}
		return Command{Name: "delete", ID: rest[0]}, nil
	default:
		return Command{}, fmt.Errorf("unknown command: %s", name)
	}
}

func Main() {
	store, err := notes.Load("notes.json")
	if err != nil {
		store = notes.New()
	}
 
	command, err := parseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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