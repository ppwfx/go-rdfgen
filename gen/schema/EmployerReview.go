package schema

type EmployerReviewTrait struct {

}

type EmployerReview struct {
	MetaTrait
	EmployerReviewTrait
	ReviewTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewEmployerReview() (x EmployerReview) {
	x.Type = "http://schema.org/EmployerReview"

	return
}
