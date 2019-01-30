package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsScholarlyArticle strings.Replacer
var strconvScholarlyArticle strconv.NumError

var basicScholarlyArticleTraitMapping = map[string]func(*schema.ScholarlyArticleTrait, []string){}
var customScholarlyArticleTraitMapping = map[string]func(*schema.ScholarlyArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ScholarlyArticle", func(ctx jsonld.Context) (interface{}) {
		return NewScholarlyArticleFromContext(ctx)
	})

}

func NewScholarlyArticleFromContext(ctx jsonld.Context) (x schema.ScholarlyArticle) {
	x.Type = "http://schema.org/ScholarlyArticle"
	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToScholarlyArticleTrait(ctx, &x.ScholarlyArticleTrait)
	MapCustomToScholarlyArticleTrait(ctx, &x.ScholarlyArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToScholarlyArticleTrait(ctx jsonld.Context, x *schema.ScholarlyArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicScholarlyArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToScholarlyArticleTrait(ctx jsonld.Context, x *schema.ScholarlyArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customScholarlyArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}