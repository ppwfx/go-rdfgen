package schema

type ActivateActionTrait struct {

}

type ActivateAction struct {
	MetaTrait
	ActivateActionTrait
	ControlActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewActivateAction() (x ActivateAction) {
	x.Type = "http://schema.org/ActivateAction"

	return
}
