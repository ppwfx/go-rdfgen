package schema

type MaximumDoseScheduleTrait struct {

}

type MaximumDoseSchedule struct {
	MetaTrait
	MaximumDoseScheduleTrait
	DoseScheduleTrait
	MedicalIntangibleTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMaximumDoseSchedule() (x MaximumDoseSchedule) {
	x.Type = "http://schema.org/MaximumDoseSchedule"

	return
}
