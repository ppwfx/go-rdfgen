package schema

type AggregateRatingTrait struct {

	// The item that is being reviewed/rated.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	ItemReviewed *Thing `json:"itemReviewed,omitempty" jsonld:"http://schema.org/itemReviewed"`

	// The count of total number of ratings.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	RatingCount *Integer `json:"ratingCount,omitempty" jsonld:"http://schema.org/ratingCount"`

	// The count of total number of reviews.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	ReviewCount *Integer `json:"reviewCount,omitempty" jsonld:"http://schema.org/reviewCount"`

}

type AggregateRating struct {
	MetaTrait
	AggregateRatingTrait
	RatingTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewAggregateRating() (x AggregateRating) {
	x.Type = "http://schema.org/AggregateRating"

	return
}
