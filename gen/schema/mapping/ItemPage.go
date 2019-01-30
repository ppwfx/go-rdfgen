package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsItemPage strings.Replacer
var strconvItemPage strconv.NumError

var basicItemPageTraitMapping = map[string]func(*schema.ItemPageTrait, []string){}
var customItemPageTraitMapping = map[string]func(*schema.ItemPageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ItemPage", func(ctx jsonld.Context) (interface{}) {
		return NewItemPageFromContext(ctx)
	})

}

func NewItemPageFromContext(ctx jsonld.Context) (x schema.ItemPage) {
	x.Type = "http://schema.org/ItemPage"
	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToItemPageTrait(ctx, &x.ItemPageTrait)
	MapCustomToItemPageTrait(ctx, &x.ItemPageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToItemPageTrait(ctx jsonld.Context, x *schema.ItemPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicItemPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToItemPageTrait(ctx jsonld.Context, x *schema.ItemPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customItemPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}