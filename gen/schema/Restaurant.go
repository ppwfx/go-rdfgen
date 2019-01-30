package schema

type RestaurantTrait struct {

}

type Restaurant struct {
	MetaTrait
	RestaurantTrait
	FoodEstablishmentTrait
	LocalBusinessTrait
	PlaceTrait
	ThingTrait
	OrganizationTrait
	AdditionalTrait
}

func NewRestaurant() (x Restaurant) {
	x.Type = "http://schema.org/Restaurant"

	return
}
