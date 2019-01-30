package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOutletStore strings.Replacer
var strconvOutletStore strconv.NumError

var basicOutletStoreTraitMapping = map[string]func(*schema.OutletStoreTrait, []string){}
var customOutletStoreTraitMapping = map[string]func(*schema.OutletStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OutletStore", func(ctx jsonld.Context) (interface{}) {
		return NewOutletStoreFromContext(ctx)
	})

}

func NewOutletStoreFromContext(ctx jsonld.Context) (x schema.OutletStore) {
	x.Type = "http://schema.org/OutletStore"
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


	MapBasicToOutletStoreTrait(ctx, &x.OutletStoreTrait)
	MapCustomToOutletStoreTrait(ctx, &x.OutletStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOutletStoreTrait(ctx jsonld.Context, x *schema.OutletStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOutletStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOutletStoreTrait(ctx jsonld.Context, x *schema.OutletStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOutletStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}