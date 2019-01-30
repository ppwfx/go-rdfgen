package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPawnShop strings.Replacer
var strconvPawnShop strconv.NumError

var basicPawnShopTraitMapping = map[string]func(*schema.PawnShopTrait, []string){}
var customPawnShopTraitMapping = map[string]func(*schema.PawnShopTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PawnShop", func(ctx jsonld.Context) (interface{}) {
		return NewPawnShopFromContext(ctx)
	})

}

func NewPawnShopFromContext(ctx jsonld.Context) (x schema.PawnShop) {
	x.Type = "http://schema.org/PawnShop"
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


	MapBasicToPawnShopTrait(ctx, &x.PawnShopTrait)
	MapCustomToPawnShopTrait(ctx, &x.PawnShopTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPawnShopTrait(ctx jsonld.Context, x *schema.PawnShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPawnShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPawnShopTrait(ctx jsonld.Context, x *schema.PawnShopTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPawnShopTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}