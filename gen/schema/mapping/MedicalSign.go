package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalSign strings.Replacer
var strconvMedicalSign strconv.NumError

var basicMedicalSignTraitMapping = map[string]func(*schema.MedicalSignTrait, []string){}
var customMedicalSignTraitMapping = map[string]func(*schema.MedicalSignTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalSign", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalSignFromContext(ctx)
	})




	customMedicalSignTraitMapping["http://schema.org/identifyingExam"] = func(x *schema.MedicalSignTrait, ctx jsonld.Context, s string){
		var y schema.PhysicalExam
		if strings.HasPrefix(s, "_:") {
			y = NewPhysicalExamFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPhysicalExam()
			y.Id = s
		}

		x.IdentifyingExam = &y

		return
	}

	customMedicalSignTraitMapping["http://schema.org/identifyingTest"] = func(x *schema.MedicalSignTrait, ctx jsonld.Context, s string){
		var y schema.MedicalTest
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalTestFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalTest()
			y.Id = s
		}

		x.IdentifyingTest = &y

		return
	}
}

func NewMedicalSignFromContext(ctx jsonld.Context) (x schema.MedicalSign) {
	x.Type = "http://schema.org/MedicalSign"
	MapBasicToMedicalSignOrSymptomTrait(ctx, &x.MedicalSignOrSymptomTrait)
	MapCustomToMedicalSignOrSymptomTrait(ctx, &x.MedicalSignOrSymptomTrait)

	MapBasicToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)
	MapCustomToMedicalConditionTrait(ctx, &x.MedicalConditionTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalSignTrait(ctx, &x.MedicalSignTrait)
	MapCustomToMedicalSignTrait(ctx, &x.MedicalSignTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalSignTrait(ctx jsonld.Context, x *schema.MedicalSignTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalSignTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalSignTrait(ctx jsonld.Context, x *schema.MedicalSignTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalSignTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}