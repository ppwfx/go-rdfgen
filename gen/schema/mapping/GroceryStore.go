package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGroceryStore strings.Replacer
var strconvGroceryStore strconv.NumError

var basicGroceryStoreTraitMapping = map[string]func(*schema.GroceryStoreTrait, []string){}
var customGroceryStoreTraitMapping = map[string]func(*schema.GroceryStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GroceryStore", func(ctx jsonld.Context) (interface{}) {
		return NewGroceryStoreFromContext(ctx)
	})

}

func NewGroceryStoreFromContext(ctx jsonld.Context) (x schema.GroceryStore) {
	x.Type = "http://schema.org/GroceryStore"
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


	MapBasicToGroceryStoreTrait(ctx, &x.GroceryStoreTrait)
	MapCustomToGroceryStoreTrait(ctx, &x.GroceryStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGroceryStoreTrait(ctx jsonld.Context, x *schema.GroceryStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGroceryStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGroceryStoreTrait(ctx jsonld.Context, x *schema.GroceryStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGroceryStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}