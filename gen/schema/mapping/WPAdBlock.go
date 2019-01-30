package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWPAdBlock strings.Replacer
var strconvWPAdBlock strconv.NumError

var basicWPAdBlockTraitMapping = map[string]func(*schema.WPAdBlockTrait, []string){}
var customWPAdBlockTraitMapping = map[string]func(*schema.WPAdBlockTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WPAdBlock", func(ctx jsonld.Context) (interface{}) {
		return NewWPAdBlockFromContext(ctx)
	})

}

func NewWPAdBlockFromContext(ctx jsonld.Context) (x schema.WPAdBlock) {
	x.Type = "http://schema.org/WPAdBlock"
	MapBasicToWebPageElementTrait(ctx, &x.WebPageElementTrait)
	MapCustomToWebPageElementTrait(ctx, &x.WebPageElementTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWPAdBlockTrait(ctx, &x.WPAdBlockTrait)
	MapCustomToWPAdBlockTrait(ctx, &x.WPAdBlockTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWPAdBlockTrait(ctx jsonld.Context, x *schema.WPAdBlockTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWPAdBlockTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWPAdBlockTrait(ctx jsonld.Context, x *schema.WPAdBlockTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWPAdBlockTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}