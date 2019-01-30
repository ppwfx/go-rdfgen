package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCheckoutPage strings.Replacer
var strconvCheckoutPage strconv.NumError

var basicCheckoutPageTraitMapping = map[string]func(*schema.CheckoutPageTrait, []string){}
var customCheckoutPageTraitMapping = map[string]func(*schema.CheckoutPageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CheckoutPage", func(ctx jsonld.Context) (interface{}) {
		return NewCheckoutPageFromContext(ctx)
	})

}

func NewCheckoutPageFromContext(ctx jsonld.Context) (x schema.CheckoutPage) {
	x.Type = "http://schema.org/CheckoutPage"
	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCheckoutPageTrait(ctx, &x.CheckoutPageTrait)
	MapCustomToCheckoutPageTrait(ctx, &x.CheckoutPageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCheckoutPageTrait(ctx jsonld.Context, x *schema.CheckoutPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCheckoutPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCheckoutPageTrait(ctx jsonld.Context, x *schema.CheckoutPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCheckoutPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}