package schema

type GatedResidenceCommunityTrait struct {

}

type GatedResidenceCommunity struct {
	MetaTrait
	GatedResidenceCommunityTrait
	ResidenceTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewGatedResidenceCommunity() (x GatedResidenceCommunity) {
	x.Type = "http://schema.org/GatedResidenceCommunity"

	return
}
