package schema

type BedAndBreakfastTrait struct {

}

type BedAndBreakfast struct {
	MetaTrait
	BedAndBreakfastTrait
	LodgingBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewBedAndBreakfast() (x BedAndBreakfast) {
	x.Type = "http://schema.org/BedAndBreakfast"

	return
}
