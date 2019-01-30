package schema

type CampgroundTrait struct {

}

type Campground struct {
	MetaTrait
	CampgroundTrait
	LodgingBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	CivicStructureTrait
	AdditionalTrait
}

func NewCampground() (x Campground) {
	x.Type = "http://schema.org/Campground"

	return
}
