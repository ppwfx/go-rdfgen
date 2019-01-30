package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsCity strings.Replacer
var strconvCity strconv.NumError

var basicCityTraitMapping = map[string]func(*schema.CityTrait, []string){}
var customCityTraitMapping = map[string]func(*schema.CityTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/City", func(ctx jsonld.Context) (interface{}) {
		return NewCityFromContext(ctx)
	})

}

func NewCityFromContext(ctx jsonld.Context) (x schema.City) {
	x.Type = "http://schema.org/City"
	MapBasicToAdministrativeAreaTrait(ctx, &x.AdministrativeAreaTrait)
	MapCustomToAdministrativeAreaTrait(ctx, &x.AdministrativeAreaTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToCityTrait(ctx, &x.CityTrait)
	MapCustomToCityTrait(ctx, &x.CityTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToCityTrait(ctx jsonld.Context, x *schema.CityTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicCityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToCityTrait(ctx jsonld.Context, x *schema.CityTrait) () {
	for k, v := range ctx.Current {
		f, ok := customCityTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}