package schema

type CheckActionTrait struct {

}

type CheckAction struct {
	MetaTrait
	CheckActionTrait
	FindActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewCheckAction() (x CheckAction) {
	x.Type = "http://schema.org/CheckAction"

	return
}
