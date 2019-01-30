package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHealthInsurancePlan strings.Replacer
var strconvHealthInsurancePlan strconv.NumError

var basicHealthInsurancePlanTraitMapping = map[string]func(*schema.HealthInsurancePlanTrait, []string){}
var customHealthInsurancePlanTraitMapping = map[string]func(*schema.HealthInsurancePlanTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HealthInsurancePlan", func(ctx jsonld.Context) (interface{}) {
		return NewHealthInsurancePlanFromContext(ctx)
	})



	basicHealthInsurancePlanTraitMapping["http://schema.org/healthPlanDrugTier"] = func(x *schema.HealthInsurancePlanTrait, s []string) {
		x.HealthPlanDrugTier = s[0]
	}


	basicHealthInsurancePlanTraitMapping["http://schema.org/benefitsSummaryUrl"] = func(x *schema.HealthInsurancePlanTrait, s []string) {
		x.BenefitsSummaryUrl = s[0]
	}


	basicHealthInsurancePlanTraitMapping["http://schema.org/healthPlanDrugOption"] = func(x *schema.HealthInsurancePlanTrait, s []string) {
		x.HealthPlanDrugOption = s[0]
	}


	basicHealthInsurancePlanTraitMapping["http://schema.org/healthPlanId"] = func(x *schema.HealthInsurancePlanTrait, s []string) {
		x.HealthPlanId = s[0]
	}


	basicHealthInsurancePlanTraitMapping["http://schema.org/healthPlanMarketingUrl"] = func(x *schema.HealthInsurancePlanTrait, s []string) {
		x.HealthPlanMarketingUrl = s[0]
	}





	customHealthInsurancePlanTraitMapping["http://schema.org/contactPoint"] = func(x *schema.HealthInsurancePlanTrait, ctx jsonld.Context, s string){
		var y schema.ContactPoint
		if strings.HasPrefix(s, "_:") {
			y = NewContactPointFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewContactPoint()
			y.Id = s
		}

		x.ContactPoint = &y

		return
	}

	customHealthInsurancePlanTraitMapping["http://schema.org/includesHealthPlanFormulary"] = func(x *schema.HealthInsurancePlanTrait, ctx jsonld.Context, s string){
		var y schema.HealthPlanFormulary
		if strings.HasPrefix(s, "_:") {
			y = NewHealthPlanFormularyFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewHealthPlanFormulary()
			y.Id = s
		}

		x.IncludesHealthPlanFormulary = &y

		return
	}

	customHealthInsurancePlanTraitMapping["http://schema.org/includesHealthPlanNetwork"] = func(x *schema.HealthInsurancePlanTrait, ctx jsonld.Context, s string){
		var y schema.HealthPlanNetwork
		if strings.HasPrefix(s, "_:") {
			y = NewHealthPlanNetworkFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewHealthPlanNetwork()
			y.Id = s
		}

		x.IncludesHealthPlanNetwork = &y

		return
	}

	customHealthInsurancePlanTraitMapping["http://schema.org/usesHealthPlanIdStandard"] = func(x *schema.HealthInsurancePlanTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.UsesHealthPlanIdStandard, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.UsesHealthPlanIdStandard = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.UsesHealthPlanIdStandard = s
		}
	}
}

func NewHealthInsurancePlanFromContext(ctx jsonld.Context) (x schema.HealthInsurancePlan) {
	x.Type = "http://schema.org/HealthInsurancePlan"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHealthInsurancePlanTrait(ctx, &x.HealthInsurancePlanTrait)
	MapCustomToHealthInsurancePlanTrait(ctx, &x.HealthInsurancePlanTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHealthInsurancePlanTrait(ctx jsonld.Context, x *schema.HealthInsurancePlanTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHealthInsurancePlanTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHealthInsurancePlanTrait(ctx jsonld.Context, x *schema.HealthInsurancePlanTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHealthInsurancePlanTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}