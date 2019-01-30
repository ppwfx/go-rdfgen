package schema

type CookActionTrait struct {

	// A sub property of location. The specific food establishment where the action occurred.
	//
	// RangeIncludes:
	// - http://schema.org/Place
	// - http://schema.org/FoodEstablishment
	//
	FoodEstablishment interface{} `json:"foodEstablishment,omitempty" jsonld:"http://schema.org/foodEstablishment"`

	// A sub property of location. The specific food event where the action occurred.
	//
	// RangeIncludes:
	// - http://schema.org/FoodEvent
	//
	FoodEvent *FoodEvent `json:"foodEvent,omitempty" jsonld:"http://schema.org/foodEvent"`

	// A sub property of instrument. The recipe/instructions used to perform the action.
	//
	// RangeIncludes:
	// - http://schema.org/Recipe
	//
	Recipe *Recipe `json:"recipe,omitempty" jsonld:"http://schema.org/recipe"`

}

type CookAction struct {
	MetaTrait
	CookActionTrait
	CreateActionTrait
	ActionTrait
	ThingTrait
	AdditionalTrait
}

func NewCookAction() (x CookAction) {
	x.Type = "http://schema.org/CookAction"

	return
}
