package schema

type CheckOutActionTrait struct {

}

type CheckOutAction struct {
	MetaTrait
	CheckOutActionTrait
	CommunicateActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewCheckOutAction() (x CheckOutAction) {
	x.Type = "http://schema.org/CheckOutAction"

	return
}
