package schema

type MedicalWebPageTrait struct {

	// An aspect of medical practice that is considered on the page, such as 'diagnosis', 'treatment', 'causes', 'prognosis', 'etiology', 'epidemiology', etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Aspect string `json:"aspect,omitempty" jsonld:"http://schema.org/aspect"`

}

type MedicalWebPage struct {
	MetaTrait
	MedicalWebPageTrait
	WebPageTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewMedicalWebPage() (x MedicalWebPage) {
	x.Type = "http://schema.org/MedicalWebPage"

	return
}
