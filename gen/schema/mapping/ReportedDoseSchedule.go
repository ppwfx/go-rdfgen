package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsReportedDoseSchedule strings.Replacer
var strconvReportedDoseSchedule strconv.NumError

var basicReportedDoseScheduleTraitMapping = map[string]func(*schema.ReportedDoseScheduleTrait, []string){}
var customReportedDoseScheduleTraitMapping = map[string]func(*schema.ReportedDoseScheduleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ReportedDoseSchedule", func(ctx jsonld.Context) (interface{}) {
		return NewReportedDoseScheduleFromContext(ctx)
	})

}

func NewReportedDoseScheduleFromContext(ctx jsonld.Context) (x schema.ReportedDoseSchedule) {
	x.Type = "http://schema.org/ReportedDoseSchedule"
	MapBasicToDoseScheduleTrait(ctx, &x.DoseScheduleTrait)
	MapCustomToDoseScheduleTrait(ctx, &x.DoseScheduleTrait)

	MapBasicToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)
	MapCustomToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToReportedDoseScheduleTrait(ctx, &x.ReportedDoseScheduleTrait)
	MapCustomToReportedDoseScheduleTrait(ctx, &x.ReportedDoseScheduleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToReportedDoseScheduleTrait(ctx jsonld.Context, x *schema.ReportedDoseScheduleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicReportedDoseScheduleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToReportedDoseScheduleTrait(ctx jsonld.Context, x *schema.ReportedDoseScheduleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customReportedDoseScheduleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}