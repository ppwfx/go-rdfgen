package schema

type BedTypeTrait struct {

}

type BedType struct {
	MetaTrait
	BedTypeTrait
	QualitativeValueTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewBedType() (x BedType) {
	x.Type = "http://schema.org/BedType"

	return
}
