package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsShoeStore strings.Replacer
var strconvShoeStore strconv.NumError

var basicShoeStoreTraitMapping = map[string]func(*schema.ShoeStoreTrait, []string){}
var customShoeStoreTraitMapping = map[string]func(*schema.ShoeStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ShoeStore", func(ctx jsonld.Context) (interface{}) {
		return NewShoeStoreFromContext(ctx)
	})

}

func NewShoeStoreFromContext(ctx jsonld.Context) (x schema.ShoeStore) {
	x.Type = "http://schema.org/ShoeStore"
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


	MapBasicToShoeStoreTrait(ctx, &x.ShoeStoreTrait)
	MapCustomToShoeStoreTrait(ctx, &x.ShoeStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToShoeStoreTrait(ctx jsonld.Context, x *schema.ShoeStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicShoeStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToShoeStoreTrait(ctx jsonld.Context, x *schema.ShoeStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customShoeStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}