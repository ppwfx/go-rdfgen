package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSiteNavigationElement strings.Replacer
var strconvSiteNavigationElement strconv.NumError

var basicSiteNavigationElementTraitMapping = map[string]func(*schema.SiteNavigationElementTrait, []string){}
var customSiteNavigationElementTraitMapping = map[string]func(*schema.SiteNavigationElementTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SiteNavigationElement", func(ctx jsonld.Context) (interface{}) {
		return NewSiteNavigationElementFromContext(ctx)
	})

}

func NewSiteNavigationElementFromContext(ctx jsonld.Context) (x schema.SiteNavigationElement) {
	x.Type = "http://schema.org/SiteNavigationElement"
	MapBasicToWebPageElementTrait(ctx, &x.WebPageElementTrait)
	MapCustomToWebPageElementTrait(ctx, &x.WebPageElementTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSiteNavigationElementTrait(ctx, &x.SiteNavigationElementTrait)
	MapCustomToSiteNavigationElementTrait(ctx, &x.SiteNavigationElementTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSiteNavigationElementTrait(ctx jsonld.Context, x *schema.SiteNavigationElementTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSiteNavigationElementTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSiteNavigationElementTrait(ctx jsonld.Context, x *schema.SiteNavigationElementTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSiteNavigationElementTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}