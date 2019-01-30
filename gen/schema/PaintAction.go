package schema

type PaintActionTrait struct {

}

type PaintAction struct {
	MetaTrait
	PaintActionTrait
	CreateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewPaintAction() (x PaintAction) {
	x.Type = "http://schema.org/PaintAction"

	return
}
