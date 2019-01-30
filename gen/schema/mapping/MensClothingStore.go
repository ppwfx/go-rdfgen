package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMensClothingStore strings.Replacer
var strconvMensClothingStore strconv.NumError

var basicMensClothingStoreTraitMapping = map[string]func(*schema.MensClothingStoreTrait, []string){}
var customMensClothingStoreTraitMapping = map[string]func(*schema.MensClothingStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MensClothingStore", func(ctx jsonld.Context) (interface{}) {
		return NewMensClothingStoreFromContext(ctx)
	})

}

func NewMensClothingStoreFromContext(ctx jsonld.Context) (x schema.MensClothingStore) {
	x.Type = "http://schema.org/MensClothingStore"
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


	MapBasicToMensClothingStoreTrait(ctx, &x.MensClothingStoreTrait)
	MapCustomToMensClothingStoreTrait(ctx, &x.MensClothingStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMensClothingStoreTrait(ctx jsonld.Context, x *schema.MensClothingStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMensClothingStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMensClothingStoreTrait(ctx jsonld.Context, x *schema.MensClothingStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMensClothingStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}