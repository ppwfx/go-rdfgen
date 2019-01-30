package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCreditCard strings.Replacer
var strconvCreditCard strconv.NumError

var basicCreditCardTraitMapping = map[string]func(*schema.CreditCardTrait, []string){}
var customCreditCardTraitMapping = map[string]func(*schema.CreditCardTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CreditCard", func(ctx jsonld.Context) (interface{}) {
		return NewCreditCardFromContext(ctx)
	})



	customCreditCardTraitMapping["http://schema.org/monthlyMinimumRepaymentAmount"] = func(x *schema.CreditCardTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.MonthlyMinimumRepaymentAmount, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.MonthlyMinimumRepaymentAmount = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.MonthlyMinimumRepaymentAmount = s
		}
	}
}

func NewCreditCardFromContext(ctx jsonld.Context) (x schema.CreditCard) {
	x.Type = "http://schema.org/CreditCard"
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

	MapBasicToPaymentCardTrait(ctx, &x.PaymentCardTrait)
	MapCustomToPaymentCardTrait(ctx, &x.PaymentCardTrait)

	MapBasicToPaymentMethodTrait(ctx, &x.PaymentMethodTrait)
	MapCustomToPaymentMethodTrait(ctx, &x.PaymentMethodTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)


	MapBasicToCreditCardTrait(ctx, &x.CreditCardTrait)
	MapCustomToCreditCardTrait(ctx, &x.CreditCardTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCreditCardTrait(ctx jsonld.Context, x *schema.CreditCardTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCreditCardTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCreditCardTrait(ctx jsonld.Context, x *schema.CreditCardTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCreditCardTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}