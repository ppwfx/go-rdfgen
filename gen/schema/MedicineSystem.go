package schema

type MedicineSystemTrait struct {

}

type MedicineSystem struct {
	MetaTrait
	MedicineSystemTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicineSystem() (x MedicineSystem) {
	x.Type = "http://schema.org/MedicineSystem"

	return
}
