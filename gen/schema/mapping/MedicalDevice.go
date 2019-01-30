package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalDevice strings.Replacer
var strconvMedicalDevice strconv.NumError

var basicMedicalDeviceTraitMapping = map[string]func(*schema.MedicalDeviceTrait, []string){}
var customMedicalDeviceTraitMapping = map[string]func(*schema.MedicalDeviceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalDevice", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalDeviceFromContext(ctx)
	})


	basicMedicalDeviceTraitMapping["http://schema.org/preOp"] = func(x *schema.MedicalDeviceTrait, s []string) {
		x.PreOp = s[0]
	}


	basicMedicalDeviceTraitMapping["http://schema.org/procedure"] = func(x *schema.MedicalDeviceTrait, s []string) {
		x.Procedure = s[0]
	}



	basicMedicalDeviceTraitMapping["http://schema.org/postOp"] = func(x *schema.MedicalDeviceTrait, s []string) {
		x.PostOp = s[0]
	}






	customMedicalDeviceTraitMapping["http://schema.org/contraindication"] = func(x *schema.MedicalDeviceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Contraindication, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Contraindication = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Contraindication = s
		}
	}

	customMedicalDeviceTraitMapping["http://schema.org/purpose"] = func(x *schema.MedicalDeviceTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Purpose, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Purpose = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Purpose = s
		}
	}

	customMedicalDeviceTraitMapping["http://schema.org/indication"] = func(x *schema.MedicalDeviceTrait, ctx jsonld.Context, s string){
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

	customMedicalDeviceTraitMapping["http://schema.org/adverseOutcome"] = func(x *schema.MedicalDeviceTrait, ctx jsonld.Context, s string){
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

	customMedicalDeviceTraitMapping["http://schema.org/seriousAdverseOutcome"] = func(x *schema.MedicalDeviceTrait, ctx jsonld.Context, s string){
		var y schema.MedicalEntity
		if strings.HasPrefix(s, "_:") {
			y = NewMedicalEntityFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMedicalEntity()
			y.Id = s
		}

		x.SeriousAdverseOutcome = &y

		return
	}
}

func NewMedicalDeviceFromContext(ctx jsonld.Context) (x schema.MedicalDevice) {
	x.Type = "http://schema.org/MedicalDevice"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalDeviceTrait(ctx, &x.MedicalDeviceTrait)
	MapCustomToMedicalDeviceTrait(ctx, &x.MedicalDeviceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalDeviceTrait(ctx jsonld.Context, x *schema.MedicalDeviceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalDeviceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalDeviceTrait(ctx jsonld.Context, x *schema.MedicalDeviceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalDeviceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}