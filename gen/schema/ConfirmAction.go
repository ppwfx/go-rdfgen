package schema

type ConfirmActionTrait struct {

}

type ConfirmAction struct {
	MetaTrait
	ConfirmActionTrait
	InformActionTrait
	CommunicateActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewConfirmAction() (x ConfirmAction) {
	x.Type = "http://schema.org/ConfirmAction"

	return
}
