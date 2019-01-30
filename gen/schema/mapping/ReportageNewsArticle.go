package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReportageNewsArticle strings.Replacer
var strconvReportageNewsArticle strconv.NumError

var basicReportageNewsArticleTraitMapping = map[string]func(*schema.ReportageNewsArticleTrait, []string){}
var customReportageNewsArticleTraitMapping = map[string]func(*schema.ReportageNewsArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReportageNewsArticle", func(ctx jsonld.Context) (interface{}) {
		return NewReportageNewsArticleFromContext(ctx)
	})

}

func NewReportageNewsArticleFromContext(ctx jsonld.Context) (x schema.ReportageNewsArticle) {
	x.Type = "http://schema.org/ReportageNewsArticle"
	MapBasicToNewsArticleTrait(ctx, &x.NewsArticleTrait)
	MapCustomToNewsArticleTrait(ctx, &x.NewsArticleTrait)

	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReportageNewsArticleTrait(ctx, &x.ReportageNewsArticleTrait)
	MapCustomToReportageNewsArticleTrait(ctx, &x.ReportageNewsArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReportageNewsArticleTrait(ctx jsonld.Context, x *schema.ReportageNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReportageNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReportageNewsArticleTrait(ctx jsonld.Context, x *schema.ReportageNewsArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReportageNewsArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}