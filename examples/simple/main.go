package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ARTSYS-H/crow/examples/simple/person"
	"github.com/ARTSYS-H/crow/pkg/crow"
)

func main() {
	app := crow.New("Demo", "This is a demo of Crow CLI.")
	person := person.New()
	err := app.AddCommand(person, "modify value of person")
	if err != nil {
		log.Fatalln(err)
	}
	err = app.Execute(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
