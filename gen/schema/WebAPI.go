package schema

type WebAPITrait struct {

	// Further documentation describing the Web API in more detail.
	//
	// RangeIncludes:
	// - http://schema.org/CreativeWork
	// - http://schema.org/URL
	//
	Documentation interface{} `json:"documentation,omitempty" jsonld:"http://schema.org/documentation"`

}

type WebAPI struct {
	MetaTrait
	WebAPITrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewWebAPI() (x WebAPI) {
	x.Type = "http://schema.org/WebAPI"

	return
}
