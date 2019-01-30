package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAboutPage strings.Replacer
var strconvAboutPage strconv.NumError

var basicAboutPageTraitMapping = map[string]func(*schema.AboutPageTrait, []string){}
var customAboutPageTraitMapping = map[string]func(*schema.AboutPageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AboutPage", func(ctx jsonld.Context) (interface{}) {
		return NewAboutPageFromContext(ctx)
	})
}

func NewAboutPageFromContext(ctx jsonld.Context) (x schema.AboutPage) {
	x.Type = "http://schema.org/AboutPage"
	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAboutPageTrait(ctx, &x.AboutPageTrait)
	MapCustomToAboutPageTrait(ctx, &x.AboutPageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAboutPageTrait(ctx jsonld.Context, x *schema.AboutPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAboutPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAboutPageTrait(ctx jsonld.Context, x *schema.AboutPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAboutPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}