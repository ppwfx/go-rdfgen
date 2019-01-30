package schema

type DeliveryChargeSpecificationTrait struct {

	// The geographic area where a service or offered item is provided.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/AdministrativeArea
	// - http://schema.org/GeoShape
	// - http://schema.org/Place
	//
	AreaServed interface{} `json:"areaServed,omitempty" jsonld:"http://schema.org/areaServed"`

	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.<br/><br/>\n\nSee also <a class=\"localLink\" href=\"http://schema.org/ineligibleRegion\">ineligibleRegion</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/GeoShape
	// - http://schema.org/Place
	//
	EligibleRegion interface{} `json:"eligibleRegion,omitempty" jsonld:"http://schema.org/eligibleRegion"`

	// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.<br/><br/>\n\nSee also <a class=\"localLink\" href=\"http://schema.org/eligibleRegion\">eligibleRegion</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/GeoShape
	// - http://schema.org/Place
	//
	IneligibleRegion interface{} `json:"ineligibleRegion,omitempty" jsonld:"http://schema.org/ineligibleRegion"`

	// The delivery method(s) to which the delivery charge or payment charge specification applies.
	//
	// RangeIncludes:
	// - http://schema.org/DeliveryMethod
	//
	AppliesToDeliveryMethod *DeliveryMethod `json:"appliesToDeliveryMethod,omitempty" jsonld:"http://schema.org/appliesToDeliveryMethod"`

}

type DeliveryChargeSpecification struct {
	MetaTrait
	DeliveryChargeSpecificationTrait
	PriceSpecificationTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDeliveryChargeSpecification() (x DeliveryChargeSpecification) {
	x.Type = "http://schema.org/DeliveryChargeSpecification"

	return
}
