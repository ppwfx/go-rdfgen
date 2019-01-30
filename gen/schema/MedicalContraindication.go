package schema

type MedicalContraindicationTrait struct {

}

type MedicalContraindication struct {
	MetaTrait
	MedicalContraindicationTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalContraindication() (x MedicalContraindication) {
	x.Type = "http://schema.org/MedicalContraindication"

	return
}
