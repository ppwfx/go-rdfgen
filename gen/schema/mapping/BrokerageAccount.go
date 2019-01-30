package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBrokerageAccount strings.Replacer
var strconvBrokerageAccount strconv.NumError

var basicBrokerageAccountTraitMapping = map[string]func(*schema.BrokerageAccountTrait, []string){}
var customBrokerageAccountTraitMapping = map[string]func(*schema.BrokerageAccountTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BrokerageAccount", func(ctx jsonld.Context) (interface{}) {
		return NewBrokerageAccountFromContext(ctx)
	})

}

func NewBrokerageAccountFromContext(ctx jsonld.Context) (x schema.BrokerageAccount) {
	x.Type = "http://schema.org/BrokerageAccount"
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


	MapBasicToBrokerageAccountTrait(ctx, &x.BrokerageAccountTrait)
	MapCustomToBrokerageAccountTrait(ctx, &x.BrokerageAccountTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBrokerageAccountTrait(ctx jsonld.Context, x *schema.BrokerageAccountTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBrokerageAccountTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBrokerageAccountTrait(ctx jsonld.Context, x *schema.BrokerageAccountTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBrokerageAccountTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}