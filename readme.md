# jot

A tiny command-line notes app. Write notes, file them in notebooks, tag them, pin the important ones, and search across everything.

Built in Go as a learning project, working through *The Deeper Love of Go* by John Arundel.

## Project layout

```
jot/
├── notes/
│   ├── notes.go       — Note and Store types, all methods
│   └── notes_test.go  — tests
└── cmd/jot/
    ├── main.go        — the program you run
    └── main_test.go   — tests for argument parsing
```

## Usage

Each run of the program is one command.

```
go run ./cmd/jot/ <command> [args]
```

Notes are saved to `notes.json` in the directory you run the program from.

### add

Add a new note.

```
go run ./cmd/jot/ add <title> <body>
```

```
go run ./cmd/jot/ add "Go maps" "Maps are great."
```

```
added note 1
  #1  Go maps  []
```

### list

List all notes.

```
go run ./cmd/jot/ list
```

```
* #1  Go maps  [Go] #go
  #2  Sourdough  [Cooking]
  #3  Go slices  [Go]
```

A `*` at the start means the note is pinned.

### view

View the full body of a note.

```
go run ./cmd/jot/ view 1
```

```
  #1  Go maps  [Go] #go
---
Maps are great.
```

### search

Search across note titles and bodies. Case-insensitive.

```
go run ./cmd/jot/ search flour
```

```
  #2  Sourdough  [Cooking]
```

### notebook

List all notes in a notebook.

```
go run ./cmd/jot/ notebook Go
```

```
  #1  Go maps  [Go] #go
  #3  Go slices  [Go]
```

### pin

Pin a note so it appears at the top of the list.

```
go run ./cmd/jot/ pin 1
```

```
pinned note 1
* #1  Go maps  [Go] #go
```

### tag

Add a tag to a note. Adding the same tag twice has no effect.

```
go run ./cmd/jot/ tag 1 important
```

```
tagged note 1 with important
  #1  Go maps  [Go] #important
```

## Running the tests

```
go test ./...
```
