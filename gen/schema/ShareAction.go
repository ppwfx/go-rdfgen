package schema

type ShareActionTrait struct {

}

type ShareAction struct {
	MetaTrait
	ShareActionTrait
	CommunicateActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewShareAction() (x ShareAction) {
	x.Type = "http://schema.org/ShareAction"

	return
}
