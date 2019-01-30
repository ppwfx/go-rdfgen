package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsFAQPage strings.Replacer
var strconvFAQPage strconv.NumError

var basicFAQPageTraitMapping = map[string]func(*schema.FAQPageTrait, []string){}
var customFAQPageTraitMapping = map[string]func(*schema.FAQPageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/FAQPage", func(ctx jsonld.Context) (interface{}) {
		return NewFAQPageFromContext(ctx)
	})

}

func NewFAQPageFromContext(ctx jsonld.Context) (x schema.FAQPage) {
	x.Type = "http://schema.org/FAQPage"
	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToFAQPageTrait(ctx, &x.FAQPageTrait)
	MapCustomToFAQPageTrait(ctx, &x.FAQPageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToFAQPageTrait(ctx jsonld.Context, x *schema.FAQPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicFAQPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToFAQPageTrait(ctx jsonld.Context, x *schema.FAQPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customFAQPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}