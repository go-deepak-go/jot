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

func TestGetNote_FindsNoteInStoreByID(t *testing.T) {
	t.Parallel()
	store := getTestStore()
	want := notes.Note{
		ID: "1",
		Title: "Go maps",
		Body: "Maps are great.",
		Notebook: "",
		Tags: []string{"go", "learning"},
		Pinned: true,
	}
	got, ok := store.GetNote("1")
	if !ok {
		t.Fatal("Note not found!")
	}
	if want != got {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetNote_ReturnsFalseWhenNoteNotFound(t *testing.T) {
	t.Parallel()
	store = getTestStore()
	_, ok := store.GetNote("nonexistent ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}