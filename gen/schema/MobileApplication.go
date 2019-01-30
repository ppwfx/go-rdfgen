package schema

type MobileApplicationTrait struct {

	// Specifies specific carrier(s) requirements for the application (e.g. an application may only work on a specific carrier network).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CarrierRequirements string `json:"carrierRequirements,omitempty" jsonld:"http://schema.org/carrierRequirements"`

}

type MobileApplication struct {
	MetaTrait
	MobileApplicationTrait
	SoftwareApplicationTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMobileApplication() (x MobileApplication) {
	x.Type = "http://schema.org/MobileApplication"

	return
}
