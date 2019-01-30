package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsBusOrCoach strings.Replacer
var strconvBusOrCoach strconv.NumError

var basicBusOrCoachTraitMapping = map[string]func(*schema.BusOrCoachTrait, []string){}
var customBusOrCoachTraitMapping = map[string]func(*schema.BusOrCoachTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/BusOrCoach", func(ctx jsonld.Context) (interface{}) {
		return NewBusOrCoachFromContext(ctx)
	})


	basicBusOrCoachTraitMapping["http://schema.org/acrissCode"] = func(x *schema.BusOrCoachTrait, s []string) {
		x.AcrissCode = s[0]
	}



	customBusOrCoachTraitMapping["http://schema.org/roofLoad"] = func(x *schema.BusOrCoachTrait, ctx jsonld.Context, s string){
		var y schema.QuantitativeValue
		if strings.HasPrefix(s, "_:") {
			y = NewQuantitativeValueFromContext(ctx.SetCurrent(s))
			delete(ctx.Subjects, s)
		} else {
			y := schema.NewQuantitativeValue()
			y.Id = s
		}

		x.RoofLoad = &y

		return
	}
}

func NewBusOrCoachFromContext(ctx jsonld.Context) (x schema.BusOrCoach) {
	x.Type = "http://schema.org/BusOrCoach"
	MapBasicToVehicleTrait(ctx, &x.VehicleTrait)
	MapCustomToVehicleTrait(ctx, &x.VehicleTrait)

	MapBasicToProductTrait(ctx, &x.ProductTrait)
	MapCustomToProductTrait(ctx, &x.ProductTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToBusOrCoachTrait(ctx, &x.BusOrCoachTrait)
	MapCustomToBusOrCoachTrait(ctx, &x.BusOrCoachTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToBusOrCoachTrait(ctx jsonld.Context, x *schema.BusOrCoachTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicBusOrCoachTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToBusOrCoachTrait(ctx jsonld.Context, x *schema.BusOrCoachTrait) () {
	for k, v := range ctx.Current {
		f, ok := customBusOrCoachTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}