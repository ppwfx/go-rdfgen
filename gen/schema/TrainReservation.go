package schema

type TrainReservationTrait struct {

}

type TrainReservation struct {
	MetaTrait
	TrainReservationTrait
	ReservationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewTrainReservation() (x TrainReservation) {
	x.Type = "http://schema.org/TrainReservation"

	return
}
