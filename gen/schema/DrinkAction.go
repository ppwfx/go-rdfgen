package schema

type DrinkActionTrait struct {

}

type DrinkAction struct {
	MetaTrait
	DrinkActionTrait
	ConsumeActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewDrinkAction() (x DrinkAction) {
	x.Type = "http://schema.org/DrinkAction"

	return
}
