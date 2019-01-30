package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCityHall strings.Replacer
var strconvCityHall strconv.NumError

var basicCityHallTraitMapping = map[string]func(*schema.CityHallTrait, []string){}
var customCityHallTraitMapping = map[string]func(*schema.CityHallTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/CityHall", func(ctx jsonld.Context) (interface{}) {
		return NewCityHallFromContext(ctx)
	})

}

func NewCityHallFromContext(ctx jsonld.Context) (x schema.CityHall) {
	x.Type = "http://schema.org/CityHall"
	MapBasicToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)
	MapCustomToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCityHallTrait(ctx, &x.CityHallTrait)
	MapCustomToCityHallTrait(ctx, &x.CityHallTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCityHallTrait(ctx jsonld.Context, x *schema.CityHallTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCityHallTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCityHallTrait(ctx jsonld.Context, x *schema.CityHallTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCityHallTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}