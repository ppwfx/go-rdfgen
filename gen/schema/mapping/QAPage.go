package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsQAPage strings.Replacer
var strconvQAPage strconv.NumError

var basicQAPageTraitMapping = map[string]func(*schema.QAPageTrait, []string){}
var customQAPageTraitMapping = map[string]func(*schema.QAPageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/QAPage", func(ctx jsonld.Context) (interface{}) {
		return NewQAPageFromContext(ctx)
	})

}

func NewQAPageFromContext(ctx jsonld.Context) (x schema.QAPage) {
	x.Type = "http://schema.org/QAPage"
	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToQAPageTrait(ctx, &x.QAPageTrait)
	MapCustomToQAPageTrait(ctx, &x.QAPageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToQAPageTrait(ctx jsonld.Context, x *schema.QAPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicQAPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToQAPageTrait(ctx jsonld.Context, x *schema.QAPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customQAPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}