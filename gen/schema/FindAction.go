package schema

type FindActionTrait struct {

}

type FindAction struct {
	MetaTrait
	FindActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewFindAction() (x FindAction) {
	x.Type = "http://schema.org/FindAction"

	return
}
