package schema

type AcceptActionTrait struct {

}

type AcceptAction struct {
	MetaTrait
	AcceptActionTrait
	AllocateActionTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewAcceptAction() (x AcceptAction) {
	x.Type = "http://schema.org/AcceptAction"

	return
}
