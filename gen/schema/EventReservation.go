package schema

type EventReservationTrait struct {

}

type EventReservation struct {
	MetaTrait
	EventReservationTrait
	ReservationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewEventReservation() (x EventReservation) {
	x.Type = "http://schema.org/EventReservation"

	return
}
