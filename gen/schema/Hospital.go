package schema

type HospitalTrait struct {

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

type Hospital struct {
	MetaTrait
	HospitalTrait
	CivicStructureTrait
	PlaceTrait
	ThingTrait
	MedicalOrganizationTrait
	OrganizationTrait
	EmergencyServiceTrait
	LocalBusinessTrait
	AdditionalTrait
}

func NewHospital() (x Hospital) {
	x.Type = "http://schema.org/Hospital"

	return
}
