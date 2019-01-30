package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBackgroundNewsArticle strings.Replacer
var strconvBackgroundNewsArticle strconv.NumError

var basicBackgroundNewsArticleTraitMapping = map[string]func(*schema.BackgroundNewsArticleTrait, []string){}
var customBackgroundNewsArticleTraitMapping = map[string]func(*schema.BackgroundNewsArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BackgroundNewsArticle", func(ctx jsonld.Context) (interface{}) {
		return NewBackgroundNewsArticleFromContext(ctx)
	})

}

func NewBackgroundNewsArticleFromContext(ctx jsonld.Context) (x schema.BackgroundNewsArticle) {
	x.Type = "http://schema.org/BackgroundNewsArticle"
	MapBasicToNewsArticleTrait(ctx, &x.NewsArticleTrait)
	MapCustomToNewsArticleTrait(ctx, &x.NewsArticleTrait)

	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBackgroundNewsArticleTrait(ctx, &x.BackgroundNewsArticleTrait)
	MapCustomToBackgroundNewsArticleTrait(ctx, &x.BackgroundNewsArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBackgroundNewsArticleTrait(ctx jsonld.Context, x *schema.BackgroundNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBackgroundNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBackgroundNewsArticleTrait(ctx jsonld.Context, x *schema.BackgroundNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBackgroundNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}