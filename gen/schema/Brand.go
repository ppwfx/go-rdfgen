package schema

type BrandTrait struct {

	// The overall rating, based on a collection of reviews or ratings, of the item.
	//
	// RangeIncludes:
	// - http://schema.org/AggregateRating
	//
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty" jsonld:"http://schema.org/aggregateRating"`

	// A review of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Review
	//
	Review *Review `json:"review,omitempty" jsonld:"http://schema.org/review"`

	// An associated logo.
	//
	// RangeIncludes:
	// - http://schema.org/ImageObject
	// - http://schema.org/URL
	//
	Logo interface{} `json:"logo,omitempty" jsonld:"http://schema.org/logo"`

}

type Brand struct {
	MetaTrait
	BrandTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBrand() (x Brand) {
	x.Type = "http://schema.org/Brand"

	return
}
