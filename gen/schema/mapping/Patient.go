package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPatient strings.Replacer
var strconvPatient strconv.NumError

var basicPatientTraitMapping = map[string]func(*schema.PatientTrait, []string){}
var customPatientTraitMapping = map[string]func(*schema.PatientTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Patient", func(ctx jsonld.Context) (interface{}) {
		return NewPatientFromContext(ctx)
	})





	customPatientTraitMapping["http://schema.org/healthCondition"] = func(x *schema.PatientTrait, ctx jsonld.Context, s string){
		var y schema.MedicalCondition
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalConditionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalCondition()
			y.Id = s
		}

		x.HealthCondition = &y

		return
	}

	customPatientTraitMapping["http://schema.org/diagnosis"] = func(x *schema.PatientTrait, ctx jsonld.Context, s string){
		var y schema.MedicalCondition
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalConditionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalCondition()
			y.Id = s
		}

		x.Diagnosis = &y

		return
	}

	customPatientTraitMapping["http://schema.org/drug"] = func(x *schema.PatientTrait, ctx jsonld.Context, s string){
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
}

func NewPatientFromContext(ctx jsonld.Context) (x schema.Patient) {
	x.Type = "http://schema.org/Patient"
	MapBasicToPersonTrait(ctx, &x.PersonTrait)
	MapCustomToPersonTrait(ctx, &x.PersonTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToMedicalAudienceTrait(ctx, &x.MedicalAudienceTrait)
	MapCustomToMedicalAudienceTrait(ctx, &x.MedicalAudienceTrait)

	MapBasicToAudienceTrait(ctx, &x.AudienceTrait)
	MapCustomToAudienceTrait(ctx, &x.AudienceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)
	MapCustomToMedicalEnumerationTrait(ctx, &x.MedicalEnumerationTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToPeopleAudienceTrait(ctx, &x.PeopleAudienceTrait)
	MapCustomToPeopleAudienceTrait(ctx, &x.PeopleAudienceTrait)


	MapBasicToPatientTrait(ctx, &x.PatientTrait)
	MapCustomToPatientTrait(ctx, &x.PatientTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPatientTrait(ctx jsonld.Context, x *schema.PatientTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPatientTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPatientTrait(ctx jsonld.Context, x *schema.PatientTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPatientTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}