package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDoseSchedule strings.Replacer
var strconvDoseSchedule strconv.NumError

var basicDoseScheduleTraitMapping = map[string]func(*schema.DoseScheduleTrait, []string){}
var customDoseScheduleTraitMapping = map[string]func(*schema.DoseScheduleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DoseSchedule", func(ctx jsonld.Context) (interface{}) {
		return NewDoseScheduleFromContext(ctx)
	})


	basicDoseScheduleTraitMapping["http://schema.org/targetPopulation"] = func(x *schema.DoseScheduleTrait, s []string) {
		x.TargetPopulation = s[0]
	}


	basicDoseScheduleTraitMapping["http://schema.org/doseUnit"] = func(x *schema.DoseScheduleTrait, s []string) {
		x.DoseUnit = s[0]
	}


	basicDoseScheduleTraitMapping["http://schema.org/frequency"] = func(x *schema.DoseScheduleTrait, s []string) {
		x.Frequency = s[0]
	}



	customDoseScheduleTraitMapping["http://schema.org/doseValue"] = func(x *schema.DoseScheduleTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DoseValue, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DoseValue = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DoseValue = s
		}
	}
}

func NewDoseScheduleFromContext(ctx jsonld.Context) (x schema.DoseSchedule) {
	x.Type = "http://schema.org/DoseSchedule"
	MapBasicToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)
	MapCustomToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDoseScheduleTrait(ctx, &x.DoseScheduleTrait)
	MapCustomToDoseScheduleTrait(ctx, &x.DoseScheduleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDoseScheduleTrait(ctx jsonld.Context, x *schema.DoseScheduleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDoseScheduleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDoseScheduleTrait(ctx jsonld.Context, x *schema.DoseScheduleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDoseScheduleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}