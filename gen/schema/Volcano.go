package schema

type VolcanoTrait struct {

}

type Volcano struct {
	MetaTrait
	VolcanoTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewVolcano() (x Volcano) {
	x.Type = "http://schema.org/Volcano"

	return
}
