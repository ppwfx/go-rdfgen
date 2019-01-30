package schema

type TouristInformationCenterTrait struct {

}

type TouristInformationCenter struct {
	MetaTrait
	TouristInformationCenterTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewTouristInformationCenter() (x TouristInformationCenter) {
	x.Type = "http://schema.org/TouristInformationCenter"

	return
}
