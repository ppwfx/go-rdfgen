package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWPFooter strings.Replacer
var strconvWPFooter strconv.NumError

var basicWPFooterTraitMapping = map[string]func(*schema.WPFooterTrait, []string){}
var customWPFooterTraitMapping = map[string]func(*schema.WPFooterTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WPFooter", func(ctx jsonld.Context) (interface{}) {
		return NewWPFooterFromContext(ctx)
	})

}

func NewWPFooterFromContext(ctx jsonld.Context) (x schema.WPFooter) {
	x.Type = "http://schema.org/WPFooter"
	MapBasicToWebPageElementTrait(ctx, &x.WebPageElementTrait)
	MapCustomToWebPageElementTrait(ctx, &x.WebPageElementTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWPFooterTrait(ctx, &x.WPFooterTrait)
	MapCustomToWPFooterTrait(ctx, &x.WPFooterTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWPFooterTrait(ctx jsonld.Context, x *schema.WPFooterTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWPFooterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWPFooterTrait(ctx jsonld.Context, x *schema.WPFooterTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWPFooterTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}