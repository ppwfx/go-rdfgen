package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWebApplication strings.Replacer
var strconvWebApplication strconv.NumError

var basicWebApplicationTraitMapping = map[string]func(*schema.WebApplicationTrait, []string){}
var customWebApplicationTraitMapping = map[string]func(*schema.WebApplicationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WebApplication", func(ctx jsonld.Context) (interface{}) {
		return NewWebApplicationFromContext(ctx)
	})


	basicWebApplicationTraitMapping["http://schema.org/browserRequirements"] = func(x *schema.WebApplicationTrait, s []string) {
		x.BrowserRequirements = s[0]
	}

}

func NewWebApplicationFromContext(ctx jsonld.Context) (x schema.WebApplication) {
	x.Type = "http://schema.org/WebApplication"
	MapBasicToSoftwareApplicationTrait(ctx, &x.SoftwareApplicationTrait)
	MapCustomToSoftwareApplicationTrait(ctx, &x.SoftwareApplicationTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWebApplicationTrait(ctx, &x.WebApplicationTrait)
	MapCustomToWebApplicationTrait(ctx, &x.WebApplicationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWebApplicationTrait(ctx jsonld.Context, x *schema.WebApplicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWebApplicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWebApplicationTrait(ctx jsonld.Context, x *schema.WebApplicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWebApplicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}