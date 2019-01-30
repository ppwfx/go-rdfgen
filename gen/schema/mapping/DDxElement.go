package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDDxElement strings.Replacer
var strconvDDxElement strconv.NumError

var basicDDxElementTraitMapping = map[string]func(*schema.DDxElementTrait, []string){}
var customDDxElementTraitMapping = map[string]func(*schema.DDxElementTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DDxElement", func(ctx jsonld.Context) (interface{}) {
		return NewDDxElementFromContext(ctx)
	})




	customDDxElementTraitMapping["http://schema.org/diagnosis"] = func(x *schema.DDxElementTrait, ctx jsonld.Context, s string){
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

	customDDxElementTraitMapping["http://schema.org/distinguishingSign"] = func(x *schema.DDxElementTrait, ctx jsonld.Context, s string){
		var y schema.MedicalSignOrSymptom
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalSignOrSymptomFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalSignOrSymptom()
			y.Id = s
		}

		x.DistinguishingSign = &y

		return
	}
}

func NewDDxElementFromContext(ctx jsonld.Context) (x schema.DDxElement) {
	x.Type = "http://schema.org/DDxElement"
	MapBasicToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)
	MapCustomToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDDxElementTrait(ctx, &x.DDxElementTrait)
	MapCustomToDDxElementTrait(ctx, &x.DDxElementTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDDxElementTrait(ctx jsonld.Context, x *schema.DDxElementTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDDxElementTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDDxElementTrait(ctx jsonld.Context, x *schema.DDxElementTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDDxElementTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}