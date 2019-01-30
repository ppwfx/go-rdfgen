package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAskPublicNewsArticle strings.Replacer
var strconvAskPublicNewsArticle strconv.NumError

var basicAskPublicNewsArticleTraitMapping = map[string]func(*schema.AskPublicNewsArticleTrait, []string){}
var customAskPublicNewsArticleTraitMapping = map[string]func(*schema.AskPublicNewsArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AskPublicNewsArticle", func(ctx jsonld.Context) (interface{}) {
		return NewAskPublicNewsArticleFromContext(ctx)
	})

}

func NewAskPublicNewsArticleFromContext(ctx jsonld.Context) (x schema.AskPublicNewsArticle) {
	x.Type = "http://schema.org/AskPublicNewsArticle"
	MapBasicToNewsArticleTrait(ctx, &x.NewsArticleTrait)
	MapCustomToNewsArticleTrait(ctx, &x.NewsArticleTrait)

	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAskPublicNewsArticleTrait(ctx, &x.AskPublicNewsArticleTrait)
	MapCustomToAskPublicNewsArticleTrait(ctx, &x.AskPublicNewsArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAskPublicNewsArticleTrait(ctx jsonld.Context, x *schema.AskPublicNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAskPublicNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAskPublicNewsArticleTrait(ctx jsonld.Context, x *schema.AskPublicNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAskPublicNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}