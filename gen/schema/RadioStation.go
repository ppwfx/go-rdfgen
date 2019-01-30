package schema

type RadioStationTrait struct {

}

type RadioStation struct {
	MetaTrait
	RadioStationTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewRadioStation() (x RadioStation) {
	x.Type = "http://schema.org/RadioStation"

	return
}
