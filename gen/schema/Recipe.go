package schema

type RecipeTrait struct {

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

	// The method of cooking, such as Frying, Steaming, ...
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	CookingMethod string `json:"cookingMethod,omitempty" jsonld:"http://schema.org/cookingMethod"`

	// A single ingredient used in the recipe, e.g. sugar, flour or garlic.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	RecipeIngredient string `json:"recipeIngredient,omitempty" jsonld:"http://schema.org/recipeIngredient"`

	// A step in making the recipe, in the form of a single item (document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/CreativeWork
	// - http://schema.org/ItemList
	//
	RecipeInstructions interface{} `json:"recipeInstructions,omitempty" jsonld:"http://schema.org/recipeInstructions"`

	// The cuisine of the recipe (for example, French or Ethiopian).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	RecipeCuisine string `json:"recipeCuisine,omitempty" jsonld:"http://schema.org/recipeCuisine"`

	// The category of the recipe—for example, appetizer, entree, etc.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	RecipeCategory string `json:"recipeCategory,omitempty" jsonld:"http://schema.org/recipeCategory"`

	// The quantity produced by the recipe (for example, number of people served, number of servings, etc).
	//
	// RangeIncludes:
	// - http://schema.org/Text
	// - http://schema.org/QuantitativeValue
	//
	RecipeYield interface{} `json:"recipeYield,omitempty" jsonld:"http://schema.org/recipeYield"`

	// A single ingredient used in the recipe, e.g. sugar, flour or garlic.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	Ingredients string `json:"ingredients,omitempty" jsonld:"http://schema.org/ingredients"`

	// The time it takes to actually cook the dish, in <a href=\"http://en.wikipedia.org/wiki/ISO_8601\">ISO 8601 duration format</a>.
	//
	// RangeIncludes:
	// - http://schema.org/Duration
	//
	CookTime *Duration `json:"cookTime,omitempty" jsonld:"http://schema.org/cookTime"`

}

type Recipe struct {
	MetaTrait
	RecipeTrait
	HowToTrait
	CreativeWorkTrait
	ThingTrait
	AdditionalTrait
}

func NewRecipe() (x Recipe) {
	x.Type = "http://schema.org/Recipe"

	return
}
