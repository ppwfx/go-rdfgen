package schema

type HotelTrait struct {

}

type Hotel struct {
	MetaTrait
	HotelTrait
	LodgingBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewHotel() (x Hotel) {
	x.Type = "http://schema.org/Hotel"

	return
}
