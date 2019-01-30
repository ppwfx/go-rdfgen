package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalRiskScore strings.Replacer
var strconvMedicalRiskScore strconv.NumError

var basicMedicalRiskScoreTraitMapping = map[string]func(*schema.MedicalRiskScoreTrait, []string){}
var customMedicalRiskScoreTraitMapping = map[string]func(*schema.MedicalRiskScoreTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalRiskScore", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalRiskScoreFromContext(ctx)
	})


	basicMedicalRiskScoreTraitMapping["http://schema.org/algorithm"] = func(x *schema.MedicalRiskScoreTrait, s []string) {
		x.Algorithm = s[0]
	}

}

func NewMedicalRiskScoreFromContext(ctx jsonld.Context) (x schema.MedicalRiskScore) {
	x.Type = "http://schema.org/MedicalRiskScore"
	MapBasicToMedicalRiskEstimatorTrait(ctx, &x.MedicalRiskEstimatorTrait)
	MapCustomToMedicalRiskEstimatorTrait(ctx, &x.MedicalRiskEstimatorTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalRiskScoreTrait(ctx, &x.MedicalRiskScoreTrait)
	MapCustomToMedicalRiskScoreTrait(ctx, &x.MedicalRiskScoreTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalRiskScoreTrait(ctx jsonld.Context, x *schema.MedicalRiskScoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalRiskScoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalRiskScoreTrait(ctx jsonld.Context, x *schema.MedicalRiskScoreTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalRiskScoreTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}