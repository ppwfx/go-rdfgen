package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOpinionNewsArticle strings.Replacer
var strconvOpinionNewsArticle strconv.NumError

var basicOpinionNewsArticleTraitMapping = map[string]func(*schema.OpinionNewsArticleTrait, []string){}
var customOpinionNewsArticleTraitMapping = map[string]func(*schema.OpinionNewsArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OpinionNewsArticle", func(ctx jsonld.Context) (interface{}) {
		return NewOpinionNewsArticleFromContext(ctx)
	})

}

func NewOpinionNewsArticleFromContext(ctx jsonld.Context) (x schema.OpinionNewsArticle) {
	x.Type = "http://schema.org/OpinionNewsArticle"
	MapBasicToNewsArticleTrait(ctx, &x.NewsArticleTrait)
	MapCustomToNewsArticleTrait(ctx, &x.NewsArticleTrait)

	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOpinionNewsArticleTrait(ctx, &x.OpinionNewsArticleTrait)
	MapCustomToOpinionNewsArticleTrait(ctx, &x.OpinionNewsArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOpinionNewsArticleTrait(ctx jsonld.Context, x *schema.OpinionNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOpinionNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOpinionNewsArticleTrait(ctx jsonld.Context, x *schema.OpinionNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOpinionNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}