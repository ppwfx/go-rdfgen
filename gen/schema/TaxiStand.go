package schema

type TaxiStandTrait struct {

}

type TaxiStand struct {
	MetaTrait
	TaxiStandTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewTaxiStand() (x TaxiStand) {
	x.Type = "http://schema.org/TaxiStand"

	return
}
