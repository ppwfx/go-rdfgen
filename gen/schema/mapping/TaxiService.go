package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTaxiService strings.Replacer
var strconvTaxiService strconv.NumError

var basicTaxiServiceTraitMapping = map[string]func(*schema.TaxiServiceTrait, []string){}
var customTaxiServiceTraitMapping = map[string]func(*schema.TaxiServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/TaxiService", func(ctx jsonld.Context) (interface{}) {
		return NewTaxiServiceFromContext(ctx)
	})

}

func NewTaxiServiceFromContext(ctx jsonld.Context) (x schema.TaxiService) {
	x.Type = "http://schema.org/TaxiService"
	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTaxiServiceTrait(ctx, &x.TaxiServiceTrait)
	MapCustomToTaxiServiceTrait(ctx, &x.TaxiServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTaxiServiceTrait(ctx jsonld.Context, x *schema.TaxiServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTaxiServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTaxiServiceTrait(ctx jsonld.Context, x *schema.TaxiServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTaxiServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}