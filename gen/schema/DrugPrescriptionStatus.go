package schema

type DrugPrescriptionStatusTrait struct {

}

type DrugPrescriptionStatus struct {
	MetaTrait
	DrugPrescriptionStatusTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDrugPrescriptionStatus() (x DrugPrescriptionStatus) {
	x.Type = "http://schema.org/DrugPrescriptionStatus"

	return
}
