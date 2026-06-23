package notes

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
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

type Store struct {
	notes map[string]Note
	nextID int
}

type storeJSON struct {
    Notes  map[string]Note `json:"notes"`
    NextID int             `json:"nextID"`
}

func New() Store {
	return Store{
		notes: map[string]Note{},
	}
}

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
	note, ok := s.notes[ID]
	return note ,ok
}

func (s *Store) AddNote(note Note) (Note, error) {
	if len(strings.TrimSpace(note.Title)) == 0 {
		return Note{}, fmt.Errorf("title cannot be empty")
	}
	s.nextID++
	note.ID = fmt.Sprintf("%d", s.nextID)
	s.notes[note.ID] = note
	return note, nil
}

func (s Store) GetAllNotes() []Note {
	return slices.Collect(maps.Values(s.notes))
}

func (s Store) AddTag(ID string, tag string) error {
	note, ok := s.notes[ID]
	if !ok {
		return fmt.Errorf("no note with ID %s", ID)
	}
	for _, existingTag := range note.Tags {
		if existingTag == tag {
			return nil
		}
	}
	note.Tags = append(note.Tags, tag)
	s.notes[ID] = note
	return nil
}

func (s Store) Pin(ID string) error {
	note, ok := s.notes[ID]
	if !ok {
		return fmt.Errorf("no note with ID %s", ID)
	}
	note.Pinned = true
	s.notes[ID] = note
	return nil
}

func (s Store) Search(query string) []Note {
	query = strings.ToLower(query)
	var matches []Note
	for _, note := range s.notes {
		haystack := strings.ToLower(note.Title + " " + note.Body)
		if strings.Contains(haystack, query) {
			matches = append(matches, note)
		}
	}
	return matches
}

func (s Store) InNotebook(notebook string) []Note {
	var result []Note
	for _, note := range s.notes {
		if note.Notebook == notebook {
			result = append(result, note)
		}
	}
	return result
}

func (s Store) Save(path string) error {
	data, err := json.Marshal(storeJSON{
		Notes: s.notes,
		NextID: s.nextID,
	})
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Load(path string) (Store, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Store{}, err
	}
	var sj storeJSON
	err = json.Unmarshal(data, &sj)
	if err != nil {
		return Store{}, err
	}
	if sj.Notes == nil {
		sj.Notes = map[string]Note{}
	}
	return Store{
		notes: sj.Notes,
		nextID: sj.NextID,
	}, nil
}

func (s Store) Delete(ID string) error {
	_, ok := s.notes[ID]
	if !ok {
		return fmt.Errorf("no note with ID %s", ID)
	}
	delete(s.notes, ID)
	return nil
}
