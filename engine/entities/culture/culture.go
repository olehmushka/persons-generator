package culture

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	js "persons_generator/core/storage/json_storage"
	"persons_generator/core/tools"
	we "persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/language"
	pm "persons_generator/engine/probability_machine"

	"github.com/google/uuid"
)

type Culture struct {
	ID    uuid.UUID  `json:"id"`
	Proto []*Culture `json:"proto"`
	Name  string     `json:"name"`

	Abstuct         *AbstructCulture   `json:"abstract"`
	Root            *Root              `json:"root"`
	Language        *language.Language `json:"language"`
	Ethos           *Ethos             `json:"ethos"`
	Traditions      []*Tradition       `json:"traditions"`
	GenderDominance *g.Dominance       `json:"gender_dominance"`
	MartialCustom   g.Acceptance       `json:"martial_custom"`

	storageFolderName string
}

func New(cfg Config, preferred *Preference) (*Culture, error) {
	proto, err := getProtoCultures(preferred)
	if err != nil {
		return nil, err
	}
	if len(proto) == 0 {
		return nil, we.New(http.StatusInternalServerError, nil, "proto cultures can not be zero")
	}

	return NewWithProto(cfg, proto)
}

func getProtoCultures(preferred *Preference) ([]*Culture, error) {
	if preferred == nil {
		return GetRandomProtoCultures(1, 3)
	}

	switch preferred.Kind {
	case StrictPrefKind:
		return getProtoCulturesForStrictKind(preferred)
	case HybridPrefKind:
		return getProtoCulturesForHybridKind(preferred)
	default:
		return nil, we.New(http.StatusInternalServerError, nil, fmt.Sprintf("unexpected preference kind (kind_name=%s)", preferred.Kind))
	}
}

func getProtoCulturesForStrictKind(preferred *Preference) ([]*Culture, error) {
	if preferred == nil {
		return nil, we.New(http.StatusInternalServerError, nil, "preferred can not be nil")
	}
	if preferred.Kind == StrictPrefKind && len(preferred.Names) != preferred.Amount {
		return nil, we.New(http.StatusInternalServerError, nil, fmt.Sprintf("for strict kind of preference number preferred names (%d) can not be not equal to amount (%d)", len(preferred.Names), preferred.Amount))
	}

	out, err := GetProtoCulturesByPreferred(preferred.Names)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func getProtoCulturesForHybridKind(preferred *Preference) ([]*Culture, error) {
	if preferred == nil {
		return nil, we.New(http.StatusInternalServerError, nil, "preferred can not be nil")
	}

	out, err := GetProtoCulturesByPreferred(preferred.Names)
	if err != nil {
		return nil, err
	}
	if len(out) == preferred.Amount {
		return out, nil
	}
	for i := 0; i < 20; i++ {
		cs, err := tools.RandomValuesOfSlice(pm.RandFloat64, AllCultures, preferred.Amount-len(out))
		if err != nil {
			return nil, err
		}
		out = append(out, cs...)
		out = UniqueCultures(out)
		if len(out) == preferred.Amount {
			break
		}
	}

	return out, nil
}

func NewWithProto(cfg Config, proto []*Culture) (*Culture, error) {
	c := &Culture{
		ID:    uuid.New(),
		Proto: proto,

		storageFolderName: cfg.StorageFolderName,
	}
	c.GenderDominance = getGenderDominance(c.Proto)
	c.MartialCustom = getMartialCustom(c.Proto)
	l, err := language.New(getLanguageNamesFromProto(c.Proto))
	if err != nil {
		return nil, err
	}
	c.Language = l
	name, err := c.Language.GetCultureName()
	if err != nil {
		return nil, err
	}
	c.Name = name
	c.Ethos = getEthos(c.Proto)
	t, err := getTraditions(c.Proto)
	if err != nil {
		return nil, err
	}
	c.Traditions = t
	r, err := getRoot(c.Proto)
	if err != nil {
		return nil, err
	}
	c.Root = r

	return c, nil
}

func NewMany(cfg Config, amount int, preferred []*Preference) ([]*Culture, error) {
	cultures := make([]*Culture, amount)
	if amount < len(preferred) {
		return nil, we.New(http.StatusInternalServerError, nil, fmt.Sprintf("amount (%d) can not be less than length of preferred (length=%d)", amount, len(preferred)))
	}
	for i := range cultures {
		var p *Preference
		if len(preferred) > i {
			p = preferred[i]
		}
		c, err := New(cfg, p)
		if err != nil {
			return nil, err
		}
		cultures[i] = c
	}

	return cultures, nil
}

func (c *Culture) Print() {
	if c == nil {
		return
	}

	fmt.Printf("Culture (name=%s)\n", c.Name)
	c.Root.Print()
	if len(c.Proto) > 0 {
		fmt.Printf("Proto cultures (culture_name=%s)\n", c.Name)
		for _, p := range c.Proto {
			if p == nil {
				continue
			}
			fmt.Printf(" - %s\n", p.Name)
		}
	}
	c.Language.Print()
	c.Ethos.Print()
	c.PrintTraditions()
	c.GenderDominance.Print()
	c.PrintMartialCustom()
}

func getLanguageNamesFromProto(proto []*Culture) []string {
	names := make([]string, 0, len(proto))
	for _, p := range proto {
		if p == nil {
			continue
		}
		if p.Language == nil {
			continue
		}
		names = append(names, p.Language.Name)
	}

	return tools.Unique(names)
}

func GetProtoCulturesByPreferred(names []string) ([]*Culture, error) {
	out := make([]*Culture, len(names))
	for i := range out {
		found := GetProtoCultureByPreferredName(names[i])
		if found == nil {
			return nil, we.New(http.StatusInternalServerError, nil, fmt.Sprintf("can not found proto culture by name (name=%s)", names[i]))
		}
		out[i] = found
	}

	return out, nil
}

func GetProtoCultureByPreferredName(prefName string) *Culture {
	for _, c := range AllCultures {
		if tools.ContainString(c.Name, prefName) {
			return c
		}
	}

	return nil
}

func UniqueCultures(cultures []*Culture) []*Culture {
	if len(cultures) <= 1 {
		return cultures
	}

	preOut := make(map[string]*Culture)
	for _, c := range cultures {
		preOut[c.Name] = c
	}

	out := make([]*Culture, 0, len(preOut))
	for _, c := range preOut {
		out = append(out, c)
	}

	return out
}

func GetRandomProtoCultures(min, max int) ([]*Culture, error) {
	if max > len(AllCultures) {
		return nil, we.New(http.StatusInternalServerError, nil, fmt.Sprintf("number of all cultures can not be less than max for random generation (%d)", max))
	}
	amount, err := pm.RandIntInRange(min, max)
	if err != nil {
		return nil, err
	}
	var out = make([]*Culture, 0, amount)
	for i := 0; i < 20; i++ {
		cs, err := tools.RandomValuesOfSlice(pm.RandFloat64, AllCultures, amount-len(out))
		if err != nil {
			return nil, err
		}
		out = append(out, cs...)
		out = UniqueCultures(out)
		if len(out) == amount {
			break
		}
	}

	return out, nil
}

func GetCultureByName(name string, list []*Culture) *Culture {
	if name == "" || len(list) == 0 {
		return nil
	}

	return tools.Search(list, func(c *Culture) string { return c.Name }, name)
}

func GetCulturesByName(name string, list []*Culture) []*Culture {
	if name == "" || len(list) == 0 {
		return nil
	}

	return tools.SearchMany(list, func(c *Culture) string { return c.Name }, name)
}

func MapCultureNames(cultures []*Culture) []string {
	if len(cultures) == 0 {
		return []string{}
	}

	out := make([]string, len(cultures))
	for i := range out {
		out[i] = cultures[i].Name
	}

	return out
}

func (c *Culture) Save() error {
	if c == nil {
		return we.New(http.StatusInternalServerError, nil, "can not save nil culture")
	}

	b, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	return js.
		New(js.Config{StorageFolderName: c.storageFolderName}).
		Store(strings.Join([]string{"culture", c.ID.String()}, "_")+".json", b)
}

func ReadByID(storageFolderName string, id uuid.UUID) (*Culture, error) {
	filename := strings.Join([]string{"culture", id.String()}, "_") + ".json"
	b, err := js.
		New(js.Config{StorageFolderName: storageFolderName}).
		Get(filename)
	if err != nil {
		return nil, err
	}

	var out Culture
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
