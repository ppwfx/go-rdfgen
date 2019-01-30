package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWholesaleStore strings.Replacer
var strconvWholesaleStore strconv.NumError

var basicWholesaleStoreTraitMapping = map[string]func(*schema.WholesaleStoreTrait, []string){}
var customWholesaleStoreTraitMapping = map[string]func(*schema.WholesaleStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WholesaleStore", func(ctx jsonld.Context) (interface{}) {
		return NewWholesaleStoreFromContext(ctx)
	})

}

func NewWholesaleStoreFromContext(ctx jsonld.Context) (x schema.WholesaleStore) {
	x.Type = "http://schema.org/WholesaleStore"
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


	MapBasicToWholesaleStoreTrait(ctx, &x.WholesaleStoreTrait)
	MapCustomToWholesaleStoreTrait(ctx, &x.WholesaleStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWholesaleStoreTrait(ctx jsonld.Context, x *schema.WholesaleStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWholesaleStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWholesaleStoreTrait(ctx jsonld.Context, x *schema.WholesaleStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWholesaleStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}