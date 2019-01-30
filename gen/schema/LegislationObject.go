package schema

type LegislationObjectTrait struct {

	// The legal value of this legislation file. The same legislation can be written in multiple files with different legal values. Typically a digitally signed PDF have a \"stronger\" legal value than the HTML file of the same act.
	//
	// RangeIncludes:
	// - http://schema.org/LegalValueLevel
	//
	LegislationLegalValue *LegalValueLevel `json:"legislationLegalValue,omitempty" jsonld:"http://schema.org/legislationLegalValue"`

}

type LegislationObject struct {
	MetaTrait
	LegislationObjectTrait
	MediaObjectTrait
	CreativeWorkTrait
	ThingTrait
	LegislationTrait
	AdditionalTrait
}

func NewLegislationObject() (x LegislationObject) {
	x.Type = "http://schema.org/LegislationObject"

	return
}
