package schema

type FlightTrait struct {

	// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Seller interface{} `json:"seller,omitempty" jsonld:"http://schema.org/seller"`

	// 'carrier' is an out-dated term indicating the 'provider' for parcel delivery and flights.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	Carrier *Organization `json:"carrier,omitempty" jsonld:"http://schema.org/carrier"`

	// The unique identifier for a flight including the airline IATA code. For example, if describing United flight 110, where the IATA code for United is 'UA', the flightNumber is 'UA110'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	FlightNumber string `json:"flightNumber,omitempty" jsonld:"http://schema.org/flightNumber"`

	// The time when a passenger can check into the flight online.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	WebCheckinTime *DateTime `json:"webCheckinTime,omitempty" jsonld:"http://schema.org/webCheckinTime"`

	// The distance of the flight.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Distance
	//
	FlightDistance interface{} `json:"flightDistance,omitempty" jsonld:"http://schema.org/flightDistance"`

	// Identifier of the flight's departure gate.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	DepartureGate string `json:"departureGate,omitempty" jsonld:"http://schema.org/departureGate"`

	// Identifier of the flight's arrival terminal.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ArrivalTerminal string `json:"arrivalTerminal,omitempty" jsonld:"http://schema.org/arrivalTerminal"`

	// The kind of aircraft (e.g., \"Boeing 747\").
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Vehicle
	//
	Aircraft interface{} `json:"aircraft,omitempty" jsonld:"http://schema.org/aircraft"`

	// Identifier of the flight's departure terminal.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	DepartureTerminal string `json:"departureTerminal,omitempty" jsonld:"http://schema.org/departureTerminal"`

	// Identifier of the flight's arrival gate.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ArrivalGate string `json:"arrivalGate,omitempty" jsonld:"http://schema.org/arrivalGate"`

	// Description of the meals that will be provided or available for purchase.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	MealService string `json:"mealService,omitempty" jsonld:"http://schema.org/mealService"`

	// The estimated time the flight will take.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Duration
	//
	EstimatedFlightDuration interface{} `json:"estimatedFlightDuration,omitempty" jsonld:"http://schema.org/estimatedFlightDuration"`

	// The airport where the flight terminates.
	//
	// RangeIncludes:
	// - http://schema.org/Airport
	//
	ArrivalAirport *Airport `json:"arrivalAirport,omitempty" jsonld:"http://schema.org/arrivalAirport"`

	// The type of boarding policy used by the airline (e.g. zone-based or group-based).
	//
	// RangeIncludes:
	// - http://schema.org/BoardingPolicyType
	//
	BoardingPolicy *BoardingPolicyType `json:"boardingPolicy,omitempty" jsonld:"http://schema.org/boardingPolicy"`

	// The airport where the flight originates.
	//
	// RangeIncludes:
	// - http://schema.org/Airport
	//
	DepartureAirport *Airport `json:"departureAirport,omitempty" jsonld:"http://schema.org/departureAirport"`

}

type Flight struct {
	MetaTrait
	FlightTrait
	TripTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewFlight() (x Flight) {
	x.Type = "http://schema.org/Flight"

	return
}
