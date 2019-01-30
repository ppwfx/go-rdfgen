package schema

type EngineSpecificationTrait struct {

	// The volume swept by all of the pistons inside the cylinders of an internal combustion engine in a single movement. <br/><br/>\n\nTypical unit code(s): CMQ for cubic centimeter, LTR for liters, INQ for cubic inches\n* Note 1: You can link to information about how the given value has been determined using the <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a> property.\n* Note 2: You can use <a class=\"localLink\" href=\"http://schema.org/minValue\">minValue</a> and <a class=\"localLink\" href=\"http://schema.org/maxValue\">maxValue</a> to indicate ranges.
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	EngineDisplacement *QuantitativeValue `json:"engineDisplacement,omitempty" jsonld:"http://schema.org/engineDisplacement"`

	// The power of the vehicle's engine.\n    Typical unit code(s): KWT for kilowatt, BHP for brake horsepower, N12 for metric horsepower (PS, with 1 PS = 735,49875 W)<br/><br/>\n\n<ul>\n<li>Note 1: There are many different ways of measuring an engine's power. For an overview, see  <a href=\"http://en.wikipedia.org/wiki/Horsepower#Engine_power_test_codes\">http://en.wikipedia.org/wiki/Horsepower#Engine<em>power</em>test_codes</a>.</li>\n<li>Note 2: You can link to information about how the given value has been determined using the <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a> property.</li>\n<li>Note 3: You can use <a class=\"localLink\" href=\"http://schema.org/minValue\">minValue</a> and <a class=\"localLink\" href=\"http://schema.org/maxValue\">maxValue</a> to indicate ranges.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	EnginePower *QuantitativeValue `json:"enginePower,omitempty" jsonld:"http://schema.org/enginePower"`

	// The type of engine or engines powering the vehicle.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	// - http://schema.org/QualitativeValue
	//
	EngineType interface{} `json:"engineType,omitempty" jsonld:"http://schema.org/engineType"`

	// The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	// - http://schema.org/QualitativeValue
	//
	FuelType interface{} `json:"fuelType,omitempty" jsonld:"http://schema.org/fuelType"`

	// The torque (turning force) of the vehicle's engine.<br/><br/>\n\nTypical unit code(s): NU for newton metre (N m), F17 for pound-force per foot, or F48 for pound-force per inch<br/><br/>\n\n<ul>\n<li>Note 1: You can link to information about how the given value has been determined (e.g. reference RPM) using the <a class=\"localLink\" href=\"http://schema.org/valueReference\">valueReference</a> property.</li>\n<li>Note 2: You can use <a class=\"localLink\" href=\"http://schema.org/minValue\">minValue</a> and <a class=\"localLink\" href=\"http://schema.org/maxValue\">maxValue</a> to indicate ranges.</li>\n</ul>\n
	//
	// RangeIncludes:
	// - http://schema.org/QuantitativeValue
	//
	Torque *QuantitativeValue `json:"torque,omitempty" jsonld:"http://schema.org/torque"`

}

type EngineSpecification struct {
	MetaTrait
	EngineSpecificationTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewEngineSpecification() (x EngineSpecification) {
	x.Type = "http://schema.org/EngineSpecification"

	return
}
