package schema

type SteeringPositionValueTrait struct {

}

type SteeringPositionValue struct {
	MetaTrait
	SteeringPositionValueTrait
	QualitativeValueTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewSteeringPositionValue() (x SteeringPositionValue) {
	x.Type = "http://schema.org/SteeringPositionValue"

	return
}
