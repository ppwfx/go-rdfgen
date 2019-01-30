package schema

type RegisterActionTrait struct {

}

type RegisterAction struct {
	MetaTrait
	RegisterActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewRegisterAction() (x RegisterAction) {
	x.Type = "http://schema.org/RegisterAction"

	return
}
