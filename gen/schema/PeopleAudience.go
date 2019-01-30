package schema

type PeopleAudienceTrait struct {

	// Specifying the health condition(s) of a patient, medical study, or other target audience.
	//
	// RangeIncludes:
	// - http://schema.org/MedicalCondition
	//
	HealthCondition *MedicalCondition `json:"healthCondition,omitempty" jsonld:"http://schema.org/healthCondition"`

	// Audiences defined by a person's gender.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	RequiredGender string `json:"requiredGender,omitempty" jsonld:"http://schema.org/requiredGender"`

	// The gender of the person or audience.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	SuggestedGender string `json:"suggestedGender,omitempty" jsonld:"http://schema.org/suggestedGender"`

	// Audiences defined by a person's maximum age.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	RequiredMaxAge *Integer `json:"requiredMaxAge,omitempty" jsonld:"http://schema.org/requiredMaxAge"`

	// Audiences defined by a person's minimum age.
	//
	// RangeIncludes:
	// - http://schema.org/Integer
	//
	RequiredMinAge *Integer `json:"requiredMinAge,omitempty" jsonld:"http://schema.org/requiredMinAge"`

	// Maximal age recommended for viewing content.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	SuggestedMaxAge float64 `json:"suggestedMaxAge,omitempty" jsonld:"http://schema.org/suggestedMaxAge"`

	// Minimal age recommended for viewing content.
	//
	// RangeIncludes:
	// - http://schema.org/Number
	//
	SuggestedMinAge float64 `json:"suggestedMinAge,omitempty" jsonld:"http://schema.org/suggestedMinAge"`

}

type PeopleAudience struct {
	MetaTrait
	PeopleAudienceTrait
	AudienceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPeopleAudience() (x PeopleAudience) {
	x.Type = "http://schema.org/PeopleAudience"

	return
}
