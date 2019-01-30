package schema

type TrainStationTrait struct {

}

type TrainStation struct {
	MetaTrait
	TrainStationTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewTrainStation() (x TrainStation) {
	x.Type = "http://schema.org/TrainStation"

	return
}
