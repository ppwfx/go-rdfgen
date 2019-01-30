package schema

type FoodServiceTrait struct {

}

type FoodService struct {
	MetaTrait
	FoodServiceTrait
	ServiceTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewFoodService() (x FoodService) {
	x.Type = "http://schema.org/FoodService"

	return
}
