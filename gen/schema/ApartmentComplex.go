package schema

type ApartmentComplexTrait struct {

}

type ApartmentComplex struct {
	MetaTrait
	ApartmentComplexTrait
	ResidenceTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewApartmentComplex() (x ApartmentComplex) {
	x.Type = "http://schema.org/ApartmentComplex"

	return
}
