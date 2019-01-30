package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsParkingFacility strings.Replacer
var strconvParkingFacility strconv.NumError

var basicParkingFacilityTraitMapping = map[string]func(*schema.ParkingFacilityTrait, []string){}
var customParkingFacilityTraitMapping = map[string]func(*schema.ParkingFacilityTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/ParkingFacility", func(ctx jsonld.Context) (interface{}) {
		return NewParkingFacilityFromContext(ctx)
	})

}

func NewParkingFacilityFromContext(ctx jsonld.Context) (x schema.ParkingFacility) {
	x.Type = "http://schema.org/ParkingFacility"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToParkingFacilityTrait(ctx, &x.ParkingFacilityTrait)
	MapCustomToParkingFacilityTrait(ctx, &x.ParkingFacilityTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToParkingFacilityTrait(ctx jsonld.Context, x *schema.ParkingFacilityTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicParkingFacilityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToParkingFacilityTrait(ctx jsonld.Context, x *schema.ParkingFacilityTrait) () {
	for k, v := range ctx.Current {
		f, ok := customParkingFacilityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}