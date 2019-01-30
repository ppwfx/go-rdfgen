package schema

type ContactPointTrait struct {

	// The hours during which this service or contact is available.
	//
	// RangeIncludes:
	// - http://schema.org/OpeningHoursSpecification
	//
	HoursAvailable *OpeningHoursSpecification `json:"hoursAvailable,omitempty" jsonld:"http://schema.org/hoursAvailable"`

	// The geographic area where a service or offered item is provided.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/AdministrativeArea
	// - http://schema.org/GeoShape
	// - http://schema.org/Place
	//
	AreaServed interface{} `json:"areaServed,omitempty" jsonld:"http://schema.org/areaServed"`

	// A language someone may use with or at the item, service or place. Please use one of the language codes from the <a href=\"http://tools.ietf.org/html/bcp47\">IETF BCP 47 standard</a>. See also <a class=\"localLink\" href=\"http://schema.org/inLanguage\">inLanguage</a>
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Language
	//
	AvailableLanguage interface{} `json:"availableLanguage,omitempty" jsonld:"http://schema.org/availableLanguage"`

	// An option available on this contact point (e.g. a toll-free number or support for hearing-impaired callers).
	//
	// RangeIncludes:
	// - http://schema.org/ContactPointOption
	//
	ContactOption *ContactPointOption `json:"contactOption,omitempty" jsonld:"http://schema.org/contactOption"`

	// A person or organization can have different contact points, for different purposes. For example, a sales contact point, a PR contact point and so on. This property is used to specify the kind of contact point.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ContactType string `json:"contactType,omitempty" jsonld:"http://schema.org/contactType"`

	// Email address.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Email string `json:"email,omitempty" jsonld:"http://schema.org/email"`

	// The fax number.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	FaxNumber string `json:"faxNumber,omitempty" jsonld:"http://schema.org/faxNumber"`

	// The product or service this support contact point is related to (such as product support for a particular product line). This can be a specific product or product line (e.g. \"iPhone\") or a general category of products or services (e.g. \"smartphones\").
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Product
	//
	ProductSupported interface{} `json:"productSupported,omitempty" jsonld:"http://schema.org/productSupported"`

	// The geographic area where the service is provided.
	//
	// RangeIncludes:
	// - http://schema.org/GeoShape
	// - http://schema.org/AdministrativeArea
	// - http://schema.org/Place
	//
	ServiceArea interface{} `json:"serviceArea,omitempty" jsonld:"http://schema.org/serviceArea"`

	// The telephone number.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Telephone string `json:"telephone,omitempty" jsonld:"http://schema.org/telephone"`

}

type ContactPoint struct {
	MetaTrait
	ContactPointTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewContactPoint() (x ContactPoint) {
	x.Type = "http://schema.org/ContactPoint"

	return
}
