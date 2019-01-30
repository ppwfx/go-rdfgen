package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMass strings.Replacer
var strconvMass strconv.NumError

var basicMassTraitMapping = map[string]func(*schema.MassTrait, []string){}
var customMassTraitMapping = map[string]func(*schema.MassTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Mass", func(ctx jsonld.Context) (interface{}) {
		return NewMassFromContext(ctx)
	})

}

func NewMassFromContext(ctx jsonld.Context) (x schema.Mass) {
	x.Type = "http://schema.org/Mass"
	MapBasicToQuantityTrait(ctx, &x.QuantityTrait)
	MapCustomToQuantityTrait(ctx, &x.QuantityTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMassTrait(ctx, &x.MassTrait)
	MapCustomToMassTrait(ctx, &x.MassTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMassTrait(ctx jsonld.Context, x *schema.MassTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMassTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMassTrait(ctx jsonld.Context, x *schema.MassTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMassTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}