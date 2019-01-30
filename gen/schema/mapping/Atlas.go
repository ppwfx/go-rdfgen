package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAtlas strings.Replacer
var strconvAtlas strconv.NumError

var basicAtlasTraitMapping = map[string]func(*schema.AtlasTrait, []string){}
var customAtlasTraitMapping = map[string]func(*schema.AtlasTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Atlas", func(ctx jsonld.Context) (interface{}) {
		return NewAtlasFromContext(ctx)
	})

}

func NewAtlasFromContext(ctx jsonld.Context) (x schema.Atlas) {
	x.Type = "http://schema.org/Atlas"
	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAtlasTrait(ctx, &x.AtlasTrait)
	MapCustomToAtlasTrait(ctx, &x.AtlasTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAtlasTrait(ctx jsonld.Context, x *schema.AtlasTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAtlasTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAtlasTrait(ctx jsonld.Context, x *schema.AtlasTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAtlasTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}