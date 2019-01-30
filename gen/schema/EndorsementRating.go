package schema

type EndorsementRatingTrait struct {

}

type EndorsementRating struct {
	MetaTrait
	EndorsementRatingTrait
	RatingTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewEndorsementRating() (x EndorsementRating) {
	x.Type = "http://schema.org/EndorsementRating"

	return
}
