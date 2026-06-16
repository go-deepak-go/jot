package main

import "testing"

func TestParseArgs_ReturnsErrorWithNoArgs(t *testing.T) {
	t.Parallel()
	_, err := parseArgs([]string{"jot"})
	if err == nil {
		t.Fatal("want error when no command given, got nil")
	}
}

func TestParseArgs_ReturnsErrorForUnknownCommand(t *testing.T) {
	t.Parallel()
	_, err := parseArgs([]string{"jot", "fly"})
	if err == nil {
		t.Fatal("want error for unknown command, got nil")
	}
}

func TestParseArgs_List(t *testing.T) {
	t.Parallel()
	cmd, err := parseArgs([]string{"jot", "list"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if cmd.Name != "list" {
		t.Fatalf("want name 'list', got %q", cmd.Name)
	}
}

func TestParseArgs_ViewReturnsErrorWithNoID(t *testing.T) {
	t.Parallel()
	_, err := parseArgs([]string{"jot", "view"})
	if err == nil {
		t.Fatal("want error when view has no ID, got nil")
	}
}

func TestParseArgs_ViewReturnsCommandWithID(t *testing.T) {
	t.Parallel()
	cmd, err := parseArgs([]string{"jot", "view", "1"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if cmd.Name != "view" {
		t.Fatalf("want name 'view', got %q", cmd.Name)
	}
	if cmd.ID != "1" {
		t.Fatalf("want ID '1', got %q", cmd.ID)
	}
}

func TestParseArgs_SearchReturnsErrorWithNoQuery(t *testing.T) {
	t.Parallel()
	_, err := parseArgs([]string{"jot", "search"})
	if err == nil {
		t.Fatal("want error when search has no query, got nil")
	}
}

func TestParseArgs_SearchReturnsCommandWithQuery(t *testing.T) {
	t.Parallel()
	cmd, err := parseArgs([]string{"jot", "search", "flour"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if cmd.Query != "flour" {
		t.Fatalf("want query 'flour', got %q", cmd.Query)
	}
}

func TestParseArgs_NotebookReturnsErrorWithNoName(t *testing.T) {
	t.Parallel()
	_, err := parseArgs([]string{"jot", "notebook"})
	if err == nil {
		t.Fatal("want error when notebook has no name, got nil")
	}
}

func TestParseArgs_NotebookReturnsCommandWithName(t *testing.T) {
	t.Parallel()
	cmd, err := parseArgs([]string{"jot", "notebook", "Go"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if cmd.Notebook != "Go" {
		t.Fatalf("want notebook 'Go', got %q", cmd.Notebook)
	}
}

func TestParseArgs_PinReturnsErrorWithNoID(t *testing.T) {
	t.Parallel()
	_, err := parseArgs([]string{"jot", "pin"})
	if err == nil {
		t.Fatal("want error when pin has no ID, got nil")
	}
}

func TestParseArgs_PinReturnsCommandWithID(t *testing.T) {
	t.Parallel()
	cmd, err := parseArgs([]string{"jot", "pin", "2"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if cmd.ID != "2" {
		t.Fatalf("want ID '2', got %q", cmd.ID)
	}
}

func TestParseArgs_TagReturnsErrorWithNoArgs(t *testing.T) {
	t.Parallel()
	_, err := parseArgs([]string{"jot", "tag"})
	if err == nil {
		t.Fatal("want error when tag has no args, got nil")
	}
}

func TestParseArgs_TagReturnsErrorWithOnlyID(t *testing.T) {
	t.Parallel()
	_, err := parseArgs([]string{"jot", "tag", "1"})
	if err == nil {
		t.Fatal("want error when tag has no tag name, got nil")
	}
}

func TestParseArgs_TagReturnsCommandWithIDAndTag(t *testing.T) {
	t.Parallel()
	cmd, err := parseArgs([]string{"jot", "tag", "1", "important"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if cmd.ID != "1" {
		t.Fatalf("want ID '1', got %q", cmd.ID)
	}
	if cmd.Tag != "important" {
		t.Fatalf("want tag 'important', got %q", cmd.Tag)
	}
}