package schema

type MarryActionTrait struct {

}

type MarryAction struct {
	MetaTrait
	MarryActionTrait
	InteractActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewMarryAction() (x MarryAction) {
	x.Type = "http://schema.org/MarryAction"

	return
}
