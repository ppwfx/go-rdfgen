package schema

type EntryPointTrait struct {

	// An application that can complete the request.
	//
	// RangeIncludes:
	// - http://schema.org/SoftwareApplication
	//
	ActionApplication *SoftwareApplication `json:"actionApplication,omitempty" jsonld:"http://schema.org/actionApplication"`

	// The high level platform(s) where the Action can be performed for the given URL. To specify a specific application or operating system instance, use actionApplication.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/URL
	//
	ActionPlatform interface{} `json:"actionPlatform,omitempty" jsonld:"http://schema.org/actionPlatform"`

	// An application that can complete the request.
	//
	// RangeIncludes:
	// - http://schema.org/SoftwareApplication
	//
	Application *SoftwareApplication `json:"application,omitempty" jsonld:"http://schema.org/application"`

	// The supported content type(s) for an EntryPoint response.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ContentType string `json:"contentType,omitempty" jsonld:"http://schema.org/contentType"`

	// The supported encoding type(s) for an EntryPoint request.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	EncodingType string `json:"encodingType,omitempty" jsonld:"http://schema.org/encodingType"`

	// An HTTP method that specifies the appropriate HTTP method for a request to an HTTP EntryPoint. Values are capitalized strings as used in HTTP.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	HttpMethod string `json:"httpMethod,omitempty" jsonld:"http://schema.org/httpMethod"`

	// An url template (RFC6570) that will be used to construct the target of the execution of the action.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	UrlTemplate string `json:"urlTemplate,omitempty" jsonld:"http://schema.org/urlTemplate"`

}

type EntryPoint struct {
	MetaTrait
	EntryPointTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewEntryPoint() (x EntryPoint) {
	x.Type = "http://schema.org/EntryPoint"

	return
}
