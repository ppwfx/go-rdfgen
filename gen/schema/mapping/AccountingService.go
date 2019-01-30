package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsAccountingService strings.Replacer
var strconvAccountingService strconv.NumError

var basicAccountingServiceTraitMapping = map[string]func(*schema.AccountingServiceTrait, []string){}
var customAccountingServiceTraitMapping = map[string]func(*schema.AccountingServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/AccountingService", func(ctx jsonld.Context) (interface{}) {
		return NewAccountingServiceFromContext(ctx)
	})

}

func NewAccountingServiceFromContext(ctx jsonld.Context) (x schema.AccountingService) {
	x.Type = "http://schema.org/AccountingService"
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


	MapBasicToAccountingServiceTrait(ctx, &x.AccountingServiceTrait)
	MapCustomToAccountingServiceTrait(ctx, &x.AccountingServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToAccountingServiceTrait(ctx jsonld.Context, x *schema.AccountingServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicAccountingServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToAccountingServiceTrait(ctx jsonld.Context, x *schema.AccountingServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customAccountingServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}