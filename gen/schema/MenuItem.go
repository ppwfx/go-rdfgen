package schema

type MenuItemTrait struct {

	// An offer to provide this item&#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	//
	// RangeIncludes:
	// - http://schema.org/Offer
	//
	Offers *Offer `json:"offers,omitempty" jsonld:"http://schema.org/offers"`

	// Additional menu item(s) such as a side dish of salad or side order of fries that can be added to this menu item. Additionally it can be a menu section containing allowed add-on menu items for this menu item.
	//
	// RangeIncludes:
	// - http://schema.org/MenuItem
	// - http://schema.org/MenuSection
	//
	MenuAddOn interface{} `json:"menuAddOn,omitempty" jsonld:"http://schema.org/menuAddOn"`

	// Nutrition information about the recipe or menu item.
	//
	// RangeIncludes:
	// - http://schema.org/NutritionInformation
	//
	Nutrition *NutritionInformation `json:"nutrition,omitempty" jsonld:"http://schema.org/nutrition"`

	// Indicates a dietary restriction or guideline for which this recipe or menu item is suitable, e.g. diabetic, halal etc.
	//
	// RangeIncludes:
	// - http://schema.org/RestrictedDiet
	//
	SuitableForDiet *RestrictedDiet `json:"suitableForDiet,omitempty" jsonld:"http://schema.org/suitableForDiet"`

}

type MenuItem struct {
	MetaTrait
	MenuItemTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewMenuItem() (x MenuItem) {
	x.Type = "http://schema.org/MenuItem"

	return
}
