package schema

type ReportedDoseScheduleTrait struct {

}

type ReportedDoseSchedule struct {
	MetaTrait
	ReportedDoseScheduleTrait
	DoseScheduleTrait
	MedicalIntangibleTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewReportedDoseSchedule() (x ReportedDoseSchedule) {
	x.Type = "http://schema.org/ReportedDoseSchedule"

	return
}
