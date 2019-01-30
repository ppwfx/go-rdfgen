package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGardenStore strings.Replacer
var strconvGardenStore strconv.NumError

var basicGardenStoreTraitMapping = map[string]func(*schema.GardenStoreTrait, []string){}
var customGardenStoreTraitMapping = map[string]func(*schema.GardenStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GardenStore", func(ctx jsonld.Context) (interface{}) {
		return NewGardenStoreFromContext(ctx)
	})

}

func NewGardenStoreFromContext(ctx jsonld.Context) (x schema.GardenStore) {
	x.Type = "http://schema.org/GardenStore"
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


	MapBasicToGardenStoreTrait(ctx, &x.GardenStoreTrait)
	MapCustomToGardenStoreTrait(ctx, &x.GardenStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGardenStoreTrait(ctx jsonld.Context, x *schema.GardenStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGardenStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGardenStoreTrait(ctx jsonld.Context, x *schema.GardenStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGardenStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}