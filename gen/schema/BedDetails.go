package schema

type BedDetailsTrait struct {

	// The quantity of the given bed type available in the HotelRoom, Suite, House, or Apartment.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	NumberOfBeds float64 `json:"numberOfBeds,omitempty" jsonld:"http://schema.org/numberOfBeds"`

	// The type of bed to which the BedDetail refers, i.e. the type of bed available in the quantity indicated by quantity.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/BedType
	//
	TypeOfBed interface{} `json:"typeOfBed,omitempty" jsonld:"http://schema.org/typeOfBed"`

}

type BedDetails struct {
	MetaTrait
	BedDetailsTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBedDetails() (x BedDetails) {
	x.Type = "http://schema.org/BedDetails"

	return
}
