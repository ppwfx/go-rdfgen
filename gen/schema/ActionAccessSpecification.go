package schema

type ActionAccessSpecificationTrait struct {

	// Indicates if use of the media require a subscription  (either paid or free). Allowed values are <code>true</code> or <code>false</code> (note that an earlier version had 'yes', 'no').
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	// - http://schema.org/MediaSubscription
	//
	RequiresSubscription interface{} `json:"requiresSubscription,omitempty" jsonld:"http://schema.org/requiresSubscription"`

	// The end of the availability of the product or service included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	AvailabilityEnds *DateTime `json:"availabilityEnds,omitempty" jsonld:"http://schema.org/availabilityEnds"`

	// The beginning of the availability of the product or service included in the offer.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	AvailabilityStarts *DateTime `json:"availabilityStarts,omitempty" jsonld:"http://schema.org/availabilityStarts"`

	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.<br/><br/>\n\nSee also <a class=\"localLink\" href=\"http://schema.org/ineligibleRegion\">ineligibleRegion</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/GeoShape
	// - http://schema.org/Place
	//
	EligibleRegion interface{} `json:"eligibleRegion,omitempty" jsonld:"http://schema.org/eligibleRegion"`

	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	// - http://schema.org/PhysicalActivityCategory
	//
	Category interface{} `json:"category,omitempty" jsonld:"http://schema.org/category"`

	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf,omitempty" jsonld:"http://schema.org/expectsAcceptanceOf"`

}

type ActionAccessSpecification struct {
	MetaTrait
	ActionAccessSpecificationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewActionAccessSpecification() (x ActionAccessSpecification) {
	x.Type = "http://schema.org/ActionAccessSpecification"

	return
}
