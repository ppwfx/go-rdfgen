package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFastFoodRestaurant strings.Replacer
var strconvFastFoodRestaurant strconv.NumError

var basicFastFoodRestaurantTraitMapping = map[string]func(*schema.FastFoodRestaurantTrait, []string){}
var customFastFoodRestaurantTraitMapping = map[string]func(*schema.FastFoodRestaurantTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FastFoodRestaurant", func(ctx jsonld.Context) (interface{}) {
		return NewFastFoodRestaurantFromContext(ctx)
	})

}

func NewFastFoodRestaurantFromContext(ctx jsonld.Context) (x schema.FastFoodRestaurant) {
	x.Type = "http://schema.org/FastFoodRestaurant"
	MapBasicToFoodEstablishmentTrait(ctx, &x.FoodEstablishmentTrait)
	MapCustomToFoodEstablishmentTrait(ctx, &x.FoodEstablishmentTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToFastFoodRestaurantTrait(ctx, &x.FastFoodRestaurantTrait)
	MapCustomToFastFoodRestaurantTrait(ctx, &x.FastFoodRestaurantTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFastFoodRestaurantTrait(ctx jsonld.Context, x *schema.FastFoodRestaurantTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFastFoodRestaurantTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFastFoodRestaurantTrait(ctx jsonld.Context, x *schema.FastFoodRestaurantTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFastFoodRestaurantTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}