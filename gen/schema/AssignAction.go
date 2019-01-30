package schema

type AssignActionTrait struct {

}

type AssignAction struct {
	MetaTrait
	AssignActionTrait
	AllocateActionTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewAssignAction() (x AssignAction) {
	x.Type = "http://schema.org/AssignAction"

	return
}
