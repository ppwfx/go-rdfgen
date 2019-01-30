package schema

type AddActionTrait struct {

}

type AddAction struct {
	MetaTrait
	AddActionTrait
	UpdateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewAddAction() (x AddAction) {
	x.Type = "http://schema.org/AddAction"

	return
}
