package schema

type MedicalOrganizationTrait struct {

	// Name or unique ID of network. (Networks are often reused across different insurance plans).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HealthPlanNetworkId string `json:"healthPlanNetworkId,omitempty" jsonld:"http://schema.org/healthPlanNetworkId"`

	// Whether the provider is accepting new patients.
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	IsAcceptingNewPatients bool `json:"isAcceptingNewPatients,omitempty" jsonld:"http://schema.org/isAcceptingNewPatients"`

	// A medical specialty of the provider.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalSpecialty
	//
	MedicalSpecialty *MedicalSpecialty `json:"medicalSpecialty,omitempty" jsonld:"http://schema.org/medicalSpecialty"`

}

type MedicalOrganization struct {
	MetaTrait
	MedicalOrganizationTrait
	OrganizationTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalOrganization() (x MedicalOrganization) {
	x.Type = "http://schema.org/MedicalOrganization"

	return
}
