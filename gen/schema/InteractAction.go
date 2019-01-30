package schema

type InteractActionTrait struct {

}

type InteractAction struct {
	MetaTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewInteractAction() (x InteractAction) {
	x.Type = "http://schema.org/InteractAction"

	return
}
