package schema

type FastFoodRestaurantTrait struct {

}

type FastFoodRestaurant struct {
	MetaTrait
	FastFoodRestaurantTrait
	FoodEstablishmentTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewFastFoodRestaurant() (x FastFoodRestaurant) {
	x.Type = "http://schema.org/FastFoodRestaurant"

	return
}
