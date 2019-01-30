package schema

type MedicalTestTrait struct {

	// Range of acceptable values for a typical patient, when applicable.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/MedicalEnumeration
	//
	NormalRange interface{} `json:"normalRange,omitempty" jsonld:"http://schema.org/normalRange"`

	// Drugs that affect the test's results.
	//
	// RangeIncludes:
	// - http://schema.org/Drug
	//
	AffectedBy *Drug `json:"affectedBy,omitempty" jsonld:"http://schema.org/affectedBy"`

	// A sign detected by the test.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalSign
	//
	SignDetected *MedicalSign `json:"signDetected,omitempty" jsonld:"http://schema.org/signDetected"`

	// A condition the test is used to diagnose.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalCondition
	//
	UsedToDiagnose *MedicalCondition `json:"usedToDiagnose,omitempty" jsonld:"http://schema.org/usedToDiagnose"`

	// Device used to perform the test.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalDevice
	//
	UsesDevice *MedicalDevice `json:"usesDevice,omitempty" jsonld:"http://schema.org/usesDevice"`

}

type MedicalTest struct {
	MetaTrait
	MedicalTestTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalTest() (x MedicalTest) {
	x.Type = "http://schema.org/MedicalTest"

	return
}
