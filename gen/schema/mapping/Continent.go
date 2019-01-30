package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsContinent strings.Replacer
var strconvContinent strconv.NumError

var basicContinentTraitMapping = map[string]func(*schema.ContinentTrait, []string){}
var customContinentTraitMapping = map[string]func(*schema.ContinentTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/Continent", func(ctx jsonld.Context) (interface{}) {
		return NewContinentFromContext(ctx)
	})

}

func NewContinentFromContext(ctx jsonld.Context) (x schema.Continent) {
	x.Type = "http://schema.org/Continent"
	MapBasicToLandformTrait(ctx, &x.LandformTrait)
	MapCustomToLandformTrait(ctx, &x.LandformTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToContinentTrait(ctx, &x.ContinentTrait)
	MapCustomToContinentTrait(ctx, &x.ContinentTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToContinentTrait(ctx jsonld.Context, x *schema.ContinentTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicContinentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToContinentTrait(ctx jsonld.Context, x *schema.ContinentTrait) () {
	for k, v := range ctx.Current {
		f, ok := customContinentTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}