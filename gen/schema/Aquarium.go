package schema

type AquariumTrait struct {

}

type Aquarium struct {
	MetaTrait
	AquariumTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewAquarium() (x Aquarium) {
	x.Type = "http://schema.org/Aquarium"

	return
}
