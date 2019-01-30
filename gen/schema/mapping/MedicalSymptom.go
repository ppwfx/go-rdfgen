package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalSymptom strings.Replacer
var strconvMedicalSymptom strconv.NumError

var basicMedicalSymptomTraitMapping = map[string]func(*schema.MedicalSymptomTrait, []string){}
var customMedicalSymptomTraitMapping = map[string]func(*schema.MedicalSymptomTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalSymptom", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalSymptomFromContext(ctx)
	})

}

func NewMedicalSymptomFromContext(ctx jsonld.Context) (x schema.MedicalSymptom) {
	x.Type = "http://schema.org/MedicalSymptom"
	MapBasicToMedicalSignOrSymptomTrait(ctx, &x.MedicalSignOrSymptomTrait)
	MapCustomToMedicalSignOrSymptomTrait(ctx, &x.MedicalSignOrSymptomTrait)

	MapBasicToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)
	MapCustomToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalSymptomTrait(ctx, &x.MedicalSymptomTrait)
	MapCustomToMedicalSymptomTrait(ctx, &x.MedicalSymptomTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalSymptomTrait(ctx jsonld.Context, x *schema.MedicalSymptomTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalSymptomTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalSymptomTrait(ctx jsonld.Context, x *schema.MedicalSymptomTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalSymptomTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}