package orchestrator

import (
	"fmt"

	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
)

func (o *Orchestrator) ShowReligions() {
	fmt.Println()
	for _, r := range o.religions {
		r.Print()
	}
	o.w.PrintLocationReligions()
	fmt.Println()
}

func (o *Orchestrator) CreateReligion(c *culture.Culture) (*religion.Religion, error) {
	return religion.New(c)
}

func (o *Orchestrator) CreateReligions(c []*culture.Culture) ([]*religion.Religion, error) {
	out := make([]*religion.Religion, len(c))
	for i := range out {
		r, err := o.CreateReligion(c[i])
		if err != nil {
			return nil, err
		}
		out[i] = r
	}

	return out, nil
}
