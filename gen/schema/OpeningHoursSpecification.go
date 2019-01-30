package schema

type OpeningHoursSpecificationTrait struct {

	// The closing hour of the place or service on the given day(s) of the week.
	//
	// RangeIncludes:
	// - http://schema.org/Time
	//
	Closes *Time `json:"closes,omitempty" jsonld:"http://schema.org/closes"`

	// The day of the week for which these opening hours are valid.
	//
	// RangeIncludes:
	// - http://schema.org/DayOfWeek
	//
	DayOfWeek *DayOfWeek `json:"dayOfWeek,omitempty" jsonld:"http://schema.org/dayOfWeek"`

	// The opening hour of the place or service on the given day(s) of the week.
	//
	// RangeIncludes:
	// - http://schema.org/Time
	//
	Opens *Time `json:"opens,omitempty" jsonld:"http://schema.org/opens"`

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

}

type OpeningHoursSpecification struct {
	MetaTrait
	OpeningHoursSpecificationTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewOpeningHoursSpecification() (x OpeningHoursSpecification) {
	x.Type = "http://schema.org/OpeningHoursSpecification"

	return
}
