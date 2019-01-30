package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsRepaymentSpecification strings.Replacer
var strconvRepaymentSpecification strconv.NumError

var basicRepaymentSpecificationTraitMapping = map[string]func(*schema.RepaymentSpecificationTrait, []string){}
var customRepaymentSpecificationTraitMapping = map[string]func(*schema.RepaymentSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/RepaymentSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewRepaymentSpecificationFromContext(ctx)
	})





	basicRepaymentSpecificationTraitMapping["http://schema.org/loanPaymentFrequency"] = func(x *schema.RepaymentSpecificationTrait, s []string) {
		var err error
		x.LoanPaymentFrequency, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	basicRepaymentSpecificationTraitMapping["http://schema.org/numberOfLoanPayments"] = func(x *schema.RepaymentSpecificationTrait, s []string) {
		var err error
		x.NumberOfLoanPayments, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			println(err.Error())
		}
	}


	customRepaymentSpecificationTraitMapping["http://schema.org/downPayment"] = func(x *schema.RepaymentSpecificationTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.DownPayment, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.DownPayment = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.DownPayment = s
		}
	}

	customRepaymentSpecificationTraitMapping["http://schema.org/earlyPrepaymentPenalty"] = func(x *schema.RepaymentSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.MonetaryAmount
		if strings.HasPrefix(s, "_:") {
			y = NewMonetaryAmountFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMonetaryAmount()
			y.Id = s
		}

		x.EarlyPrepaymentPenalty = &y

		return
	}

	customRepaymentSpecificationTraitMapping["http://schema.org/loanPaymentAmount"] = func(x *schema.RepaymentSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.MonetaryAmount
		if strings.HasPrefix(s, "_:") {
			y = NewMonetaryAmountFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMonetaryAmount()
			y.Id = s
		}

		x.LoanPaymentAmount = &y

		return
	}
}

func NewRepaymentSpecificationFromContext(ctx jsonld.Context) (x schema.RepaymentSpecification) {
	x.Type = "http://schema.org/RepaymentSpecification"
	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToRepaymentSpecificationTrait(ctx, &x.RepaymentSpecificationTrait)
	MapCustomToRepaymentSpecificationTrait(ctx, &x.RepaymentSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToRepaymentSpecificationTrait(ctx jsonld.Context, x *schema.RepaymentSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicRepaymentSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToRepaymentSpecificationTrait(ctx jsonld.Context, x *schema.RepaymentSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customRepaymentSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}