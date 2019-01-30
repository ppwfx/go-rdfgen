package schema

type BusStationTrait struct {

}

type BusStation struct {
	MetaTrait
	BusStationTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewBusStation() (x BusStation) {
	x.Type = "http://schema.org/BusStation"

	return
}
