package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCafeOrCoffeeShop strings.Replacer
var strconvCafeOrCoffeeShop strconv.NumError

var basicCafeOrCoffeeShopTraitMapping = map[string]func(*schema.CafeOrCoffeeShopTrait, []string){}
var customCafeOrCoffeeShopTraitMapping = map[string]func(*schema.CafeOrCoffeeShopTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CafeOrCoffeeShop", func(ctx jsonld.Context) (interface{}) {
		return NewCafeOrCoffeeShopFromContext(ctx)
	})

}

func NewCafeOrCoffeeShopFromContext(ctx jsonld.Context) (x schema.CafeOrCoffeeShop) {
	x.Type = "http://schema.org/CafeOrCoffeeShop"
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


	MapBasicToCafeOrCoffeeShopTrait(ctx, &x.CafeOrCoffeeShopTrait)
	MapCustomToCafeOrCoffeeShopTrait(ctx, &x.CafeOrCoffeeShopTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCafeOrCoffeeShopTrait(ctx jsonld.Context, x *schema.CafeOrCoffeeShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCafeOrCoffeeShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCafeOrCoffeeShopTrait(ctx jsonld.Context, x *schema.CafeOrCoffeeShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCafeOrCoffeeShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}