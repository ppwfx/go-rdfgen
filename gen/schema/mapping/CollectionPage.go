package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCollectionPage strings.Replacer
var strconvCollectionPage strconv.NumError

var basicCollectionPageTraitMapping = map[string]func(*schema.CollectionPageTrait, []string){}
var customCollectionPageTraitMapping = map[string]func(*schema.CollectionPageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CollectionPage", func(ctx jsonld.Context) (interface{}) {
		return NewCollectionPageFromContext(ctx)
	})

}

func NewCollectionPageFromContext(ctx jsonld.Context) (x schema.CollectionPage) {
	x.Type = "http://schema.org/CollectionPage"
	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCollectionPageTrait(ctx, &x.CollectionPageTrait)
	MapCustomToCollectionPageTrait(ctx, &x.CollectionPageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCollectionPageTrait(ctx jsonld.Context, x *schema.CollectionPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCollectionPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCollectionPageTrait(ctx jsonld.Context, x *schema.CollectionPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCollectionPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}