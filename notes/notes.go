package notes

import (
	"fmt"
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