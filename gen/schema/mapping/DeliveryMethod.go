package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDeliveryMethod strings.Replacer
var strconvDeliveryMethod strconv.NumError

var basicDeliveryMethodTraitMapping = map[string]func(*schema.DeliveryMethodTrait, []string){}
var customDeliveryMethodTraitMapping = map[string]func(*schema.DeliveryMethodTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/DeliveryMethod", func(ctx jsonld.Context) (interface{}) {
		return NewDeliveryMethodFromContext(ctx)
	})

}

func NewDeliveryMethodFromContext(ctx jsonld.Context) (x schema.DeliveryMethod) {
	x.Type = "http://schema.org/DeliveryMethod"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDeliveryMethodTrait(ctx, &x.DeliveryMethodTrait)
	MapCustomToDeliveryMethodTrait(ctx, &x.DeliveryMethodTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDeliveryMethodTrait(ctx jsonld.Context, x *schema.DeliveryMethodTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDeliveryMethodTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDeliveryMethodTrait(ctx jsonld.Context, x *schema.DeliveryMethodTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDeliveryMethodTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}