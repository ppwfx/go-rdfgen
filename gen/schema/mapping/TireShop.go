package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTireShop strings.Replacer
var strconvTireShop strconv.NumError

var basicTireShopTraitMapping = map[string]func(*schema.TireShopTrait, []string){}
var customTireShopTraitMapping = map[string]func(*schema.TireShopTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TireShop", func(ctx jsonld.Context) (interface{}) {
		return NewTireShopFromContext(ctx)
	})

}

func NewTireShopFromContext(ctx jsonld.Context) (x schema.TireShop) {
	x.Type = "http://schema.org/TireShop"
	MapBasicToStoreTrait(ctx, &x.StoreTrait)
	MapCustomToStoreTrait(ctx, &x.StoreTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToTireShopTrait(ctx, &x.TireShopTrait)
	MapCustomToTireShopTrait(ctx, &x.TireShopTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTireShopTrait(ctx jsonld.Context, x *schema.TireShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTireShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTireShopTrait(ctx jsonld.Context, x *schema.TireShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTireShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}