package schema

type BusReservationTrait struct {

}

type BusReservation struct {
	MetaTrait
	BusReservationTrait
	ReservationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBusReservation() (x BusReservation) {
	x.Type = "http://schema.org/BusReservation"

	return
}
