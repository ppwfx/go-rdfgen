package schema

type SuiteTrait struct {

	// The type of bed or beds included in the accommodation. For the single case of just one bed of a certain type, you use bed directly with a text.\n      If you want to indicate the quantity of a certain kind of bed, use an instance of BedDetails. For more detailed information, use the amenityFeature property.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/BedDetails
	// - http://schema.org/BedType
	//
	Bed interface{} `json:"bed,omitempty" jsonld:"http://schema.org/bed"`

	// The allowed total occupancy for the accommodation in persons (including infants etc). For individual accommodations, this is not necessarily the legal maximum but defines the permitted usage as per the contractual agreement (e.g. a double room used by a single person).\nTypical unit code(s): C62 for person
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	Occupancy *QuantitativeValue `json:"occupancy,omitempty" jsonld:"http://schema.org/occupancy"`

	// The number of rooms (excluding bathrooms and closets) of the accommodation or lodging business.\nTypical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	NumberOfRooms interface{} `json:"numberOfRooms,omitempty" jsonld:"http://schema.org/numberOfRooms"`

}

type Suite struct {
	MetaTrait
	SuiteTrait
	AccommodationTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewSuite() (x Suite) {
	x.Type = "http://schema.org/Suite"

	return
}
