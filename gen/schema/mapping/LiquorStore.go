package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLiquorStore strings.Replacer
var strconvLiquorStore strconv.NumError

var basicLiquorStoreTraitMapping = map[string]func(*schema.LiquorStoreTrait, []string){}
var customLiquorStoreTraitMapping = map[string]func(*schema.LiquorStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LiquorStore", func(ctx jsonld.Context) (interface{}) {
		return NewLiquorStoreFromContext(ctx)
	})

}

func NewLiquorStoreFromContext(ctx jsonld.Context) (x schema.LiquorStore) {
	x.Type = "http://schema.org/LiquorStore"
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


	MapBasicToLiquorStoreTrait(ctx, &x.LiquorStoreTrait)
	MapCustomToLiquorStoreTrait(ctx, &x.LiquorStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLiquorStoreTrait(ctx jsonld.Context, x *schema.LiquorStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLiquorStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLiquorStoreTrait(ctx jsonld.Context, x *schema.LiquorStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLiquorStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}