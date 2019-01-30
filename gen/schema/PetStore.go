package schema

type PetStoreTrait struct {

}

type PetStore struct {
	MetaTrait
	PetStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewPetStore() (x PetStore) {
	x.Type = "http://schema.org/PetStore"

	return
}
