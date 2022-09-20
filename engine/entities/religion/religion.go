package religion

import (
	"encoding/json"
	"fmt"
	"strings"

	js "persons_generator/core/storage/json_storage"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	we "persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/gender"
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
	t, err := NewType(r)
	if err != nil {
		return nil, err
	}
	r.Type = t
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

func NewReferences(cfg Config, amount int, cultures []*culture.Culture) ([]*CultureReference, error) {
	if amount > len(cultures) {
		return newRefForCultureNumberLess(cfg, amount, cultures)
	}
	if amount == len(cultures) {
		return newRefForCultureNumberEqual(cfg, amount, cultures)
	}
	if amount < len(cultures) {
		return newRefForCultureNumberGreater(cfg, amount, cultures)
	}

	return nil, we.NewBadRequestError(nil, fmt.Sprintf("can not create religion_culture reference for amount=%d, cultures length=%d", amount, len(cultures)))
}

func newRefForCultureNumberLess(cfg Config, amount int, cultures []*culture.Culture) ([]*CultureReference, error) {
	var (
		sampleReligions      = make([]*Religion, amount)
		religions            = make([]*Religion, 0, amount)
		chunkSampleReligions = tools.ChunkFor(sampleReligions, len(cultures))
		cultureReligionMap   = make(map[string][]string, amount)
	)

	for i := range chunkSampleReligions {
		for range chunkSampleReligions[i] {
			c := cultures[i]
			r, err := New(cfg, c)
			if err != nil {
				return nil, err
			}
			religions = append(religions, r)
			val, ok := cultureReligionMap[c.Name]
			if !ok || len(val) == 0 {
				cultureReligionMap[c.Name] = []string{r.Name}
			} else {
				cultureReligionMap[c.Name] = append(cultureReligionMap[c.Name], r.Name)
			}
		}
	}

	refs := make([]*CultureReference, 0, amount)
	for cultureName, religionNames := range cultureReligionMap {
		c := culture.GetCultureByName(cultureName, cultures)
		if c == nil {
			continue
		}
		for _, religionName := range religionNames {
			r := GetReligionByName(religionName, religions)
			if r == nil {
				continue
			}
			refs = append(refs, &CultureReference{
				Culture:  c,
				Religion: r,
			})
		}
	}

	return refs, nil
}

func newRefForCultureNumberEqual(cfg Config, amount int, cultures []*culture.Culture) ([]*CultureReference, error) {
	var (
		religions          = make([]*Religion, 0, amount)
		cultureReligionMap = make(map[string]string, amount)
	)

	for _, c := range cultures {
		r, err := New(cfg, c)
		if err != nil {
			return nil, err
		}
		religions = append(religions, r)
		cultureReligionMap[c.Name] = r.Name
	}

	refs := make([]*CultureReference, 0, amount)
	for cultureName, religionName := range cultureReligionMap {
		c := culture.GetCultureByName(cultureName, cultures)
		if c == nil {
			continue
		}
		r := GetReligionByName(religionName, religions)
		if r == nil {
			continue
		}
		refs = append(refs, &CultureReference{
			Culture:  c,
			Religion: r,
		})
	}

	return refs, nil
}

func newRefForCultureNumberGreater(cfg Config, amount int, cultures []*culture.Culture) ([]*CultureReference, error) {
	var (
		religions          = make([]*Religion, 0, amount)
		culturesChunks     = tools.ChunkFor(cultures, amount)
		cultureReligionMap = make(map[string]string, amount)
	)

	for _, chunk := range culturesChunks {
		hybridCulture, err := culture.NewWithProto(culture.Config{StorageFolderName: cfg.StorageFolderName}, chunk)
		if err != nil {
			return nil, err
		}
		r, err := New(cfg, hybridCulture)
		if err != nil {
			return nil, err
		}
		religions = append(religions, r)
		for _, c := range chunk {
			cultureReligionMap[c.Name] = r.Name
		}
	}

	refs := make([]*CultureReference, 0, amount)
	for cultureName, religionName := range cultureReligionMap {
		c := culture.GetCultureByName(cultureName, cultures)
		if c == nil {
			continue
		}
		r := GetReligionByName(religionName, religions)
		if r == nil {
			continue
		}
		refs = append(refs, &CultureReference{
			Culture:  c,
			Religion: r,
		})
	}

	return refs, nil
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
		return nil, we.NewBadRequestError(nil, fmt.Sprintf("amount (%d) can not be less than length of preferred (length=%d)", amount, len(preferred)))
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

func (r *Religion) IsZero() bool {
	return r == nil
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

func (r *Religion) GetAdulteryAcceptance(sex gender.Sex) Acceptance {
	if r == nil {
		return ""
	}

	for _, t := range r.Theology.Taboos.Taboos {
		if sex.IsMale() && t.Name == MaleAdulteryTabooName {
			return t.Acceptance
		}
		if sex.IsFemale() && t.Name == FemaleAdulteryTabooName {
			return t.Acceptance
		}
	}

	return ""
}

func (r *Religion) GetDivorceAcceptance() Acceptance {
	if r == nil {
		return ""
	}

	switch r.Theology.MarriageTradition.Divorce {
	case AlwaysAllowed:
		return Accepted
	case MustBeApproved:
		return Shunned
	case Disallowed:
		return Criminal
	}

	return ""
}

func (r *Religion) GetAcceptedNumberSpounces(sex gender.Sex) int {
	if r == nil {
		return 0
	}

	return r.Theology.MarriageTradition.GetAcceptedNumberSpounces(sex)
}

/* ********************************************************************************************************** */

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
		return we.NewBadRequestError(nil, "can not save nil culture")
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

func UniqueReligions(religions []*Religion) []*Religion {
	if len(religions) <= 1 {
		return religions
	}

	preOut := make(map[string]*Religion)
	for _, c := range religions {
		preOut[c.Name] = c
	}

	out := make([]*Religion, 0, len(preOut))
	for _, c := range preOut {
		out = append(out, c)
	}

	return out
}

func GetReligionSimilarityCoef(r1, r2 *Religion) (float64, error) {
	if r1.IsZero() && r2.IsZero() {
		return 1, nil
	}
	if r1.IsZero() || r2.IsZero() {
		return 0, wrapped_error.NewInternalServerError(nil, "can not compare religions if one of it is <nil>")
	}
	if r1.ID == r2.ID {
		return 1, nil
	}

	similarityTraits := []struct {
		enable bool
		value  float64
		coef   float64
	}{
		{
			enable: IsTypesEqual(r1.Type, r2.Type),
			value:  1,
			coef:   0.3,
		},
		{
			enable: true,
			value:  1 - g.GetDelta(r1.GenderDominance, r2.GenderDominance),
			coef:   0.1,
		},
		{
			enable: true,
			value:  GetDoctrineSimilarityCoef(r1.Doctrine, r2.Doctrine),
			coef:   0.15,
		},
		{
			enable: true,
			value:  GetAttributesSimilarityCoef(r1.Attributes, r2.Attributes),
			coef:   0.15,
		},
		{
			enable: true,
			value:  GetTheologySimilarityCoef(r1.Theology, r2.Theology),
			coef:   0.5,
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
