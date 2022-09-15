package traits

var MarriageTraits = []*Trait{
	MarriedTrait,
	DivorcedTrait,
	AdultererTrait,
}

var (
	MarriedTrait = &Trait{
		Name: "married",
	}
	DivorcedTrait = &Trait{
		Name: "divorced",
	}
	AdultererTrait = &Trait{
		Name: "adulterer",
	}
)
