package schema

type AutoPartsStoreTrait struct {

}

type AutoPartsStore struct {
	MetaTrait
	AutoPartsStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AutomotiveBusinessTrait
	AdditionalTrait
}

func NewAutoPartsStore() (x AutoPartsStore) {
	x.Type = "http://schema.org/AutoPartsStore"

	return
}
