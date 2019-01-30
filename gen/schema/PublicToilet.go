package schema

type PublicToiletTrait struct {

}

type PublicToilet struct {
	MetaTrait
	PublicToiletTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewPublicToilet() (x PublicToilet) {
	x.Type = "http://schema.org/PublicToilet"

	return
}
