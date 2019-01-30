package schema

type ReviewActionTrait struct {

	// A sub property of result. The review that resulted in the performing of the action.
	//
	// RangeIncludes:
	// - http://schema.org/Review
	//
	ResultReview *Review `json:"resultReview,omitempty" jsonld:"http://schema.org/resultReview"`

}

type ReviewAction struct {
	MetaTrait
	ReviewActionTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewReviewAction() (x ReviewAction) {
	x.Type = "http://schema.org/ReviewAction"

	return
}
