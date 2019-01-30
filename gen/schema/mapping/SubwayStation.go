package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsSubwayStation strings.Replacer
var strconvSubwayStation strconv.NumError

var basicSubwayStationTraitMapping = map[string]func(*schema.SubwayStationTrait, []string){}
var customSubwayStationTraitMapping = map[string]func(*schema.SubwayStationTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/SubwayStation", func(ctx jsonld.Context) (interface{}) {
		return NewSubwayStationFromContext(ctx)
	})

}

func NewSubwayStationFromContext(ctx jsonld.Context) (x schema.SubwayStation) {
	x.Type = "http://schema.org/SubwayStation"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToSubwayStationTrait(ctx, &x.SubwayStationTrait)
	MapCustomToSubwayStationTrait(ctx, &x.SubwayStationTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToSubwayStationTrait(ctx jsonld.Context, x *schema.SubwayStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicSubwayStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToSubwayStationTrait(ctx jsonld.Context, x *schema.SubwayStationTrait) () {
	for k, v := range ctx.Current {
		f, ok := customSubwayStationTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}