package schema

type RentalCarReservationTrait struct {

	// When a taxi will pickup a passenger or a rental car can be picked up.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	PickupTime *DateTime `json:"pickupTime,omitempty" jsonld:"http://schema.org/pickupTime"`

	// When a rental car can be dropped off.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	DropoffTime *DateTime `json:"dropoffTime,omitempty" jsonld:"http://schema.org/dropoffTime"`

	// Where a taxi will pick up a passenger or a rental car can be picked up.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	PickupLocation *Place `json:"pickupLocation,omitempty" jsonld:"http://schema.org/pickupLocation"`

	// Where a rental car can be dropped off.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	DropoffLocation *Place `json:"dropoffLocation,omitempty" jsonld:"http://schema.org/dropoffLocation"`

}

type RentalCarReservation struct {
	MetaTrait
	RentalCarReservationTrait
	ReservationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewRentalCarReservation() (x RentalCarReservation) {
	x.Type = "http://schema.org/RentalCarReservation"

	return
}
