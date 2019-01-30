package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAnalysisNewsArticle strings.Replacer
var strconvAnalysisNewsArticle strconv.NumError

var basicAnalysisNewsArticleTraitMapping = map[string]func(*schema.AnalysisNewsArticleTrait, []string){}
var customAnalysisNewsArticleTraitMapping = map[string]func(*schema.AnalysisNewsArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AnalysisNewsArticle", func(ctx jsonld.Context) (interface{}) {
		return NewAnalysisNewsArticleFromContext(ctx)
	})

}

func NewAnalysisNewsArticleFromContext(ctx jsonld.Context) (x schema.AnalysisNewsArticle) {
	x.Type = "http://schema.org/AnalysisNewsArticle"
	MapBasicToNewsArticleTrait(ctx, &x.NewsArticleTrait)
	MapCustomToNewsArticleTrait(ctx, &x.NewsArticleTrait)

	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToAnalysisNewsArticleTrait(ctx, &x.AnalysisNewsArticleTrait)
	MapCustomToAnalysisNewsArticleTrait(ctx, &x.AnalysisNewsArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAnalysisNewsArticleTrait(ctx jsonld.Context, x *schema.AnalysisNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAnalysisNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAnalysisNewsArticleTrait(ctx jsonld.Context, x *schema.AnalysisNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAnalysisNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}