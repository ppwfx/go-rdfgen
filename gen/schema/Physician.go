package schema

type PhysicianTrait struct {

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

	// A hospital with which the physician or office is affiliated.
	//
	// RangeIncludes:
	// - http://schema.org/Hospital
	//
	HospitalAffiliation *Hospital `json:"hospitalAffiliation,omitempty" jsonld:"http://schema.org/hospitalAffiliation"`

}

type Physician struct {
	MetaTrait
	PhysicianTrait
	MedicalOrganizationTrait
	OrganizationTrait
	ThingTrait
	MedicalBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	AdditionalTrait
}

func NewPhysician() (x Physician) {
	x.Type = "http://schema.org/Physician"

	return
}
