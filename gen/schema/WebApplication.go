package schema

type WebApplicationTrait struct {

	// Specifies browser requirements in human-readable text. For example, 'requires HTML5 support'.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	BrowserRequirements string `json:"browserRequirements,omitempty" jsonld:"http://schema.org/browserRequirements"`

}

type WebApplication struct {
	MetaTrait
	WebApplicationTrait
	SoftwareApplicationTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewWebApplication() (x WebApplication) {
	x.Type = "http://schema.org/WebApplication"

	return
}
