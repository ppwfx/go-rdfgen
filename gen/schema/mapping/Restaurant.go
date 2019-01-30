package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRestaurant strings.Replacer
var strconvRestaurant strconv.NumError

var basicRestaurantTraitMapping = map[string]func(*schema.RestaurantTrait, []string){}
var customRestaurantTraitMapping = map[string]func(*schema.RestaurantTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Restaurant", func(ctx jsonld.Context) (interface{}) {
		return NewRestaurantFromContext(ctx)
	})

}

func NewRestaurantFromContext(ctx jsonld.Context) (x schema.Restaurant) {
	x.Type = "http://schema.org/Restaurant"
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


	MapBasicToRestaurantTrait(ctx, &x.RestaurantTrait)
	MapCustomToRestaurantTrait(ctx, &x.RestaurantTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRestaurantTrait(ctx jsonld.Context, x *schema.RestaurantTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRestaurantTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRestaurantTrait(ctx jsonld.Context, x *schema.RestaurantTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRestaurantTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}