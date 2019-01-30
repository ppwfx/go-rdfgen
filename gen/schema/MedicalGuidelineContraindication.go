package schema

type MedicalGuidelineContraindicationTrait struct {

}

type MedicalGuidelineContraindication struct {
	MetaTrait
	MedicalGuidelineContraindicationTrait
	MedicalGuidelineTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalGuidelineContraindication() (x MedicalGuidelineContraindication) {
	x.Type = "http://schema.org/MedicalGuidelineContraindication"

	return
}
