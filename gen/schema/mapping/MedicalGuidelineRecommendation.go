package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalGuidelineRecommendation strings.Replacer
var strconvMedicalGuidelineRecommendation strconv.NumError

var basicMedicalGuidelineRecommendationTraitMapping = map[string]func(*schema.MedicalGuidelineRecommendationTrait, []string){}
var customMedicalGuidelineRecommendationTraitMapping = map[string]func(*schema.MedicalGuidelineRecommendationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalGuidelineRecommendation", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalGuidelineRecommendationFromContext(ctx)
	})


	basicMedicalGuidelineRecommendationTraitMapping["http://schema.org/recommendationStrength"] = func(x *schema.MedicalGuidelineRecommendationTrait, s []string) {
		x.RecommendationStrength = s[0]
	}

}

func NewMedicalGuidelineRecommendationFromContext(ctx jsonld.Context) (x schema.MedicalGuidelineRecommendation) {
	x.Type = "http://schema.org/MedicalGuidelineRecommendation"
	MapBasicToMedicalGuidelineTrait(ctx, &x.MedicalGuidelineTrait)
	MapCustomToMedicalGuidelineTrait(ctx, &x.MedicalGuidelineTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalGuidelineRecommendationTrait(ctx, &x.MedicalGuidelineRecommendationTrait)
	MapCustomToMedicalGuidelineRecommendationTrait(ctx, &x.MedicalGuidelineRecommendationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalGuidelineRecommendationTrait(ctx jsonld.Context, x *schema.MedicalGuidelineRecommendationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalGuidelineRecommendationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalGuidelineRecommendationTrait(ctx jsonld.Context, x *schema.MedicalGuidelineRecommendationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalGuidelineRecommendationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}