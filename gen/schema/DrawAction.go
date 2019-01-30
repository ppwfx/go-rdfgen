package schema

type DrawActionTrait struct {

}

type DrawAction struct {
	MetaTrait
	DrawActionTrait
	CreateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewDrawAction() (x DrawAction) {
	x.Type = "http://schema.org/DrawAction"

	return
}
