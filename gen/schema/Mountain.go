package schema

type MountainTrait struct {

}

type Mountain struct {
	MetaTrait
	MountainTrait
	LandformTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewMountain() (x Mountain) {
	x.Type = "http://schema.org/Mountain"

	return
}
