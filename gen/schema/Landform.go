package schema

type LandformTrait struct {

}

type Landform struct {
	MetaTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewLandform() (x Landform) {
	x.Type = "http://schema.org/Landform"

	return
}
