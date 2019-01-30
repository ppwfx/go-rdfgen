package schema

type SubscribeActionTrait struct {

}

type SubscribeAction struct {
	MetaTrait
	SubscribeActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewSubscribeAction() (x SubscribeAction) {
	x.Type = "http://schema.org/SubscribeAction"

	return
}
