package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSatiricalArticle strings.Replacer
var strconvSatiricalArticle strconv.NumError

var basicSatiricalArticleTraitMapping = map[string]func(*schema.SatiricalArticleTrait, []string){}
var customSatiricalArticleTraitMapping = map[string]func(*schema.SatiricalArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SatiricalArticle", func(ctx jsonld.Context) (interface{}) {
		return NewSatiricalArticleFromContext(ctx)
	})

}

func NewSatiricalArticleFromContext(ctx jsonld.Context) (x schema.SatiricalArticle) {
	x.Type = "http://schema.org/SatiricalArticle"
	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSatiricalArticleTrait(ctx, &x.SatiricalArticleTrait)
	MapCustomToSatiricalArticleTrait(ctx, &x.SatiricalArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSatiricalArticleTrait(ctx jsonld.Context, x *schema.SatiricalArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSatiricalArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSatiricalArticleTrait(ctx jsonld.Context, x *schema.SatiricalArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSatiricalArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}