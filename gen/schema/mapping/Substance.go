package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSubstance strings.Replacer
var strconvSubstance strconv.NumError

var basicSubstanceTraitMapping = map[string]func(*schema.SubstanceTrait, []string){}
var customSubstanceTraitMapping = map[string]func(*schema.SubstanceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Substance", func(ctx jsonld.Context) (interface{}) {
		return NewSubstanceFromContext(ctx)
	})


	basicSubstanceTraitMapping["http://schema.org/activeIngredient"] = func(x *schema.SubstanceTrait, s []string) {
		x.ActiveIngredient = s[0]
	}



	customSubstanceTraitMapping["http://schema.org/maximumIntake"] = func(x *schema.SubstanceTrait, ctx jsonld.Context, s string){
		var y schema.MaximumDoseSchedule
		if strings.HasPrefix(s, "_:") {
			y = NewMaximumDoseScheduleFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMaximumDoseSchedule()
			y.Id = s
		}

		x.MaximumIntake = &y

		return
	}
}

func NewSubstanceFromContext(ctx jsonld.Context) (x schema.Substance) {
	x.Type = "http://schema.org/Substance"
	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSubstanceTrait(ctx, &x.SubstanceTrait)
	MapCustomToSubstanceTrait(ctx, &x.SubstanceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSubstanceTrait(ctx jsonld.Context, x *schema.SubstanceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSubstanceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSubstanceTrait(ctx jsonld.Context, x *schema.SubstanceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSubstanceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}