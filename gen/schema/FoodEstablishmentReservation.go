package schema

type FoodEstablishmentReservationTrait struct {

	// The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to <em>December</em>.<br/><br/>\n\nNote that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	EndTime *DateTime `json:"endTime,omitempty" jsonld:"http://schema.org/endTime"`

	// The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from <em>January</em> to December.<br/><br/>\n\nNote that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	StartTime *DateTime `json:"startTime,omitempty" jsonld:"http://schema.org/startTime"`

	// Number of people the reservation should accommodate.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Integer
	//
	PartySize interface{} `json:"partySize,omitempty" jsonld:"http://schema.org/partySize"`

}

type FoodEstablishmentReservation struct {
	MetaTrait
	FoodEstablishmentReservationTrait
	ReservationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewFoodEstablishmentReservation() (x FoodEstablishmentReservation) {
	x.Type = "http://schema.org/FoodEstablishmentReservation"

	return
}
