package schema

type BoardingPolicyTypeTrait struct {

}

type BoardingPolicyType struct {
	MetaTrait
	BoardingPolicyTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBoardingPolicyType() (x BoardingPolicyType) {
	x.Type = "http://schema.org/BoardingPolicyType"

	return
}
