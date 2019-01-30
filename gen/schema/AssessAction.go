package schema

type AssessActionTrait struct {

}

type AssessAction struct {
	MetaTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewAssessAction() (x AssessAction) {
	x.Type = "http://schema.org/AssessAction"

	return
}
