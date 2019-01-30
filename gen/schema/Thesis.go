package schema

type ThesisTrait struct {

	// Qualification, candidature, degree, application that Thesis supports.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	InSupportOf string `json:"inSupportOf,omitempty" jsonld:"http://schema.org/inSupportOf"`

}

type Thesis struct {
	MetaTrait
	ThesisTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewThesis() (x Thesis) {
	x.Type = "http://schema.org/Thesis"

	return
}
