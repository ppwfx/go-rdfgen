package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInvestmentFund strings.Replacer
var strconvInvestmentFund strconv.NumError

var basicInvestmentFundTraitMapping = map[string]func(*schema.InvestmentFundTrait, []string){}
var customInvestmentFundTraitMapping = map[string]func(*schema.InvestmentFundTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InvestmentFund", func(ctx jsonld.Context) (interface{}) {
		return NewInvestmentFundFromContext(ctx)
	})

}

func NewInvestmentFundFromContext(ctx jsonld.Context) (x schema.InvestmentFund) {
	x.Type = "http://schema.org/InvestmentFund"
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


	MapBasicToInvestmentFundTrait(ctx, &x.InvestmentFundTrait)
	MapCustomToInvestmentFundTrait(ctx, &x.InvestmentFundTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInvestmentFundTrait(ctx jsonld.Context, x *schema.InvestmentFundTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInvestmentFundTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInvestmentFundTrait(ctx jsonld.Context, x *schema.InvestmentFundTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInvestmentFundTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}