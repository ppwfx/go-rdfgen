package schema

type SelfStorageTrait struct {

}

type SelfStorage struct {
	MetaTrait
	SelfStorageTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewSelfStorage() (x SelfStorage) {
	x.Type = "http://schema.org/SelfStorage"

	return
}
