package orchestrator

import (
	"persons_generator/entities"
	newReligion "persons_generator/entities/religion/religion"
	"persons_generator/world"
)

type Orchestrator struct {
	w         *entities.World
	cultures  []*entities.Culture
	religions []*newReligion.Religion
	// humans    []*human.Human
}

func New(cfg *Config) *Orchestrator {
	w := world.FillWorld(&entities.World{
		Size: cfg.WorldSize,
	})
	r := newReligion.NewReligions(cfg.ReligionNumber)
	// hs := human.NewHumans(&human.NewHumansParams{
	// 	Cultures: map[string]*human.NewHumanParams{
	// 		culture.Babylonian.Name: {Count: 10},
	// 		culture.Egyptian.Name:   {Count: 10},
	// 	},
	// })

	return &Orchestrator{
		w:         w,
		religions: r,
		// humans:    hs,
	}
}

func (o *Orchestrator) Orchestrate() {
}

func (o *Orchestrator) ShowReligions() {
	// for _, h := range o.humans {
	// 	fmt.Printf(" - %s\n", h.GetFullName())
	// }
	for _, r := range o.religions {
		r.Print()
	}
}
