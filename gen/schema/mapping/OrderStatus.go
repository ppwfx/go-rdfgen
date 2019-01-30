package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsOrderStatus strings.Replacer
var strconvOrderStatus strconv.NumError

var basicOrderStatusTraitMapping = map[string]func(*schema.OrderStatusTrait, []string){}
var customOrderStatusTraitMapping = map[string]func(*schema.OrderStatusTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/OrderStatus", func(ctx jsonld.Context) (interface{}) {
		return NewOrderStatusFromContext(ctx)
	})

}

func NewOrderStatusFromContext(ctx jsonld.Context) (x schema.OrderStatus) {
	x.Type = "http://schema.org/OrderStatus"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToOrderStatusTrait(ctx, &x.OrderStatusTrait)
	MapCustomToOrderStatusTrait(ctx, &x.OrderStatusTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToOrderStatusTrait(ctx jsonld.Context, x *schema.OrderStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicOrderStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToOrderStatusTrait(ctx jsonld.Context, x *schema.OrderStatusTrait) () {
	for k, v := range ctx.Current {
		f, ok := customOrderStatusTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}