package schema

type QuantitativeValueDistributionTrait struct {

	// The duration of the item (movie, audio recording, event, etc.) in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 date format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	Duration *Duration `json:"duration,omitempty" jsonld:"http://schema.org/duration"`

	// The median value.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	Median float64 `json:"median,omitempty" jsonld:"http://schema.org/median"`

	// The 10th percentile value.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	Percentile10 float64 `json:"percentile10,omitempty" jsonld:"http://schema.org/percentile10"`

	// The 25th percentile value.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	Percentile25 float64 `json:"percentile25,omitempty" jsonld:"http://schema.org/percentile25"`

	// The 75th percentile value.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	Percentile75 float64 `json:"percentile75,omitempty" jsonld:"http://schema.org/percentile75"`

	// The 90th percentile value.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	Percentile90 float64 `json:"percentile90,omitempty" jsonld:"http://schema.org/percentile90"`

}

type QuantitativeValueDistribution struct {
	MetaTrait
	QuantitativeValueDistributionTrait
	QuantitativeValueTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewQuantitativeValueDistribution() (x QuantitativeValueDistribution) {
	x.Type = "http://schema.org/QuantitativeValueDistribution"

	return
}
