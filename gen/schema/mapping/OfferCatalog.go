package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOfferCatalog strings.Replacer
var strconvOfferCatalog strconv.NumError

var basicOfferCatalogTraitMapping = map[string]func(*schema.OfferCatalogTrait, []string){}
var customOfferCatalogTraitMapping = map[string]func(*schema.OfferCatalogTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OfferCatalog", func(ctx jsonld.Context) (interface{}) {
		return NewOfferCatalogFromContext(ctx)
	})

}

func NewOfferCatalogFromContext(ctx jsonld.Context) (x schema.OfferCatalog) {
	x.Type = "http://schema.org/OfferCatalog"
	MapBasicToItemListTrait(ctx, &x.ItemListTrait)
	MapCustomToItemListTrait(ctx, &x.ItemListTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOfferCatalogTrait(ctx, &x.OfferCatalogTrait)
	MapCustomToOfferCatalogTrait(ctx, &x.OfferCatalogTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOfferCatalogTrait(ctx jsonld.Context, x *schema.OfferCatalogTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOfferCatalogTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOfferCatalogTrait(ctx jsonld.Context, x *schema.OfferCatalogTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOfferCatalogTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}