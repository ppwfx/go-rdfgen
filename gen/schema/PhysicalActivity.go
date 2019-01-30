package schema

type PhysicalActivityTrait struct {

	// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Thing
	// - http://schema.org/PhysicalActivityCategory
	//
	Category interface{} `json:"category,omitempty" jsonld:"http://schema.org/category"`

	// Changes in the normal mechanical, physical, and biochemical functions that are associated with this activity or condition.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Pathophysiology string `json:"pathophysiology,omitempty" jsonld:"http://schema.org/pathophysiology"`

	// The characteristics of associated patients, such as age, gender, race etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Epidemiology string `json:"epidemiology,omitempty" jsonld:"http://schema.org/epidemiology"`

	// The anatomy of the underlying organ system or structures associated with this entity.
	//
	// RangeIncludes:
	// - http://schema.org/AnatomicalStructure
	// - http://schema.org/AnatomicalSystem
	// - http://schema.org/SuperficialAnatomy
	//
	AssociatedAnatomy interface{} `json:"associatedAnatomy,omitempty" jsonld:"http://schema.org/associatedAnatomy"`

}

type PhysicalActivity struct {
	MetaTrait
	PhysicalActivityTrait
	LifestyleModificationTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewPhysicalActivity() (x PhysicalActivity) {
	x.Type = "http://schema.org/PhysicalActivity"

	return
}
