package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMedicalConditionStage strings.Replacer
var strconvMedicalConditionStage strconv.NumError

var basicMedicalConditionStageTraitMapping = map[string]func(*schema.MedicalConditionStageTrait, []string){}
var customMedicalConditionStageTraitMapping = map[string]func(*schema.MedicalConditionStageTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MedicalConditionStage", func(ctx jsonld.Context) (interface{}) {
		return NewMedicalConditionStageFromContext(ctx)
	})


	basicMedicalConditionStageTraitMapping["http://schema.org/subStageSuffix"] = func(x *schema.MedicalConditionStageTrait, s []string) {
		x.SubStageSuffix = s[0]
	}


	basicMedicalConditionStageTraitMapping["http://schema.org/stageAsNumber"] = func(x *schema.MedicalConditionStageTrait, s []string) {
		var err error
		x.StageAsNumber, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}

}

func NewMedicalConditionStageFromContext(ctx jsonld.Context) (x schema.MedicalConditionStage) {
	x.Type = "http://schema.org/MedicalConditionStage"
	MapBasicToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)
	MapCustomToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMedicalConditionStageTrait(ctx, &x.MedicalConditionStageTrait)
	MapCustomToMedicalConditionStageTrait(ctx, &x.MedicalConditionStageTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMedicalConditionStageTrait(ctx jsonld.Context, x *schema.MedicalConditionStageTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMedicalConditionStageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMedicalConditionStageTrait(ctx jsonld.Context, x *schema.MedicalConditionStageTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMedicalConditionStageTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}