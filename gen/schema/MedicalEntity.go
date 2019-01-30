package schema

type MedicalEntityTrait struct {

	// A medical code for the entity, taken from a controlled vocabulary or ontology such as ICD-9, DiseasesDB, MeSH, SNOMED-CT, RxNorm, etc.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalCode
	//
	Code *MedicalCode `json:"code,omitempty" jsonld:"http://schema.org/code"`

	// A medical guideline related to this entity.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalGuideline
	//
	Guideline *MedicalGuideline `json:"guideline,omitempty" jsonld:"http://schema.org/guideline"`

	// The drug or supplement's legal status, including any controlled substance schedules that apply.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/DrugLegalStatus
	// - http://schema.org/MedicalEnumeration
	//
	LegalStatus interface{} `json:"legalStatus,omitempty" jsonld:"http://schema.org/legalStatus"`

	// The system of medicine that includes this MedicalEntity, for example 'evidence-based', 'homeopathic', 'chiropractic', etc.
	//
	// RangeIncludes:
	// - http://schema.org/MedicineSystem
	//
	MedicineSystem *MedicineSystem `json:"medicineSystem,omitempty" jsonld:"http://schema.org/medicineSystem"`

	// If applicable, the organization that officially recognizes this entity as part of its endorsed system of medicine.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	RecognizingAuthority *Organization `json:"recognizingAuthority,omitempty" jsonld:"http://schema.org/recognizingAuthority"`

	// If applicable, a medical specialty in which this entity is relevant.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalSpecialty
	//
	RelevantSpecialty *MedicalSpecialty `json:"relevantSpecialty,omitempty" jsonld:"http://schema.org/relevantSpecialty"`

	// A medical study or trial related to this entity.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalStudy
	//
	Study *MedicalStudy `json:"study,omitempty" jsonld:"http://schema.org/study"`

}

type MedicalEntity struct {
	MetaTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalEntity() (x MedicalEntity) {
	x.Type = "http://schema.org/MedicalEntity"

	return
}
