package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDrugStrength strings.Replacer
var strconvDrugStrength strconv.NumError

var basicDrugStrengthTraitMapping = map[string]func(*schema.DrugStrengthTrait, []string){}
var customDrugStrengthTraitMapping = map[string]func(*schema.DrugStrengthTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DrugStrength", func(ctx jsonld.Context) (interface{}) {
		return NewDrugStrengthFromContext(ctx)
	})


	basicDrugStrengthTraitMapping["http://schema.org/activeIngredient"] = func(x *schema.DrugStrengthTrait, s []string) {
		x.ActiveIngredient = s[0]
	}


	basicDrugStrengthTraitMapping["http://schema.org/strengthUnit"] = func(x *schema.DrugStrengthTrait, s []string) {
		x.StrengthUnit = s[0]
	}



	basicDrugStrengthTraitMapping["http://schema.org/strengthValue"] = func(x *schema.DrugStrengthTrait, s []string) {
		var err error
		x.StrengthValue, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}



	customDrugStrengthTraitMapping["http://schema.org/maximumIntake"] = func(x *schema.DrugStrengthTrait, ctx jsonld.Context, s string){
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

	customDrugStrengthTraitMapping["http://schema.org/availableIn"] = func(x *schema.DrugStrengthTrait, ctx jsonld.Context, s string){
		var y schema.AdministrativeArea
		if strings.HasPrefix(s, "_:") {
			y = NewAdministrativeAreaFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewAdministrativeArea()
			y.Id = s
		}

		x.AvailableIn = &y

		return
	}
}

func NewDrugStrengthFromContext(ctx jsonld.Context) (x schema.DrugStrength) {
	x.Type = "http://schema.org/DrugStrength"
	MapBasicToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)
	MapCustomToMedicalIntangibleTrait(ctx, &x.MedicalIntangibleTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDrugStrengthTrait(ctx, &x.DrugStrengthTrait)
	MapCustomToDrugStrengthTrait(ctx, &x.DrugStrengthTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDrugStrengthTrait(ctx jsonld.Context, x *schema.DrugStrengthTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDrugStrengthTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDrugStrengthTrait(ctx jsonld.Context, x *schema.DrugStrengthTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDrugStrengthTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}