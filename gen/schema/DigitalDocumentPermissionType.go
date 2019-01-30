package schema

type DigitalDocumentPermissionTypeTrait struct {

}

type DigitalDocumentPermissionType struct {
	MetaTrait
	DigitalDocumentPermissionTypeTrait
	EnumerationTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDigitalDocumentPermissionType() (x DigitalDocumentPermissionType) {
	x.Type = "http://schema.org/DigitalDocumentPermissionType"

	return
}
