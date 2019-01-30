package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBookStore strings.Replacer
var strconvBookStore strconv.NumError

var basicBookStoreTraitMapping = map[string]func(*schema.BookStoreTrait, []string){}
var customBookStoreTraitMapping = map[string]func(*schema.BookStoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BookStore", func(ctx jsonld.Context) (interface{}) {
		return NewBookStoreFromContext(ctx)
	})

}

func NewBookStoreFromContext(ctx jsonld.Context) (x schema.BookStore) {
	x.Type = "http://schema.org/BookStore"
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


	MapBasicToBookStoreTrait(ctx, &x.BookStoreTrait)
	MapCustomToBookStoreTrait(ctx, &x.BookStoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBookStoreTrait(ctx jsonld.Context, x *schema.BookStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBookStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBookStoreTrait(ctx jsonld.Context, x *schema.BookStoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBookStoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}