package notes

type Note struct {
	ID       string
	Title    string
	Body     string
	Notebook string
	Tags     []string
	Pinned   bool
}