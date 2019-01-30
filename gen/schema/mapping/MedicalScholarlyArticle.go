package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalScholarlyArticle strings.Replacer
var strconvMedicalScholarlyArticle strconv.NumError

var basicMedicalScholarlyArticleTraitMapping = map[string]func(*schema.MedicalScholarlyArticleTrait, []string){}
var customMedicalScholarlyArticleTraitMapping = map[string]func(*schema.MedicalScholarlyArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalScholarlyArticle", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalScholarlyArticleFromContext(ctx)
	})


	basicMedicalScholarlyArticleTraitMapping["http://schema.org/publicationType"] = func(x *schema.MedicalScholarlyArticleTrait, s []string) {
		x.PublicationType = s[0]
	}

}

func NewMedicalScholarlyArticleFromContext(ctx jsonld.Context) (x schema.MedicalScholarlyArticle) {
	x.Type = "http://schema.org/MedicalScholarlyArticle"
	MapBasicToScholarlyArticleTrait(ctx, &x.ScholarlyArticleTrait)
	MapCustomToScholarlyArticleTrait(ctx, &x.ScholarlyArticleTrait)

	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalScholarlyArticleTrait(ctx, &x.MedicalScholarlyArticleTrait)
	MapCustomToMedicalScholarlyArticleTrait(ctx, &x.MedicalScholarlyArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalScholarlyArticleTrait(ctx jsonld.Context, x *schema.MedicalScholarlyArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalScholarlyArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalScholarlyArticleTrait(ctx jsonld.Context, x *schema.MedicalScholarlyArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalScholarlyArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}