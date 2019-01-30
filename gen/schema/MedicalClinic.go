package schema

type MedicalClinicTrait struct {

	// A medical specialty of the provider.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalSpecialty
	//
	MedicalSpecialty *MedicalSpecialty `json:"medicalSpecialty,omitempty" jsonld:"http://schema.org/medicalSpecialty"`

	// A medical service available from this provider.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalProcedure
	// - http://schema.org/MedicalTest
	// - http://schema.org/MedicalTherapy
	//
	AvailableService interface{} `json:"availableService,omitempty" jsonld:"http://schema.org/availableService"`

}

type MedicalClinic struct {
	MetaTrait
	MedicalClinicTrait
	MedicalOrganizationTrait
	OrganizationTrait
	ThingTrait
	MedicalBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	AdditionalTrait
}

func NewMedicalClinic() (x MedicalClinic) {
	x.Type = "http://schema.org/MedicalClinic"

	return
}
