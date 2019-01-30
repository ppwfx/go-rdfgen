package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMaximumDoseSchedule strings.Replacer
var strconvMaximumDoseSchedule strconv.NumError

var basicMaximumDoseScheduleTraitMapping = map[string]func(*schema.MaximumDoseScheduleTrait, []string){}
var customMaximumDoseScheduleTraitMapping = map[string]func(*schema.MaximumDoseScheduleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MaximumDoseSchedule", func(ctx jsonld.Context) (interface{}) {
		return NewMaximumDoseScheduleFromContext(ctx)
	})

}

func NewMaximumDoseScheduleFromContext(ctx jsonld.Context) (x schema.MaximumDoseSchedule) {
	x.Type = "http://schema.org/MaximumDoseSchedule"
	MapBasicToDoseScheduleTrait(ctx, &x.DoseScheduleTrait)
	MapCustomToDoseScheduleTrait(ctx, &x.DoseScheduleTrait)

	MapBasicToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)
	MapCustomToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMaximumDoseScheduleTrait(ctx, &x.MaximumDoseScheduleTrait)
	MapCustomToMaximumDoseScheduleTrait(ctx, &x.MaximumDoseScheduleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMaximumDoseScheduleTrait(ctx jsonld.Context, x *schema.MaximumDoseScheduleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMaximumDoseScheduleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMaximumDoseScheduleTrait(ctx jsonld.Context, x *schema.MaximumDoseScheduleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMaximumDoseScheduleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}