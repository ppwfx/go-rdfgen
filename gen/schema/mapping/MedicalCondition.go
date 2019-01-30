package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalCondition strings.Replacer
var strconvMedicalCondition strconv.NumError

var basicMedicalConditionTraitMapping = map[string]func(*schema.MedicalConditionTrait, []string){}
var customMedicalConditionTraitMapping = map[string]func(*schema.MedicalConditionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalCondition", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalConditionFromContext(ctx)
	})



	basicMedicalConditionTraitMapping["http://schema.org/subtype"] = func(x *schema.MedicalConditionTrait, s []string) {
		x.Subtype = s[0]
	}


	basicMedicalConditionTraitMapping["http://schema.org/naturalProgression"] = func(x *schema.MedicalConditionTrait, s []string) {
		x.NaturalProgression = s[0]
	}


	basicMedicalConditionTraitMapping["http://schema.org/pathophysiology"] = func(x *schema.MedicalConditionTrait, s []string) {
		x.Pathophysiology = s[0]
	}


	basicMedicalConditionTraitMapping["http://schema.org/possibleComplication"] = func(x *schema.MedicalConditionTrait, s []string) {
		x.PossibleComplication = s[0]
	}


	basicMedicalConditionTraitMapping["http://schema.org/expectedPrognosis"] = func(x *schema.MedicalConditionTrait, s []string) {
		x.ExpectedPrognosis = s[0]
	}


	basicMedicalConditionTraitMapping["http://schema.org/epidemiology"] = func(x *schema.MedicalConditionTrait, s []string) {
		x.Epidemiology = s[0]
	}













	customMedicalConditionTraitMapping["http://schema.org/status"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Status, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Status = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Status = s
		}
	}

	customMedicalConditionTraitMapping["http://schema.org/associatedAnatomy"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.AssociatedAnatomy, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.AssociatedAnatomy = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.AssociatedAnatomy = s
		}
	}

	customMedicalConditionTraitMapping["http://schema.org/drug"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
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

	customMedicalConditionTraitMapping["http://schema.org/cause"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
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

	customMedicalConditionTraitMapping["http://schema.org/differentialDiagnosis"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
		var y schema.DDxElement
		if strings.HasPrefix(s, "_:") {
			y = NewDDxElementFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDDxElement()
			y.Id = s
		}

		x.DifferentialDiagnosis = &y

		return
	}

	customMedicalConditionTraitMapping["http://schema.org/possibleTreatment"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
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

	customMedicalConditionTraitMapping["http://schema.org/primaryPrevention"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
		var y schema.MedicalTherapy
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalTherapyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalTherapy()
			y.Id = s
		}

		x.PrimaryPrevention = &y

		return
	}

	customMedicalConditionTraitMapping["http://schema.org/riskFactor"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
		var y schema.MedicalRiskFactor
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalRiskFactorFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalRiskFactor()
			y.Id = s
		}

		x.RiskFactor = &y

		return
	}

	customMedicalConditionTraitMapping["http://schema.org/secondaryPrevention"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
		var y schema.MedicalTherapy
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalTherapyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalTherapy()
			y.Id = s
		}

		x.SecondaryPrevention = &y

		return
	}

	customMedicalConditionTraitMapping["http://schema.org/signOrSymptom"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
		var y schema.MedicalSignOrSymptom
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalSignOrSymptomFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalSignOrSymptom()
			y.Id = s
		}

		x.SignOrSymptom = &y

		return
	}

	customMedicalConditionTraitMapping["http://schema.org/stage"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
		var y schema.MedicalConditionStage
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalConditionStageFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalConditionStage()
			y.Id = s
		}

		x.Stage = &y

		return
	}

	customMedicalConditionTraitMapping["http://schema.org/typicalTest"] = func(x *schema.MedicalConditionTrait, ctx jsonld.Context, s string){
		var y schema.MedicalTest
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalTestFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalTest()
			y.Id = s
		}

		x.TypicalTest = &y

		return
	}
}

func NewMedicalConditionFromContext(ctx jsonld.Context) (x schema.MedicalCondition) {
	x.Type = "http://schema.org/MedicalCondition"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)
	MapCustomToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalConditionTrait(ctx jsonld.Context, x *schema.MedicalConditionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalConditionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalConditionTrait(ctx jsonld.Context, x *schema.MedicalConditionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalConditionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}