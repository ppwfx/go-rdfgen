package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTherapeuticProcedure strings.Replacer
var strconvTherapeuticProcedure strconv.NumError

var basicTherapeuticProcedureTraitMapping = map[string]func(*schema.TherapeuticProcedureTrait, []string){}
var customTherapeuticProcedureTraitMapping = map[string]func(*schema.TherapeuticProcedureTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TherapeuticProcedure", func(ctx jsonld.Context) (interface{}) {
		return NewTherapeuticProcedureFromContext(ctx)
	})






	customTherapeuticProcedureTraitMapping["http://schema.org/drug"] = func(x *schema.TherapeuticProcedureTrait, ctx jsonld.Context, s string){
		var y schema.Drug
		if strings.HasPrefix(s, "_:") {
			y = NewDrugFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDrug()
			y.Id = s
		}

		x.Drug = &y

		return
	}

	customTherapeuticProcedureTraitMapping["http://schema.org/doseSchedule"] = func(x *schema.TherapeuticProcedureTrait, ctx jsonld.Context, s string){
		var y schema.DoseSchedule
		if strings.HasPrefix(s, "_:") {
			y = NewDoseScheduleFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDoseSchedule()
			y.Id = s
		}

		x.DoseSchedule = &y

		return
	}

	customTherapeuticProcedureTraitMapping["http://schema.org/indication"] = func(x *schema.TherapeuticProcedureTrait, ctx jsonld.Context, s string){
		var y schema.MedicalIndication
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalIndicationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalIndication()
			y.Id = s
		}

		x.Indication = &y

		return
	}

	customTherapeuticProcedureTraitMapping["http://schema.org/adverseOutcome"] = func(x *schema.TherapeuticProcedureTrait, ctx jsonld.Context, s string){
		var y schema.MedicalEntity
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalEntityFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalEntity()
			y.Id = s
		}

		x.AdverseOutcome = &y

		return
	}
}

func NewTherapeuticProcedureFromContext(ctx jsonld.Context) (x schema.TherapeuticProcedure) {
	x.Type = "http://schema.org/TherapeuticProcedure"
	MapBasicToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)
	MapCustomToMedicalProcedureTrait(ctx, &x.MedicalProcedureTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTherapeuticProcedureTrait(ctx, &x.TherapeuticProcedureTrait)
	MapCustomToTherapeuticProcedureTrait(ctx, &x.TherapeuticProcedureTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTherapeuticProcedureTrait(ctx jsonld.Context, x *schema.TherapeuticProcedureTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTherapeuticProcedureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTherapeuticProcedureTrait(ctx jsonld.Context, x *schema.TherapeuticProcedureTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTherapeuticProcedureTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}