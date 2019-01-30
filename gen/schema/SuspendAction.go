package schema

type SuspendActionTrait struct {

}

type SuspendAction struct {
	MetaTrait
	SuspendActionTrait
	ControlActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewSuspendAction() (x SuspendAction) {
	x.Type = "http://schema.org/SuspendAction"

	return
}
