package schema

type ApplyActionTrait struct {

}

type ApplyAction struct {
	MetaTrait
	ApplyActionTrait
	OrganizeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewApplyAction() (x ApplyAction) {
	x.Type = "http://schema.org/ApplyAction"

	return
}
