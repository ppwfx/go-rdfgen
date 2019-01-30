package schema

type TaxiReservationTrait struct {

	// When a taxi will pickup a passenger or a rental car can be picked up.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	PickupTime *DateTime `json:"pickupTime,omitempty" jsonld:"http://schema.org/pickupTime"`

	// Number of people the reservation should accommodate.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Integer
	//
	PartySize interface{} `json:"partySize,omitempty" jsonld:"http://schema.org/partySize"`

	// Where a taxi will pick up a passenger or a rental car can be picked up.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	PickupLocation *Place `json:"pickupLocation,omitempty" jsonld:"http://schema.org/pickupLocation"`

}

type TaxiReservation struct {
	MetaTrait
	TaxiReservationTrait
	ReservationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewTaxiReservation() (x TaxiReservation) {
	x.Type = "http://schema.org/TaxiReservation"

	return
}
