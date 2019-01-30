package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCookAction strings.Replacer
var strconvCookAction strconv.NumError

var basicCookActionTraitMapping = map[string]func(*schema.CookActionTrait, []string){}
var customCookActionTraitMapping = map[string]func(*schema.CookActionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CookAction", func(ctx jsonld.Context) (interface{}) {
		return NewCookActionFromContext(ctx)
	})





	customCookActionTraitMapping["http://schema.org/foodEstablishment"] = func(x *schema.CookActionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.FoodEstablishment, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.FoodEstablishment = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.FoodEstablishment = s
		}
	}

	customCookActionTraitMapping["http://schema.org/foodEvent"] = func(x *schema.CookActionTrait, ctx jsonld.Context, s string){
		var y schema.FoodEvent
		if strings.HasPrefix(s, "_:") {
			y = NewFoodEventFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewFoodEvent()
			y.Id = s
		}

		x.FoodEvent = &y

		return
	}

	customCookActionTraitMapping["http://schema.org/recipe"] = func(x *schema.CookActionTrait, ctx jsonld.Context, s string){
		var y schema.Recipe
		if strings.HasPrefix(s, "_:") {
			y = NewRecipeFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewRecipe()
			y.Id = s
		}

		x.Recipe = &y

		return
	}
}

func NewCookActionFromContext(ctx jsonld.Context) (x schema.CookAction) {
	x.Type = "http://schema.org/CookAction"
	MapBasicToCreateActionTrait(ctx, &x.CreateActionTrait)
	MapCustomToCreateActionTrait(ctx, &x.CreateActionTrait)

	MapBasicToActionTrait(ctx, &x.ActionTrait)
	MapCustomToActionTrait(ctx, &x.ActionTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCookActionTrait(ctx, &x.CookActionTrait)
	MapCustomToCookActionTrait(ctx, &x.CookActionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCookActionTrait(ctx jsonld.Context, x *schema.CookActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCookActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCookActionTrait(ctx jsonld.Context, x *schema.CookActionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCookActionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}