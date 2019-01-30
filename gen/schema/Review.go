package schema

type ReviewTrait struct {

	// This Review or Rating is relevant to this part or facet of the itemReviewed.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ReviewAspect string `json:"reviewAspect,omitempty" jsonld:"http://schema.org/reviewAspect"`

	// The item that is being reviewed/rated.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	//
	ItemReviewed *Thing `json:"itemReviewed,omitempty" jsonld:"http://schema.org/itemReviewed"`

	// The actual body of the review.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ReviewBody string `json:"reviewBody,omitempty" jsonld:"http://schema.org/reviewBody"`

	// The rating given in this review. Note that reviews can themselves be rated. The <code>reviewRating</code> applies to rating given by the review. The <a class=\"localLink\" href=\"http://schema.org/aggregateRating\">aggregateRating</a> property applies to the review itself, as a creative work.
	//
	// RangeIncludes:
	// - http://schema.org/Rating
	//
	ReviewRating *Rating `json:"reviewRating,omitempty" jsonld:"http://schema.org/reviewRating"`

}

type Review struct {
	MetaTrait
	ReviewTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewReview() (x Review) {
	x.Type = "http://schema.org/Review"

	return
}
