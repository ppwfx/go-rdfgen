package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsToyStore strings.Replacer
var strconvToyStore strconv.NumError

var basicToyStoreTraitMapping = map[string]func(*schema.ToyStoreTrait, []string){}
var customToyStoreTraitMapping = map[string]func(*schema.ToyStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ToyStore", func(ctx jsonld.Context) (interface{}) {
		return NewToyStoreFromContext(ctx)
	})

}

func NewToyStoreFromContext(ctx jsonld.Context) (x schema.ToyStore) {
	x.Type = "http://schema.org/ToyStore"
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


	MapBasicToToyStoreTrait(ctx, &x.ToyStoreTrait)
	MapCustomToToyStoreTrait(ctx, &x.ToyStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToToyStoreTrait(ctx jsonld.Context, x *schema.ToyStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicToyStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToToyStoreTrait(ctx jsonld.Context, x *schema.ToyStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customToyStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}