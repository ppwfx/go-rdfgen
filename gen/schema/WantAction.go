package schema

type WantActionTrait struct {

}

type WantAction struct {
	MetaTrait
	WantActionTrait
	ReactActionTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewWantAction() (x WantAction) {
	x.Type = "http://schema.org/WantAction"

	return
}
