package schema

type PhysicalExamTrait struct {

}

type PhysicalExam struct {
	MetaTrait
	PhysicalExamTrait
	MedicalEnumerationTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	MedicalProcedureTrait
	MedicalEntityTrait
	AdditionalTrait
}

func NewPhysicalExam() (x PhysicalExam) {
	x.Type = "http://schema.org/PhysicalExam"

	return
}
