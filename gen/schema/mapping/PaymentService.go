package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPaymentService strings.Replacer
var strconvPaymentService strconv.NumError

var basicPaymentServiceTraitMapping = map[string]func(*schema.PaymentServiceTrait, []string){}
var customPaymentServiceTraitMapping = map[string]func(*schema.PaymentServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PaymentService", func(ctx jsonld.Context) (interface{}) {
		return NewPaymentServiceFromContext(ctx)
	})

}

func NewPaymentServiceFromContext(ctx jsonld.Context) (x schema.PaymentService) {
	x.Type = "http://schema.org/PaymentService"
	MapBasicToFinancialProductTrait(ctx, &x.FinancialProductTrait)
	MapCustomToFinancialProductTrait(ctx, &x.FinancialProductTrait)

	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPaymentServiceTrait(ctx, &x.PaymentServiceTrait)
	MapCustomToPaymentServiceTrait(ctx, &x.PaymentServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPaymentServiceTrait(ctx jsonld.Context, x *schema.PaymentServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPaymentServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPaymentServiceTrait(ctx jsonld.Context, x *schema.PaymentServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPaymentServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}