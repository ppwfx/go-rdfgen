package schema

type GovernmentOfficeTrait struct {

}

type GovernmentOffice struct {
	MetaTrait
	GovernmentOfficeTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewGovernmentOffice() (x GovernmentOffice) {
	x.Type = "http://schema.org/GovernmentOffice"

	return
}
