package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsNutritionInformation strings.Replacer
var strconvNutritionInformation strconv.NumError

var basicNutritionInformationTraitMapping = map[string]func(*schema.NutritionInformationTrait, []string){}
var customNutritionInformationTraitMapping = map[string]func(*schema.NutritionInformationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/NutritionInformation", func(ctx jsonld.Context) (interface{}) {
		return NewNutritionInformationFromContext(ctx)
	})









	basicNutritionInformationTraitMapping["http://schema.org/servingSize"] = func(x *schema.NutritionInformationTrait, s []string) {
		x.ServingSize = s[0]
	}






	customNutritionInformationTraitMapping["http://schema.org/calories"] = func(x *schema.NutritionInformationTrait, ctx jsonld.Context, s string){
		var y schema.Energy
		if strings.HasPrefix(s, "_:") {
			y = NewEnergyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewEnergy()
			y.Id = s
		}

		x.Calories = &y

		return
	}

	customNutritionInformationTraitMapping["http://schema.org/carbohydrateContent"] = func(x *schema.NutritionInformationTrait, ctx jsonld.Context, s string){
		var y schema.Mass
		if strings.HasPrefix(s, "_:") {
			y = NewMassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMass()
			y.Id = s
		}

		x.CarbohydrateContent = &y

		return
	}

	customNutritionInformationTraitMapping["http://schema.org/cholesterolContent"] = func(x *schema.NutritionInformationTrait, ctx jsonld.Context, s string){
		var y schema.Mass
		if strings.HasPrefix(s, "_:") {
			y = NewMassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMass()
			y.Id = s
		}

		x.CholesterolContent = &y

		return
	}

	customNutritionInformationTraitMapping["http://schema.org/fatContent"] = func(x *schema.NutritionInformationTrait, ctx jsonld.Context, s string){
		var y schema.Mass
		if strings.HasPrefix(s, "_:") {
			y = NewMassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMass()
			y.Id = s
		}

		x.FatContent = &y

		return
	}

	customNutritionInformationTraitMapping["http://schema.org/fiberContent"] = func(x *schema.NutritionInformationTrait, ctx jsonld.Context, s string){
		var y schema.Mass
		if strings.HasPrefix(s, "_:") {
			y = NewMassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMass()
			y.Id = s
		}

		x.FiberContent = &y

		return
	}

	customNutritionInformationTraitMapping["http://schema.org/proteinContent"] = func(x *schema.NutritionInformationTrait, ctx jsonld.Context, s string){
		var y schema.Mass
		if strings.HasPrefix(s, "_:") {
			y = NewMassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMass()
			y.Id = s
		}

		x.ProteinContent = &y

		return
	}

	customNutritionInformationTraitMapping["http://schema.org/saturatedFatContent"] = func(x *schema.NutritionInformationTrait, ctx jsonld.Context, s string){
		var y schema.Mass
		if strings.HasPrefix(s, "_:") {
			y = NewMassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMass()
			y.Id = s
		}

		x.SaturatedFatContent = &y

		return
	}

	customNutritionInformationTraitMapping["http://schema.org/sodiumContent"] = func(x *schema.NutritionInformationTrait, ctx jsonld.Context, s string){
		var y schema.Mass
		if strings.HasPrefix(s, "_:") {
			y = NewMassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMass()
			y.Id = s
		}

		x.SodiumContent = &y

		return
	}

	customNutritionInformationTraitMapping["http://schema.org/sugarContent"] = func(x *schema.NutritionInformationTrait, ctx jsonld.Context, s string){
		var y schema.Mass
		if strings.HasPrefix(s, "_:") {
			y = NewMassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMass()
			y.Id = s
		}

		x.SugarContent = &y

		return
	}

	customNutritionInformationTraitMapping["http://schema.org/transFatContent"] = func(x *schema.NutritionInformationTrait, ctx jsonld.Context, s string){
		var y schema.Mass
		if strings.HasPrefix(s, "_:") {
			y = NewMassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMass()
			y.Id = s
		}

		x.TransFatContent = &y

		return
	}

	customNutritionInformationTraitMapping["http://schema.org/unsaturatedFatContent"] = func(x *schema.NutritionInformationTrait, ctx jsonld.Context, s string){
		var y schema.Mass
		if strings.HasPrefix(s, "_:") {
			y = NewMassFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMass()
			y.Id = s
		}

		x.UnsaturatedFatContent = &y

		return
	}
}

func NewNutritionInformationFromContext(ctx jsonld.Context) (x schema.NutritionInformation) {
	x.Type = "http://schema.org/NutritionInformation"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToNutritionInformationTrait(ctx, &x.NutritionInformationTrait)
	MapCustomToNutritionInformationTrait(ctx, &x.NutritionInformationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToNutritionInformationTrait(ctx jsonld.Context, x *schema.NutritionInformationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicNutritionInformationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToNutritionInformationTrait(ctx jsonld.Context, x *schema.NutritionInformationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customNutritionInformationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}