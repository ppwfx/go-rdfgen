package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsShoppingCenter strings.Replacer
var strconvShoppingCenter strconv.NumError

var basicShoppingCenterTraitMapping = map[string]func(*schema.ShoppingCenterTrait, []string){}
var customShoppingCenterTraitMapping = map[string]func(*schema.ShoppingCenterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ShoppingCenter", func(ctx jsonld.Context) (interface{}) {
		return NewShoppingCenterFromContext(ctx)
	})

}

func NewShoppingCenterFromContext(ctx jsonld.Context) (x schema.ShoppingCenter) {
	x.Type = "http://schema.org/ShoppingCenter"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToShoppingCenterTrait(ctx, &x.ShoppingCenterTrait)
	MapCustomToShoppingCenterTrait(ctx, &x.ShoppingCenterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToShoppingCenterTrait(ctx jsonld.Context, x *schema.ShoppingCenterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicShoppingCenterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToShoppingCenterTrait(ctx jsonld.Context, x *schema.ShoppingCenterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customShoppingCenterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}