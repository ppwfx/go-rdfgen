package schema

type PostOfficeTrait struct {

}

type PostOffice struct {
	MetaTrait
	PostOfficeTrait
	GovernmentOfficeTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewPostOffice() (x PostOffice) {
	x.Type = "http://schema.org/PostOffice"

	return
}
