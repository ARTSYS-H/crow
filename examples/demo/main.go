package main

import (
	"fmt"
	"os"

	"github.com/ARTSYS-H/crow/examples/demo/author"
	"github.com/ARTSYS-H/crow/examples/demo/editor"
	"github.com/ARTSYS-H/crow/pkg/crow"
)

func main() {
	app := crow.New("demo", "Demo is a demo application.")
	author := author.New()
	editor := editor.New()
	app.AddCommmand(author, "Description of the Author")
	app.AddCommmand(editor, "Description of the Editor")
	err := app.Execute(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
