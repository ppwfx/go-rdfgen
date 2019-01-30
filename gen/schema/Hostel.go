package schema

type HostelTrait struct {

}

type Hostel struct {
	MetaTrait
	HostelTrait
	LodgingBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewHostel() (x Hostel) {
	x.Type = "http://schema.org/Hostel"

	return
}
