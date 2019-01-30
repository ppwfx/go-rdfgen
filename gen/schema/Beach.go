package schema

type BeachTrait struct {

}

type Beach struct {
	MetaTrait
	BeachTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewBeach() (x Beach) {
	x.Type = "http://schema.org/Beach"

	return
}
