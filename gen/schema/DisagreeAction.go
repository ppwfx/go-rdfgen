package schema

type DisagreeActionTrait struct {

}

type DisagreeAction struct {
	MetaTrait
	DisagreeActionTrait
	ReactActionTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewDisagreeAction() (x DisagreeAction) {
	x.Type = "http://schema.org/DisagreeAction"

	return
}
