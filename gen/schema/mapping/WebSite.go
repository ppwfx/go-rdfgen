package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWebSite strings.Replacer
var strconvWebSite strconv.NumError

var basicWebSiteTraitMapping = map[string]func(*schema.WebSiteTrait, []string){}
var customWebSiteTraitMapping = map[string]func(*schema.WebSiteTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WebSite", func(ctx jsonld.Context) (interface{}) {
		return NewWebSiteFromContext(ctx)
	})


	basicWebSiteTraitMapping["http://schema.org/issn"] = func(x *schema.WebSiteTrait, s []string) {
		x.Issn = s[0]
	}

}

func NewWebSiteFromContext(ctx jsonld.Context) (x schema.WebSite) {
	x.Type = "http://schema.org/WebSite"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWebSiteTrait(ctx, &x.WebSiteTrait)
	MapCustomToWebSiteTrait(ctx, &x.WebSiteTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWebSiteTrait(ctx jsonld.Context, x *schema.WebSiteTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWebSiteTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWebSiteTrait(ctx jsonld.Context, x *schema.WebSiteTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWebSiteTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}