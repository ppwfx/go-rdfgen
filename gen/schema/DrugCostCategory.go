package schema

type DrugCostCategoryTrait struct {

}

type DrugCostCategory struct {
	MetaTrait
	DrugCostCategoryTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDrugCostCategory() (x DrugCostCategory) {
	x.Type = "http://schema.org/DrugCostCategory"

	return
}
