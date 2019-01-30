package schema

type DigitalDocumentTrait struct {

	// A permission related to the access to this document (e.g. permission to read or write an electronic document). For a public document, specify a grantee with an Audience with audienceType equal to \"public\".
	//
	// RangeIncludes:
	// - http://schema.org/DigitalDocumentPermission
	//
	HasDigitalDocumentPermission *DigitalDocumentPermission `json:"hasDigitalDocumentPermission,omitempty" jsonld:"http://schema.org/hasDigitalDocumentPermission"`

}

type DigitalDocument struct {
	MetaTrait
	DigitalDocumentTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewDigitalDocument() (x DigitalDocument) {
	x.Type = "http://schema.org/DigitalDocument"

	return
}
