package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFoodEvent strings.Replacer
var strconvFoodEvent strconv.NumError

var basicFoodEventTraitMapping = map[string]func(*schema.FoodEventTrait, []string){}
var customFoodEventTraitMapping = map[string]func(*schema.FoodEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FoodEvent", func(ctx jsonld.Context) (interface{}) {
		return NewFoodEventFromContext(ctx)
	})

}

func NewFoodEventFromContext(ctx jsonld.Context) (x schema.FoodEvent) {
	x.Type = "http://schema.org/FoodEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFoodEventTrait(ctx, &x.FoodEventTrait)
	MapCustomToFoodEventTrait(ctx, &x.FoodEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFoodEventTrait(ctx jsonld.Context, x *schema.FoodEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFoodEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFoodEventTrait(ctx jsonld.Context, x *schema.FoodEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFoodEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}