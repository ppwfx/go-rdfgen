package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReviewNewsArticle strings.Replacer
var strconvReviewNewsArticle strconv.NumError

var basicReviewNewsArticleTraitMapping = map[string]func(*schema.ReviewNewsArticleTrait, []string){}
var customReviewNewsArticleTraitMapping = map[string]func(*schema.ReviewNewsArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReviewNewsArticle", func(ctx jsonld.Context) (interface{}) {
		return NewReviewNewsArticleFromContext(ctx)
	})

}

func NewReviewNewsArticleFromContext(ctx jsonld.Context) (x schema.ReviewNewsArticle) {
	x.Type = "http://schema.org/ReviewNewsArticle"
	MapBasicToNewsArticleTrait(ctx, &x.NewsArticleTrait)
	MapCustomToNewsArticleTrait(ctx, &x.NewsArticleTrait)

	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToCriticReviewTrait(ctx, &x.CriticReviewTrait)
	MapCustomToCriticReviewTrait(ctx, &x.CriticReviewTrait)

	MapBasicToReviewTrait(ctx, &x.ReviewTrait)
	MapCustomToReviewTrait(ctx, &x.ReviewTrait)


	MapBasicToReviewNewsArticleTrait(ctx, &x.ReviewNewsArticleTrait)
	MapCustomToReviewNewsArticleTrait(ctx, &x.ReviewNewsArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReviewNewsArticleTrait(ctx jsonld.Context, x *schema.ReviewNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReviewNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReviewNewsArticleTrait(ctx jsonld.Context, x *schema.ReviewNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReviewNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}