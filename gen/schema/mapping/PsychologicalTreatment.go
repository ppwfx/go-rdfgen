package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPsychologicalTreatment strings.Replacer
var strconvPsychologicalTreatment strconv.NumError

var basicPsychologicalTreatmentTraitMapping = map[string]func(*schema.PsychologicalTreatmentTrait, []string){}
var customPsychologicalTreatmentTraitMapping = map[string]func(*schema.PsychologicalTreatmentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PsychologicalTreatment", func(ctx jsonld.Context) (interface{}) {
		return NewPsychologicalTreatmentFromContext(ctx)
	})

}

func NewPsychologicalTreatmentFromContext(ctx jsonld.Context) (x schema.PsychologicalTreatment) {
	x.Type = "http://schema.org/PsychologicalTreatment"
	MapBasicToTherapeuticProcedureTrait(ctx, &x.TherapeuticProcedureTrait)
	MapCustomToTherapeuticProcedureTrait(ctx, &x.TherapeuticProcedureTrait)

	MapBasicToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)
	MapCustomToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPsychologicalTreatmentTrait(ctx, &x.PsychologicalTreatmentTrait)
	MapCustomToPsychologicalTreatmentTrait(ctx, &x.PsychologicalTreatmentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPsychologicalTreatmentTrait(ctx jsonld.Context, x *schema.PsychologicalTreatmentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPsychologicalTreatmentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPsychologicalTreatmentTrait(ctx jsonld.Context, x *schema.PsychologicalTreatmentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPsychologicalTreatmentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}