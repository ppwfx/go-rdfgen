package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsContactPage strings.Replacer
var strconvContactPage strconv.NumError

var basicContactPageTraitMapping = map[string]func(*schema.ContactPageTrait, []string){}
var customContactPageTraitMapping = map[string]func(*schema.ContactPageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ContactPage", func(ctx jsonld.Context) (interface{}) {
		return NewContactPageFromContext(ctx)
	})

}

func NewContactPageFromContext(ctx jsonld.Context) (x schema.ContactPage) {
	x.Type = "http://schema.org/ContactPage"
	MapBasicToWebPageTrait(ctx, &x.WebPageTrait)
	MapCustomToWebPageTrait(ctx, &x.WebPageTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToContactPageTrait(ctx, &x.ContactPageTrait)
	MapCustomToContactPageTrait(ctx, &x.ContactPageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToContactPageTrait(ctx jsonld.Context, x *schema.ContactPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicContactPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToContactPageTrait(ctx jsonld.Context, x *schema.ContactPageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customContactPageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}