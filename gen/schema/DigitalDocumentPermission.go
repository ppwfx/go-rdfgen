package schema

type DigitalDocumentPermissionTrait struct {

	// The person, organization, contact point, or audience that has been granted this permission.
	//
	// RangeIncludes:
	// - http://schema.org/ContactPoint
	// - http://schema.org/Person
	// - http://schema.org/Organization
	// - http://schema.org/Audience
	//
	Grantee interface{} `json:"grantee,omitempty" jsonld:"http://schema.org/grantee"`

	// The type of permission granted the person, organization, or audience.
	//
	// RangeIncludes:
	// - http://schema.org/DigitalDocumentPermissionType
	//
	PermissionType *DigitalDocumentPermissionType `json:"permissionType,omitempty" jsonld:"http://schema.org/permissionType"`

}

type DigitalDocumentPermission struct {
	MetaTrait
	DigitalDocumentPermissionTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewDigitalDocumentPermission() (x DigitalDocumentPermission) {
	x.Type = "http://schema.org/DigitalDocumentPermission"

	return
}
