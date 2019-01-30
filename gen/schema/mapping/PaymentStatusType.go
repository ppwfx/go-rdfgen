package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPaymentStatusType strings.Replacer
var strconvPaymentStatusType strconv.NumError

var basicPaymentStatusTypeTraitMapping = map[string]func(*schema.PaymentStatusTypeTrait, []string){}
var customPaymentStatusTypeTraitMapping = map[string]func(*schema.PaymentStatusTypeTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PaymentStatusType", func(ctx jsonld.Context) (interface{}) {
		return NewPaymentStatusTypeFromContext(ctx)
	})

}

func NewPaymentStatusTypeFromContext(ctx jsonld.Context) (x schema.PaymentStatusType) {
	x.Type = "http://schema.org/PaymentStatusType"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPaymentStatusTypeTrait(ctx, &x.PaymentStatusTypeTrait)
	MapCustomToPaymentStatusTypeTrait(ctx, &x.PaymentStatusTypeTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPaymentStatusTypeTrait(ctx jsonld.Context, x *schema.PaymentStatusTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPaymentStatusTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPaymentStatusTypeTrait(ctx jsonld.Context, x *schema.PaymentStatusTypeTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPaymentStatusTypeTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}