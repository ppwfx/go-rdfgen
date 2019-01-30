package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsParcelService strings.Replacer
var strconvParcelService strconv.NumError

var basicParcelServiceTraitMapping = map[string]func(*schema.ParcelServiceTrait, []string){}
var customParcelServiceTraitMapping = map[string]func(*schema.ParcelServiceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ParcelService", func(ctx jsonld.Context) (interface{}) {
		return NewParcelServiceFromContext(ctx)
	})

}

func NewParcelServiceFromContext(ctx jsonld.Context) (x schema.ParcelService) {
	x.Type = "http://schema.org/ParcelService"
	MapBasicToDeliveryMethodTrait(ctx, &x.DeliveryMethodTrait)
	MapCustomToDeliveryMethodTrait(ctx, &x.DeliveryMethodTrait)

	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToParcelServiceTrait(ctx, &x.ParcelServiceTrait)
	MapCustomToParcelServiceTrait(ctx, &x.ParcelServiceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToParcelServiceTrait(ctx jsonld.Context, x *schema.ParcelServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicParcelServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToParcelServiceTrait(ctx jsonld.Context, x *schema.ParcelServiceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customParcelServiceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}