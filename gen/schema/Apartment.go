package schema

type ApartmentTrait struct {

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

type Apartment struct {
	MetaTrait
	ApartmentTrait
	AccommodationTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewApartment() (x Apartment) {
	x.Type = "http://schema.org/Apartment"

	return
}
