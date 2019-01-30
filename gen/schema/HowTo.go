package schema

type HowToTrait struct {

	// A sub-property of instrument. A supply consumed when performing instructions or a direction.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/HowToSupply
	//
	Supply interface{} `json:"supply,omitempty" jsonld:"http://schema.org/supply"`

	// A single step item (as HowToStep, text, document, video, etc.) or a HowToSection (originally misnamed 'steps'; 'step' is preferred).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/CreativeWork
	// - http://schema.org/ItemList
	//
	Steps interface{} `json:"steps,omitempty" jsonld:"http://schema.org/steps"`

	// The quantity that results by performing instructions. For example, a paper airplane, 10 personalized candles.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/QuantitativeValue
	//
	Yield interface{} `json:"yield,omitempty" jsonld:"http://schema.org/yield"`

	// The estimated cost of the supply or supplies consumed when performing instructions.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/MonetaryAmount
	//
	EstimatedCost interface{} `json:"estimatedCost,omitempty" jsonld:"http://schema.org/estimatedCost"`

	// A single step item (as HowToStep, text, document, video, etc.) or a HowToSection.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/CreativeWork
	// - http://schema.org/HowToSection
	// - http://schema.org/HowToStep
	//
	Step interface{} `json:"step,omitempty" jsonld:"http://schema.org/step"`

	// A sub property of instrument. An object used (but not consumed) when performing instructions or a direction.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/HowToTool
	//
	Tool interface{} `json:"tool,omitempty" jsonld:"http://schema.org/tool"`

	// The length of time it takes to perform instructions or a direction (not including time to prepare the supplies), in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 duration format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	PerformTime *Duration `json:"performTime,omitempty" jsonld:"http://schema.org/performTime"`

	// The length of time it takes to prepare the items to be used in instructions or a direction, in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 duration format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	PrepTime *Duration `json:"prepTime,omitempty" jsonld:"http://schema.org/prepTime"`

	// The total time required to perform instructions or a direction (including time to prepare the supplies), in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 duration format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	TotalTime *Duration `json:"totalTime,omitempty" jsonld:"http://schema.org/totalTime"`

}

type HowTo struct {
	MetaTrait
	HowToTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewHowTo() (x HowTo) {
	x.Type = "http://schema.org/HowTo"

	return
}
