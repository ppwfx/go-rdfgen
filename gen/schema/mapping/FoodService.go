package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFoodService strings.Replacer
var strconvFoodService strconv.NumError

var basicFoodServiceTraitMapping = map[string]func(*schema.FoodServiceTrait, []string){}
var customFoodServiceTraitMapping = map[string]func(*schema.FoodServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FoodService", func(ctx jsonld.Context) (interface{}) {
		return NewFoodServiceFromContext(ctx)
	})

}

func NewFoodServiceFromContext(ctx jsonld.Context) (x schema.FoodService) {
	x.Type = "http://schema.org/FoodService"
	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFoodServiceTrait(ctx, &x.FoodServiceTrait)
	MapCustomToFoodServiceTrait(ctx, &x.FoodServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFoodServiceTrait(ctx jsonld.Context, x *schema.FoodServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFoodServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFoodServiceTrait(ctx jsonld.Context, x *schema.FoodServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFoodServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}