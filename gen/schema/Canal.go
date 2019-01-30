package schema

type CanalTrait struct {

}

type Canal struct {
	MetaTrait
	CanalTrait
	BodyOfWaterTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewCanal() (x Canal) {
	x.Type = "http://schema.org/Canal"

	return
}
