package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalRiskEstimator strings.Replacer
var strconvMedicalRiskEstimator strconv.NumError

var basicMedicalRiskEstimatorTraitMapping = map[string]func(*schema.MedicalRiskEstimatorTrait, []string){}
var customMedicalRiskEstimatorTraitMapping = map[string]func(*schema.MedicalRiskEstimatorTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalRiskEstimator", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalRiskEstimatorFromContext(ctx)
	})




	customMedicalRiskEstimatorTraitMapping["http://schema.org/estimatesRiskOf"] = func(x *schema.MedicalRiskEstimatorTrait, ctx jsonld.Context, s string){
		var y schema.MedicalEntity
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalEntityFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalEntity()
			y.Id = s
		}

		x.EstimatesRiskOf = &y

		return
	}

	customMedicalRiskEstimatorTraitMapping["http://schema.org/includedRiskFactor"] = func(x *schema.MedicalRiskEstimatorTrait, ctx jsonld.Context, s string){
		var y schema.MedicalRiskFactor
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalRiskFactorFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalRiskFactor()
			y.Id = s
		}

		x.IncludedRiskFactor = &y

		return
	}
}

func NewMedicalRiskEstimatorFromContext(ctx jsonld.Context) (x schema.MedicalRiskEstimator) {
	x.Type = "http://schema.org/MedicalRiskEstimator"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalRiskEstimatorTrait(ctx, &x.MedicalRiskEstimatorTrait)
	MapCustomToMedicalRiskEstimatorTrait(ctx, &x.MedicalRiskEstimatorTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalRiskEstimatorTrait(ctx jsonld.Context, x *schema.MedicalRiskEstimatorTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalRiskEstimatorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalRiskEstimatorTrait(ctx jsonld.Context, x *schema.MedicalRiskEstimatorTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalRiskEstimatorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}