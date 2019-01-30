package schema

type EmbassyTrait struct {

}

type Embassy struct {
	MetaTrait
	EmbassyTrait
	GovernmentBuildingTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewEmbassy() (x Embassy) {
	x.Type = "http://schema.org/Embassy"

	return
}
