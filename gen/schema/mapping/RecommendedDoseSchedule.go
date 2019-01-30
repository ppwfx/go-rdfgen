package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRecommendedDoseSchedule strings.Replacer
var strconvRecommendedDoseSchedule strconv.NumError

var basicRecommendedDoseScheduleTraitMapping = map[string]func(*schema.RecommendedDoseScheduleTrait, []string){}
var customRecommendedDoseScheduleTraitMapping = map[string]func(*schema.RecommendedDoseScheduleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RecommendedDoseSchedule", func(ctx jsonld.Context) (interface{}) {
		return NewRecommendedDoseScheduleFromContext(ctx)
	})

}

func NewRecommendedDoseScheduleFromContext(ctx jsonld.Context) (x schema.RecommendedDoseSchedule) {
	x.Type = "http://schema.org/RecommendedDoseSchedule"
	MapBasicToDoseScheduleTrait(ctx, &x.DoseScheduleTrait)
	MapCustomToDoseScheduleTrait(ctx, &x.DoseScheduleTrait)

	MapBasicToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)
	MapCustomToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRecommendedDoseScheduleTrait(ctx, &x.RecommendedDoseScheduleTrait)
	MapCustomToRecommendedDoseScheduleTrait(ctx, &x.RecommendedDoseScheduleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRecommendedDoseScheduleTrait(ctx jsonld.Context, x *schema.RecommendedDoseScheduleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRecommendedDoseScheduleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRecommendedDoseScheduleTrait(ctx jsonld.Context, x *schema.RecommendedDoseScheduleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRecommendedDoseScheduleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}