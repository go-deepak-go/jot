package notes_test

import (
	"cmp"
	"path/filepath"
	"slices"
	"testing"

	notes "github.com/go-deepak-go/jot"
	gocmp "github.com/google/go-cmp/cmp"
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
	store := notes.New()
	store.AddNote(notes.Note{
			Title: "Go maps",
			Body: "Maps are great.",
			Notebook: "",
			Tags: []string{"go", "learning"},
			Pinned: true,
	})
	store.AddNote(notes.Note{
			Title: "Go Slices",
			Body: "Slices are great.",
			Notebook: "",
			Tags: []string{"go", "learning"},
			Pinned: false,
	})
	return store
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
	// if want.ID != got.ID ||
	// 	want.Title != got.Title ||
	// 	want.Body != got.Body ||
	// 	want.Notebook != got.Notebook ||
	// 	want.Pinned != got.Pinned ||
	// 	!slices.Equal(want.Tags, got.Tags) {
	// 	t.Fatalf("want %#v, got %#v", want, got)
	// }
	if !gocmp.Equal(want, got) {
		t.Fatalf("%s", gocmp.Diff(want, got))
	}
}

func TestGetNote_ReturnsFalseWhenNoteNotFound(t *testing.T) {
	t.Parallel()
	store := getTestStore()
	_, ok := store.GetNote("nonexistent ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddNote_AddsNoteToStore(t *testing.T) {
	t.Parallel()
	store := getTestStore()
	note, err := store.AddNote(notes.Note{
        Title:    "Making a notes app",
        Body:     "test",
        Notebook: "",
        Tags:     []string{"go", "learning"},
        Pinned:   true,
    })
	if err != nil {
        t.Fatalf("unexpected error: %s", err)
    }
    _, ok := store.GetNote(note.ID)
    if !ok {
        t.Fatal("added note not found in the store!")
    }
}

func TestAddNote_RejectsEmptyTitle(t *testing.T) {
    t.Parallel()
    store := getTestStore()
    _, err := store.AddNote(notes.Note{Title: ""})
    if err == nil {
        t.Fatal("want error for empty title, got nil")
    }
}

func TestGetAllNotes_ReturnAllNotes(t *testing.T) {
	t.Parallel()
	store := getTestStore()
	want := []notes.Note{
		{
			ID: "1",
			Title: "Go maps",
			Body: "Maps are great.",
			Notebook: "",
			Tags: []string{"go", "learning"},
			Pinned: true,
		},
		{
			ID: "2",
			Title: "Go Slices",
			Body: "Slices are great.",
			Notebook: "",
			Tags: []string{"go", "learning"},
			Pinned: false,
		},
	}
	got := store.GetAllNotes()
	slices.SortFunc(got, func(a, b notes.Note) int {
		return cmp.Compare(a.Title, b.Title)
	})
	slices.SortFunc(want, func(a, b notes.Note) int {
		return cmp.Compare(a.Title, b.Title)
	})
	// if len(want) != len(got) {
	// 	t.Fatalf("got different lists: want %#v, got: %#v", want, got)
	// }
	// for index := range got {
	// 	if want[index].ID != got[index].ID ||
	// 	want[index].Title != got[index].Title ||
	// 	want[index].Body != got[index].Body ||
	// 	want[index].Notebook != got[index].Notebook ||
	// 	want[index].Pinned != got[index].Pinned ||
	// 	!slices.Equal(want[index].Tags, got[index].Tags) {
	// 		t.Fatalf("want %#v, got %#v", want[index], got[index])
	// 	}
	// }
	if !gocmp.Equal(want, got) {
		t.Fatalf("%s", gocmp.Diff(want, got))
	}
}

func TestAddTag_AddsTagToNote(t *testing.T) {
	t.Parallel()
	store := getTestStore()
	err := store.AddTag("1", "maps")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	got, _ := store.GetNote("1")
	if !slices.Contains(got.Tags, "maps") {
		t.Fatal("want tag 'maps' to be present, but it wasn't")
	}
}

func TestAddTag_IgnoresDuplicates(t *testing.T) {
	t.Parallel()
	store := getTestStore()
	store.AddTag("1", "go")
	store.AddTag("1", "go")
	got, _ := store.GetNote("1")
	count := 0
	for _, tag := range got.Tags {
		if tag == "go" {
			count++
		}
	}
	if count != 1 {
		t.Fatalf("want tag 'go' exactly once, got %d times", count)
	}
}

func TestAddTag_ReturnsErrorForMissingNote(t *testing.T) {
	t.Parallel()
	store := getTestStore()
	err := store.AddTag("nonexistent", "go")
	if err == nil {
		t.Fatal("want error for nonexistent ID, got nil")
	}
}

func TestPin_PinsNote(t *testing.T) {
	t.Parallel()
	store := getTestStore()
	err := store.Pin("2")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	got, _ := store.GetNote("2")
	if !got.Pinned {
		t.Fatal("want note to be pinned, but it wasn't")
	}
}

func TestPin_ReturnsErrorForMissingNote(t *testing.T) {
	t.Parallel()
	store := getTestStore()
	err := store.Pin("nonexistent")
	if err == nil {
		t.Fatal("want error for nonexistent ID, got nil")
	}
}

func TestSearch_FindsByTitle(t *testing.T) {
	t.Parallel()
	store := getTestStore()
	got := store.Search("maps")
	if len(got) != 1 {
		t.Fatalf("want 1 result, got %d", len(got))
	}
	if got[0].ID != "1" {
		t.Fatalf("want note ID 1, got %s", got[0].ID)
	}
}

func TestSearch_IsCaseInsensitive(t *testing.T) {
	t.Parallel()
	store := getTestStore()
	got := store.Search("MAPS")
	if len(got) != 1 {
		t.Fatalf("want 1 result for uppercase query, got %d", len(got))
	}
}

func TestSearch_ReturnsEmptyWhenNoMatch(t * testing.T) {
	t.Parallel()
	store := getTestStore()
	got := store.Search("python")
	if len(got) != 0 {
		t.Fatalf("want 0 results, got %d", len(got))
	}
}

func TestInNotebook_ReturnsNotesInNotebook(t *testing.T) {
    t.Parallel()
    store := notes.New()
    store.AddNote(notes.Note{Title: "A", Notebook: "Go"})
    store.AddNote(notes.Note{Title: "B", Notebook: "Cooking"})
    store.AddNote(notes.Note{Title: "C", Notebook: "Go"})
    got := store.InNotebook("Go")
    if len(got) != 2 {
        t.Fatalf("want 2 notes in Go notebook, got %d", len(got))
    }
}

func TestSaveAndLoad_RoundTrip(t *testing.T) {
    t.Parallel()
    store := getTestStore()
    dir := t.TempDir()
    path := filepath.Join(dir, "notes.json")
    err := store.Save(path)
    if err != nil {
        t.Fatalf("unexpected error saving: %s", err)
    }
    loaded, err := notes.Load(path)
    if err != nil {
        t.Fatalf("unexpected error loading: %s", err)
    }
    want := store.GetAllNotes()
    got := loaded.GetAllNotes()
    slices.SortFunc(want, func(a, b notes.Note) int {
        return cmp.Compare(a.ID, b.ID)
    })
    slices.SortFunc(got, func(a, b notes.Note) int {
        return cmp.Compare(a.ID, b.ID)
    })
    if !gocmp.Equal(want, got) {
        t.Fatalf("%s", gocmp.Diff(want, got))
    }
}