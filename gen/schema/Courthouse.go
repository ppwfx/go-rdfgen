package schema

type CourthouseTrait struct {

}

type Courthouse struct {
	MetaTrait
	CourthouseTrait
	GovernmentBuildingTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewCourthouse() (x Courthouse) {
	x.Type = "http://schema.org/Courthouse"

	return
}
