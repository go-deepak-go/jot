package main

type Command struct {
	Name     string
	ID       string
	Query    string
	Tag      string
	Notebook string
	Title    string
	Body     string
}

func main() {
	Main()
}
