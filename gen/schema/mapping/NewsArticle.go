package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsNewsArticle strings.Replacer
var strconvNewsArticle strconv.NumError

var basicNewsArticleTraitMapping = map[string]func(*schema.NewsArticleTrait, []string){}
var customNewsArticleTraitMapping = map[string]func(*schema.NewsArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/NewsArticle", func(ctx jsonld.Context) (interface{}) {
		return NewNewsArticleFromContext(ctx)
	})


	basicNewsArticleTraitMapping["http://schema.org/dateline"] = func(x *schema.NewsArticleTrait, s []string) {
		x.Dateline = s[0]
	}


	basicNewsArticleTraitMapping["http://schema.org/printColumn"] = func(x *schema.NewsArticleTrait, s []string) {
		x.PrintColumn = s[0]
	}


	basicNewsArticleTraitMapping["http://schema.org/printEdition"] = func(x *schema.NewsArticleTrait, s []string) {
		x.PrintEdition = s[0]
	}


	basicNewsArticleTraitMapping["http://schema.org/printPage"] = func(x *schema.NewsArticleTrait, s []string) {
		x.PrintPage = s[0]
	}


	basicNewsArticleTraitMapping["http://schema.org/printSection"] = func(x *schema.NewsArticleTrait, s []string) {
		x.PrintSection = s[0]
	}

}

func NewNewsArticleFromContext(ctx jsonld.Context) (x schema.NewsArticle) {
	x.Type = "http://schema.org/NewsArticle"
	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToNewsArticleTrait(ctx, &x.NewsArticleTrait)
	MapCustomToNewsArticleTrait(ctx, &x.NewsArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToNewsArticleTrait(ctx jsonld.Context, x *schema.NewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToNewsArticleTrait(ctx jsonld.Context, x *schema.NewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}