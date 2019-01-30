package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalSignOrSymptom strings.Replacer
var strconvMedicalSignOrSymptom strconv.NumError

var basicMedicalSignOrSymptomTraitMapping = map[string]func(*schema.MedicalSignOrSymptomTrait, []string){}
var customMedicalSignOrSymptomTraitMapping = map[string]func(*schema.MedicalSignOrSymptomTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalSignOrSymptom", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalSignOrSymptomFromContext(ctx)
	})




	customMedicalSignOrSymptomTraitMapping["http://schema.org/cause"] = func(x *schema.MedicalSignOrSymptomTrait, ctx jsonld.Context, s string){
		var y schema.MedicalCause
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalCauseFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalCause()
			y.Id = s
		}

		x.Cause = &y

		return
	}

	customMedicalSignOrSymptomTraitMapping["http://schema.org/possibleTreatment"] = func(x *schema.MedicalSignOrSymptomTrait, ctx jsonld.Context, s string){
		var y schema.MedicalTherapy
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalTherapyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalTherapy()
			y.Id = s
		}

		x.PossibleTreatment = &y

		return
	}
}

func NewMedicalSignOrSymptomFromContext(ctx jsonld.Context) (x schema.MedicalSignOrSymptom) {
	x.Type = "http://schema.org/MedicalSignOrSymptom"
	MapBasicToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)
	MapCustomToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalSignOrSymptomTrait(ctx, &x.MedicalSignOrSymptomTrait)
	MapCustomToMedicalSignOrSymptomTrait(ctx, &x.MedicalSignOrSymptomTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalSignOrSymptomTrait(ctx jsonld.Context, x *schema.MedicalSignOrSymptomTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalSignOrSymptomTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalSignOrSymptomTrait(ctx jsonld.Context, x *schema.MedicalSignOrSymptomTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalSignOrSymptomTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}