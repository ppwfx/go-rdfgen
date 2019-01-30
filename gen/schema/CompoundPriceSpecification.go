package schema

type CompoundPriceSpecificationTrait struct {

	// This property links to all <a class=\"localLink\" href=\"http://schema.org/UnitPriceSpecification\">UnitPriceSpecification</a> nodes that apply in parallel for the <a class=\"localLink\" href=\"http://schema.org/CompoundPriceSpecification\">CompoundPriceSpecification</a> node.
	//
	// RangeIncludes:
	// - http://schema.org/UnitPriceSpecification
	//
	PriceComponent *UnitPriceSpecification `json:"priceComponent,omitempty" jsonld:"http://schema.org/priceComponent"`

}

type CompoundPriceSpecification struct {
	MetaTrait
	CompoundPriceSpecificationTrait
	PriceSpecificationTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewCompoundPriceSpecification() (x CompoundPriceSpecification) {
	x.Type = "http://schema.org/CompoundPriceSpecification"

	return
}
