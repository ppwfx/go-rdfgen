package schema

type DrugLegalStatusTrait struct {

	// The location in which the status applies.
	//
	// RangeIncludes:
	// - http://schema.org/AdministrativeArea
	//
	ApplicableLocation *AdministrativeArea `json:"applicableLocation,omitempty" jsonld:"http://schema.org/applicableLocation"`

}

type DrugLegalStatus struct {
	MetaTrait
	DrugLegalStatusTrait
	MedicalIntangibleTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewDrugLegalStatus() (x DrugLegalStatus) {
	x.Type = "http://schema.org/DrugLegalStatus"

	return
}
