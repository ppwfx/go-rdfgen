package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMortgageLoan strings.Replacer
var strconvMortgageLoan strconv.NumError

var basicMortgageLoanTraitMapping = map[string]func(*schema.MortgageLoanTrait, []string){}
var customMortgageLoanTraitMapping = map[string]func(*schema.MortgageLoanTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MortgageLoan", func(ctx jsonld.Context) (interface{}) {
		return NewMortgageLoanFromContext(ctx)
	})



	basicMortgageLoanTraitMapping["http://schema.org/domiciledMortgage"] = func(x *schema.MortgageLoanTrait, s []string) {
		var err error
		x.DomiciledMortgage, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	customMortgageLoanTraitMapping["http://schema.org/loanMortgageMandateAmount"] = func(x *schema.MortgageLoanTrait, ctx jsonld.Context, s string){
		var y schema.MonetaryAmount
		if strings.HasPrefix(s, "_:") {
			y = NewMonetaryAmountFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMonetaryAmount()
			y.Id = s
		}

		x.LoanMortgageMandateAmount = &y

		return
	}
}

func NewMortgageLoanFromContext(ctx jsonld.Context) (x schema.MortgageLoan) {
	x.Type = "http://schema.org/MortgageLoan"
	MapBasicToLoanOrCreditTrait(ctx, &x.LoanOrCreditTrait)
	MapCustomToLoanOrCreditTrait(ctx, &x.LoanOrCreditTrait)

	MapBasicToFinancialProductTrait(ctx, &x.FinancialProductTrait)
	MapCustomToFinancialProductTrait(ctx, &x.FinancialProductTrait)

	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMortgageLoanTrait(ctx, &x.MortgageLoanTrait)
	MapCustomToMortgageLoanTrait(ctx, &x.MortgageLoanTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMortgageLoanTrait(ctx jsonld.Context, x *schema.MortgageLoanTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMortgageLoanTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMortgageLoanTrait(ctx jsonld.Context, x *schema.MortgageLoanTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMortgageLoanTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}