package person

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/human"
	"persons_generator/engine/entities/person/traits"
	pm "persons_generator/engine/probability_machine"
)

// Action methods
func (p *Person) HaveSex(partner *Person, year int) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "can not <nil> have sex with someone else")
	}
	if partner == nil {
		return wrapped_error.NewInternalServerError(nil, "can not person have sex with <nil>")
	}

	// check if it is homosexual sex
	if p.Human.Sex == partner.Human.Sex {
		if err := p.AppendTrait(traits.HomosexualistTrait); err != nil {
			return wrapped_error.NewInternalServerError(nil, "can not append homosexualist trait into curr person")
		}
		if err := partner.AppendTrait(traits.HomosexualistTrait); err != nil {
			return wrapped_error.NewInternalServerError(nil, "can not append homosexualist trait into partner")
		}
	}

	// process human side sex for curr person
	if err := p.Human.HaveSex(partner.Human); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not have sex for curr person")
	}
	if err := p.AppendHaveSexEvent(year, partner); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not append have_sex event for curr person")
	}
	// process human side sex for partner
	if err := partner.Human.HaveSex(p.Human); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not have sex for partner")
	}
	if err := partner.AppendHaveSexEvent(year, p); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not append have_sex event for patner")
	}

	// process traits
	if err := p.ProcessLustfulTrait(); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not process lustful trait for curr person")
	}
	if err := partner.ProcessLustfulTrait(); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not process lustful trait for partner")
	}

	return nil
}

func (p *Person) GetMarried(partner *Person, year int) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "<nil> can not get married with person")
	}
	if partner == nil {
		return wrapped_error.NewInternalServerError(nil, "can not get married for <nil> partner")
	}

	// process marriage for curr person
	p.Spouces = UniquePersons(append(p.Spouces, partner))
	if err := p.AppendGetMarriedEvent(year, partner); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not append get_married event for curr person")
	}
	if err := p.AppendTrait(traits.MarriedTrait); err != nil {
		return wrapped_error.NewInternalServerError(nil, "can not append married trait into curr person")
	}

	// process marriage for partner
	partner.Spouces = UniquePersons(append(partner.Spouces, p))
	if err := partner.AppendGetMarriedEvent(year, p); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not append get_married event for partner")
	}
	if err := partner.AppendTrait(traits.MarriedTrait); err != nil {
		return wrapped_error.NewInternalServerError(nil, "can not append married trait into partner")
	}

	return nil
}

func (p *Person) Divorce(partner *Person, year int) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "<nil> can not divorce with person")
	}
	if partner == nil {
		return wrapped_error.NewInternalServerError(nil, "can not divorce for <nil> partner")
	}

	// check if perople are married
	areMarried, err := p.AreMarried(partner)
	if err != nil {
		return wrapped_error.NewInternalServerError(nil, "can not check if people are married for divorce")
	}
	if !areMarried {
		return wrapped_error.NewInternalServerError(nil, "can not divorce curr person & partner if they are not married")
	}

	// process divorce for curr person
	p.Spouces = RemovePersonFromSliceByID(p.Spouces, partner.ID)
	if err := p.AppendDivorcedEvent(year, partner); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not append divorced event for curr person")
	}
	if err := p.AppendTrait(traits.DivorcedTrait); err != nil {
		return wrapped_error.NewInternalServerError(nil, "can not append divorced trait into curr person")
	}

	// process divorce for partner
	partner.Spouces = RemovePersonFromSliceByID(partner.Spouces, p.ID)
	if err := partner.AppendDivorcedEvent(year, p); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not append divorced event for partner")
	}
	if err := partner.AppendTrait(traits.DivorcedTrait); err != nil {
		return wrapped_error.NewInternalServerError(nil, "can not append divorced trait into partner")
	}

	return nil
}

func (p *Person) IncreaseAge(year int) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "<nil> can not increase age")
	}

	age := p.Human.IncrementAge()
	p.Metadata.WishGetMarriedCoef = GetWishGetMarriedCoef(p.Human.Sex, age)
	p.Metadata.DeathCoef = GetDeathCoef(age)

	return nil
}

func (p *Person) TryDie(year int) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "<nil> can not die")
	}

	doesDie, err := pm.GetRandomBool(p.Metadata.DeathCoef)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not get random bool for try_die method")
	}
	if !doesDie {
		return nil
	}
	if err := p.Die(year); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not die in try_die method")
	}

	return nil
}

func (p *Person) GiveBirth(year int) ([]*Person, error) {
	if p == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "<nil> can not give born")
	}

	children := make([]*Person, 0, 2)
	for _, h := range p.Human.GiveBirth() {
		child, err := New(h, p.Culture, p.Religion, year)
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not give birth child")
		}
		children = append(children, child)
	}

	return children, nil
}

// Checker methods

func (p *Person) GetSympathicCoef(partner *Person, religionsSimilarity, culturesSimilarity map[string]float64) (float64, error) {
	if p == nil {
		return 0, wrapped_error.NewInternalServerError(nil, "<nil> can not check if someone is sympathic for person")
	}
	if partner == nil {
		return 0, wrapped_error.NewInternalServerError(nil, "can not check if <nil> peron is sympathic")
	}

	religionSimilarityCoef, ok := religionsSimilarity[fmt.Sprintf("%s:%s", p.Religion.ID.String(), partner.Religion.ID.String())]
	if !ok {
		return 0, wrapped_error.NewInternalServerError(nil, "can not find religion similarity coef")
	}
	cultureSimilarityCoef, ok := culturesSimilarity[fmt.Sprintf("%s:%s", p.Culture.ID.String(), partner.Culture.ID.String())]
	if !ok {
		return 0, wrapped_error.NewInternalServerError(nil, "can not find culture similarity coef")
	}

	coefs := []struct {
		coef  float64
		value float64
	}{
		{
			coef:  1,
			value: human.GetSimilarityCoef(p.Human, partner.Human),
		},
		{
			coef:  1.5,
			value: cultureSimilarityCoef,
		},
		{
			coef:  2,
			value: religionSimilarityCoef,
		},
	}
	var sum float64
	for _, coef := range coefs {
		sum += coef.coef * coef.value
	}

	return sum / float64(len(coefs)), nil
}

func (p *Person) IsMarried() (bool, error) {
	if p == nil {
		return false, wrapped_error.NewInternalServerError(nil, "<nil> can not check if person is married")
	}

	return len(p.Spouces) > 0, nil
}

func (p *Person) AreMarried(partner *Person) (bool, error) {
	if p == nil {
		return false, wrapped_error.NewInternalServerError(nil, "<nil> can not check if someone is married with person")
	}
	if partner == nil {
		return false, wrapped_error.NewInternalServerError(nil, "can not check if peron is married with <nil>")
	}

	if len(p.Spouces) == 0 || len(partner.Spouces) == 0 {
		return false, nil
	}
	var isPersonMarriedWithPartner bool
	for _, spounce := range p.Spouces {
		if spounce.ID == p.ID {
			isPersonMarriedWithPartner = true
		}
	}

	var isPartnerMarriedWithPerson bool
	for _, spounce := range partner.Spouces {
		if spounce.ID == p.ID {
			isPartnerMarriedWithPerson = true
		}
	}
	if isPersonMarriedWithPartner != isPartnerMarriedWithPerson {
		return false, wrapped_error.NewInternalServerError(nil, "curr person and partner have dissynchronized spounces")
	}

	return isPersonMarriedWithPartner, nil
}

func (p *Person) CanBePartners(partner *Person) (bool, error) {
	if p == nil {
		return false, wrapped_error.NewInternalServerError(nil, "<nil> can not check if someone can get married with person")
	}
	if partner == nil {
		return false, wrapped_error.NewInternalServerError(nil, "can not check if <nil> peron is available for getting married")
	}

	// check if any of curr person or partner is asexual
	if p.Human.Psycho.SexualOrientation.IsAsexual() || partner.Human.Psycho.SexualOrientation.IsAsexual() {
		return false, nil
	}
	// check if any of curr person or partner is heterosexual
	if p.Human.Sex != partner.Human.Sex &&
		(p.Human.Psycho.SexualOrientation.IsHeterosexual() || p.Human.Psycho.SexualOrientation.IsBisexual()) &&
		(partner.Human.Psycho.SexualOrientation.IsHeterosexual() || partner.Human.Psycho.SexualOrientation.IsBisexual()) {
		return true, nil
	}
	// check if any of curr person or partner is homosexual
	if p.Human.Sex == partner.Human.Sex &&
		(p.Human.Psycho.SexualOrientation.IsHomosexual() || p.Human.Psycho.SexualOrientation.IsBisexual()) &&
		(partner.Human.Psycho.SexualOrientation.IsHomosexual() || partner.Human.Psycho.SexualOrientation.IsBisexual()) {
		return true, nil
	}

	return false, nil
}

func (p *Person) DoesWantBeMarried(partner *Person, distance coordinate.ComplexDistance, religionsSimilarity, culturesSimilarity map[string]float64) (bool, error) {
	canBePartners, err := p.CanBePartners(partner)
	if err != nil {
		return false, wrapped_error.NewInternalServerError(err, "can not check if can be partners for does_want_be_married method")
	}
	if len(religionsSimilarity) == 0 {
		return false, wrapped_error.NewInternalServerError(err, "can not check if can be partners for religions_similarity is empty")
	}
	if len(culturesSimilarity) == 0 {
		return false, wrapped_error.NewInternalServerError(err, "can not check if can be partners for cultures_similarity is empty")
	}
	if !canBePartners {
		return false, nil
	}

	sympathicCoef, err := p.GetSympathicCoef(partner, religionsSimilarity, culturesSimilarity)
	if err != nil {
		return false, wrapped_error.NewInternalServerError(err, "can not get sympathic coef for does_want_be_married method")
	}

	var prob float64
	coefs := []struct {
		coef  float64
		value float64
	}{
		{
			coef:  0.2,
			value: sympathicCoef,
		},
		{
			coef:  0.3,
			value: GetMarriageDistanceCoef(distance),
		},
		{
			coef:  0.5,
			value: p.Metadata.WishGetMarriedCoef,
		},
	}
	for _, coef := range coefs {
		prob += coef.coef * coef.value
	}

	want, err := pm.GetRandomBool(prob)
	if err != nil {
		return false, wrapped_error.NewInternalServerError(err, "can not get random bool for does_want_be_married method")
	}

	return want, nil
}

func (p *Person) IsPregnant() (bool, error) {
	if p == nil {
		return false, wrapped_error.NewInternalServerError(nil, "<nil> can not be checked if it is pregnant")
	}

	return p.Human.IsPregnant(), nil
}

func (p *Person) Age() (int, error) {
	if p == nil {
		return 0, wrapped_error.NewInternalServerError(nil, "<nil> can not has age")
	}
	if p.Human == nil {
		return 0, wrapped_error.NewInternalServerError(nil, "person with <nil> human can not has age")
	}

	return p.Human.Age, nil
}

func (p *Person) Sex() (gender.Sex, error) {
	if p == nil {
		return "", wrapped_error.NewInternalServerError(nil, "<nil> can not has sex")
	}
	if p.Human == nil {
		return "", wrapped_error.NewInternalServerError(nil, "person with <nil> human can not has sex")
	}

	return p.Human.Sex, nil
}

// Chronology methods

func (p *Person) HadSex(year int) (bool, error) {
	if p == nil {
		return false, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("<nil> can not check if person has sex in %d year", year))
	}

	for _, r := range p.Chronology.Events {
		if r.Year == year && r.Name == HaveSexEventName {
			return true, nil
		}
	}

	return false, nil
}

func (p *Person) IsDead() (bool, error) {
	if p == nil {
		return false, wrapped_error.NewInternalServerError(nil, "<nil> can not check if person is dead")
	}

	return p.Chronology.DeathYear != -1, nil
}

func (p *Person) Die(year int) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, fmt.Sprintf("<nil> can not die in %d year", year))
	}

	p.Chronology.DeathYear = year

	return nil
}

// Traits methods

func (p *Person) AppendTrait(t *traits.Trait) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "can not appent trait into <nil>")
	}

	p.Traits = traits.Unique(append(p.Traits, t))

	return nil
}

func (p *Person) IsHomosexualist() (bool, error) {
	if p == nil {
		return false, wrapped_error.NewInternalServerError(nil, "<nil> can not check if person is homosexualist")
	}
	for _, t := range p.Traits {
		if t.Name == traits.HomosexualistTrait.Name {
			return true, nil
		}
	}

	return false, nil
}

func (p *Person) IsLustful() (bool, error) {
	if p == nil {
		return false, wrapped_error.NewInternalServerError(nil, "<nil> can not check if person is lustful")
	}

	for _, t := range p.Traits {
		if t.Name == traits.LustfulTrait.Name {
			return true, nil
		}
	}

	return false, nil
}

func (p *Person) ProcessLustfulTrait() error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "<nil> can not process lustful trait")
	}

	isLustful, err := p.IsLustful()
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can check if person is lustful by *Person entity")
	}
	if isLustful {
		return nil
	}

	isLustfulByChron, errByChron := p.Chronology.IsLustful()
	if errByChron != nil {
		return wrapped_error.NewInternalServerError(errByChron, "can not chech if person is lustful by chronology")
	}
	if !isLustfulByChron {
		return nil
	}
	if err := p.AppendTrait(traits.LustfulTrait); err != nil {
		return wrapped_error.NewInternalServerError(nil, "can not append lustful trait into person")
	}

	return nil
}

// Event methods

func (p *Person) appendEvent(e Event) error {
	if p == nil {
		return wrapped_error.NewInternalServerError(nil, "can not append event for <nil> *Person")
	}

	p.Chronology.Events = append(p.Chronology.Events, e)

	return nil
}

func (p *Person) AppendHaveSexEvent(year int, partner *Person) error {
	if err := p.appendEvent(Event{
		Year:  year,
		Name:  HaveSexEventName,
		Value: partner,
	}); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not append have_sex event")
	}

	return nil
}

func (p *Person) AppendGetMarriedEvent(year int, partner *Person) error {
	if err := p.appendEvent(Event{
		Year:  year,
		Name:  GetMarriedEventName,
		Value: partner,
	}); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not append get_married event")
	}

	return nil
}

func (p *Person) AppendDivorcedEvent(year int, partner *Person) error {
	if err := p.appendEvent(Event{
		Year:  year,
		Name:  DivorcedEventName,
		Value: partner,
	}); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not append divorced event")
	}

	return nil
}

// Print

func (p *Person) Print() {
	if p == nil {
		return
	}

	body, _ := json.Marshal(p.Human.Body)
	psycho, _ := json.Marshal(p.Human.Psycho)
	fmt.Printf(`{
		"id": "%s",
		"own_name": "%s",
		"culture_name": "%s",
		"religion_name": "%s",
		"sex": "%s",
		"age": %d,
		"body": %s,
		"psycho": %s
	}`+"\n", p.ID.String(), p.OwnName, p.Culture.Name, p.Religion.Name, p.Human.Sex.String(), p.Human.Age, body, psycho)
}
