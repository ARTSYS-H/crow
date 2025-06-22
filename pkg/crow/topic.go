package crow

import "fmt"

type Topic struct {
	Short string
	Long  string
}

func (t *Topic) String() string {
	return fmt.Sprint(t.Long)
}
