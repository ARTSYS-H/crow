package editor

import "fmt"

type Editor struct {
	Name string `help:"Name of the Editor"`
	City string `flag:"c" help:"City of the Editor"`
}

func New() *Editor {
	return &Editor{
		Name: "Crow",
		City: "Dijon",
	}
}

func (e *Editor) Run() error {
	fmt.Printf("The residence of %s is in %s.\n", e.Name, e.City)
	return nil
}
