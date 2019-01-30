package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsJewelryStore strings.Replacer
var strconvJewelryStore strconv.NumError

var basicJewelryStoreTraitMapping = map[string]func(*schema.JewelryStoreTrait, []string){}
var customJewelryStoreTraitMapping = map[string]func(*schema.JewelryStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/JewelryStore", func(ctx jsonld.Context) (interface{}) {
		return NewJewelryStoreFromContext(ctx)
	})

}

func NewJewelryStoreFromContext(ctx jsonld.Context) (x schema.JewelryStore) {
	x.Type = "http://schema.org/JewelryStore"
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


	MapBasicToJewelryStoreTrait(ctx, &x.JewelryStoreTrait)
	MapCustomToJewelryStoreTrait(ctx, &x.JewelryStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToJewelryStoreTrait(ctx jsonld.Context, x *schema.JewelryStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicJewelryStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToJewelryStoreTrait(ctx jsonld.Context, x *schema.JewelryStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customJewelryStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}