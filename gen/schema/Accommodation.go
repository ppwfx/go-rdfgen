package schema

type AccommodationTrait struct {

	// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	//
	// RangeIncludes:
	// - http://schema.org/LocationFeatureSpecification
	//
	AmenityFeature *LocationFeatureSpecification `json:"amenityFeature,omitempty" jsonld:"http://schema.org/amenityFeature"`

	// Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Boolean
	//
	PetsAllowed interface{} `json:"petsAllowed,omitempty" jsonld:"http://schema.org/petsAllowed"`

	// Indications regarding the permitted usage of the accommodation.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PermittedUsage string `json:"permittedUsage,omitempty" jsonld:"http://schema.org/permittedUsage"`

	// The number of rooms (excluding bathrooms and closets) of the accommodation or lodging business.\nTypical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	NumberOfRooms interface{} `json:"numberOfRooms,omitempty" jsonld:"http://schema.org/numberOfRooms"`

	// The size of the accommodation, e.g. in square meter or squarefoot.\nTypical unit code(s): MTK for square meter, FTK for square foot, or YDK for square yard
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	FloorSize *QuantitativeValue `json:"floorSize,omitempty" jsonld:"http://schema.org/floorSize"`

}

type Accommodation struct {
	MetaTrait
	AccommodationTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewAccommodation() (x Accommodation) {
	x.Type = "http://schema.org/Accommodation"

	return
}
