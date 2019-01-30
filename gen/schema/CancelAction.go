package schema

type CancelActionTrait struct {

}

type CancelAction struct {
	MetaTrait
	CancelActionTrait
	PlanActionTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewCancelAction() (x CancelAction) {
	x.Type = "http://schema.org/CancelAction"

	return
}
