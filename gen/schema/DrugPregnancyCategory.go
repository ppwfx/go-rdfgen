package schema

type DrugPregnancyCategoryTrait struct {

}

type DrugPregnancyCategory struct {
	MetaTrait
	DrugPregnancyCategoryTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDrugPregnancyCategory() (x DrugPregnancyCategory) {
	x.Type = "http://schema.org/DrugPregnancyCategory"

	return
}
