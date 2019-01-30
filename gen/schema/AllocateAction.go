package schema

type AllocateActionTrait struct {

	// A goal towards an action is taken. Can be concrete or abstract.
	//
	// RangeIncludes:
	// - http://schema.org/Thing
	// - http://schema.org/MedicalDevicePurpose
	//
	Purpose interface{} `json:"purpose,omitempty" jsonld:"http://schema.org/purpose"`

}

type AllocateAction struct {
	MetaTrait
	AllocateActionTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewAllocateAction() (x AllocateAction) {
	x.Type = "http://schema.org/AllocateAction"

	return
}
