package schema

type WebSiteTrait struct {

	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Issn string `json:"issn,omitempty" jsonld:"http://schema.org/issn"`

}

type WebSite struct {
	MetaTrait
	WebSiteTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewWebSite() (x WebSite) {
	x.Type = "http://schema.org/WebSite"

	return
}
