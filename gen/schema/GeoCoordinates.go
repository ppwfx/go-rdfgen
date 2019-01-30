package schema

type GeoCoordinatesTrait struct {

	// Physical address of the item.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/PostalAddress
	//
	Address interface{} `json:"address,omitempty" jsonld:"http://schema.org/address"`

	// The country. For example, USA. You can also provide the two-letter <a href=\"http://en.wikipedia.org/wiki/ISO_3166-1\">ISO 3166-1 alpha-2 country code</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Country
	//
	AddressCountry interface{} `json:"addressCountry,omitempty" jsonld:"http://schema.org/addressCountry"`

	// The elevation of a location (<a href=\"https://en.wikipedia.org/wiki/World_Geodetic_System\">WGS 84</a>).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	Elevation interface{} `json:"elevation,omitempty" jsonld:"http://schema.org/elevation"`

	// The latitude of a location. For example <code>37.42242</code> (<a href=\"https://en.wikipedia.org/wiki/World_Geodetic_System\">WGS 84</a>).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	Latitude interface{} `json:"latitude,omitempty" jsonld:"http://schema.org/latitude"`

	// The longitude of a location. For example <code>-122.08585</code> (<a href=\"https://en.wikipedia.org/wiki/World_Geodetic_System\">WGS 84</a>).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Number
	//
	Longitude interface{} `json:"longitude,omitempty" jsonld:"http://schema.org/longitude"`

	// The postal code. For example, 94043.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	PostalCode string `json:"postalCode,omitempty" jsonld:"http://schema.org/postalCode"`

}

type GeoCoordinates struct {
	MetaTrait
	GeoCoordinatesTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewGeoCoordinates() (x GeoCoordinates) {
	x.Type = "http://schema.org/GeoCoordinates"

	return
}
