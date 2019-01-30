package schema

type LandmarksOrHistoricalBuildingsTrait struct {

}

type LandmarksOrHistoricalBuildings struct {
	MetaTrait
	LandmarksOrHistoricalBuildingsTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewLandmarksOrHistoricalBuildings() (x LandmarksOrHistoricalBuildings) {
	x.Type = "http://schema.org/LandmarksOrHistoricalBuildings"

	return
}
