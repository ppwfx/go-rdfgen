package schema

type UserReviewTrait struct {

}

type UserReview struct {
	MetaTrait
	UserReviewTrait
	ReviewTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewUserReview() (x UserReview) {
	x.Type = "http://schema.org/UserReview"

	return
}
