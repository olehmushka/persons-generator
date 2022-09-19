package culture

import (
	"encoding/json"
	"fmt"
	"strings"

	js "persons_generator/core/storage/json_storage"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
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
		return nil, we.NewBadRequestError(nil, "proto cultures can not be zero")
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
		return nil, we.NewBadRequestError(nil, fmt.Sprintf("unexpected preference kind (kind_name=%s)", preferred.Kind))
	}
}

func getProtoCulturesForStrictKind(preferred *Preference) ([]*Culture, error) {
	if preferred == nil {
		return nil, we.NewBadRequestError(nil, "preferred can not be nil")
	}
	if preferred.Kind == StrictPrefKind && len(preferred.Names) != preferred.Amount {
		return nil, we.NewBadRequestError(nil, fmt.Sprintf("for strict kind of preference number preferred names (%d) can not be not equal to amount (%d)", len(preferred.Names), preferred.Amount))
	}

	out, err := GetProtoCulturesByPreferred(preferred.Names)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func getProtoCulturesForHybridKind(preferred *Preference) ([]*Culture, error) {
	if preferred == nil {
		return nil, we.NewBadRequestError(nil, "preferred can not be nil")
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
	c.Ethos = getRandomEthosFromCultures(c.Proto)
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
	acs, err := getAbstractCulture(c.Root, c.Proto)
	if err != nil {
		return nil, err
	}
	c.Abstuct = acs

	return c, nil
}

func NewHybrid(cfg Config, amount int, parent []*Culture) (*Culture, error) {
	if len(parent) > amount {
		return nil, we.NewBadRequestError(nil, fmt.Sprintf("preferred parent cultures number (%d) can not be greater than required amount for hybrid culture (%d)", len(parent), amount))
	}
	cultures := make([]*Culture, amount)
	for i := range cultures {
		var c *Culture
		if i < len(parent) {
			c = parent[i]
		} else {
			generated, err := New(cfg, nil)
			if err != nil {
				return nil, err
			}
			c = generated
		}
		cultures[i] = c
	}

	return NewWithProto(cfg, cultures)
}

func NewMany(cfg Config, amount int, preferred []*Preference) ([]*Culture, error) {
	cultures := make([]*Culture, amount)
	if amount < len(preferred) {
		return nil, we.NewBadRequestError(nil, fmt.Sprintf("amount (%d) can not be less than length of preferred (length=%d)", amount, len(preferred)))
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

func (c *Culture) IsZero() bool {
	return c == nil
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
			return nil, we.NewBadRequestError(nil, fmt.Sprintf("can not found proto culture by name (name=%s)", names[i]))
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
		return nil, we.NewBadRequestError(nil, fmt.Sprintf("number of all cultures can not be less than max for random generation (%d)", max))
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
		return we.NewBadRequestError(nil, "can not save nil culture")
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

func GetProtosSimilarityCoef(c1, c2 []*Culture) float64 {
	if len(c1) == 0 && len(c2) == 0 {
		return 1
	}
	if len(c1) == 0 || len(c2) == 0 {
		return 0
	}

	sameProtoCultures := tools.GetCrossOfSlices(c1, c1, func(c1, c2 *Culture) bool {
		if c1 == nil && c2 == nil {
			return true
		}
		if c1 == nil || c2 == nil {
			return false
		}

		return c1.Name == c2.Name
	})
	averageProtoCulturesLength := float64(len(c1)+len(c2)) / 2

	return float64(len(sameProtoCultures)) / averageProtoCulturesLength
}

func GetCultureSimilarityCoef(c1, c2 *Culture) (float64, error) {
	if c1.IsZero() && c2.IsZero() {
		return 1, nil
	}
	if c1.IsZero() || c2.IsZero() {
		return 0, wrapped_error.NewInternalServerError(nil, "can not compare cultures if one of it is <nil>")
	}
	if c1.ID == c2.ID {
		return 1, nil
	}

	languageSimilarityCoef, err := language.GetLanguageSimilarityCoef(c1.Language, c2.Language)
	if err != nil {
		return 0, wrapped_error.NewInternalServerError(err, "can not get language similarity coef")
	}

	similarityTraits := []struct {
		enable bool
		value  float64
		coef   float64
	}{
		{
			enable: true,
			value:  GetProtosSimilarityCoef(c1.Proto, c2.Proto),
			coef:   0.1,
		},
		{
			enable: true,
			value:  GetAbstructCultureSimilarityCoef(c1.Abstuct, c2.Abstuct),
			coef:   0.1,
		},
		{
			enable: IsRootEqual(c1.Root, c2.Root),
			value:  1,
			coef:   0.1,
		},
		{
			enable: true,
			value:  languageSimilarityCoef,
			coef:   0.3,
		},
		{
			enable: true,
			value:  GetEthosSimilarityCoef(c1.Ethos, c2.Ethos),
			coef:   0.15,
		},
		{
			enable: true,
			value:  GetTraditionsSimilarityCoef(c1.Traditions, c2.Traditions),
			coef:   0.25,
		},
		{
			enable: true,
			value:  1 - g.GetDelta(c1.GenderDominance, c2.GenderDominance),
			coef:   0.05,
		},
		{
			enable: c1.MartialCustom == c2.MartialCustom,
			value:  1,
			coef:   0.05,
		},
	}

	var out float64
	for _, t := range similarityTraits {
		if t.enable {
			out += t.value * t.coef
		}
	}

	return out, nil
}
