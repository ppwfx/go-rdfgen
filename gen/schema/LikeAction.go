package schema

type LikeActionTrait struct {

}

type LikeAction struct {
	MetaTrait
	LikeActionTrait
	ReactActionTrait
	AssessActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewLikeAction() (x LikeAction) {
	x.Type = "http://schema.org/LikeAction"

	return
}
