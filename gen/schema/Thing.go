package schema

type ThingTrait struct {

	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	AdditionalType string `json:"additionalType,omitempty" jsonld:"http://schema.org/additionalType"`

	// An alias for the item.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AlternateName string `json:"alternateName,omitempty" jsonld:"http://schema.org/alternateName"`

	// A description of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Description string `json:"description,omitempty" jsonld:"http://schema.org/description"`

	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	DisambiguatingDescription string `json:"disambiguatingDescription,omitempty" jsonld:"http://schema.org/disambiguatingDescription"`

	// The identifier property represents any kind of identifier for any kind of <a class=\"localLink\" href=\"http://schema.org/Thing\">Thing</a>, such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See <a href=\"/docs/datamodel.html#identifierBg\">background notes</a> for more details.
	//
	// RangeIncludes:
	// - http://schema.org/PropertyValue
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	Identifier interface{} `json:"identifier,omitempty" jsonld:"http://schema.org/identifier"`

	// An image of the item. This can be a <a class=\"localLink\" href=\"http://schema.org/URL\">URL</a> or a fully described <a class=\"localLink\" href=\"http://schema.org/ImageObject\">ImageObject</a>.
	//
	// RangeIncludes:
	// - http://schema.org/ImageObject
	// - http://schema.org/URL
	//
	Image interface{} `json:"image,omitempty" jsonld:"http://schema.org/image"`

	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See <a href=\"/docs/datamodel.html#mainEntityBackground\">background notes</a> for details.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	MainEntityOfPage interface{} `json:"mainEntityOfPage,omitempty" jsonld:"http://schema.org/mainEntityOfPage"`

	// The name of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Name string `json:"name,omitempty" jsonld:"http://schema.org/name"`

	// Indicates a potential Action, which describes an idealized action in which this thing would play an 'object' role.
	//
	// RangeIncludes:
	// - http://schema.org/Action
	//
	PotentialAction *Action `json:"potentialAction,omitempty" jsonld:"http://schema.org/potentialAction"`

	// URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	SameAs string `json:"sameAs,omitempty" jsonld:"http://schema.org/sameAs"`

	// A CreativeWork or Event about this Thing..
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/Event
	//
	SubjectOf interface{} `json:"subjectOf,omitempty" jsonld:"http://schema.org/subjectOf"`

	// URL of the item.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	Url string `json:"url,omitempty" jsonld:"http://schema.org/url"`

}

type Thing struct {
	MetaTrait
	ThingTrait
	AdditionalTrait
}

func NewThing() (x Thing) {
	x.Type = "http://schema.org/Thing"

	return
}
