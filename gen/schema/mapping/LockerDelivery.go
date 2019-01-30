package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLockerDelivery strings.Replacer
var strconvLockerDelivery strconv.NumError

var basicLockerDeliveryTraitMapping = map[string]func(*schema.LockerDeliveryTrait, []string){}
var customLockerDeliveryTraitMapping = map[string]func(*schema.LockerDeliveryTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LockerDelivery", func(ctx jsonld.Context) (interface{}) {
		return NewLockerDeliveryFromContext(ctx)
	})

}

func NewLockerDeliveryFromContext(ctx jsonld.Context) (x schema.LockerDelivery) {
	x.Type = "http://schema.org/LockerDelivery"
	MapBasicToDeliveryMethodTrait(ctx, &x.DeliveryMethodTrait)
	MapCustomToDeliveryMethodTrait(ctx, &x.DeliveryMethodTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLockerDeliveryTrait(ctx, &x.LockerDeliveryTrait)
	MapCustomToLockerDeliveryTrait(ctx, &x.LockerDeliveryTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLockerDeliveryTrait(ctx jsonld.Context, x *schema.LockerDeliveryTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLockerDeliveryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLockerDeliveryTrait(ctx jsonld.Context, x *schema.LockerDeliveryTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLockerDeliveryTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}