package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsWPSideBar strings.Replacer
var strconvWPSideBar strconv.NumError

var basicWPSideBarTraitMapping = map[string]func(*schema.WPSideBarTrait, []string){}
var customWPSideBarTraitMapping = map[string]func(*schema.WPSideBarTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/WPSideBar", func(ctx jsonld.Context) (interface{}) {
		return NewWPSideBarFromContext(ctx)
	})

}

func NewWPSideBarFromContext(ctx jsonld.Context) (x schema.WPSideBar) {
	x.Type = "http://schema.org/WPSideBar"
	MapBasicToWebPageElementTrait(ctx, &x.WebPageElementTrait)
	MapCustomToWebPageElementTrait(ctx, &x.WebPageElementTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToWPSideBarTrait(ctx, &x.WPSideBarTrait)
	MapCustomToWPSideBarTrait(ctx, &x.WPSideBarTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToWPSideBarTrait(ctx jsonld.Context, x *schema.WPSideBarTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicWPSideBarTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToWPSideBarTrait(ctx jsonld.Context, x *schema.WPSideBarTrait) () {
	for k, v := range ctx.Current {
		f, ok := customWPSideBarTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}