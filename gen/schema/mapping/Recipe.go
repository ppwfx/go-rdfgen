package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRecipe strings.Replacer
var strconvRecipe strconv.NumError

var basicRecipeTraitMapping = map[string]func(*schema.RecipeTrait, []string){}
var customRecipeTraitMapping = map[string]func(*schema.RecipeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Recipe", func(ctx jsonld.Context) (interface{}) {
		return NewRecipeFromContext(ctx)
	})




	basicRecipeTraitMapping["http://schema.org/cookingMethod"] = func(x *schema.RecipeTrait, s []string) {
		x.CookingMethod = s[0]
	}


	basicRecipeTraitMapping["http://schema.org/recipeIngredient"] = func(x *schema.RecipeTrait, s []string) {
		x.RecipeIngredient = s[0]
	}



	basicRecipeTraitMapping["http://schema.org/recipeCuisine"] = func(x *schema.RecipeTrait, s []string) {
		x.RecipeCuisine = s[0]
	}


	basicRecipeTraitMapping["http://schema.org/recipeCategory"] = func(x *schema.RecipeTrait, s []string) {
		x.RecipeCategory = s[0]
	}



	basicRecipeTraitMapping["http://schema.org/ingredients"] = func(x *schema.RecipeTrait, s []string) {
		x.Ingredients = s[0]
	}



	customRecipeTraitMapping["http://schema.org/nutrition"] = func(x *schema.RecipeTrait, ctx jsonld.Context, s string){
		var y schema.NutritionInformation
		if strings.HasPrefix(s, "_:") {
			y = NewNutritionInformationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewNutritionInformation()
			y.Id = s
		}

		x.Nutrition = &y

		return
	}

	customRecipeTraitMapping["http://schema.org/suitableForDiet"] = func(x *schema.RecipeTrait, ctx jsonld.Context, s string){
		var y schema.RestrictedDiet
		if strings.HasPrefix(s, "_:") {
			y = NewRestrictedDietFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewRestrictedDiet()
			y.Id = s
		}

		x.SuitableForDiet = &y

		return
	}

	customRecipeTraitMapping["http://schema.org/recipeInstructions"] = func(x *schema.RecipeTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RecipeInstructions, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RecipeInstructions = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RecipeInstructions = s
		}
	}

	customRecipeTraitMapping["http://schema.org/recipeYield"] = func(x *schema.RecipeTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RecipeYield, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RecipeYield = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RecipeYield = s
		}
	}

	customRecipeTraitMapping["http://schema.org/cookTime"] = func(x *schema.RecipeTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.CookTime = &y

		return
	}
}

func NewRecipeFromContext(ctx jsonld.Context) (x schema.Recipe) {
	x.Type = "http://schema.org/Recipe"
	MapBasicToHowToTrait(ctx, &x.HowToTrait)
	MapCustomToHowToTrait(ctx, &x.HowToTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRecipeTrait(ctx, &x.RecipeTrait)
	MapCustomToRecipeTrait(ctx, &x.RecipeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRecipeTrait(ctx jsonld.Context, x *schema.RecipeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRecipeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRecipeTrait(ctx jsonld.Context, x *schema.RecipeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRecipeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}