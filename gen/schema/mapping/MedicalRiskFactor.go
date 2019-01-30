package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalRiskFactor strings.Replacer
var strconvMedicalRiskFactor strconv.NumError

var basicMedicalRiskFactorTraitMapping = map[string]func(*schema.MedicalRiskFactorTrait, []string){}
var customMedicalRiskFactorTraitMapping = map[string]func(*schema.MedicalRiskFactorTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalRiskFactor", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalRiskFactorFromContext(ctx)
	})



	customMedicalRiskFactorTraitMapping["http://schema.org/increasesRiskOf"] = func(x *schema.MedicalRiskFactorTrait, ctx jsonld.Context, s string){
		var y schema.MedicalEntity
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalEntityFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalEntity()
			y.Id = s
		}

		x.IncreasesRiskOf = &y

		return
	}
}

func NewMedicalRiskFactorFromContext(ctx jsonld.Context) (x schema.MedicalRiskFactor) {
	x.Type = "http://schema.org/MedicalRiskFactor"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalRiskFactorTrait(ctx, &x.MedicalRiskFactorTrait)
	MapCustomToMedicalRiskFactorTrait(ctx, &x.MedicalRiskFactorTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalRiskFactorTrait(ctx jsonld.Context, x *schema.MedicalRiskFactorTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalRiskFactorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalRiskFactorTrait(ctx jsonld.Context, x *schema.MedicalRiskFactorTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalRiskFactorTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}