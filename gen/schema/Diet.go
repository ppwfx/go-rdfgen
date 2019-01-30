package schema

type DietTrait struct {

	// Medical expert advice related to the plan.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ExpertConsiderations string `json:"expertConsiderations,omitempty" jsonld:"http://schema.org/expertConsiderations"`

	// Descriptive information establishing the overarching theory/philosophy of the plan. May include the rationale for the name, the population where the plan first came to prominence, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Overview string `json:"overview,omitempty" jsonld:"http://schema.org/overview"`

	// Nutritional information specific to the dietary plan. May include dietary recommendations on what foods to avoid, what foods to consume, and specific alterations/deviations from the USDA or other regulatory body's approved dietary guidelines.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	DietFeatures string `json:"dietFeatures,omitempty" jsonld:"http://schema.org/dietFeatures"`

	// Specific physiologic benefits associated to the plan.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PhysiologicalBenefits string `json:"physiologicalBenefits,omitempty" jsonld:"http://schema.org/physiologicalBenefits"`

	// Specific physiologic risks associated to the diet plan.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Risks string `json:"risks,omitempty" jsonld:"http://schema.org/risks"`

	// People or organizations that endorse the plan.
	//
	// RangeIncludes:
	// - http://schema.org/Person
	// - http://schema.org/Organization
	//
	Endorsers interface{} `json:"endorsers,omitempty" jsonld:"http://schema.org/endorsers"`

}

type Diet struct {
	MetaTrait
	DietTrait
	CreativeWorkTrait
	ThingTrait
	LifestyleModificationTrait
	MedicalEntityTrait
	AdditionalTrait
}

func NewDiet() (x Diet) {
	x.Type = "http://schema.org/Diet"

	return
}
