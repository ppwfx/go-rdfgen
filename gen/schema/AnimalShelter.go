package schema

type AnimalShelterTrait struct {

}

type AnimalShelter struct {
	MetaTrait
	AnimalShelterTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewAnimalShelter() (x AnimalShelter) {
	x.Type = "http://schema.org/AnimalShelter"

	return
}
