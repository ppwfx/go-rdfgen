package schema

type PostalAddressTrait struct {

	// The country. For example, USA. You can also provide the two-letter <a href=\"http://en.wikipedia.org/wiki/ISO_3166-1\">ISO 3166-1 alpha-2 country code</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Country
	//
	AddressCountry interface{} `json:"addressCountry,omitempty" jsonld:"http://schema.org/addressCountry"`

	// The postal code. For example, 94043.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PostalCode string `json:"postalCode,omitempty" jsonld:"http://schema.org/postalCode"`

	// The locality. For example, Mountain View.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AddressLocality string `json:"addressLocality,omitempty" jsonld:"http://schema.org/addressLocality"`

	// The region. For example, CA.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	AddressRegion string `json:"addressRegion,omitempty" jsonld:"http://schema.org/addressRegion"`

	// The post office box number for PO box addresses.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PostOfficeBoxNumber string `json:"postOfficeBoxNumber,omitempty" jsonld:"http://schema.org/postOfficeBoxNumber"`

	// The street address. For example, 1600 Amphitheatre Pkwy.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	StreetAddress string `json:"streetAddress,omitempty" jsonld:"http://schema.org/streetAddress"`

}

type PostalAddress struct {
	MetaTrait
	PostalAddressTrait
	ContactPointTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewPostalAddress() (x PostalAddress) {
	x.Type = "http://schema.org/PostalAddress"

	return
}
