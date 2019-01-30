package schema

type AudienceTrait struct {

	// The target group associated with a given audience (e.g. veterans, car owners, musicians, etc.).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AudienceType string `json:"audienceType,omitempty" jsonld:"http://schema.org/audienceType"`

	// The geographic area associated with the audience.
	//
	// RangeIncludes:
	// - http://schema.org/AdministrativeArea
	//
	GeographicArea *AdministrativeArea `json:"geographicArea,omitempty" jsonld:"http://schema.org/geographicArea"`

}

type Audience struct {
	MetaTrait
	AudienceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewAudience() (x Audience) {
	x.Type = "http://schema.org/Audience"

	return
}
