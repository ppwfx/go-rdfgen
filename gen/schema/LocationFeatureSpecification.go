package schema

type LocationFeatureSpecificationTrait struct {

	// The date when the item becomes valid.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ValidFrom *DateTime `json:"validFrom,omitempty" jsonld:"http://schema.org/validFrom"`

	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	ValidThrough *DateTime `json:"validThrough,omitempty" jsonld:"http://schema.org/validThrough"`

	// The hours during which this service or contact is available.
	//
	// RangeIncludes:
	// - http://schema.org/OpeningHoursSpecification
	//
	HoursAvailable *OpeningHoursSpecification `json:"hoursAvailable,omitempty" jsonld:"http://schema.org/hoursAvailable"`

}

type LocationFeatureSpecification struct {
	MetaTrait
	LocationFeatureSpecificationTrait
	PropertyValueTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewLocationFeatureSpecification() (x LocationFeatureSpecification) {
	x.Type = "http://schema.org/LocationFeatureSpecification"

	return
}
