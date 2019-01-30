package schema

type EducationalAudienceTrait struct {

	// An educationalRole of an EducationalAudience.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	EducationalRole string `json:"educationalRole,omitempty" jsonld:"http://schema.org/educationalRole"`

}

type EducationalAudience struct {
	MetaTrait
	EducationalAudienceTrait
	AudienceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewEducationalAudience() (x EducationalAudience) {
	x.Type = "http://schema.org/EducationalAudience"

	return
}
