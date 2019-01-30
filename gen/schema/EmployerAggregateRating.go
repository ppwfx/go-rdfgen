package schema

type EmployerAggregateRatingTrait struct {

}

type EmployerAggregateRating struct {
	MetaTrait
	EmployerAggregateRatingTrait
	AggregateRatingTrait
	RatingTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewEmployerAggregateRating() (x EmployerAggregateRating) {
	x.Type = "http://schema.org/EmployerAggregateRating"

	return
}
