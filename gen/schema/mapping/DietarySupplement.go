package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDietarySupplement strings.Replacer
var strconvDietarySupplement strconv.NumError

var basicDietarySupplementTraitMapping = map[string]func(*schema.DietarySupplementTrait, []string){}
var customDietarySupplementTraitMapping = map[string]func(*schema.DietarySupplementTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DietarySupplement", func(ctx jsonld.Context) (interface{}) {
		return NewDietarySupplementFromContext(ctx)
	})




	basicDietarySupplementTraitMapping["http://schema.org/safetyConsideration"] = func(x *schema.DietarySupplementTrait, s []string) {
		x.SafetyConsideration = s[0]
	}


	basicDietarySupplementTraitMapping["http://schema.org/targetPopulation"] = func(x *schema.DietarySupplementTrait, s []string) {
		x.TargetPopulation = s[0]
	}


	basicDietarySupplementTraitMapping["http://schema.org/activeIngredient"] = func(x *schema.DietarySupplementTrait, s []string) {
		x.ActiveIngredient = s[0]
	}


	basicDietarySupplementTraitMapping["http://schema.org/nonProprietaryName"] = func(x *schema.DietarySupplementTrait, s []string) {
		x.NonProprietaryName = s[0]
	}


	basicDietarySupplementTraitMapping["http://schema.org/mechanismOfAction"] = func(x *schema.DietarySupplementTrait, s []string) {
		x.MechanismOfAction = s[0]
	}


	basicDietarySupplementTraitMapping["http://schema.org/background"] = func(x *schema.DietarySupplementTrait, s []string) {
		x.Background = s[0]
	}


	basicDietarySupplementTraitMapping["http://schema.org/proprietaryName"] = func(x *schema.DietarySupplementTrait, s []string) {
		x.ProprietaryName = s[0]
	}


	basicDietarySupplementTraitMapping["http://schema.org/isProprietary"] = func(x *schema.DietarySupplementTrait, s []string) {
		var err error
		x.IsProprietary, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}




	customDietarySupplementTraitMapping["http://schema.org/manufacturer"] = func(x *schema.DietarySupplementTrait, ctx jsonld.Context, s string){
		var y schema.Organization
		if strings.HasPrefix(s, "_:") {
			y = NewOrganizationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewOrganization()
			y.Id = s
		}

		x.Manufacturer = &y

		return
	}

	customDietarySupplementTraitMapping["http://schema.org/legalStatus"] = func(x *schema.DietarySupplementTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.LegalStatus, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.LegalStatus = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.LegalStatus = s
		}
	}

	customDietarySupplementTraitMapping["http://schema.org/maximumIntake"] = func(x *schema.DietarySupplementTrait, ctx jsonld.Context, s string){
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

	customDietarySupplementTraitMapping["http://schema.org/recommendedIntake"] = func(x *schema.DietarySupplementTrait, ctx jsonld.Context, s string){
		var y schema.RecommendedDoseSchedule
		if strings.HasPrefix(s, "_:") {
			y = NewRecommendedDoseScheduleFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewRecommendedDoseSchedule()
			y.Id = s
		}

		x.RecommendedIntake = &y

		return
	}
}

func NewDietarySupplementFromContext(ctx jsonld.Context) (x schema.DietarySupplement) {
	x.Type = "http://schema.org/DietarySupplement"
	MapBasicToSubstanceTrait(ctx, &x.SubstanceTrait)
	MapCustomToSubstanceTrait(ctx, &x.SubstanceTrait)

	MapBasicToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)
	MapCustomToMedicalEntityTrait(ctx, &x.MedicalEntityTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDietarySupplementTrait(ctx, &x.DietarySupplementTrait)
	MapCustomToDietarySupplementTrait(ctx, &x.DietarySupplementTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDietarySupplementTrait(ctx jsonld.Context, x *schema.DietarySupplementTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDietarySupplementTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDietarySupplementTrait(ctx jsonld.Context, x *schema.DietarySupplementTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDietarySupplementTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}