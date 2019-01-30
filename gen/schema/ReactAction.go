package schema

type ReactActionTrait struct {

}

type ReactAction struct {
	MetaTrait
	ReactActionTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewReactAction() (x ReactAction) {
	x.Type = "http://schema.org/ReactAction"

	return
}
