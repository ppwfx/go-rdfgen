package schema

type RejectActionTrait struct {

}

type RejectAction struct {
	MetaTrait
	RejectActionTrait
	AllocateActionTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewRejectAction() (x RejectAction) {
	x.Type = "http://schema.org/RejectAction"

	return
}
