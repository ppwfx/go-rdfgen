package schema

type EatActionTrait struct {

}

type EatAction struct {
	MetaTrait
	EatActionTrait
	ConsumeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewEatAction() (x EatAction) {
	x.Type = "http://schema.org/EatAction"

	return
}
