package schema

type BloodTestTrait struct {

}

type BloodTest struct {
	MetaTrait
	BloodTestTrait
	MedicalTestTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewBloodTest() (x BloodTest) {
	x.Type = "http://schema.org/BloodTest"

	return
}
