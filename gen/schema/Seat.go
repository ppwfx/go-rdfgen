package schema

type SeatTrait struct {

	// The location of the reserved seat (e.g., 27).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SeatNumber string `json:"seatNumber,omitempty" jsonld:"http://schema.org/seatNumber"`

	// The row location of the reserved seat (e.g., B).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SeatRow string `json:"seatRow,omitempty" jsonld:"http://schema.org/seatRow"`

	// The section location of the reserved seat (e.g. Orchestra).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SeatSection string `json:"seatSection,omitempty" jsonld:"http://schema.org/seatSection"`

	// The type/class of the seat.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/QualitativeValue
	//
	SeatingType interface{} `json:"seatingType,omitempty" jsonld:"http://schema.org/seatingType"`

}

type Seat struct {
	MetaTrait
	SeatTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewSeat() (x Seat) {
	x.Type = "http://schema.org/Seat"

	return
}
