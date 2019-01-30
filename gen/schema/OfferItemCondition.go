package schema

type OfferItemConditionTrait struct {

}

type OfferItemCondition struct {
	MetaTrait
	OfferItemConditionTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewOfferItemCondition() (x OfferItemCondition) {
	x.Type = "http://schema.org/OfferItemCondition"

	return
}
