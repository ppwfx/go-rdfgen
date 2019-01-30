package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsPond strings.Replacer
var strconvPond strconv.NumError

var basicPondTraitMapping = map[string]func(*schema.PondTrait, []string){}
var customPondTraitMapping = map[string]func(*schema.PondTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Pond", func(ctx jsonld.Context) (interface{}) {
		return NewPondFromContext(ctx)
	})

}

func NewPondFromContext(ctx jsonld.Context) (x schema.Pond) {
	x.Type = "http://schema.org/Pond"
	MapBasicToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)
	MapCustomToBodyOfWaterTrait(ctx, &x.BodyOfWaterTrait)

	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToPondTrait(ctx, &x.PondTrait)
	MapCustomToPondTrait(ctx, &x.PondTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToPondTrait(ctx jsonld.Context, x *schema.PondTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicPondTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToPondTrait(ctx jsonld.Context, x *schema.PondTrait) () {
	for k, v := range ctx.Current {
		f, ok := customPondTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}