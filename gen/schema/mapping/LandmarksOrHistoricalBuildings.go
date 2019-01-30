package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLandmarksOrHistoricalBuildings strings.Replacer
var strconvLandmarksOrHistoricalBuildings strconv.NumError

var basicLandmarksOrHistoricalBuildingsTraitMapping = map[string]func(*schema.LandmarksOrHistoricalBuildingsTrait, []string){}
var customLandmarksOrHistoricalBuildingsTraitMapping = map[string]func(*schema.LandmarksOrHistoricalBuildingsTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LandmarksOrHistoricalBuildings", func(ctx jsonld.Context) (interface{}) {
		return NewLandmarksOrHistoricalBuildingsFromContext(ctx)
	})

}

func NewLandmarksOrHistoricalBuildingsFromContext(ctx jsonld.Context) (x schema.LandmarksOrHistoricalBuildings) {
	x.Type = "http://schema.org/LandmarksOrHistoricalBuildings"
	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLandmarksOrHistoricalBuildingsTrait(ctx, &x.LandmarksOrHistoricalBuildingsTrait)
	MapCustomToLandmarksOrHistoricalBuildingsTrait(ctx, &x.LandmarksOrHistoricalBuildingsTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLandmarksOrHistoricalBuildingsTrait(ctx jsonld.Context, x *schema.LandmarksOrHistoricalBuildingsTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLandmarksOrHistoricalBuildingsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLandmarksOrHistoricalBuildingsTrait(ctx jsonld.Context, x *schema.LandmarksOrHistoricalBuildingsTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLandmarksOrHistoricalBuildingsTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}