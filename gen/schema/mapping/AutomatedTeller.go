package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAutomatedTeller strings.Replacer
var strconvAutomatedTeller strconv.NumError

var basicAutomatedTellerTraitMapping = map[string]func(*schema.AutomatedTellerTrait, []string){}
var customAutomatedTellerTraitMapping = map[string]func(*schema.AutomatedTellerTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AutomatedTeller", func(ctx jsonld.Context) (interface{}) {
		return NewAutomatedTellerFromContext(ctx)
	})

}

func NewAutomatedTellerFromContext(ctx jsonld.Context) (x schema.AutomatedTeller) {
	x.Type = "http://schema.org/AutomatedTeller"
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


	MapBasicToAutomatedTellerTrait(ctx, &x.AutomatedTellerTrait)
	MapCustomToAutomatedTellerTrait(ctx, &x.AutomatedTellerTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAutomatedTellerTrait(ctx jsonld.Context, x *schema.AutomatedTellerTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAutomatedTellerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAutomatedTellerTrait(ctx jsonld.Context, x *schema.AutomatedTellerTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAutomatedTellerTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}