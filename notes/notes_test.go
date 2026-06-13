package notes_test

import (
	"jot/notes"
	"testing"
)

func TestSummary_ReturnsNoteSummary(t *testing.T) {
	t.Parallel()
	note := notes.Note{
		ID: "1",
		Title: "ABC",
		Body: "qwerty",
		Notebook: "Test",
		Tags: []string{"go", "learning"},
		Pinned: true,
	}
	want := "* #1  ABC  [Test] #go #learning"
	got := note.Summary()
	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func getTestStore() notes.Store {
	return notes.Store{
		"1": {
			ID: "1",
			Title: "Go maps",
			Body: "Maps are great.",
			Notebook: "",
			Tags: []string{"go", "learning"},
			Pinned: true,
		},
		"2": {
			ID: "2",
			Title: "Go Slices",
			Body: "Slices are great.",
			Notebook: "",
			Tags: []string{"go", "learning"},
			Pinned: false,
		},
	}
}

