package schema

type HinduTempleTrait struct {

}

type HinduTemple struct {
	MetaTrait
	HinduTempleTrait
	PlaceOfWorshipTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewHinduTemple() (x HinduTemple) {
	x.Type = "http://schema.org/HinduTemple"

	return
}
