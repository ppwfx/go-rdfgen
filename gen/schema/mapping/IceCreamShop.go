package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsIceCreamShop strings.Replacer
var strconvIceCreamShop strconv.NumError

var basicIceCreamShopTraitMapping = map[string]func(*schema.IceCreamShopTrait, []string){}
var customIceCreamShopTraitMapping = map[string]func(*schema.IceCreamShopTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/IceCreamShop", func(ctx jsonld.Context) (interface{}) {
		return NewIceCreamShopFromContext(ctx)
	})

}

func NewIceCreamShopFromContext(ctx jsonld.Context) (x schema.IceCreamShop) {
	x.Type = "http://schema.org/IceCreamShop"
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


	MapBasicToIceCreamShopTrait(ctx, &x.IceCreamShopTrait)
	MapCustomToIceCreamShopTrait(ctx, &x.IceCreamShopTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToIceCreamShopTrait(ctx jsonld.Context, x *schema.IceCreamShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicIceCreamShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToIceCreamShopTrait(ctx jsonld.Context, x *schema.IceCreamShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := customIceCreamShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}