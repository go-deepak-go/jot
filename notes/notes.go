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

func (s Store) AddTag(ID string, tag string) error {
	note, ok := s[ID]
	if !ok {
		return fmt.Errorf("no note with ID %s", ID)
	}
	for _, existingTag := range note.Tags {
		if existingTag == tag {
			return nil
		}
	}
	note.Tags = append(note.Tags, tag)
	s[ID] = note
	return nil
}

func (s Store) Pin(ID string) error {
	note, ok := s[ID]
	if !ok {
		return fmt.Errorf("no note with ID %s", ID)
	}
	note.Pinned = true
	s[ID] = note
	return nil
}

func (s Store) Search(query string) []Note {
	query = strings.ToLower(query)
	var matches []Note
	for _, note := range s {
		haystack := strings.ToLower(note.Title + " " + note.Body)
		if strings.Contains(haystack, query) {
			matches = append(matches, note)
		}
	}
	return matches
}

func (s Store) InNotebook(notebook string) []Note {
	var result []Note
	for _, note := range s {
		if note.Notebook == notebook {
			result = append(result, note)
		}
	}
	return result
}