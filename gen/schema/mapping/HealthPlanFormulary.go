package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHealthPlanFormulary strings.Replacer
var strconvHealthPlanFormulary strconv.NumError

var basicHealthPlanFormularyTraitMapping = map[string]func(*schema.HealthPlanFormularyTrait, []string){}
var customHealthPlanFormularyTraitMapping = map[string]func(*schema.HealthPlanFormularyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HealthPlanFormulary", func(ctx jsonld.Context) (interface{}) {
		return NewHealthPlanFormularyFromContext(ctx)
	})


	basicHealthPlanFormularyTraitMapping["http://schema.org/healthPlanCostSharing"] = func(x *schema.HealthPlanFormularyTrait, s []string) {
		var err error
		x.HealthPlanCostSharing, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	basicHealthPlanFormularyTraitMapping["http://schema.org/healthPlanDrugTier"] = func(x *schema.HealthPlanFormularyTrait, s []string) {
		x.HealthPlanDrugTier = s[0]
	}


	basicHealthPlanFormularyTraitMapping["http://schema.org/offersPrescriptionByMail"] = func(x *schema.HealthPlanFormularyTrait, s []string) {
		var err error
		x.OffersPrescriptionByMail, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}

}

func NewHealthPlanFormularyFromContext(ctx jsonld.Context) (x schema.HealthPlanFormulary) {
	x.Type = "http://schema.org/HealthPlanFormulary"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHealthPlanFormularyTrait(ctx, &x.HealthPlanFormularyTrait)
	MapCustomToHealthPlanFormularyTrait(ctx, &x.HealthPlanFormularyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHealthPlanFormularyTrait(ctx jsonld.Context, x *schema.HealthPlanFormularyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHealthPlanFormularyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHealthPlanFormularyTrait(ctx jsonld.Context, x *schema.HealthPlanFormularyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHealthPlanFormularyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}