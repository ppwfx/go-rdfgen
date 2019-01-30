package schema

type LodgingBusinessTrait struct {

	// An intended audience, i.e. a group for whom something was created.
	//
	// RangeIncludes:
	// - http://schema.org/Audience
	//
	Audience *Audience `json:"audience,omitempty" jsonld:"http://schema.org/audience"`

	// A language someone may use with or at the item, service or place. Please use one of the language codes from the <a href=\"http://tools.ietf.org/html/bcp47\">IETF BCP 47 standard</a>. See also <a class=\"localLink\" href=\"http://schema.org/inLanguage\">inLanguage</a>
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Language
	//
	AvailableLanguage interface{} `json:"availableLanguage,omitempty" jsonld:"http://schema.org/availableLanguage"`

	// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	//
	// RangeIncludes:
	// - http://schema.org/LocationFeatureSpecification
	//
	AmenityFeature *LocationFeatureSpecification `json:"amenityFeature,omitempty" jsonld:"http://schema.org/amenityFeature"`

	// The earliest someone may check into a lodging establishment.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	CheckinTime *DateTime `json:"checkinTime,omitempty" jsonld:"http://schema.org/checkinTime"`

	// The latest someone may check out of a lodging establishment.
	//
	// RangeIncludes:
	// - http://schema.org/DateTime
	//
	CheckoutTime *DateTime `json:"checkoutTime,omitempty" jsonld:"http://schema.org/checkoutTime"`

	// Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Boolean
	//
	PetsAllowed interface{} `json:"petsAllowed,omitempty" jsonld:"http://schema.org/petsAllowed"`

	// An official rating for a lodging business or food establishment, e.g. from national associations or standards bodies. Use the author property to indicate the rating organization, e.g. as an Organization with name such as (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
	//
	// RangeIncludes:
	// - http://schema.org/Rating
	//
	StarRating *Rating `json:"starRating,omitempty" jsonld:"http://schema.org/starRating"`

}

type LodgingBusiness struct {
	MetaTrait
	LodgingBusinessTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewLodgingBusiness() (x LodgingBusiness) {
	x.Type = "http://schema.org/LodgingBusiness"

	return
}
