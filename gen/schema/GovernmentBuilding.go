package schema

type GovernmentBuildingTrait struct {

}

type GovernmentBuilding struct {
	MetaTrait
	GovernmentBuildingTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewGovernmentBuilding() (x GovernmentBuilding) {
	x.Type = "http://schema.org/GovernmentBuilding"

	return
}
