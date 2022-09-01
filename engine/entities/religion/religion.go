package religion

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	js "persons_generator/core/storage/json_storage"
	"persons_generator/core/tools"
	we "persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	g "persons_generator/engine/entities/gender"
	pm "persons_generator/engine/probability_machine"

	"github.com/google/uuid"
)

type Religion struct {
	M        Metadata          `json:"m"`
	Metadata *religionMetadata `json:"metadata"`

	ID              uuid.UUID    `json:"id"`
	Name            string       `json:"name"`
	Type            *Type        `json:"type"`
	GenderDominance *g.Dominance `json:"gender_dominance"`
	Doctrine        *Doctrine    `json:"doctrine"`
	Attributes      *Attributes  `json:"attributes"`
	Theology        *Theology    `json:"theology"`

	storageFolderName string
}

func New(cfg Config, c *culture.Culture) (*Religion, error) {
	lowBaseCoef, err := pm.RandFloat64InRange(0.45, 0.75)
	if err != nil {
		return nil, err
	}
	baseCoef, err := pm.RandFloat64InRange(0.95, 1.05)
	if err != nil {
		return nil, err
	}
	highBaseCoef, err := pm.RandFloat64InRange(1, 1.25)
	if err != nil {
		return nil, err
	}
	maxMetadataValue, err := pm.RandFloat64InRange(8, 10)
	if err != nil {
		return nil, err
	}

	r := &Religion{
		M: Metadata{
			LowBaseCoef:  lowBaseCoef,
			BaseCoef:     baseCoef,
			HighBaseCoef: highBaseCoef,

			MaxMetadataValue: maxMetadataValue,
		},

		ID: uuid.New(),

		storageFolderName: cfg.StorageFolderName,
	}
	name, err := c.Language.GetReligionName()
	if err != nil {
		return nil, err
	}
	r.Name = name
	r.Type = NewType(r)
	r.GenderDominance = c.GenderDominance
	metadata, err := r.generateMetadata()
	if err != nil {
		return nil, err
	}
	r.Metadata = metadata
	doctrine, err := NewDoctrine(r)
	if err != nil {
		return nil, err
	}
	r.Doctrine = doctrine
	attributes, err := NewAttributes(r, c)
	if err != nil {
		return nil, err
	}
	r.Attributes = attributes
	theology, err := NewTheology(r, c)
	if err != nil {
		return nil, err
	}
	r.Theology = theology

	return r, nil
}

func NewByPreferred(cfg Config, preferred *Preference) (*Religion, error) {
	c, err := getCulture(cfg, preferred)
	if err != nil {
		return nil, err
	}

	return New(cfg, c)
}

func getCulture(cfg Config, preferred *Preference) (*culture.Culture, error) {
	var (
		amount       int = 1
		prefCultures []*culture.Culture
	)
	if preferred != nil {
		amount = preferred.Amount
		prefCultures = preferred.Cultures
	}

	return culture.NewHybrid(
		culture.Config{StorageFolderName: cfg.StorageFolderName},
		amount,
		prefCultures,
	)
}

func NewMany(cfg Config, cultures []*culture.Culture) ([]*Religion, error) {
	religions := make([]*Religion, len(cultures))
	for i := range religions {
		r, err := New(cfg, cultures[i])
		if err != nil {
			return nil, err
		}
		religions[i] = r
	}

	return religions, nil
}

func NewManyByPreferred(cfg Config, amount int, preferred []*Preference) ([]*Religion, error) {
	if amount < len(preferred) {
		return nil, we.New(http.StatusInternalServerError, nil, fmt.Sprintf("amount (%d) can not be less than length of preferred (length=%d)", amount, len(preferred)))
	}

	out := make([]*Religion, amount)
	for i := range out {
		var p *Preference
		if len(preferred) > i {
			p = preferred[i]
		}
		c, err := NewByPreferred(cfg, p)
		if err != nil {
			return nil, err
		}
		out[i] = c
	}

	return out, nil
}

func (r *Religion) UpdateMetadata(rm *religionMetadata) {
	r.Metadata = rm
}

func (r *Religion) Print() {
	fmt.Printf("Religion (name=%s)\n", r.Name)
	r.Type.Print()
	r.GenderDominance.Print()
	r.Doctrine.Print()
	r.Attributes.Print()
	r.Theology.Print()

	fmt.Printf("\n\nMetadata:\n%+v\n", r.Metadata)
	fmt.Printf("=====================\n\n")
}

/* ********************************************************************************************************** */

func (r *Religion) HasReincarnation() bool {
	if r.Doctrine != nil {
		if r.Doctrine.HighGoal != nil {
			if r.Doctrine.HighGoal.ContainsReincarnation() {
				return true
			}
		}
	}

	if r.Theology != nil {
		if r.Theology.HasReincarnation() {
			return true
		}
	}

	return false
}

func GetReligionByName(name string, list []*Religion) *Religion {
	if name == "" || len(list) == 0 {
		return nil
	}

	return tools.Search(list, func(r *Religion) string { return r.Name }, name)
}

func MapReligionNames(religions []*Religion) []string {
	if len(religions) == 0 {
		return []string{}
	}

	out := make([]string, len(religions))
	for i := range out {
		out[i] = religions[i].Name
	}

	return out
}

func (r *Religion) Save() error {
	if r == nil {
		return we.New(http.StatusInternalServerError, nil, "can not save nil culture")
	}

	b, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return err
	}

	return js.
		New(js.Config{StorageFolderName: r.storageFolderName}).
		Store(strings.Join([]string{"religion", r.ID.String()}, "_")+".json", b)
}

func ReadByID(storageFolderName string, id uuid.UUID) (*Religion, error) {
	filename := strings.Join([]string{"religion", id.String()}, "_") + ".json"
	b, err := js.
		New(js.Config{StorageFolderName: storageFolderName}).
		Get(filename)
	if err != nil {
		return nil, err
	}

	var out Religion
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
