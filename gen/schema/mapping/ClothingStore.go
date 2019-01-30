package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsClothingStore strings.Replacer
var strconvClothingStore strconv.NumError

var basicClothingStoreTraitMapping = map[string]func(*schema.ClothingStoreTrait, []string){}
var customClothingStoreTraitMapping = map[string]func(*schema.ClothingStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ClothingStore", func(ctx jsonld.Context) (interface{}) {
		return NewClothingStoreFromContext(ctx)
	})

}

func NewClothingStoreFromContext(ctx jsonld.Context) (x schema.ClothingStore) {
	x.Type = "http://schema.org/ClothingStore"
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


	MapBasicToClothingStoreTrait(ctx, &x.ClothingStoreTrait)
	MapCustomToClothingStoreTrait(ctx, &x.ClothingStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToClothingStoreTrait(ctx jsonld.Context, x *schema.ClothingStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicClothingStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToClothingStoreTrait(ctx jsonld.Context, x *schema.ClothingStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customClothingStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}