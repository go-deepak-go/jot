package main

import (
	"fmt"
	"os"

	"jot/notes"
)

func main() {
	store := notes.Store{
		"1": {ID: "1", Title: "Go maps", Body: "Maps are great.", Notebook: "Go", Tags: []string{"go"}, Pinned: false},
		"2": {ID: "2", Title: "Sourdough", Body: "Mix flour and water.", Notebook: "Cooking", Tags: []string{}, Pinned: false},
		"3": {ID: "3", Title: "Go slices", Body: "Slices are fun.", Notebook: "Go", Tags: []string{}, Pinned: false},
	}

	if len(os.Args) < 2 {
		fmt.Println("usage: jot <command> [args]")
		fmt.Println("commands: list, view, search, notebook, pin, tag")
		return
	}

	command := os.Args[1]

	switch command {
	case "list":
		allNotes := store.GetAllNotes()
		for _, note := range allNotes {
			fmt.Println(note.Summary())
		}

	case "view":
		if len(os.Args) < 3 {
			fmt.Println("usage: jot view <id>")
			return
		}
		id := os.Args[2]
		note, ok := store.GetNote(id)
		if !ok {
			fmt.Println("no note with ID", id)
			return
		}
		fmt.Println(note.Summary())
		fmt.Println("---")
		fmt.Println(note.Body)

	case "search":
		// jot search <query>
		if len(os.Args) < 3 {
			fmt.Println("usage: jot search <query>")
			return
		}
		query := os.Args[2]
		matches := store.Search(query)
		if len(matches) == 0 {
			fmt.Println("no matches for:", query)
			return
		}
		for _, note := range matches {
			fmt.Println(note.Summary())
		}
 
	case "notebook":
		// jot notebook <name>
		if len(os.Args) < 3 {
			fmt.Println("usage: jot notebook <name>")
			return
		}
		name := os.Args[2]
		result := store.InNotebook(name)
		if len(result) == 0 {
			fmt.Println("no notes in notebook:", name)
			return
		}
		for _, note := range result {
			fmt.Println(note.Summary())
		}
 
	case "pin":
		// jot pin <id>
		if len(os.Args) < 3 {
			fmt.Println("usage: jot pin <id>")
			return
		}
		id := os.Args[2]
		err := store.Pin(id)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("pinned note", id)
		fmt.Println(store[id].Summary())
 
	case "tag":
		// jot tag <id> <tag>
		if len(os.Args) < 4 {
			fmt.Println("usage: jot tag <id> <tag>")
			return
		}
		id := os.Args[2]
		tag := os.Args[3]
		err := store.AddTag(id, tag)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("tagged note", id, "with", tag)
		fmt.Println(store[id].Summary())
 
	default:
		fmt.Println("unknown command:", command)
		fmt.Println("commands: list, view, search, notebook, pin, tag")
	}
}