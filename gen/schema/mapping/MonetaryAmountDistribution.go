package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMonetaryAmountDistribution strings.Replacer
var strconvMonetaryAmountDistribution strconv.NumError

var basicMonetaryAmountDistributionTraitMapping = map[string]func(*schema.MonetaryAmountDistributionTrait, []string){}
var customMonetaryAmountDistributionTraitMapping = map[string]func(*schema.MonetaryAmountDistributionTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MonetaryAmountDistribution", func(ctx jsonld.Context) (interface{}) {
		return NewMonetaryAmountDistributionFromContext(ctx)
	})


	basicMonetaryAmountDistributionTraitMapping["http://schema.org/currency"] = func(x *schema.MonetaryAmountDistributionTrait, s []string) {
		x.Currency = s[0]
	}

}

func NewMonetaryAmountDistributionFromContext(ctx jsonld.Context) (x schema.MonetaryAmountDistribution) {
	x.Type = "http://schema.org/MonetaryAmountDistribution"
	MapBasicToQuantitativeValueTrait(ctx, &x.QuantitativeValueTrait)
	MapCustomToQuantitativeValueTrait(ctx, &x.QuantitativeValueTrait)

	MapBasicToStructuredValueTrait(ctx, &x.StructuredValueTrait)
	MapCustomToStructuredValueTrait(ctx, &x.StructuredValueTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMonetaryAmountDistributionTrait(ctx, &x.MonetaryAmountDistributionTrait)
	MapCustomToMonetaryAmountDistributionTrait(ctx, &x.MonetaryAmountDistributionTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMonetaryAmountDistributionTrait(ctx jsonld.Context, x *schema.MonetaryAmountDistributionTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMonetaryAmountDistributionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMonetaryAmountDistributionTrait(ctx jsonld.Context, x *schema.MonetaryAmountDistributionTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMonetaryAmountDistributionTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}