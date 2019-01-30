package schema

type HouseTrait struct {

	// The number of rooms (excluding bathrooms and closets) of the accommodation or lodging business.\nTypical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	NumberOfRooms interface{} `json:"numberOfRooms,omitempty" jsonld:"http://schema.org/numberOfRooms"`

}

type House struct {
	MetaTrait
	HouseTrait
	AccommodationTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewHouse() (x House) {
	x.Type = "http://schema.org/House"

	return
}
