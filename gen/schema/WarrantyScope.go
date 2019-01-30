package schema

type WarrantyScopeTrait struct {

}

type WarrantyScope struct {
	MetaTrait
	WarrantyScopeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewWarrantyScope() (x WarrantyScope) {
	x.Type = "http://schema.org/WarrantyScope"

	return
}
