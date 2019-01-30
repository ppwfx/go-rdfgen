package schema

type CriticReviewTrait struct {

}

type CriticReview struct {
	MetaTrait
	CriticReviewTrait
	ReviewTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewCriticReview() (x CriticReview) {
	x.Type = "http://schema.org/CriticReview"

	return
}
