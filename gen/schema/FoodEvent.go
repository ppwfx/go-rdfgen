package schema

type FoodEventTrait struct {

}

type FoodEvent struct {
	MetaTrait
	FoodEventTrait
	EventTrait
	ThingTrait
	AdditionalTrait
}

func NewFoodEvent() (x FoodEvent) {
	x.Type = "http://schema.org/FoodEvent"

	return
}
