package schema

type DryCleaningOrLaundryTrait struct {

}

type DryCleaningOrLaundry struct {
	MetaTrait
	DryCleaningOrLaundryTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewDryCleaningOrLaundry() (x DryCleaningOrLaundry) {
	x.Type = "http://schema.org/DryCleaningOrLaundry"

	return
}
