package schema

type RatingTrait struct {

	// The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	// - http://schema.org/Person
	//
	Author interface{} `json:"author,omitempty" jsonld:"http://schema.org/author"`

	// The highest value allowed in this rating system. If bestRating is omitted, 5 is assumed.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	BestRating interface{} `json:"bestRating,omitempty" jsonld:"http://schema.org/bestRating"`

	// The rating for the content.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	RatingValue interface{} `json:"ratingValue,omitempty" jsonld:"http://schema.org/ratingValue"`

	// This Review or Rating is relevant to this part or facet of the itemReviewed.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ReviewAspect string `json:"reviewAspect,omitempty" jsonld:"http://schema.org/reviewAspect"`

	// The lowest value allowed in this rating system. If worstRating is omitted, 1 is assumed.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	WorstRating interface{} `json:"worstRating,omitempty" jsonld:"http://schema.org/worstRating"`

}

type Rating struct {
	MetaTrait
	RatingTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewRating() (x Rating) {
	x.Type = "http://schema.org/Rating"

	return
}
