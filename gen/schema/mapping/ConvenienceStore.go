package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsConvenienceStore strings.Replacer
var strconvConvenienceStore strconv.NumError

var basicConvenienceStoreTraitMapping = map[string]func(*schema.ConvenienceStoreTrait, []string){}
var customConvenienceStoreTraitMapping = map[string]func(*schema.ConvenienceStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ConvenienceStore", func(ctx jsonld.Context) (interface{}) {
		return NewConvenienceStoreFromContext(ctx)
	})

}

func NewConvenienceStoreFromContext(ctx jsonld.Context) (x schema.ConvenienceStore) {
	x.Type = "http://schema.org/ConvenienceStore"
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


	MapBasicToConvenienceStoreTrait(ctx, &x.ConvenienceStoreTrait)
	MapCustomToConvenienceStoreTrait(ctx, &x.ConvenienceStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToConvenienceStoreTrait(ctx jsonld.Context, x *schema.ConvenienceStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicConvenienceStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToConvenienceStoreTrait(ctx jsonld.Context, x *schema.ConvenienceStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customConvenienceStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}