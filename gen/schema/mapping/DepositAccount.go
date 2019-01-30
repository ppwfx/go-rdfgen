package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDepositAccount strings.Replacer
var strconvDepositAccount strconv.NumError

var basicDepositAccountTraitMapping = map[string]func(*schema.DepositAccountTrait, []string){}
var customDepositAccountTraitMapping = map[string]func(*schema.DepositAccountTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DepositAccount", func(ctx jsonld.Context) (interface{}) {
		return NewDepositAccountFromContext(ctx)
	})

}

func NewDepositAccountFromContext(ctx jsonld.Context) (x schema.DepositAccount) {
	x.Type = "http://schema.org/DepositAccount"
	MapBasicToInvestmentOrDepositTrait(ctx, &x.InvestmentOrDepositTrait)
	MapCustomToInvestmentOrDepositTrait(ctx, &x.InvestmentOrDepositTrait)

	MapBasicToFinancialProductTrait(ctx, &x.FinancialProductTrait)
	MapCustomToFinancialProductTrait(ctx, &x.FinancialProductTrait)

	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToBankAccountTrait(ctx, &x.BankAccountTrait)
	MapCustomToBankAccountTrait(ctx, &x.BankAccountTrait)


	MapBasicToDepositAccountTrait(ctx, &x.DepositAccountTrait)
	MapCustomToDepositAccountTrait(ctx, &x.DepositAccountTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDepositAccountTrait(ctx jsonld.Context, x *schema.DepositAccountTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDepositAccountTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDepositAccountTrait(ctx jsonld.Context, x *schema.DepositAccountTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDepositAccountTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}