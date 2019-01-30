package schema

type AgreeActionTrait struct {

}

type AgreeAction struct {
	MetaTrait
	AgreeActionTrait
	ReactActionTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewAgreeAction() (x AgreeAction) {
	x.Type = "http://schema.org/AgreeAction"

	return
}
