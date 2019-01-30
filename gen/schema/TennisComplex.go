package schema

type TennisComplexTrait struct {

}

type TennisComplex struct {
	MetaTrait
	TennisComplexTrait
	SportsActivityLocationTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewTennisComplex() (x TennisComplex) {
	x.Type = "http://schema.org/TennisComplex"

	return
}
