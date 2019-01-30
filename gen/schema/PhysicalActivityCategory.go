package schema

type PhysicalActivityCategoryTrait struct {

}

type PhysicalActivityCategory struct {
	MetaTrait
	PhysicalActivityCategoryTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPhysicalActivityCategory() (x PhysicalActivityCategory) {
	x.Type = "http://schema.org/PhysicalActivityCategory"

	return
}
