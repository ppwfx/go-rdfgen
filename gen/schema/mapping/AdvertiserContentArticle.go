package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAdvertiserContentArticle strings.Replacer
var strconvAdvertiserContentArticle strconv.NumError

var basicAdvertiserContentArticleTraitMapping = map[string]func(*schema.AdvertiserContentArticleTrait, []string){}
var customAdvertiserContentArticleTraitMapping = map[string]func(*schema.AdvertiserContentArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AdvertiserContentArticle", func(ctx jsonld.Context) (interface{}) {
		return NewAdvertiserContentArticleFromContext(ctx)
	})

}

func NewAdvertiserContentArticleFromContext(ctx jsonld.Context) (x schema.AdvertiserContentArticle) {
	x.Type = "http://schema.org/AdvertiserContentArticle"
	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAdvertiserContentArticleTrait(ctx, &x.AdvertiserContentArticleTrait)
	MapCustomToAdvertiserContentArticleTrait(ctx, &x.AdvertiserContentArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAdvertiserContentArticleTrait(ctx jsonld.Context, x *schema.AdvertiserContentArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAdvertiserContentArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAdvertiserContentArticleTrait(ctx jsonld.Context, x *schema.AdvertiserContentArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAdvertiserContentArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}