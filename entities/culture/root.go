package culture

import "fmt"

type Root struct {
	Name string
}

func (r *Root) Print() {
	if r == nil {
		return
	}
	fmt.Printf("Root: %s\n", r.Name)
}
