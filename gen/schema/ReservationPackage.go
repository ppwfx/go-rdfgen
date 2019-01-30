package schema

type ReservationPackageTrait struct {

	// The individual reservations included in the package. Typically a repeated property.
	//
	// RangeIncludes:
	// - http://schema.org/Reservation
	//
	SubReservation *Reservation `json:"subReservation,omitempty" jsonld:"http://schema.org/subReservation"`

}

type ReservationPackage struct {
	MetaTrait
	ReservationPackageTrait
	ReservationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewReservationPackage() (x ReservationPackage) {
	x.Type = "http://schema.org/ReservationPackage"

	return
}
