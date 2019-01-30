package schema

type ImagingTestTrait struct {

	// Imaging technique used.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalImagingTechnique
	//
	ImagingTechnique *MedicalImagingTechnique `json:"imagingTechnique,omitempty" jsonld:"http://schema.org/imagingTechnique"`

}

type ImagingTest struct {
	MetaTrait
	ImagingTestTrait
	MedicalTestTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewImagingTest() (x ImagingTest) {
	x.Type = "http://schema.org/ImagingTest"

	return
}
