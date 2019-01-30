package schema

type AirportTrait struct {

	// IATA identifier for an airline or airport.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	IataCode string `json:"iataCode,omitempty" jsonld:"http://schema.org/iataCode"`

	// ICAO identifier for an airport.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	IcaoCode string `json:"icaoCode,omitempty" jsonld:"http://schema.org/icaoCode"`

}

type Airport struct {
	MetaTrait
	AirportTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewAirport() (x Airport) {
	x.Type = "http://schema.org/Airport"

	return
}
