package schema

type FlightReservationTrait struct {

	// The priority status assigned to a passenger for security or boarding (e.g. FastTrack or Priority).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/QualitativeValue
	//
	PassengerPriorityStatus interface{} `json:"passengerPriorityStatus,omitempty" jsonld:"http://schema.org/passengerPriorityStatus"`

	// The passenger's sequence number as assigned by the airline.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PassengerSequenceNumber string `json:"passengerSequenceNumber,omitempty" jsonld:"http://schema.org/passengerSequenceNumber"`

	// The type of security screening the passenger is subject to.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SecurityScreening string `json:"securityScreening,omitempty" jsonld:"http://schema.org/securityScreening"`

	// The airline-specific indicator of boarding order / preference.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BoardingGroup string `json:"boardingGroup,omitempty" jsonld:"http://schema.org/boardingGroup"`

}

type FlightReservation struct {
	MetaTrait
	FlightReservationTrait
	ReservationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewFlightReservation() (x FlightReservation) {
	x.Type = "http://schema.org/FlightReservation"

	return
}
