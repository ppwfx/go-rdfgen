package schema

type NutritionInformationTrait struct {

	// The number of calories.
	//
	// RangeIncludes:
	// - http://schema.org/Energy
	//
	Calories *Energy `json:"calories,omitempty" jsonld:"http://schema.org/calories"`

	// The number of grams of carbohydrates.
	//
	// RangeIncludes:
	// - http://schema.org/Mass
	//
	CarbohydrateContent *Mass `json:"carbohydrateContent,omitempty" jsonld:"http://schema.org/carbohydrateContent"`

	// The number of milligrams of cholesterol.
	//
	// RangeIncludes:
	// - http://schema.org/Mass
	//
	CholesterolContent *Mass `json:"cholesterolContent,omitempty" jsonld:"http://schema.org/cholesterolContent"`

	// The number of grams of fat.
	//
	// RangeIncludes:
	// - http://schema.org/Mass
	//
	FatContent *Mass `json:"fatContent,omitempty" jsonld:"http://schema.org/fatContent"`

	// The number of grams of fiber.
	//
	// RangeIncludes:
	// - http://schema.org/Mass
	//
	FiberContent *Mass `json:"fiberContent,omitempty" jsonld:"http://schema.org/fiberContent"`

	// The number of grams of protein.
	//
	// RangeIncludes:
	// - http://schema.org/Mass
	//
	ProteinContent *Mass `json:"proteinContent,omitempty" jsonld:"http://schema.org/proteinContent"`

	// The number of grams of saturated fat.
	//
	// RangeIncludes:
	// - http://schema.org/Mass
	//
	SaturatedFatContent *Mass `json:"saturatedFatContent,omitempty" jsonld:"http://schema.org/saturatedFatContent"`

	// The serving size, in terms of the number of volume or mass.
	//
	// RangeIncludes:
	// - http://schema.org/Text
	//
	ServingSize string `json:"servingSize,omitempty" jsonld:"http://schema.org/servingSize"`

	// The number of milligrams of sodium.
	//
	// RangeIncludes:
	// - http://schema.org/Mass
	//
	SodiumContent *Mass `json:"sodiumContent,omitempty" jsonld:"http://schema.org/sodiumContent"`

	// The number of grams of sugar.
	//
	// RangeIncludes:
	// - http://schema.org/Mass
	//
	SugarContent *Mass `json:"sugarContent,omitempty" jsonld:"http://schema.org/sugarContent"`

	// The number of grams of trans fat.
	//
	// RangeIncludes:
	// - http://schema.org/Mass
	//
	TransFatContent *Mass `json:"transFatContent,omitempty" jsonld:"http://schema.org/transFatContent"`

	// The number of grams of unsaturated fat.
	//
	// RangeIncludes:
	// - http://schema.org/Mass
	//
	UnsaturatedFatContent *Mass `json:"unsaturatedFatContent,omitempty" jsonld:"http://schema.org/unsaturatedFatContent"`

}

type NutritionInformation struct {
	MetaTrait
	NutritionInformationTrait
	StructuredValueTrait
	IntangibleTrait
	ThingTrait
	AdditionalTrait
}

func NewNutritionInformation() (x NutritionInformation) {
	x.Type = "http://schema.org/NutritionInformation"

	return
}
