package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHealthPlanNetwork strings.Replacer
var strconvHealthPlanNetwork strconv.NumError

var basicHealthPlanNetworkTraitMapping = map[string]func(*schema.HealthPlanNetworkTrait, []string){}
var customHealthPlanNetworkTraitMapping = map[string]func(*schema.HealthPlanNetworkTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HealthPlanNetwork", func(ctx jsonld.Context) (interface{}) {
		return NewHealthPlanNetworkFromContext(ctx)
	})


	basicHealthPlanNetworkTraitMapping["http://schema.org/healthPlanCostSharing"] = func(x *schema.HealthPlanNetworkTrait, s []string) {
		var err error
		x.HealthPlanCostSharing, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	basicHealthPlanNetworkTraitMapping["http://schema.org/healthPlanNetworkId"] = func(x *schema.HealthPlanNetworkTrait, s []string) {
		x.HealthPlanNetworkId = s[0]
	}


	basicHealthPlanNetworkTraitMapping["http://schema.org/healthPlanNetworkTier"] = func(x *schema.HealthPlanNetworkTrait, s []string) {
		x.HealthPlanNetworkTier = s[0]
	}

}

func NewHealthPlanNetworkFromContext(ctx jsonld.Context) (x schema.HealthPlanNetwork) {
	x.Type = "http://schema.org/HealthPlanNetwork"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHealthPlanNetworkTrait(ctx, &x.HealthPlanNetworkTrait)
	MapCustomToHealthPlanNetworkTrait(ctx, &x.HealthPlanNetworkTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHealthPlanNetworkTrait(ctx jsonld.Context, x *schema.HealthPlanNetworkTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHealthPlanNetworkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHealthPlanNetworkTrait(ctx jsonld.Context, x *schema.HealthPlanNetworkTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHealthPlanNetworkTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}