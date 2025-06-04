package person

import "fmt"

type Person struct {
	Name string `help:"What's your name?"`
	Age  int    `help:"How old are you?"`
	City string `help:"Where did you live?"`
}

func New() *Person {
	return &Person{
		Name: "Lucas",
		Age:  27,
		City: "Dijon",
	}
}

func (p *Person) Run() error {
	fmt.Printf("I'm %s, %d and I live in %s.\n", p.Name, p.Age, p.City)
	return nil
}
