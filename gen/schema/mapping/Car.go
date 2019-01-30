package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCar strings.Replacer
var strconvCar strconv.NumError

var basicCarTraitMapping = map[string]func(*schema.CarTrait, []string){}
var customCarTraitMapping = map[string]func(*schema.CarTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Car", func(ctx jsonld.Context) (interface{}) {
		return NewCarFromContext(ctx)
	})


	basicCarTraitMapping["http://schema.org/acrissCode"] = func(x *schema.CarTrait, s []string) {
		x.AcrissCode = s[0]
	}



	customCarTraitMapping["http://schema.org/roofLoad"] = func(x *schema.CarTrait, ctx jsonld.Context, s string){
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

func NewCarFromContext(ctx jsonld.Context) (x schema.Car) {
	x.Type = "http://schema.org/Car"
	MapBasicToVehicleTrait(ctx, &x.VehicleTrait)
	MapCustomToVehicleTrait(ctx, &x.VehicleTrait)

	MapBasicToProductTrait(ctx, &x.ProductTrait)
	MapCustomToProductTrait(ctx, &x.ProductTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCarTrait(ctx, &x.CarTrait)
	MapCustomToCarTrait(ctx, &x.CarTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCarTrait(ctx jsonld.Context, x *schema.CarTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCarTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCarTrait(ctx jsonld.Context, x *schema.CarTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCarTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}