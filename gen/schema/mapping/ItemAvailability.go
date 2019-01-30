package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsItemAvailability strings.Replacer
var strconvItemAvailability strconv.NumError

var basicItemAvailabilityTraitMapping = map[string]func(*schema.ItemAvailabilityTrait, []string){}
var customItemAvailabilityTraitMapping = map[string]func(*schema.ItemAvailabilityTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ItemAvailability", func(ctx jsonld.Context) (interface{}) {
		return NewItemAvailabilityFromContext(ctx)
	})

}

func NewItemAvailabilityFromContext(ctx jsonld.Context) (x schema.ItemAvailability) {
	x.Type = "http://schema.org/ItemAvailability"
	MapBasicToEnumerationTrait(ctx, &x.EnumerationTrait)
	MapCustomToEnumerationTrait(ctx, &x.EnumerationTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToItemAvailabilityTrait(ctx, &x.ItemAvailabilityTrait)
	MapCustomToItemAvailabilityTrait(ctx, &x.ItemAvailabilityTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToItemAvailabilityTrait(ctx jsonld.Context, x *schema.ItemAvailabilityTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicItemAvailabilityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToItemAvailabilityTrait(ctx jsonld.Context, x *schema.ItemAvailabilityTrait) () {
	for k, v := range ctx.Current {
		f, ok := customItemAvailabilityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}