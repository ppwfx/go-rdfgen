package schema

type PathologyTestTrait struct {

	// The type of tissue sample required for the test.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	TissueSample string `json:"tissueSample,omitempty" jsonld:"http://schema.org/tissueSample"`

}

type PathologyTest struct {
	MetaTrait
	PathologyTestTrait
	MedicalTestTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewPathologyTest() (x PathologyTest) {
	x.Type = "http://schema.org/PathologyTest"

	return
}
