package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInsuranceAgency strings.Replacer
var strconvInsuranceAgency strconv.NumError

var basicInsuranceAgencyTraitMapping = map[string]func(*schema.InsuranceAgencyTrait, []string){}
var customInsuranceAgencyTraitMapping = map[string]func(*schema.InsuranceAgencyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InsuranceAgency", func(ctx jsonld.Context) (interface{}) {
		return NewInsuranceAgencyFromContext(ctx)
	})

}

func NewInsuranceAgencyFromContext(ctx jsonld.Context) (x schema.InsuranceAgency) {
	x.Type = "http://schema.org/InsuranceAgency"
	MapBasicToFinancialServiceTrait(ctx, &x.FinancialServiceTrait)
	MapCustomToFinancialServiceTrait(ctx, &x.FinancialServiceTrait)

	MapBasicToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)
	MapCustomToLocalBusinessTrait(ctx, &x.LocalBusinessTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToOrganizationTrait(ctx, &x.OrganizationTrait)
	MapCustomToOrganizationTrait(ctx, &x.OrganizationTrait)


	MapBasicToInsuranceAgencyTrait(ctx, &x.InsuranceAgencyTrait)
	MapCustomToInsuranceAgencyTrait(ctx, &x.InsuranceAgencyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInsuranceAgencyTrait(ctx jsonld.Context, x *schema.InsuranceAgencyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInsuranceAgencyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInsuranceAgencyTrait(ctx jsonld.Context, x *schema.InsuranceAgencyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInsuranceAgencyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}