package schema

type TrainTripTrait struct {

	// The platform from which the train departs.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	DeparturePlatform string `json:"departurePlatform,omitempty" jsonld:"http://schema.org/departurePlatform"`

	// The platform where the train arrives.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ArrivalPlatform string `json:"arrivalPlatform,omitempty" jsonld:"http://schema.org/arrivalPlatform"`

	// The name of the train (e.g. The Orient Express).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TrainName string `json:"trainName,omitempty" jsonld:"http://schema.org/trainName"`

	// The unique identifier for the train.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TrainNumber string `json:"trainNumber,omitempty" jsonld:"http://schema.org/trainNumber"`

	// The station where the train trip ends.
	//
	// RangeIncludes:
	// - http://schema.org/TrainStation
	//
	ArrivalStation *TrainStation `json:"arrivalStation,omitempty" jsonld:"http://schema.org/arrivalStation"`

	// The station from which the train departs.
	//
	// RangeIncludes:
	// - http://schema.org/TrainStation
	//
	DepartureStation *TrainStation `json:"departureStation,omitempty" jsonld:"http://schema.org/departureStation"`

}

type TrainTrip struct {
	MetaTrait
	TrainTripTrait
	TripTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewTrainTrip() (x TrainTrip) {
	x.Type = "http://schema.org/TrainTrip"

	return
}
