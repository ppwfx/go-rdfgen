package schema

type MovieTheaterTrait struct {

	// The number of screens in the movie theater.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	ScreenCount float64 `json:"screenCount,omitempty" jsonld:"http://schema.org/screenCount"`

}

type MovieTheater struct {
	MetaTrait
	MovieTheaterTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	EntertainmentBusinessTrait
	LocalBusinessTrait
	OrganizationTrait
	AdditionalTrait
}

func NewMovieTheater() (x MovieTheater) {
	x.Type = "http://schema.org/MovieTheater"

	return
}
