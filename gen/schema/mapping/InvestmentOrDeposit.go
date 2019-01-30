package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsInvestmentOrDeposit strings.Replacer
var strconvInvestmentOrDeposit strconv.NumError

var basicInvestmentOrDepositTraitMapping = map[string]func(*schema.InvestmentOrDepositTrait, []string){}
var customInvestmentOrDepositTraitMapping = map[string]func(*schema.InvestmentOrDepositTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/InvestmentOrDeposit", func(ctx jsonld.Context) (interface{}) {
		return NewInvestmentOrDepositFromContext(ctx)
	})



	customInvestmentOrDepositTraitMapping["http://schema.org/amount"] = func(x *schema.InvestmentOrDepositTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.Amount, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.Amount = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.Amount = s
		}
	}
}

func NewInvestmentOrDepositFromContext(ctx jsonld.Context) (x schema.InvestmentOrDeposit) {
	x.Type = "http://schema.org/InvestmentOrDeposit"
	MapBasicToFinancialProductTrait(ctx, &x.FinancialProductTrait)
	MapCustomToFinancialProductTrait(ctx, &x.FinancialProductTrait)

	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToInvestmentOrDepositTrait(ctx, &x.InvestmentOrDepositTrait)
	MapCustomToInvestmentOrDepositTrait(ctx, &x.InvestmentOrDepositTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToInvestmentOrDepositTrait(ctx jsonld.Context, x *schema.InvestmentOrDepositTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicInvestmentOrDepositTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToInvestmentOrDepositTrait(ctx jsonld.Context, x *schema.InvestmentOrDepositTrait) () {
	for k, v := range ctx.Current {
		f, ok := customInvestmentOrDepositTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}