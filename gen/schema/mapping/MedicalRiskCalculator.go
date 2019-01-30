package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalRiskCalculator strings.Replacer
var strconvMedicalRiskCalculator strconv.NumError

var basicMedicalRiskCalculatorTraitMapping = map[string]func(*schema.MedicalRiskCalculatorTrait, []string){}
var customMedicalRiskCalculatorTraitMapping = map[string]func(*schema.MedicalRiskCalculatorTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalRiskCalculator", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalRiskCalculatorFromContext(ctx)
	})

}

func NewMedicalRiskCalculatorFromContext(ctx jsonld.Context) (x schema.MedicalRiskCalculator) {
	x.Type = "http://schema.org/MedicalRiskCalculator"
	MapBasicToMedicalRiskEstimatorTrait(ctx, &x.MedicalRiskEstimatorTrait)
	MapCustomToMedicalRiskEstimatorTrait(ctx, &x.MedicalRiskEstimatorTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalRiskCalculatorTrait(ctx, &x.MedicalRiskCalculatorTrait)
	MapCustomToMedicalRiskCalculatorTrait(ctx, &x.MedicalRiskCalculatorTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalRiskCalculatorTrait(ctx jsonld.Context, x *schema.MedicalRiskCalculatorTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalRiskCalculatorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalRiskCalculatorTrait(ctx jsonld.Context, x *schema.MedicalRiskCalculatorTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalRiskCalculatorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}