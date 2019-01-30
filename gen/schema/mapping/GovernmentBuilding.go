package mapping

import (
	"strings"
	"strconv"
	"github.com/21stio/go-rdfgen/gen/schema"
	"github.com/21stio/go-rdfgen/pkg/jsonld"
)

var stringsGovernmentBuilding strings.Replacer
var strconvGovernmentBuilding strconv.NumError

var basicGovernmentBuildingTraitMapping = map[string]func(*schema.GovernmentBuildingTrait, []string){}
var customGovernmentBuildingTraitMapping = map[string]func(*schema.GovernmentBuildingTrait, jsonld.Context, string){}

func init() {
	jsonld.RegisterRdfType("http://schema.org/GovernmentBuilding", func(ctx jsonld.Context) (interface{}) {
		return NewGovernmentBuildingFromContext(ctx)
	})

}

func NewGovernmentBuildingFromContext(ctx jsonld.Context) (x schema.GovernmentBuilding) {
	x.Type = "http://schema.org/GovernmentBuilding"
	MapBasicToCivicStructureTrait(ctx, &x.CivicStructureTrait)
	MapCustomToCivicStructureTrait(ctx, &x.CivicStructureTrait)

	MapBasicToPlaceTrait(ctx, &x.PlaceTrait)
	MapCustomToPlaceTrait(ctx, &x.PlaceTrait)

	MapBasicToThingTrait(ctx, &x.ThingTrait)
	MapCustomToThingTrait(ctx, &x.ThingTrait)


	MapBasicToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)
	MapCustomToGovernmentBuildingTrait(ctx, &x.GovernmentBuildingTrait)

	MapBasicToMetaTrait(ctx, &x.MetaTrait)
	MapToAdditionalTrait(ctx, &x.AdditionalTrait)

	return
}

func MapBasicToGovernmentBuildingTrait(ctx jsonld.Context, x *schema.GovernmentBuildingTrait) () {
	for k, v := range ctx.Current {
		f, ok := basicGovernmentBuildingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, v)

		delete(ctx.Current, k)
	}

	return
}

func MapCustomToGovernmentBuildingTrait(ctx jsonld.Context, x *schema.GovernmentBuildingTrait) () {
	for k, v := range ctx.Current {
		f, ok := customGovernmentBuildingTraitMapping[k]
		if !ok {
			continue
		}
		f(x, ctx, v[0])

		delete(ctx.Current, k)
	}

	return
}