package schema

type DislikeActionTrait struct {

}

type DislikeAction struct {
	MetaTrait
	DislikeActionTrait
	ReactActionTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewDislikeAction() (x DislikeAction) {
	x.Type = "http://schema.org/DislikeAction"

	return
}
