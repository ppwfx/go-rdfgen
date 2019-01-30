package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsMotorizedBicycle strings.Replacer
var strconvMotorizedBicycle strconv.NumError

var basicMotorizedBicycleTraitMapping = map[string]func(*schema.MotorizedBicycleTrait, []string){}
var customMotorizedBicycleTraitMapping = map[string]func(*schema.MotorizedBicycleTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/MotorizedBicycle", func(ctx jsonld.Context) (interface{}) {
		return NewMotorizedBicycleFromContext(ctx)
	})

}

func NewMotorizedBicycleFromContext(ctx jsonld.Context) (x schema.MotorizedBicycle) {
	x.Type = "http://schema.org/MotorizedBicycle"
	MapBasicToVehicleTrait(ctx, &x.VehicleTrait)
	MapCustomToVehicleTrait(ctx, &x.VehicleTrait)

	MapBasicToProductTrait(ctx, &x.ProductTrait)
	MapCustomToProductTrait(ctx, &x.ProductTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToMotorizedBicycleTrait(ctx, &x.MotorizedBicycleTrait)
	MapCustomToMotorizedBicycleTrait(ctx, &x.MotorizedBicycleTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToMotorizedBicycleTrait(ctx jsonld.Context, x *schema.MotorizedBicycleTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicMotorizedBicycleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToMotorizedBicycleTrait(ctx jsonld.Context, x *schema.MotorizedBicycleTrait) () {
	for k, v := range ctx.Current {
		f, ok := customMotorizedBicycleTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}