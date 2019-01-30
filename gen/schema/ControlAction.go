package schema

type ControlActionTrait struct {

}

type ControlAction struct {
	MetaTrait
	ControlActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewControlAction() (x ControlAction) {
	x.Type = "http://schema.org/ControlAction"

	return
}
