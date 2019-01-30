package schema

type TouristDestinationTrait struct {

	// Attraction suitable for type(s) of tourist. eg. Children, visitors from a particular country, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Audience
	//
	TouristType interface{} `json:"touristType,omitempty" jsonld:"http://schema.org/touristType"`

	// Attraction located at destination.
	//
	// RangeIncludes:
	// - http://schema.org/TouristAttraction
	//
	IncludesAttraction *TouristAttraction `json:"includesAttraction,omitempty" jsonld:"http://schema.org/includesAttraction"`

}

type TouristDestination struct {
	MetaTrait
	TouristDestinationTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewTouristDestination() (x TouristDestination) {
	x.Type = "http://schema.org/TouristDestination"

	return
}
