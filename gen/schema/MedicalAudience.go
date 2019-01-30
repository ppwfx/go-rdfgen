package schema

type MedicalAudienceTrait struct {

}

type MedicalAudience struct {
	MetaTrait
	MedicalAudienceTrait
	AudienceTrait
	IntangibleTrait
	ThingTrait
	MedicalEnumerationTrait
	EnumerationTrait
	PeopleAudienceTrait
	AdditionalTrait
}

func NewMedicalAudience() (x MedicalAudience) {
	x.Type = "http://schema.org/MedicalAudience"

	return
}
