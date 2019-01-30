package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalTest strings.Replacer
var strconvMedicalTest strconv.NumError

var basicMedicalTestTraitMapping = map[string]func(*schema.MedicalTestTrait, []string){}
var customMedicalTestTraitMapping = map[string]func(*schema.MedicalTestTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalTest", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalTestFromContext(ctx)
	})







	customMedicalTestTraitMapping["http://schema.org/normalRange"] = func(x *schema.MedicalTestTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.NormalRange, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.NormalRange = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.NormalRange = s
		}
	}

	customMedicalTestTraitMapping["http://schema.org/affectedBy"] = func(x *schema.MedicalTestTrait, ctx jsonld.Context, s string){
		var y schema.Drug
		if strings.HasPrefix(s, "_:") {
			y = NewDrugFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDrug()
			y.Id = s
		}

		x.AffectedBy = &y

		return
	}

	customMedicalTestTraitMapping["http://schema.org/signDetected"] = func(x *schema.MedicalTestTrait, ctx jsonld.Context, s string){
		var y schema.MedicalSign
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalSignFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalSign()
			y.Id = s
		}

		x.SignDetected = &y

		return
	}

	customMedicalTestTraitMapping["http://schema.org/usedToDiagnose"] = func(x *schema.MedicalTestTrait, ctx jsonld.Context, s string){
		var y schema.MedicalCondition
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalConditionFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalCondition()
			y.Id = s
		}

		x.UsedToDiagnose = &y

		return
	}

	customMedicalTestTraitMapping["http://schema.org/usesDevice"] = func(x *schema.MedicalTestTrait, ctx jsonld.Context, s string){
		var y schema.MedicalDevice
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalDeviceFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalDevice()
			y.Id = s
		}

		x.UsesDevice = &y

		return
	}
}

func NewMedicalTestFromContext(ctx jsonld.Context) (x schema.MedicalTest) {
	x.Type = "http://schema.org/MedicalTest"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalTestTrait(ctx, &x.MedicalTestTrait)
	MapCustomToMedicalTestTrait(ctx, &x.MedicalTestTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalTestTrait(ctx jsonld.Context, x *schema.MedicalTestTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalTestTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalTestTrait(ctx jsonld.Context, x *schema.MedicalTestTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalTestTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}