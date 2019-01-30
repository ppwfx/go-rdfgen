package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsDistance strings.Replacer
var strconvDistance strconv.NumError

var basicDistanceTraitMapping = map[string]func(*schema.DistanceTrait, []string){}
var customDistanceTraitMapping = map[string]func(*schema.DistanceTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Distance", func(ctx jsonld.Context) (interface{}) {
		return NewDistanceFromContext(ctx)
	})

}

func NewDistanceFromContext(ctx jsonld.Context) (x schema.Distance) {
	x.Type = "http://schema.org/Distance"
	MapBasicToQuantityTrait(ctx, &x.QuantityTrait)
	MapCustomToQuantityTrait(ctx, &x.QuantityTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToDistanceTrait(ctx, &x.DistanceTrait)
	MapCustomToDistanceTrait(ctx, &x.DistanceTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToDistanceTrait(ctx jsonld.Context, x *schema.DistanceTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicDistanceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToDistanceTrait(ctx jsonld.Context, x *schema.DistanceTrait) () {
	for k, v := range ctx.Current {
		f, ok := customDistanceTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}