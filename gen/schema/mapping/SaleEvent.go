package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSaleEvent strings.Replacer
var strconvSaleEvent strconv.NumError

var basicSaleEventTraitMapping = map[string]func(*schema.SaleEventTrait, []string){}
var customSaleEventTraitMapping = map[string]func(*schema.SaleEventTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SaleEvent", func(ctx jsonld.Context) (interface{}) {
		return NewSaleEventFromContext(ctx)
	})

}

func NewSaleEventFromContext(ctx jsonld.Context) (x schema.SaleEvent) {
	x.Type = "http://schema.org/SaleEvent"
	MapBasicToEventTrait(ctx, &x.EventTrait)
	MapCustomToEventTrait(ctx, &x.EventTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSaleEventTrait(ctx, &x.SaleEventTrait)
	MapCustomToSaleEventTrait(ctx, &x.SaleEventTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSaleEventTrait(ctx jsonld.Context, x *schema.SaleEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSaleEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSaleEventTrait(ctx jsonld.Context, x *schema.SaleEventTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSaleEventTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}