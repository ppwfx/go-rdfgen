package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLoanOrCredit strings.Replacer
var strconvLoanOrCredit strconv.NumError

var basicLoanOrCreditTraitMapping = map[string]func(*schema.LoanOrCreditTrait, []string){}
var customLoanOrCreditTraitMapping = map[string]func(*schema.LoanOrCreditTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LoanOrCredit", func(ctx jsonld.Context) (interface{}) {
		return NewLoanOrCreditFromContext(ctx)
	})


	basicLoanOrCreditTraitMapping["http://schema.org/currency"] = func(x *schema.LoanOrCreditTrait, s []string) {
		x.Currency = s[0]
	}








	basicLoanOrCreditTraitMapping["http://schema.org/recourseLoan"] = func(x *schema.LoanOrCreditTrait, s []string) {
		var err error
		x.RecourseLoan, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	basicLoanOrCreditTraitMapping["http://schema.org/renegotiableLoan"] = func(x *schema.LoanOrCreditTrait, s []string) {
		var err error
		x.RenegotiableLoan, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	customLoanOrCreditTraitMapping["http://schema.org/amount"] = func(x *schema.LoanOrCreditTrait, ctx jsonld.Context, s string){
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

	customLoanOrCreditTraitMapping["http://schema.org/requiredCollateral"] = func(x *schema.LoanOrCreditTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.RequiredCollateral, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.RequiredCollateral = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.RequiredCollateral = s
		}
	}

	customLoanOrCreditTraitMapping["http://schema.org/loanType"] = func(x *schema.LoanOrCreditTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.LoanType, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.LoanType = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.LoanType = s
		}
	}

	customLoanOrCreditTraitMapping["http://schema.org/gracePeriod"] = func(x *schema.LoanOrCreditTrait, ctx jsonld.Context, s string){
		var y schema.Duration
		if strings.HasPrefix(s, "_:") {
			y = NewDurationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDuration()
			y.Id = s
		}

		x.GracePeriod = &y

		return
	}

	customLoanOrCreditTraitMapping["http://schema.org/loanRepaymentForm"] = func(x *schema.LoanOrCreditTrait, ctx jsonld.Context, s string){
		var y schema.RepaymentSpecification
		if strings.HasPrefix(s, "_:") {
			y = NewRepaymentSpecificationFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewRepaymentSpecification()
			y.Id = s
		}

		x.LoanRepaymentForm = &y

		return
	}

	customLoanOrCreditTraitMapping["http://schema.org/loanTerm"] = func(x *schema.LoanOrCreditTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.LoanTerm = &y

		return
	}
}

func NewLoanOrCreditFromContext(ctx jsonld.Context) (x schema.LoanOrCredit) {
	x.Type = "http://schema.org/LoanOrCredit"
	MapBasicToFinancialProductTrait(ctx, &x.FinancialProductTrait)
	MapCustomToFinancialProductTrait(ctx, &x.FinancialProductTrait)

	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLoanOrCreditTrait(ctx, &x.LoanOrCreditTrait)
	MapCustomToLoanOrCreditTrait(ctx, &x.LoanOrCreditTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLoanOrCreditTrait(ctx jsonld.Context, x *schema.LoanOrCreditTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLoanOrCreditTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLoanOrCreditTrait(ctx jsonld.Context, x *schema.LoanOrCreditTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLoanOrCreditTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}