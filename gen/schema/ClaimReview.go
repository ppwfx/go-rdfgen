package schema

type ClaimReviewTrait struct {

	// A short summary of the specific claims reviewed in a ClaimReview.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ClaimReviewed string `json:"claimReviewed,omitempty" jsonld:"http://schema.org/claimReviewed"`

}

type ClaimReview struct {
	MetaTrait
	ClaimReviewTrait
	ReviewTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewClaimReview() (x ClaimReview) {
	x.Type = "http://schema.org/ClaimReview"

	return
}
