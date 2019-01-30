package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBikeStore strings.Replacer
var strconvBikeStore strconv.NumError

var basicBikeStoreTraitMapping = map[string]func(*schema.BikeStoreTrait, []string){}
var customBikeStoreTraitMapping = map[string]func(*schema.BikeStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BikeStore", func(ctx jsonld.Context) (interface{}) {
		return NewBikeStoreFromContext(ctx)
	})

}

func NewBikeStoreFromContext(ctx jsonld.Context) (x schema.BikeStore) {
	x.Type = "http://schema.org/BikeStore"
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


	MapBasicToBikeStoreTrait(ctx, &x.BikeStoreTrait)
	MapCustomToBikeStoreTrait(ctx, &x.BikeStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBikeStoreTrait(ctx jsonld.Context, x *schema.BikeStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBikeStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBikeStoreTrait(ctx jsonld.Context, x *schema.BikeStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBikeStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}