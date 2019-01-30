package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsHealthPlanCostSharingSpecification strings.Replacer
var strconvHealthPlanCostSharingSpecification strconv.NumError

var basicHealthPlanCostSharingSpecificationTraitMapping = map[string]func(*schema.HealthPlanCostSharingSpecificationTrait, []string){}
var customHealthPlanCostSharingSpecificationTraitMapping = map[string]func(*schema.HealthPlanCostSharingSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/HealthPlanCostSharingSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewHealthPlanCostSharingSpecificationFromContext(ctx)
	})


	basicHealthPlanCostSharingSpecificationTraitMapping["http://schema.org/healthPlanCoinsuranceOption"] = func(x *schema.HealthPlanCostSharingSpecificationTrait, s []string) {
		x.HealthPlanCoinsuranceOption = s[0]
	}


	basicHealthPlanCostSharingSpecificationTraitMapping["http://schema.org/healthPlanCoinsuranceRate"] = func(x *schema.HealthPlanCostSharingSpecificationTrait, s []string) {
		var err error
		x.HealthPlanCoinsuranceRate, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}



	basicHealthPlanCostSharingSpecificationTraitMapping["http://schema.org/healthPlanCopayOption"] = func(x *schema.HealthPlanCostSharingSpecificationTrait, s []string) {
		x.HealthPlanCopayOption = s[0]
	}


	basicHealthPlanCostSharingSpecificationTraitMapping["http://schema.org/healthPlanPharmacyCategory"] = func(x *schema.HealthPlanCostSharingSpecificationTrait, s []string) {
		x.HealthPlanPharmacyCategory = s[0]
	}


	customHealthPlanCostSharingSpecificationTraitMapping["http://schema.org/healthPlanCopay"] = func(x *schema.HealthPlanCostSharingSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.PriceSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewPriceSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPriceSpecification()
			y.Id = s
		}

		x.HealthPlanCopay = &y

		return
	}
}

func NewHealthPlanCostSharingSpecificationFromContext(ctx jsonld.Context) (x schema.HealthPlanCostSharingSpecification) {
	x.Type = "http://schema.org/HealthPlanCostSharingSpecification"
	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToHealthPlanCostSharingSpecificationTrait(ctx, &x.HealthPlanCostSharingSpecificationTrait)
	MapCustomToHealthPlanCostSharingSpecificationTrait(ctx, &x.HealthPlanCostSharingSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToHealthPlanCostSharingSpecificationTrait(ctx jsonld.Context, x *schema.HealthPlanCostSharingSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicHealthPlanCostSharingSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToHealthPlanCostSharingSpecificationTrait(ctx jsonld.Context, x *schema.HealthPlanCostSharingSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customHealthPlanCostSharingSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}