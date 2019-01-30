package schema

type ServiceChannelTrait struct {

	// A language someone may use with or at the item, service or place. Please use one of the language codes from the <a href=\"http://tools.ietf.org/html/bcp47\">IETF BCP 47 standard</a>. See also <a class=\"localLink\" href=\"http://schema.org/inLanguage\">inLanguage</a>
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/Language
	//
	AvailableLanguage interface{} `json:"availableLanguage,omitempty" jsonld:"http://schema.org/availableLanguage"`

	// Estimated processing time for the service using this channel.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	ProcessingTime *Duration `json:"processingTime,omitempty" jsonld:"http://schema.org/processingTime"`

	// The service provided by this channel.
	//
	// RangeIncludes:
	// - http://schema.org/Service
	//
	ProvidesService *Service `json:"providesService,omitempty" jsonld:"http://schema.org/providesService"`

	// The location (e.g. civic structure, local business, etc.) where a person can go to access the service.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	//
	ServiceLocation *Place `json:"serviceLocation,omitempty" jsonld:"http://schema.org/serviceLocation"`

	// The phone number to use to access the service.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	//
	ServicePhone *ContactPoint `json:"servicePhone,omitempty" jsonld:"http://schema.org/servicePhone"`

	// The address for accessing the service by mail.
	//
	// RangeIncludes:
	// - http://schema.org/PostalAddress
	//
	ServicePostalAddress *PostalAddress `json:"servicePostalAddress,omitempty" jsonld:"http://schema.org/servicePostalAddress"`

	// The number to access the service by text message.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	//
	ServiceSmsNumber *ContactPoint `json:"serviceSmsNumber,omitempty" jsonld:"http://schema.org/serviceSmsNumber"`

	// The website to access the service.
	//
	// RangeIncludes:
	// - http://schema.org/URL
	//
	ServiceUrl string `json:"serviceUrl,omitempty" jsonld:"http://schema.org/serviceUrl"`

}

type ServiceChannel struct {
	MetaTrait
	ServiceChannelTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewServiceChannel() (x ServiceChannel) {
	x.Type = "http://schema.org/ServiceChannel"

	return
}
