package schema

type CampingPitchTrait struct {

}

type CampingPitch struct {
	MetaTrait
	CampingPitchTrait
	AccommodationTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewCampingPitch() (x CampingPitch) {
	x.Type = "http://schema.org/CampingPitch"

	return
}
