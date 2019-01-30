package schema

type TelevisionStationTrait struct {

}

type TelevisionStation struct {
	MetaTrait
	TelevisionStationTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewTelevisionStation() (x TelevisionStation) {
	x.Type = "http://schema.org/TelevisionStation"

	return
}
