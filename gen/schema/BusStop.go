package schema

type BusStopTrait struct {

}

type BusStop struct {
	MetaTrait
	BusStopTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewBusStop() (x BusStop) {
	x.Type = "http://schema.org/BusStop"

	return
}
