package schema

type CarUsageTypeTrait struct {

}

type CarUsageType struct {
	MetaTrait
	CarUsageTypeTrait
	QualitativeValueTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewCarUsageType() (x CarUsageType) {
	x.Type = "http://schema.org/CarUsageType"

	return
}
