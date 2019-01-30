package schema

type HardwareStoreTrait struct {

}

type HardwareStore struct {
	MetaTrait
	HardwareStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewHardwareStore() (x HardwareStore) {
	x.Type = "http://schema.org/HardwareStore"

	return
}
