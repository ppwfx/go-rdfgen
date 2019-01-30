package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsTaxi strings.Replacer
var strconvTaxi strconv.NumError

var basicTaxiTraitMapping = map[string]func(*schema.TaxiTrait, []string){}
var customTaxiTraitMapping = map[string]func(*schema.TaxiTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Taxi", func(ctx jsonld.Context) (interface{}) {
		return NewTaxiFromContext(ctx)
	})

}

func NewTaxiFromContext(ctx jsonld.Context) (x schema.Taxi) {
	x.Type = "http://schema.org/Taxi"
	MapBasicToServiceTrait(ctx, &x.ServiceTrait)
	MapCustomToServiceTrait(ctx, &x.ServiceTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToTaxiTrait(ctx, &x.TaxiTrait)
	MapCustomToTaxiTrait(ctx, &x.TaxiTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToTaxiTrait(ctx jsonld.Context, x *schema.TaxiTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicTaxiTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToTaxiTrait(ctx jsonld.Context, x *schema.TaxiTrait) () {
	for k, v := range ctx.Current {
		f, ok := customTaxiTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}