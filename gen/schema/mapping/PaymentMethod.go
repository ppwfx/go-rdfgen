package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPaymentMethod strings.Replacer
var strconvPaymentMethod strconv.NumError

var basicPaymentMethodTraitMapping = map[string]func(*schema.PaymentMethodTrait, []string){}
var customPaymentMethodTraitMapping = map[string]func(*schema.PaymentMethodTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PaymentMethod", func(ctx jsonld.Context) (interface{}) {
		return NewPaymentMethodFromContext(ctx)
	})

}

func NewPaymentMethodFromContext(ctx jsonld.Context) (x schema.PaymentMethod) {
	x.Type = "http://schema.org/PaymentMethod"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPaymentMethodTrait(ctx, &x.PaymentMethodTrait)
	MapCustomToPaymentMethodTrait(ctx, &x.PaymentMethodTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPaymentMethodTrait(ctx jsonld.Context, x *schema.PaymentMethodTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPaymentMethodTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPaymentMethodTrait(ctx jsonld.Context, x *schema.PaymentMethodTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPaymentMethodTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}