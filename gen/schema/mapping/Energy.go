package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsEnergy strings.Replacer
var strconvEnergy strconv.NumError

var basicEnergyTraitMapping = map[string]func(*schema.EnergyTrait, []string){}
var customEnergyTraitMapping = map[string]func(*schema.EnergyTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Energy", func(ctx jsonld.Context) (interface{}) {
		return NewEnergyFromContext(ctx)
	})

}

func NewEnergyFromContext(ctx jsonld.Context) (x schema.Energy) {
	x.Type = "http://schema.org/Energy"
	MapBasicToQuantityTrait(ctx, &x.QuantityTrait)
	MapCustomToQuantityTrait(ctx, &x.QuantityTrait)

	MapBasicToIntangibleTrait(ctx, &x.IntangibleTrait)
	MapCustomToIntangibleTrait(ctx, &x.IntangibleTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToEnergyTrait(ctx, &x.EnergyTrait)
	MapCustomToEnergyTrait(ctx, &x.EnergyTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToEnergyTrait(ctx jsonld.Context, x *schema.EnergyTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicEnergyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToEnergyTrait(ctx jsonld.Context, x *schema.EnergyTrait) () {
	for k, v := range ctx.Current {
		f, ok := customEnergyTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}