package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsStore strings.Replacer
var strconvStore strconv.NumError

var basicStoreTraitMapping = map[string]func(*schema.StoreTrait, []string){}
var customStoreTraitMapping = map[string]func(*schema.StoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Store", func(ctx jsonld.Context) (interface{}) {
		return NewStoreFromContext(ctx)
	})

}

func NewStoreFromContext(ctx jsonld.Context) (x schema.Store) {
	x.Type = "http://schema.org/Store"
	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToStoreTrait(ctx, &x.StoreTrait)
	MapCustomToStoreTrait(ctx, &x.StoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToStoreTrait(ctx jsonld.Context, x *schema.StoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToStoreTrait(ctx jsonld.Context, x *schema.StoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}