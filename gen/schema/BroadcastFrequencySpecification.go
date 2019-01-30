package schema

type BroadcastFrequencySpecificationTrait struct {

	// The frequency in MHz for a particular broadcast.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	// - http://schema.org/Number
	//
	BroadcastFrequencyValue interface{} `json:"broadcastFrequencyValue,omitempty" jsonld:"http://schema.org/broadcastFrequencyValue"`

}

type BroadcastFrequencySpecification struct {
	MetaTrait
	BroadcastFrequencySpecificationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBroadcastFrequencySpecification() (x BroadcastFrequencySpecification) {
	x.Type = "http://schema.org/BroadcastFrequencySpecification"

	return
}
