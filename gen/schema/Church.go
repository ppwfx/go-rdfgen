package schema

type ChurchTrait struct {

}

type Church struct {
	MetaTrait
	ChurchTrait
	PlaceOfWorshipTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewChurch() (x Church) {
	x.Type = "http://schema.org/Church"

	return
}
