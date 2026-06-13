package notes

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Note struct {
	ID       string
	Title    string
	Body     string
	Notebook string
	Tags     []string
	Pinned   bool
}

type Store map[string]Note

// Summary returns a one-line description of the note.
func (n Note) Summary() string {
	pin := " "
	if n.Pinned {
		pin = "*"
	}
	line := fmt.Sprintf("%s #%s  %s  [%s]", pin, n.ID, n.Title, n.Notebook)
	if len(n.Tags) > 0 {
		line += " #" + strings.Join(n.Tags, " #")
	}
	return line
}

func (s Store) GetNote(ID string) (Note, bool) {
	note, ok := s[ID]
	return note ,ok
}

func (s Store) AddNote(note Note) (Note, error) {
	if len(strings.TrimSpace(note.Title)) == 0 {
		return Note{}, fmt.Errorf("title cannot be empty")
	}
	s[note.ID] = note
	return note, nil
}

func (s Store) GetAllNotes() []Note {
	return slices.Collect(maps.Values(s))
}