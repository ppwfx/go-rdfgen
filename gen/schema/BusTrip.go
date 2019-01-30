package schema

type BusTripTrait struct {

	// The unique identifier for the bus.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BusNumber string `json:"busNumber,omitempty" jsonld:"http://schema.org/busNumber"`

	// The name of the bus (e.g. Bolt Express).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BusName string `json:"busName,omitempty" jsonld:"http://schema.org/busName"`

	// The stop or station from which the bus arrives.
	//
	// RangeIncludes:
	// - http://schema.org/BusStation
	// - http://schema.org/BusStop
	//
	ArrivalBusStop interface{} `json:"arrivalBusStop,omitempty" jsonld:"http://schema.org/arrivalBusStop"`

	// The stop or station from which the bus departs.
	//
	// RangeIncludes:
	// - http://schema.org/BusStation
	// - http://schema.org/BusStop
	//
	DepartureBusStop interface{} `json:"departureBusStop,omitempty" jsonld:"http://schema.org/departureBusStop"`

}

type BusTrip struct {
	MetaTrait
	BusTripTrait
	TripTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBusTrip() (x BusTrip) {
	x.Type = "http://schema.org/BusTrip"

	return
}
