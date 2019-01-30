package schema

type AirlineTrait struct {

	// IATA identifier for an airline or airport.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	IataCode string `json:"iataCode,omitempty" jsonld:"http://schema.org/iataCode"`

	// The type of boarding policy used by the airline (e.g. zone-based or group-based).
	//
	// RangeIncludes:
	// - http://schema.org/BoardingPolicyType
	//
	BoardingPolicy *BoardingPolicyType `json:"boardingPolicy,omitempty" jsonld:"http://schema.org/boardingPolicy"`

}

type Airline struct {
	MetaTrait
	AirlineTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewAirline() (x Airline) {
	x.Type = "http://schema.org/Airline"

	return
}
