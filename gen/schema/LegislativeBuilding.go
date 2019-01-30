package schema

type LegislativeBuildingTrait struct {

}

type LegislativeBuilding struct {
	MetaTrait
	LegislativeBuildingTrait
	GovernmentBuildingTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewLegislativeBuilding() (x LegislativeBuilding) {
	x.Type = "http://schema.org/LegislativeBuilding"

	return
}
