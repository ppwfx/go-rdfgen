package schema

type ElectronicsStoreTrait struct {

}

type ElectronicsStore struct {
	MetaTrait
	ElectronicsStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewElectronicsStore() (x ElectronicsStore) {
	x.Type = "http://schema.org/ElectronicsStore"

	return
}
