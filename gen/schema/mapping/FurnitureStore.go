package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFurnitureStore strings.Replacer
var strconvFurnitureStore strconv.NumError

var basicFurnitureStoreTraitMapping = map[string]func(*schema.FurnitureStoreTrait, []string){}
var customFurnitureStoreTraitMapping = map[string]func(*schema.FurnitureStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FurnitureStore", func(ctx jsonld.Context) (interface{}) {
		return NewFurnitureStoreFromContext(ctx)
	})

}

func NewFurnitureStoreFromContext(ctx jsonld.Context) (x schema.FurnitureStore) {
	x.Type = "http://schema.org/FurnitureStore"
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


	MapBasicToFurnitureStoreTrait(ctx, &x.FurnitureStoreTrait)
	MapCustomToFurnitureStoreTrait(ctx, &x.FurnitureStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFurnitureStoreTrait(ctx jsonld.Context, x *schema.FurnitureStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFurnitureStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFurnitureStoreTrait(ctx jsonld.Context, x *schema.FurnitureStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFurnitureStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}