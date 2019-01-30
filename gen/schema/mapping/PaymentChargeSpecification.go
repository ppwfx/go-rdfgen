package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPaymentChargeSpecification strings.Replacer
var strconvPaymentChargeSpecification strconv.NumError

var basicPaymentChargeSpecificationTraitMapping = map[string]func(*schema.PaymentChargeSpecificationTrait, []string){}
var customPaymentChargeSpecificationTraitMapping = map[string]func(*schema.PaymentChargeSpecificationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/PaymentChargeSpecification", func(ctx jsonld.Context) (interface{}) {
		return NewPaymentChargeSpecificationFromContext(ctx)
	})




	customPaymentChargeSpecificationTraitMapping["http://schema.org/appliesToDeliveryMethod"] = func(x *schema.PaymentChargeSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.DeliveryMethod
		if strings.HasPrefix(s, "_:") {
			y = NewDeliveryMethodFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewDeliveryMethod()
			y.Id = s
		}

		x.AppliesToDeliveryMethod = &y

		return
	}

	customPaymentChargeSpecificationTraitMapping["http://schema.org/appliesToPaymentMethod"] = func(x *schema.PaymentChargeSpecificationTrait, ctx jsonld.Context, s string){
		var y schema.PaymentMethod
		if strings.HasPrefix(s, "_:") {
			y = NewPaymentMethodFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewPaymentMethod()
			y.Id = s
		}

		x.AppliesToPaymentMethod = &y

		return
	}
}

func NewPaymentChargeSpecificationFromContext(ctx jsonld.Context) (x schema.PaymentChargeSpecification) {
	x.Type = "http://schema.org/PaymentChargeSpecification"
	MapBasicToPriceSpecificationTrait(ctx, &x.PriceSpecificationTrait)
	MapCustomToPriceSpecificationTrait(ctx, &x.PriceSpecificationTrait)

	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPaymentChargeSpecificationTrait(ctx, &x.PaymentChargeSpecificationTrait)
	MapCustomToPaymentChargeSpecificationTrait(ctx, &x.PaymentChargeSpecificationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPaymentChargeSpecificationTrait(ctx jsonld.Context, x *schema.PaymentChargeSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPaymentChargeSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPaymentChargeSpecificationTrait(ctx jsonld.Context, x *schema.PaymentChargeSpecificationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPaymentChargeSpecificationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}