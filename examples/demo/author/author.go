package author

import "fmt"

type Author struct {
	Name string `help:"Where the author live"`
	Age  int    `help:"Age of the author"`
}

func New() *Author {
	return &Author{
		Name: "Lucas",
		Age:  27,
	}
}

func (a *Author) Run() error {
	fmt.Printf("The Author is %s, %d.\n", a.Name, a.Age)
	return nil
}
