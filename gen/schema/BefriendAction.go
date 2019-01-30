package schema

type BefriendActionTrait struct {

}

type BefriendAction struct {
	MetaTrait
	BefriendActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewBefriendAction() (x BefriendAction) {
	x.Type = "http://schema.org/BefriendAction"

	return
}
