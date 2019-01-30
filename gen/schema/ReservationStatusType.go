package schema

type ReservationStatusTypeTrait struct {

}

type ReservationStatusType struct {
	MetaTrait
	ReservationStatusTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewReservationStatusType() (x ReservationStatusType) {
	x.Type = "http://schema.org/ReservationStatusType"

	return
}
