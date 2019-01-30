package schema

type TouristAttractionTrait struct {

	// A language someone may use with or at the item, service or place. Please use one of the language codes from the <a href=\"http://tools.ietf.org/html/bcp47\">IETF BCP 47 standard</a>. See also <a class=\"localLink\" href=\"http://schema.org/inLanguage\">inLanguage</a>
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Language
	//
	AvailableLanguage interface{} `json:"availableLanguage,omitempty" jsonld:"http://schema.org/availableLanguage"`

	// Attraction suitable for type(s) of tourist. eg. Children, visitors from a particular country, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Audience
	//
	TouristType interface{} `json:"touristType,omitempty" jsonld:"http://schema.org/touristType"`

}

type TouristAttraction struct {
	MetaTrait
	TouristAttractionTrait
	PlaceTrait
	ThingTrait
	AdditionalTrait
}

func NewTouristAttraction() (x TouristAttraction) {
	x.Type = "http://schema.org/TouristAttraction"

	return
}
