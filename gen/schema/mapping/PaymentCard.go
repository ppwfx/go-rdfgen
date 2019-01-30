package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPaymentCard strings.Replacer
var strconvPaymentCard strconv.NumError

var basicPaymentCardTraitMapping = map[string]func(*schema.PaymentCardTrait, []string){}
var customPaymentCardTraitMapping = map[string]func(*schema.PaymentCardTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PaymentCard", func(ctx jsonld.Context) (interface{}) {
		return NewPaymentCardFromContext(ctx)
	})




	basicPaymentCardTraitMapping["http://schema.org/contactlessPayment"] = func(x *schema.PaymentCardTrait, s []string) {
		var err error
		x.ContactlessPayment, err = strconv.ParseBool(s[0])
		if err != nil {
			println(err.Error())
		}
	}


	customPaymentCardTraitMapping["http://schema.org/floorLimit"] = func(x *schema.PaymentCardTrait, ctx jsonld.Context, s string){
		var y schema.MonetaryAmount
		if strings.HasPrefix(s, "_:") {
			y = NewMonetaryAmountFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewMonetaryAmount()
			y.Id = s
		}

		x.FloorLimit = &y

		return
	}

	customPaymentCardTraitMapping["http://schema.org/cashBack"] = func(x *schema.PaymentCardTrait, ctx jsonld.Context, s string){
		if strings.HasPrefix(s, "_:") {
			_, ok := ctx.Subjects[s]
			if ok {
				var err error
				x.CashBack, err = jsonld.FillByRdfType(ctx.SetCurrent(s))
				if err != nil {
					println(err.Error())
				}
				delete(ctx.Subjects, s)
			}

			_, ok = ctx.Decoded[s]
			if ok {
				x.CashBack = ctx.Decoded[s]
				delete(ctx.Decoded, s)
			}

		} else {
			x.CashBack = s
		}
	}
}

func NewPaymentCardFromContext(ctx jsonld.Context) (x schema.PaymentCard) {
	x.Type = "http://schema.org/PaymentCard"
	MapBasicToPaymentMethodTrait(ctx, &x.PaymentMethodTrait)
	MapCustomToPaymentMethodTrait(ctx, &x.PaymentMethodTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)

	MapBasicToFinancialProductTrait(ctx, &x.FinancialProductTrait)
	MapCustomToFinancialProductTrait(ctx, &x.FinancialProductTrait)

	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)


	MapBasicToPaymentCardTrait(ctx, &x.PaymentCardTrait)
	MapCustomToPaymentCardTrait(ctx, &x.PaymentCardTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPaymentCardTrait(ctx jsonld.Context, x *schema.PaymentCardTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPaymentCardTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPaymentCardTrait(ctx jsonld.Context, x *schema.PaymentCardTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPaymentCardTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}