package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCurrencyConversionService strings.Replacer
var strconvCurrencyConversionService strconv.NumError

var basicCurrencyConversionServiceTraitMapping = map[string]func(*schema.CurrencyConversionServiceTrait, []string){}
var customCurrencyConversionServiceTraitMapping = map[string]func(*schema.CurrencyConversionServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CurrencyConversionService", func(ctx jsonld.Context) (interface{}) {
		return NewCurrencyConversionServiceFromContext(ctx)
	})

}

func NewCurrencyConversionServiceFromContext(ctx jsonld.Context) (x schema.CurrencyConversionService) {
	x.Type = "http://schema.org/CurrencyConversionService"
	MapBasicToFinancialProductTrait(ctx, &x.FinancialProductTrait)
	MapCustomToFinancialProductTrait(ctx, &x.FinancialProductTrait)

	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCurrencyConversionServiceTrait(ctx, &x.CurrencyConversionServiceTrait)
	MapCustomToCurrencyConversionServiceTrait(ctx, &x.CurrencyConversionServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCurrencyConversionServiceTrait(ctx jsonld.Context, x *schema.CurrencyConversionServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCurrencyConversionServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCurrencyConversionServiceTrait(ctx jsonld.Context, x *schema.CurrencyConversionServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCurrencyConversionServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}