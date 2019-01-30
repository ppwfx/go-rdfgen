package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsLegislativeBuilding strings.Replacer
var strconvLegislativeBuilding strconv.NumError

var basicLegislativeBuildingTraitMapping = map[string]func(*schema.LegislativeBuildingTrait, []string){}
var customLegislativeBuildingTraitMapping = map[string]func(*schema.LegislativeBuildingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/LegislativeBuilding", func(ctx jsonld.Context) (interface{}) {
		return NewLegislativeBuildingFromContext(ctx)
	})

}

func NewLegislativeBuildingFromContext(ctx jsonld.Context) (x schema.LegislativeBuilding) {
	x.Type = "http://schema.org/LegislativeBuilding"
	MapBasicToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)
	MapCustomToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)

	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToLegislativeBuildingTrait(ctx, &x.LegislativeBuildingTrait)
	MapCustomToLegislativeBuildingTrait(ctx, &x.LegislativeBuildingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToLegislativeBuildingTrait(ctx jsonld.Context, x *schema.LegislativeBuildingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicLegislativeBuildingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToLegislativeBuildingTrait(ctx jsonld.Context, x *schema.LegislativeBuildingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customLegislativeBuildingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}