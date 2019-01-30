package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHomeGoodsStore strings.Replacer
var strconvHomeGoodsStore strconv.NumError

var basicHomeGoodsStoreTraitMapping = map[string]func(*schema.HomeGoodsStoreTrait, []string){}
var customHomeGoodsStoreTraitMapping = map[string]func(*schema.HomeGoodsStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HomeGoodsStore", func(ctx jsonld.Context) (interface{}) {
		return NewHomeGoodsStoreFromContext(ctx)
	})

}

func NewHomeGoodsStoreFromContext(ctx jsonld.Context) (x schema.HomeGoodsStore) {
	x.Type = "http://schema.org/HomeGoodsStore"
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


	MapBasicToHomeGoodsStoreTrait(ctx, &x.HomeGoodsStoreTrait)
	MapCustomToHomeGoodsStoreTrait(ctx, &x.HomeGoodsStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHomeGoodsStoreTrait(ctx jsonld.Context, x *schema.HomeGoodsStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHomeGoodsStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHomeGoodsStoreTrait(ctx jsonld.Context, x *schema.HomeGoodsStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHomeGoodsStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}