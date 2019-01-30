package schema

type LodgingReservationTrait struct {

	// The earliest someone may check into a lodging establishment.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	CheckinTime *DateTime `json:"checkinTime,omitempty" jsonld:"http://schema.org/checkinTime"`

	// The latest someone may check out of a lodging establishment.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	CheckoutTime *DateTime `json:"checkoutTime,omitempty" jsonld:"http://schema.org/checkoutTime"`

	// A full description of the lodging unit.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	LodgingUnitDescription string `json:"lodgingUnitDescription,omitempty" jsonld:"http://schema.org/lodgingUnitDescription"`

	// Textual description of the unit type (including suite vs. room, size of bed, etc.).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/QualitativeValue
	//
	LodgingUnitType interface{} `json:"lodgingUnitType,omitempty" jsonld:"http://schema.org/lodgingUnitType"`

	// The number of adults staying in the unit.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Integer
	//
	NumAdults interface{} `json:"numAdults,omitempty" jsonld:"http://schema.org/numAdults"`

	// The number of children staying in the unit.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Integer
	//
	NumChildren interface{} `json:"numChildren,omitempty" jsonld:"http://schema.org/numChildren"`

}

type LodgingReservation struct {
	MetaTrait
	LodgingReservationTrait
	ReservationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewLodgingReservation() (x LodgingReservation) {
	x.Type = "http://schema.org/LodgingReservation"

	return
}
