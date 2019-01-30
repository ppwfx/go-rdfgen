package schema

type TouristTripTrait struct {

	// Attraction suitable for type(s) of tourist. eg. Children, visitors from a particular country, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Audience
	//
	TouristType interface{} `json:"touristType,omitempty" jsonld:"http://schema.org/touristType"`

}

type TouristTrip struct {
	MetaTrait
	TouristTripTrait
	TripTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewTouristTrip() (x TouristTrip) {
	x.Type = "http://schema.org/TouristTrip"

	return
}
