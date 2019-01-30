package schema

type CarTrait struct {

	// The ACRISS Car Classification Code is a code used by many car rental companies, for classifying vehicles. ACRISS stands for Association of Car Rental Industry Systems and Standards.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AcrissCode string `json:"acrissCode,omitempty" jsonld:"http://schema.org/acrissCode"`

	// The permitted total weight of cargo and installations (e.g. a roof rack) on top of the vehicle.<br/><br/>\n\nTypical unit code(s): KGM for kilogram, LBR for pound<br/><br/>\n\n<ul>\n<li>Note 1: You can indicate additional information in the <a class=\"localLink\" href=\"http://schema.org/name\">name</a> of the <a class=\"localLink\" href=\"http://schema.org/QuantitativeValue\">QuantitativeValue</a> node.</li>\n<li>Note 2: You may also link to a <a class=\"localLink\" href=\"http://schema.org/QualitativeValue\">QualitativeValue</a> node that provides additional information using <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a></li>\n<li>Note 3: Note that you can use <a class=\"localLink\" href=\"http://schema.org/minValue\">minValue</a> and <a class=\"localLink\" href=\"http://schema.org/maxValue\">maxValue</a> to indicate ranges.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	RoofLoad *QuantitativeValue `json:"roofLoad,omitempty" jsonld:"http://schema.org/roofLoad"`

}

type Car struct {
	MetaTrait
	CarTrait
	VehicleTrait
	ProductTrait
	ThingTrait
	AdditionalTrait
}

func NewCar() (x Car) {
	x.Type = "http://schema.org/Car"

	return
}
