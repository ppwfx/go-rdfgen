package schema

type CityHallTrait struct {

}

type CityHall struct {
	MetaTrait
	CityHallTrait
	GovernmentBuildingTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewCityHall() (x CityHall) {
	x.Type = "http://schema.org/CityHall"

	return
}
