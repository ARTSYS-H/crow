package crow

import "fmt"

type Topic struct {
	Description string
	Content     string
}

func (t *Topic) String() string {
	return fmt.Sprintln(t.Content)
}
