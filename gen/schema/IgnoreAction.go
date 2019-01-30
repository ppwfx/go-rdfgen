package schema

type IgnoreActionTrait struct {

}

type IgnoreAction struct {
	MetaTrait
	IgnoreActionTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewIgnoreAction() (x IgnoreAction) {
	x.Type = "http://schema.org/IgnoreAction"

	return
}
