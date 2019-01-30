package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMotorcycle strings.Replacer
var strconvMotorcycle strconv.NumError

var basicMotorcycleTraitMapping = map[string]func(*schema.MotorcycleTrait, []string){}
var customMotorcycleTraitMapping = map[string]func(*schema.MotorcycleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Motorcycle", func(ctx jsonld.Context) (interface{}) {
		return NewMotorcycleFromContext(ctx)
	})

}

func NewMotorcycleFromContext(ctx jsonld.Context) (x schema.Motorcycle) {
	x.Type = "http://schema.org/Motorcycle"
	MapBasicToVehicleTrait(ctx, &x.VehicleTrait)
	MapCustomToVehicleTrait(ctx, &x.VehicleTrait)

	MapBasicToProductTrait(ctx, &x.ProductTrait)
	MapCustomToProductTrait(ctx, &x.ProductTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMotorcycleTrait(ctx, &x.MotorcycleTrait)
	MapCustomToMotorcycleTrait(ctx, &x.MotorcycleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMotorcycleTrait(ctx jsonld.Context, x *schema.MotorcycleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMotorcycleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMotorcycleTrait(ctx jsonld.Context, x *schema.MotorcycleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMotorcycleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}