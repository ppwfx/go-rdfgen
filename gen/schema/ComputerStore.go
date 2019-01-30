package schema

type ComputerStoreTrait struct {

}

type ComputerStore struct {
	MetaTrait
	ComputerStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewComputerStore() (x ComputerStore) {
	x.Type = "http://schema.org/ComputerStore"

	return
}
