package schema

type UnRegisterActionTrait struct {

}

type UnRegisterAction struct {
	MetaTrait
	UnRegisterActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewUnRegisterAction() (x UnRegisterAction) {
	x.Type = "http://schema.org/UnRegisterAction"

	return
}
