package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalGuidelineContraindication strings.Replacer
var strconvMedicalGuidelineContraindication strconv.NumError

var basicMedicalGuidelineContraindicationTraitMapping = map[string]func(*schema.MedicalGuidelineContraindicationTrait, []string){}
var customMedicalGuidelineContraindicationTraitMapping = map[string]func(*schema.MedicalGuidelineContraindicationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalGuidelineContraindication", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalGuidelineContraindicationFromContext(ctx)
	})

}

func NewMedicalGuidelineContraindicationFromContext(ctx jsonld.Context) (x schema.MedicalGuidelineContraindication) {
	x.Type = "http://schema.org/MedicalGuidelineContraindication"
	MapBasicToMedicalGuidelineTrait(ctx, &x.MedicalGuidelineTrait)
	MapCustomToMedicalGuidelineTrait(ctx, &x.MedicalGuidelineTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalGuidelineContraindicationTrait(ctx, &x.MedicalGuidelineContraindicationTrait)
	MapCustomToMedicalGuidelineContraindicationTrait(ctx, &x.MedicalGuidelineContraindicationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalGuidelineContraindicationTrait(ctx jsonld.Context, x *schema.MedicalGuidelineContraindicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalGuidelineContraindicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalGuidelineContraindicationTrait(ctx jsonld.Context, x *schema.MedicalGuidelineContraindicationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalGuidelineContraindicationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}