package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTechArticle strings.Replacer
var strconvTechArticle strconv.NumError

var basicTechArticleTraitMapping = map[string]func(*schema.TechArticleTrait, []string){}
var customTechArticleTraitMapping = map[string]func(*schema.TechArticleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TechArticle", func(ctx jsonld.Context) (interface{}) {
		return NewTechArticleFromContext(ctx)
	})


	basicTechArticleTraitMapping["http://schema.org/dependencies"] = func(x *schema.TechArticleTrait, s []string) {
		x.Dependencies = s[0]
	}


	basicTechArticleTraitMapping["http://schema.org/proficiencyLevel"] = func(x *schema.TechArticleTrait, s []string) {
		x.ProficiencyLevel = s[0]
	}

}

func NewTechArticleFromContext(ctx jsonld.Context) (x schema.TechArticle) {
	x.Type = "http://schema.org/TechArticle"
	MapBasicToArticleTrait(ctx, &x.ArticleTrait)
	MapCustomToArticleTrait(ctx, &x.ArticleTrait)

	MapBasicToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)
	MapCustomToCreativeWorkTrait(ctx, &x.CreativeWorkTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTechArticleTrait(ctx, &x.TechArticleTrait)
	MapCustomToTechArticleTrait(ctx, &x.TechArticleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTechArticleTrait(ctx jsonld.Context, x *schema.TechArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTechArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTechArticleTrait(ctx jsonld.Context, x *schema.TechArticleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTechArticleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}