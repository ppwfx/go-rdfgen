package schema

type ViewActionTrait struct {

}

type ViewAction struct {
	MetaTrait
	ViewActionTrait
	ConsumeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewViewAction() (x ViewAction) {
	x.Type = "http://schema.org/ViewAction"

	return
}
