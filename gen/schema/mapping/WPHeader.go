package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWPHeader strings.Replacer
var strconvWPHeader strconv.NumError

var basicWPHeaderTraitMapping = map[string]func(*schema.WPHeaderTrait, []string){}
var customWPHeaderTraitMapping = map[string]func(*schema.WPHeaderTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WPHeader", func(ctx jsonld.Context) (interface{}) {
		return NewWPHeaderFromContext(ctx)
	})

}

func NewWPHeaderFromContext(ctx jsonld.Context) (x schema.WPHeader) {
	x.Type = "http://schema.org/WPHeader"
	MapBasicToWebPageElementTrait(ctx, &x.WebPageElementTrait)
	MapCustomToWebPageElementTrait(ctx, &x.WebPageElementTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWPHeaderTrait(ctx, &x.WPHeaderTrait)
	MapCustomToWPHeaderTrait(ctx, &x.WPHeaderTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWPHeaderTrait(ctx jsonld.Context, x *schema.WPHeaderTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWPHeaderTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWPHeaderTrait(ctx jsonld.Context, x *schema.WPHeaderTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWPHeaderTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}