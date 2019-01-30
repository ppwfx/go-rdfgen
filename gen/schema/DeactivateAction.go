package schema

type DeactivateActionTrait struct {

}

type DeactivateAction struct {
	MetaTrait
	DeactivateActionTrait
	ControlActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewDeactivateAction() (x DeactivateAction) {
	x.Type = "http://schema.org/DeactivateAction"

	return
}
