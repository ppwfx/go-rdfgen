package schema

type DrugTrait struct {

	// The manufacturer of the product.
	//
	// RangeIncludes:
	// - http://schema.org/Organization
	//
	Manufacturer *Organization `json:"manufacturer,omitempty" jsonld:"http://schema.org/manufacturer"`

	// The drug or supplement's legal status, including any controlled substance schedules that apply.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/DrugLegalStatus
	// - http://schema.org/MedicalEnumeration
	//
	LegalStatus interface{} `json:"legalStatus,omitempty" jsonld:"http://schema.org/legalStatus"`

	// Indicates the status of drug prescription eg. local catalogs classifications or whether the drug is available by prescription or over-the-counter, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/DrugPrescriptionStatus
	//
	PrescriptionStatus interface{} `json:"prescriptionStatus,omitempty" jsonld:"http://schema.org/prescriptionStatus"`

	// An active ingredient, typically chemical compounds and/or biologic substances.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ActiveIngredient string `json:"activeIngredient,omitempty" jsonld:"http://schema.org/activeIngredient"`

	// A dosage form in which this drug/supplement is available, e.g. 'tablet', 'suspension', 'injection'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	DosageForm string `json:"dosageForm,omitempty" jsonld:"http://schema.org/dosageForm"`

	// Description of the absorption and elimination of drugs, including their concentration (pharmacokinetics, pK) and biological effects (pharmacodynamics, pD).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ClincalPharmacology string `json:"clincalPharmacology,omitempty" jsonld:"http://schema.org/clincalPharmacology"`

	// Description of the absorption and elimination of drugs, including their concentration (pharmacokinetics, pK) and biological effects (pharmacodynamics, pD).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ClinicalPharmacology string `json:"clinicalPharmacology,omitempty" jsonld:"http://schema.org/clinicalPharmacology"`

	// Any information related to overdose on a drug, including signs or symptoms, treatments, contact information for emergency response.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Overdosage string `json:"overdosage,omitempty" jsonld:"http://schema.org/overdosage"`

	// Any precaution, guidance, contraindication, etc. related to this drug's use during pregnancy.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PregnancyWarning string `json:"pregnancyWarning,omitempty" jsonld:"http://schema.org/pregnancyWarning"`

	// The unit in which the drug is measured, e.g. '5 mg tablet'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	DrugUnit string `json:"drugUnit,omitempty" jsonld:"http://schema.org/drugUnit"`

	// Any precaution, guidance, contraindication, etc. related to consumption of specific foods while taking this drug.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	FoodWarning string `json:"foodWarning,omitempty" jsonld:"http://schema.org/foodWarning"`

	// The generic name of this drug or supplement.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	NonProprietaryName string `json:"nonProprietaryName,omitempty" jsonld:"http://schema.org/nonProprietaryName"`

	// Any precaution, guidance, contraindication, etc. related to consumption of alcohol while taking this drug.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AlcoholWarning string `json:"alcoholWarning,omitempty" jsonld:"http://schema.org/alcoholWarning"`

	// The specific biochemical interaction through which this drug or supplement produces its pharmacological effect.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	MechanismOfAction string `json:"mechanismOfAction,omitempty" jsonld:"http://schema.org/mechanismOfAction"`

	// Any precaution, guidance, contraindication, etc. related to this drug's use by breastfeeding mothers.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BreastfeedingWarning string `json:"breastfeedingWarning,omitempty" jsonld:"http://schema.org/breastfeedingWarning"`

	// A route by which this drug may be administered, e.g. 'oral'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AdministrationRoute string `json:"administrationRoute,omitempty" jsonld:"http://schema.org/administrationRoute"`

	// The RxCUI drug identifier from RXNORM.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Rxcui string `json:"rxcui,omitempty" jsonld:"http://schema.org/rxcui"`

	// Any FDA or other warnings about the drug (text or URL).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Warning interface{} `json:"warning,omitempty" jsonld:"http://schema.org/warning"`

	// Proprietary name given to the diet plan, typically by its originator or creator.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ProprietaryName string `json:"proprietaryName,omitempty" jsonld:"http://schema.org/proprietaryName"`

	// Link to prescribing information for the drug.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	PrescribingInfo string `json:"prescribingInfo,omitempty" jsonld:"http://schema.org/prescribingInfo"`

	// Link to the drug's label details.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	LabelDetails string `json:"labelDetails,omitempty" jsonld:"http://schema.org/labelDetails"`

	// True if the drug is available in a generic form (regardless of name).
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	IsAvailableGenerically bool `json:"isAvailableGenerically,omitempty" jsonld:"http://schema.org/isAvailableGenerically"`

	// True if this item's name is a proprietary/brand name (vs. generic name).
	//
	// RangeIncludes:
	// - http://schema.org/Boolean
	//
	IsProprietary bool `json:"isProprietary,omitempty" jsonld:"http://schema.org/isProprietary"`

	// Recommended intake of this supplement for a given population as defined by a specific recommending authority.
	//
	// RangeIncludes:
	// - http://schema.org/MaximumDoseSchedule
	//
	MaximumIntake *MaximumDoseSchedule `json:"maximumIntake,omitempty" jsonld:"http://schema.org/maximumIntake"`

	// An available dosage strength for the drug.
	//
	// RangeIncludes:
	// - http://schema.org/DrugStrength
	//
	AvailableStrength *DrugStrength `json:"availableStrength,omitempty" jsonld:"http://schema.org/availableStrength"`

	// Cost per unit of the drug, as reported by the source being tagged.
	//
	// RangeIncludes:
	// - http://schema.org/DrugCost
	//
	Cost *DrugCost `json:"cost,omitempty" jsonld:"http://schema.org/cost"`

	// A dosing schedule for the drug for a given population, either observed, recommended, or maximum dose based on the type used.
	//
	// RangeIncludes:
	// - http://schema.org/DoseSchedule
	//
	DoseSchedule *DoseSchedule `json:"doseSchedule,omitempty" jsonld:"http://schema.org/doseSchedule"`

	// The class of drug this belongs to (e.g., statins).
	//
	// RangeIncludes:
	// - http://schema.org/DrugClass
	//
	DrugClass *DrugClass `json:"drugClass,omitempty" jsonld:"http://schema.org/drugClass"`

	// The insurance plans that cover this drug.
	//
	// RangeIncludes:
	// - http://schema.org/HealthInsurancePlan
	//
	IncludedInHealthInsurancePlan *HealthInsurancePlan `json:"includedInHealthInsurancePlan,omitempty" jsonld:"http://schema.org/includedInHealthInsurancePlan"`

	// Another drug that is known to interact with this drug in a way that impacts the effect of this drug or causes a risk to the patient. Note: disease interactions are typically captured as contraindications.
	//
	// RangeIncludes:
	// - http://schema.org/Drug
	//
	InteractingDrug *Drug `json:"interactingDrug,omitempty" jsonld:"http://schema.org/interactingDrug"`

	// Pregnancy category of this drug.
	//
	// RangeIncludes:
	// - http://schema.org/DrugPregnancyCategory
	//
	PregnancyCategory *DrugPregnancyCategory `json:"pregnancyCategory,omitempty" jsonld:"http://schema.org/pregnancyCategory"`

	// Any other drug related to this one, for example commonly-prescribed alternatives.
	//
	// RangeIncludes:
	// - http://schema.org/Drug
	//
	RelatedDrug *Drug `json:"relatedDrug,omitempty" jsonld:"http://schema.org/relatedDrug"`

}

type Drug struct {
	MetaTrait
	DrugTrait
	SubstanceTrait
	MedicalEntityTrait
	ThingTrait
	AdditionalTrait
}

func NewDrug() (x Drug) {
	x.Type = "http://schema.org/Drug"

	return
}
