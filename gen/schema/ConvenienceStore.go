package schema

type ConvenienceStoreTrait struct {

}

type ConvenienceStore struct {
	MetaTrait
	ConvenienceStoreTrait
	StoreTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewConvenienceStore() (x ConvenienceStore) {
	x.Type = "http://schema.org/ConvenienceStore"

	return
}
